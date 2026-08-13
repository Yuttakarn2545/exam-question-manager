package repository

import (
	"sync"

	"exam-question-manager/internal/model"
)

type ExamRepository interface {
	List() []model.Exam
	Insert(exam model.Exam) (model.Exam, error)
	Delete(id int64) (bool, error)
	ReplaceAll(exams []model.Exam) error
}

type InMemoryExamRepository struct {
	mu     sync.Mutex
	nextID int64
	exams  []model.Exam
}

func NewInMemoryExamRepository() *InMemoryExamRepository {
	return &InMemoryExamRepository{nextID: 1}
}

func (r *InMemoryExamRepository) List() []model.Exam {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make([]model.Exam, len(r.exams))
	copy(out, r.exams)
	return out
}

func (r *InMemoryExamRepository) Insert(exam model.Exam) (model.Exam, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	exam.ID = r.nextID
	r.nextID++
	r.exams = append(r.exams, exam)
	return exam, nil
}

func (r *InMemoryExamRepository) Delete(id int64) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, e := range r.exams {
		if e.ID == id {
			r.exams = append(r.exams[:i], r.exams[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (r *InMemoryExamRepository) ReplaceAll(exams []model.Exam) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.exams = exams
	return nil
}
