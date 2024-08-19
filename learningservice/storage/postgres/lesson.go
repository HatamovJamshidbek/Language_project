package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	pb "learn_service/genproto/learning"
	"learn_service/pkg/logger"
	"log/slog"

	"github.com/google/uuid"
)

type LearnManagementRepository interface {
	GetLanguages(*pb.GetLanguageRequest) (*pb.GetLanguageResponse, error)
	StartLearnLanguage(req *pb.StartLanguageRequest) (*pb.StartLanguageResponse, error)
	GetLessonsList(req *pb.Language) (*pb.LessonsListResponse, error)
	GetInfoLessons(req *pb.GetInfoLessonsResponse) (*pb.LessonsInfoRepository, error)
	CompleteLesson(req *pb.Lesson) (*pb.ComplateLessonResponse, error)
	GetLanguageExercises(req *pb.Language) (*pb.LanguageExerciseResponse, error)
	SubmitExerciseAnswer(req *pb.ExerciseAnswerRequest) (*pb.ExerciseAnswerResponse, error)
	GetVocabularyList(req *pb.Language) (*pb.VocabularyListResponse, error)
	AddVocabulary(req *pb.AddVocabularyRequest) (*pb.AddVocabularyResponse, error)
}
type LearnRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewLearnRepository(db *sql.DB) *LearnRepository {
	return &LearnRepository{
		Db:  db,
		Log: logger.NewLogger(),
	}
}

func (r *LearnRepository) GetLanguages(req *pb.GetLanguageRequest) (*pb.GetLanguageResponse, error) {
	offset := (req.Page - 1) * req.Limit

	q := `SELECT 
			code, 
			name, 
			flag_emojis 
		FROM 
			languages 
			LIMIT $1 
			OFFSET $2 `

	rows, err := r.Db.Query(q, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var languages []*pb.Language
	for rows.Next() {
		var lang pb.Language
		if err := rows.Scan(&lang.Code, &lang.Name, &lang.FlagEmogi); err != nil {
			return nil, err
		}
		languages = append(languages, &lang)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.GetLanguageResponse{
		Languages: languages,
	}, nil
}

// StartLearnLanguage yangi o'quv sessiyasini boshlaydi
func (r *LearnRepository) StartLearnLanguage(req *pb.StartLanguageRequest) (*pb.StartLanguageResponse, error) {
	q := `INSERT INTO languages (
					name,
					code, 
					level
				) 
			VALUES ($1, $2, $3)`

	_, err := r.Db.Exec(q, req.LanguageName, req.LanguageCode, req.InitalLevel)
	if err != nil {
		return nil, err
	}
	var (
		id    string
		title string
	)
	q1 := `
	SELECT 
    	id,
    	title,
		content
	FROM 
    	lessons
	WHERE 
    	languages_code = $1 AND
    	level = $2 
	LIMIT 1
 	`
	var resp pb.Info
	err = r.Db.QueryRow(q1, req.LanguageCode, req.InitalLevel).Scan(&id, &title, &resp)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	// Dummy javob qaytarish; kerak bo'lsa sozlang
	return &pb.StartLanguageResponse{
		Message:      "O'quv sessiyasi boshlandi",
		CurrentLavel: req.InitalLevel,
		NextLesson:   &resp,
	}, nil
}

// GetLessonsList muayyan til uchun barcha darslarni qaytaradi
func (r *LearnRepository) GetLessonsList(req *pb.Language) (*pb.LessonsListResponse, error) {
	q := `
	SELECT 
		id, 
		title, 
		level, 
		completed 
	FROM 
		lessons 
	WHERE 
		languages_code = $1 `

	rows, err := r.Db.Query(q, req.Code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*pb.Lesson
	for rows.Next() {
		var lesson pb.Lesson
		if err := rows.Scan(
			&lesson.Id,
			&lesson.Title,
			&lesson.Level,
			&lesson.Completed,
		); err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.LessonsListResponse{
		Language: req.Code,
		Lessons:  lessons,
	}, nil
}

// GetInfoLessons dars haqida batafsil ma'lumotlarni qaytaradi
func (r *LearnRepository) GetInfoLessons(req *pb.GetInfoLessonsResponse) (*pb.LessonsInfoRepository, error) {
	resp := pb.LessonsInfoRepository{}
	q := `
	select 
		id, 
		title, 
		languages_code,  
		level,
		content 
	from 
		lessons 
	where 
		id = $1`

	var content string
	err := r.Db.QueryRow(q, req.LessonId).Scan(
		&resp.Id,
		&resp.Title,
		&resp.Lenguage,
		&resp.Level,
		&content,
	)
	if err != nil {
		return nil, err
	}

	// Agar content ustuni bo'lmasa, bu qismini olib tashlang
	err = json.Unmarshal([]byte(content), &resp.Content)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	fmt.Println(resp.Content)

	// Boshqa joylarda ham lenguage o'rniga language ishlatiladi
	return &pb.LessonsInfoRepository{
		Id:       resp.Id,
		Title:    resp.Title,
		Lenguage: resp.Lenguage,
		Level:    resp.Level,
		Content:  resp.Content,
	}, nil
}

// CompleteLesson darsni tugatgan deb belgilaydi
// Progress Service ni ishlatganda xp ni berib ketamiz
func (r *LearnRepository) CompleteLesson(req *pb.Lesson) (*pb.ComplateLessonResponse, error) {
	q := `
	UPDATE 
		lessons 
	SET 
		completed = true 
	WHERE id = $1`
	_, err := r.Db.Exec(q, req.Id)
	if err != nil {
		return nil, err
	}

	// Dummy javob qaytarish; kerak bo'lsa sozlang
	return &pb.ComplateLessonResponse{
		Message:    "Dars yakunlandi",
		XpEarned:   10,
		NextLesson: &pb.Info{Id: "2", Title: "Keyingi Dars"},
	}, nil
}

// GetLanguageExercises muayyan til uchun mashqlarni qaytaradi
func (r *LearnRepository) GetLanguageExercises(req *pb.Language) (*pb.LanguageExerciseResponse, error) {
	q := `
	SELECT 
		ex.id,
		ex.type, 
		ex.question, 
		ex.correet_answer 
	FROM 
		exercises as ex
	INNER JOIN
		lessons as l
	ON
		l.id = ex.lesson_id
	WHERE 
		l.languages_code = $1 `
	rows, err := r.Db.Query(q, req.Code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []*pb.Exercises
	for rows.Next() {
		var exercise pb.Exercises
		if err := rows.Scan(
			&exercise.Id,
			&exercise.Type,
			&exercise.Question,
			&exercise.CorrectAnswer,
		); err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.LanguageExerciseResponse{
		Language:  req.Code,
		Exercises: exercises,
	}, nil
}

// SubmitExerciseAnswer kiritilgan javob to'g'ri yoki yo'qligini tekshiradi
func (r *LearnRepository) SubmitExerciseAnswer(req *pb.ExerciseAnswerRequest) (*pb.ExerciseAnswerResponse, error) {
	correct_answer := ""

	q := `
	SELECT 
		correet_answer 
	FROM 
		exercises 
	WHERE 
		id = $1
	`

	err := r.Db.QueryRow(q, req.ExerciseId).Scan(&correct_answer)
	if err != nil {
		return nil, err
	}

	isCorrect := true
	explanation := ""
	xpEarned := 0

	if correct_answer == req.Answer {
		isCorrect = true
		explanation = "Javob to'g'ri"
		xpEarned = 5
	} else {
		isCorrect = false
		explanation = "Javob xato"
		xpEarned = 0
	}

	// bu yerda exp ni bergandan keyin progres service ga beriladi
	return &pb.ExerciseAnswerResponse{
		Correct:     isCorrect,
		Explanation: explanation,
		XpEarned:    int32(xpEarned),
	}, nil
}

// GetVocabularyList muayyan til uchun lug'atni qaytaradi
func (r *LearnRepository) GetVocabularyList(req *pb.Language) (*pb.VocabularyListResponse, error) {
	q := `
	SELECT 
		word, 
		translation, 
		part_of_speech, 
		example_sentence, 
		audio_url 
	FROM 
		vocabulary 
	WHERE 
		language_code = $1
	`
	rows, err := r.Db.Query(q, req.Code)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var vocabulary pb.SlcVocabulary
	for rows.Next() {
		vocab := pb.Vocabulary{}
		if err := rows.Scan(
			&vocab.Word,
			&vocab.Translation,
			&vocab.PartOfSpeech,
			&vocab.ExampleSentence,
			&vocab.AudioUrl,
		); err != nil {
			return nil, err
		}
		vocabulary.Vocabularys = append(vocabulary.Vocabularys, &vocab)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.VocabularyListResponse{
		Language:   req.Code,
		Vocabulary: vocabulary.Vocabularys,
	}, nil
}

func (r *LearnRepository) AddVocabulary(req *pb.AddVocabularyRequest) (*pb.AddVocabularyResponse, error) {
	// Yangi lug'at so'zining ID sini yaratamiz
	fmt.Println("++++++++++++++VVVVVVVVV")
	wordId := uuid.New().String()
	if r.Db == nil {
		fmt.Println(")))))))))))))))))+++++++++++")
	}

	q := `
	INSERT INTO vocabulary (
							id, 
							word, 
							translation, 
							part_of_speech, 
							example_sentence,
							language_code) 
	VALUES 					(
							$1, 
							$2, 
							$3, 
							$4, 
							$5,
							$6)`

	_, err := r.Db.Exec(
		q,
		wordId,
		req.Word,
		req.Translation,
		req.PartOfSpeech,
		req.ExampleSentence,
		req.Languagecode,
	)
	if err != nil {
		return nil, err
	}

	return &pb.AddVocabularyResponse{
		Message: "Lug'at so'zi muvaffaqiyatli qo'shildi",
		WordId:  wordId,
	}, nil
}
