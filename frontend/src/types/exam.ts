export interface Exam {
  id: number
  number: number
  question: string
  answers: [string, string, string, string]
}

export interface ExamInput {
  question: string
  answers: [string, string, string, string]
}

export interface PagedExams {
  data: Exam[]
  page: number
  pageSize: number
  total: number
  totalPages: number
}
