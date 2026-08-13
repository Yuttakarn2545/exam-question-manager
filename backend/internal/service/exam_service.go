package service

import (
	"errors"
	"sort"
	"strings"

	"exam-question-manager/internal/model"
	"exam-question-manager/internal/repository"
)

var ErrExamNotFound = errors.New("exam not found")

const (
	DefaultPageSize = 10
	MaxPageSize     = 50

	MaxQuestionLength = 500
	MaxAnswerLength   = 200
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string { return e.Message }

type ExamService struct {
	repo repository.ExamRepository
}

func NewExamService(repo repository.ExamRepository) *ExamService {
	return &ExamService{repo: repo}
}

// List returns a page of exams ordered by their number.
func (s *ExamService) List(page, pageSize int) model.PagedExams {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	exams := s.listAll()
	total := len(exams)
	totalPages := (total + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}

	return model.PagedExams{
		Data:       exams[start:end],
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
}

func (s *ExamService) Create(input model.ExamInput) (model.Exam, error) {
	if err := validateInput(input); err != nil {
		return model.Exam{}, err
	}

	exam := model.Exam{
		Number:   len(s.listAll()) + 1,
		Question: input.Question,
		Answers:  input.Answers,
	}
	return s.repo.Insert(exam)
}

func (s *ExamService) Delete(id int64) error {
	deleted, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !deleted {
		return ErrExamNotFound
	}
	return s.renumber()
}

func (s *ExamService) listAll() []model.Exam {
	exams := s.repo.List()
	sort.Slice(exams, func(i, j int) bool { return exams[i].Number < exams[j].Number })
	return exams
}

func (s *ExamService) renumber() error {
	exams := s.listAll()
	for i := range exams {
		exams[i].Number = i + 1
	}
	return s.repo.ReplaceAll(exams)
}

func validateInput(input model.ExamInput) error {
	question := strings.TrimSpace(input.Question)
	if question == "" {
		return &ValidationError{Message: "question is required"}
	}
	if len(question) > MaxQuestionLength {
		return &ValidationError{Message: "question must be at most 500 characters"}
	}

	for i, answer := range input.Answers {
		trimmed := strings.TrimSpace(answer)
		label := "answer " + string(rune('1'+i))
		if trimmed == "" {
			return &ValidationError{Message: label + " is required"}
		}
		if len(trimmed) > MaxAnswerLength {
			return &ValidationError{Message: label + " must be at most 200 characters"}
		}
	}
	return nil
}
