package email

import (
	"auth_service/config"
	redis "auth_service/storage/redis"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

// Email funksiyasi berilgan email manziliga tasdiqlash kodini yuboradi va Redisda saqlaydi.
func Email(email string) (string, error) {
	fmt.Println(3)
	// Redis clientni yaratish
	client := redis.RedisConn()
	defer client.Close()

	// Tasdiqlash kodi yaratish
	randomNumber := rand.Intn(900000) + 100000
	code := strconv.Itoa(randomNumber)

	// Tasdiqlash kodini Redisda saqlash

	err := client.Set(context.Background(), "email:"+email, code, time.Minute*5).Err()
	if err != nil {
		return "", err
	}

	err = client.Set(context.Background(), "e:email", email, time.Minute*5).Err()
	if err != nil {
		return "", err
	}
	// Emailni redisga saqlash
	// err := client.Set(context.Background(), code, email, time.Minute*5).Err()
	// if err != nil {
	// 	return "", err
	// }

	// Redisdan kalitni tekshirish
	code1, err := client.Get(context.Background(), "email:"+email).Result()
	if err != nil {
		return "", err
	}
	fmt.Println(code1, "sadfasdfasdfadsfsadfasd")
	// Emailga kod yuborish
	err = SendCode(email, code1)

	if err != nil {
		return "", err
	}

	return code1, nil
}

// SendCode funksiyasi berilgan email manziliga tasdiqlash kodini yuboradi.
func SendCode(email string, code string) error {
	// Email yuboruvchi ma'lumotlari
	from := config.Load().EMAIL        // Shaxsiy emailni kiriting
	password := config.Load().PASSWORD // Email parolini kiriting

	// Qabul qiluvchi email manzili
	to := []string{
		email,
	}

	// SMTP server konfiguratsiyasi
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// SMTP autentifikatsiyasi
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Email body ni tayyorlash
	var body bytes.Buffer

	// MIME sarlavhalari

	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", code)))

	// Emailni yuborish
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
		return err
	}

	return nil
}
