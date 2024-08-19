package model

type GetLessonsListRequest struct {
	Code string `json:"code"`
}

type GetInfoLessonsRequest struct {
	Id string `json:"id"`
}

type GetLanguageExercisesRequest struct {
	Code string `json:"code"`
}

type GetVocabularyListRequest struct {
	Code string `json:"code"`
}

type UpdateUser struct {
	FullName      string `json:"full_name`
	LeanguageCode string `json:language_code`
}

type ExerciseAnswerRequest struct {
	Answer string `json:"answer"`
}
