# Exam Question Manager

A small full-stack application for managing multiple-choice exam questions, built to demonstrate **Vue 3 + TypeScript frontend development, Go/Fiber REST API design, validation, persistence, and layered backend architecture**.

## Why this project matters

This repository is a compact example of how I structure a full-stack feature from UI through API and persistence.

The backend follows:

```text
handler → service → repository
```

That separation keeps transport, business rules, and storage concerns independent and makes the persistence layer replaceable without rewriting the service or HTTP layer.

## Tech Stack

### Frontend
- Vue 3
- TypeScript
- Vite
- Tailwind CSS 4
- vue-tsc type checking

### Backend
- Go
- Fiber
- Repository abstraction
- JSON-file persistence
- Environment-based configuration

## Features

- List exam questions
- Add a question with 4 multiple-choice answers
- Delete questions
- Automatic sequential renumbering after deletion
- Pagination through `page` and `pageSize`
- Input validation
  - Question ≤ 500 characters
  - Answers ≤ 200 characters
- Structured API errors
- Persistence that survives server restarts
- Local UI state updates after add/delete without a full list re-fetch

## Project Structure

```text
backend/
  cmd/server/              application entry point
  internal/handler/        HTTP layer
  internal/service/        business logic
  internal/repository/     persistence abstraction + implementation
  data/                    local persistence

frontend/
  src/                     Vue application
```

## Architecture Note

The repository layer is defined behind an interface. Moving from JSON-file persistence to PostgreSQL or another database would only require another repository implementation; the service and handler layers can remain unchanged.

## Run Locally

### Backend

```bash
cd backend
go run ./cmd/server
```

Backend runs on:

```text
http://localhost:8080
```

API prefix:

```text
/api/v1
```

Configuration is read from `backend/.env`, including `PORT` and `DATA_FILE`.

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend runs on:

```text
http://localhost:5173
```

The API URL and page size can be configured through `frontend/.env`.

## Verification

Frontend type checking and build:

```bash
cd frontend
npm run type-check
npm run build
```

Backend:

```bash
cd backend
go test ./...
```

## Documented Assumption

The source specification does not provide a field for explicitly selecting the correct answer, so **Answer 1 is treated as the correct answer**. This is documented instead of silently inventing additional product behavior.

## Engineering Takeaways

This project demonstrates:

- Full-stack feature delivery
- REST API design
- Layered Go architecture
- Type-safe Vue development
- Validation and structured error handling
- State-management decisions that avoid unnecessary requests
- Designing code so infrastructure can be replaced later

## More

Portfolio: https://www.webbyyu.net  
LinkedIn: https://www.linkedin.com/in/yuttakan-phunkhlang/
