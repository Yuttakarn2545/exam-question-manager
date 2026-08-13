# Exam Question Manager

Full-stack app for managing multiple-choice exam questions (spec pages IT 08-1 / IT 08-2). Vue 3 + TypeScript frontend, Go (Fiber) backend.

## Structure

```
backend/     Go REST API (Fiber) — layered: handler → service → repository
frontend/    Vue 3 + TypeScript + Tailwind CSS
```

## Features

- List / add / delete exam questions (question + 4 multiple-choice answers)
- Deleting a question renumbers the remaining questions sequentially
- Pagination (`page`, `pageSize` query params)
- JSON-file persistence (`backend/data/exams.json`, created on first write — survives restarts)
- Structured API errors: `{"error": {"code": "...", "message": "..."}}`
- Input length validation (question ≤ 500 chars, answers ≤ 200 chars)

## Running locally

### Backend

```bash
cd backend
go run ./cmd/server
```

Runs on `http://localhost:8080`, API under `/api/v1`. Config via `backend/.env` (`PORT`, `DATA_FILE`).

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Runs on `http://localhost:5173`. Config via `frontend/.env` (`VITE_API_URL`, `VITE_PAGE_SIZE`).

## Design notes

- `ExamRepository` (`backend/internal/repository`) is an interface — swapping the JSON-file store for a real database only requires a new implementation, no changes to `service`/`handler`.
- The frontend avoids a full list re-fetch after add/delete; it updates local state directly from the API response instead.
- "Answer 1" is treated as the correct answer — the original spec's form has no field to mark which answer is correct, so this is a documented assumption.
