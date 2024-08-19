package service

import (
	"context"
	"fmt"
	pb "learn_service/genproto/learning"
	"learn_service/storage/postgres"
	"log/slog"
)

type LearnManagementService interface {
	GetLanguages(ctx context.Context, req *pb.GetLanguageRequest) (*pb.GetLanguageResponse, error)
	StartLearnLanguage(ctx context.Context, req *pb.StartLanguageRequest) (*pb.StartLanguageResponse, error)
	GetLessonsList(ctx context.Context, req *pb.Language) (*pb.LessonsListResponse, error)
	GetInfoLessons(ctx context.Context, req *pb.GetInfoLessonsResponse) (*pb.LessonsInfoRepository, error)
	CompleteLesson(ctx context.Context, req *pb.Lesson) (*pb.ComplateLessonResponse, error)
	GetLanguageExercises(ctx context.Context, req *pb.Language) (*pb.LanguageExerciseResponse, error)
	SubmitExerciseAnswer(ctx context.Context, req *pb.ExerciseAnswerRequest) (*pb.ExerciseAnswerResponse, error)
	GetVocabularyList(ctx context.Context, req *pb.Language) (*pb.VocabularyListResponse, error)
	AddVocabulary(ctx context.Context, req *pb.AddVocabularyRequest) (*pb.AddVocabularyResponse, error)
}

type LearningService struct {
	pb.UnimplementedLearningServiceServer
	LearnRepo postgres.LearnManagementRepository
	Logger    *slog.Logger
}

func NewLearningService(learnRepo postgres.LearnManagementRepository, logger *slog.Logger) LearnManagementService {
	return &LearningService{LearnRepo: learnRepo, Logger: logger}
}

func (s *LearningService) GetLanguages(ctx context.Context, req *pb.GetLanguageRequest) (*pb.GetLanguageResponse, error) {
	resp, err := s.LearnRepo.GetLanguages(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) StartLearnLanguage(ctx context.Context, req *pb.StartLanguageRequest) (*pb.StartLanguageResponse, error) {
	resp, err := s.LearnRepo.StartLearnLanguage(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetLessonsList(ctx context.Context, req *pb.Language) (*pb.LessonsListResponse, error) {
	resp, err := s.LearnRepo.GetLessonsList(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetInfoLessons(ctx context.Context, req *pb.GetInfoLessonsResponse) (*pb.LessonsInfoRepository, error) {
	resp, err := s.LearnRepo.GetInfoLessons(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) CompleteLesson(ctx context.Context, req *pb.Lesson) (*pb.ComplateLessonResponse, error) {
	resp, err := s.LearnRepo.CompleteLesson(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *LearningService) GetLanguageExercises(ctx context.Context, req *pb.Language) (*pb.LanguageExerciseResponse, error) {
	resp, err := s.LearnRepo.GetLanguageExercises(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) SubmitExerciseAnswer(ctx context.Context, req *pb.ExerciseAnswerRequest) (*pb.ExerciseAnswerResponse, error) {
	resp, err := s.LearnRepo.SubmitExerciseAnswer(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetVocabularyList(ctx context.Context, req *pb.Language) (*pb.VocabularyListResponse, error) {
	resp, err := s.LearnRepo.GetVocabularyList(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) AddVocabulary(ctx context.Context, req *pb.AddVocabularyRequest) (*pb.AddVocabularyResponse, error) {
	fmt.Println("+++++++++++++++")
	resp, err := s.LearnRepo.AddVocabulary(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
