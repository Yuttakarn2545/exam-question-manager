import type { Exam, ExamInput, PagedExams } from '../types/exam'

const BASE_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080/api/v1'

interface ApiErrorBody {
  error?: { code?: string; message?: string }
}

async function extractErrorMessage(res: Response): Promise<string> {
  try {
    const body = (await res.json()) as ApiErrorBody
    if (body?.error?.message) return body.error.message
  } catch {
    // response wasn't JSON — fall back below
  }
  return res.statusText
}

async function handle<T>(res: Response): Promise<T> {
  if (!res.ok) {
    throw new Error(await extractErrorMessage(res))
  }
  return res.json() as Promise<T>
}

export function fetchExams(page = 1, pageSize = 10): Promise<PagedExams> {
  return fetch(`${BASE_URL}/exams?page=${page}&pageSize=${pageSize}`).then((res) =>
    handle<PagedExams>(res),
  )
}

export function createExam(input: ExamInput): Promise<Exam> {
  return fetch(`${BASE_URL}/exams`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(input),
  }).then((res) => handle<Exam>(res))
}

export function deleteExam(id: number): Promise<void> {
  return fetch(`${BASE_URL}/exams/${id}`, { method: 'DELETE' }).then(async (res) => {
    if (!res.ok) throw new Error(await extractErrorMessage(res))
  })
}
