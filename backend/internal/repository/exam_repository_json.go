package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"exam-question-manager/internal/model"
)

// JSONFileExamRepository keeps exams in memory for fast reads and writes the
// full snapshot to a JSON file on every mutation, so data survives a restart.
// Same read/write behavior as InMemoryExamRepository — just backed by a file.
type JSONFileExamRepository struct {
	mu     sync.Mutex
	path   string
	nextID int64
	exams  []model.Exam
}

type jsonFileContents struct {
	NextID int64        `json:"nextId"`
	Exams  []model.Exam `json:"exams"`
}

// NewJSONFileExamRepository loads existing data from path if present,
// or starts empty if the file doesn't exist yet.
func NewJSONFileExamRepository(path string) (*JSONFileExamRepository, error) {
	r := &JSONFileExamRepository{path: path, nextID: 1}

	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return r, nil
	}
	if err != nil {
		return nil, err
	}

	var contents jsonFileContents
	if err := json.Unmarshal(raw, &contents); err != nil {
		return nil, err
	}

	r.exams = contents.Exams
	r.nextID = contents.NextID
	if r.nextID < 1 {
		r.nextID = 1
	}
	return r, nil
}

func (r *JSONFileExamRepository) List() []model.Exam {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make([]model.Exam, len(r.exams))
	copy(out, r.exams)
	return out
}

func (r *JSONFileExamRepository) Insert(exam model.Exam) (model.Exam, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	exam.ID = r.nextID
	r.nextID++
	r.exams = append(r.exams, exam)

	if err := r.persistLocked(); err != nil {
		return model.Exam{}, err
	}
	return exam, nil
}

func (r *JSONFileExamRepository) Delete(id int64) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, e := range r.exams {
		if e.ID == id {
			r.exams = append(r.exams[:i], r.exams[i+1:]...)
			if err := r.persistLocked(); err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (r *JSONFileExamRepository) ReplaceAll(exams []model.Exam) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.exams = exams
	return r.persistLocked()
}

// persistLocked writes the current in-memory state to disk.
// Caller must already hold r.mu.
func (r *JSONFileExamRepository) persistLocked() error {
	contents := jsonFileContents{NextID: r.nextID, Exams: r.exams}

	raw, err := json.MarshalIndent(contents, "", "  ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(r.path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(r.path, raw, 0o644)
}
