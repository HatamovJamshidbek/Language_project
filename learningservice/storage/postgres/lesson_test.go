package postgres

import (
	"fmt"
	pb "learn_service/genproto/learning"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLanguages(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	s := NewLearnRepository(db)

	req := pb.GetLanguageRequest{
		Page:  1,
		Limit: 4,
	}

	res, err := s.GetLanguages(&req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestStartLearnLanguage(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := NewLearnRepository(db)

	req := pb.StartLanguageRequest{
		LanguageCode: "FR",
		InitalLevel:  "Beginner",
	}

	res, err := r.StartLearnLanguage(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.NextLesson.Title)
}

func TestGetLessonsList(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}

	r := NewLearnRepository(db)

	req := pb.Language{
		Code: "FR",
	}

	res, err := r.GetLessonsList(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Lessons)
}

// func TestGetInfoLessons(t *testing.T) {
// 	db, err := ConnectionDb()
// 	if err != nil {
// 		panic(err)
// 	}

// 	r := NewLearnRepository(db)

// 	req := pb.GetInfoLessonsResponse{
// 		LessonId: "282148e2-68d6-4908-a3a3-a7bf32fc2af5",
// 	}

// 	res, err := r.GetInfoLessons(&req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(res)

// }

func TestGetLanguageExercises(t *testing.T) {
	db, err := ConnectionDb()
	assert.NoError(t, err)

	defer db.Close()

	r := NewLearnRepository(db)

	req := pb.Language{
		Code: "UZ",
	}

	res, err := r.GetLanguageExercises(&req)
	assert.NoError(t, err)

	assert.NotEmpty(t, res)
}

func TestSubmitExerciseAnswer(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := NewLearnRepository(db)

	req := pb.ExerciseAnswerRequest{
		Answer:     "Jupiter",
		ExerciseId: "1cb6d9a4-963d-41b4-89c2-30d04ded626d",
	}

	res, err := r.SubmitExerciseAnswer(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestGetVocabularyList(t *testing.T) {

	//  bu kod ishlaydi lkn databaseda molumot yoq ligi uchun to`g'ridan to'g'ri ishlatolmaymiz
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}

	r := NewLearnRepository(db)
	req := pb.Language{
		Code: "EN",
	}

	res, err := r.GetVocabularyList(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestAddVocabulary(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}

	r := NewLearnRepository(db)

	req := pb.AddVocabularyRequest{
		Word:            "Baguette",
		Translation:     "French bread",
		PartOfSpeech:    "noun",
		ExampleSentence: "Jachète une baguette à la boulangerie.",
		Languagecode:    "FR",
	}

	res, err := r.AddVocabulary(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
