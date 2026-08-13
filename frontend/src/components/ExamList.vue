<template>
  <section class="w-full max-w-md overflow-hidden rounded bg-white shadow">
    <header class="bg-green-500 py-2 text-center font-semibold text-white">IT 08-1</header>

    <div class="p-4">
      <button
        class="rounded bg-green-500 px-4 py-1.5 text-sm text-white hover:bg-green-600"
        type="button"
        @click="emit('add')"
      >
        เพิ่มข้อสอบ
      </button>

      <p v-if="loading" class="mt-4 text-gray-500">กำลังโหลด...</p>
      <p v-else-if="error" class="mt-4 text-red-500">{{ error }}</p>
      <p v-else-if="exams.length === 0" class="mt-4 text-gray-500">ยังไม่มีข้อสอบ</p>

      <ol v-else class="mt-4 list-none">
        <li v-for="exam in exams" :key="exam.id" class="border-t border-gray-200 py-3">
          <div class="mb-1.5 flex items-center justify-between gap-2 font-semibold">
            <span>{{ exam.number }}. {{ exam.question }}</span>
            <button
              class="rounded bg-red-500 px-3 py-0.5 text-xs text-white hover:bg-red-600"
              type="button"
              @click="onDelete(exam.id)"
            >
              ลบ
            </button>
          </div>

          <div
            v-for="(answer, index) in exam.answers"
            :key="index"
            class="flex items-center gap-2 pl-4 text-gray-600"
          >
            <input type="radio" disabled :checked="index === 0" />
            <label>{{ answer }}</label>
          </div>
        </li>
      </ol>

      <div v-if="totalPages > 1" class="mt-4 flex items-center justify-center gap-3 text-gray-600">
        <button
          class="rounded border border-gray-300 px-3 py-1 text-xs disabled:cursor-not-allowed disabled:opacity-40"
          type="button"
          :disabled="page <= 1"
          @click="goToPage(page - 1)"
        >
          ก่อนหน้า
        </button>
        <span class="text-xs">หน้า {{ page }} / {{ totalPages }}</span>
        <button
          class="rounded border border-gray-300 px-3 py-1 text-xs disabled:cursor-not-allowed disabled:opacity-40"
          type="button"
          :disabled="page >= totalPages"
          @click="goToPage(page + 1)"
        >
          ถัดไป
        </button>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { deleteExam, fetchExams } from '../services/examApi'
import type { Exam } from '../types/exam'

const emit = defineEmits<{ add: [] }>()

const PAGE_SIZE = Number(import.meta.env.VITE_PAGE_SIZE ?? 10)

const exams = ref<Exam[]>([])
const page = ref(1)
const total = ref(0)
const totalPages = ref(1)
const loading = ref(false)
const error = ref('')

async function load() {
  loading.value = true
  error.value = ''
  try {
    const result = await fetchExams(page.value, PAGE_SIZE)
    exams.value = result.data
    total.value = result.total
    totalPages.value = result.totalPages
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'โหลดข้อมูลไม่สำเร็จ'
  } finally {
    loading.value = false
  }
}

function goToPage(target: number) {
  page.value = target
  load()
}

async function onDelete(id: number) {
  const deleted = exams.value.find((exam) => exam.id === id)
  if (!deleted) return

  try {
    await deleteExam(id)
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'ลบข้อมูลไม่สำเร็จ'
    return
  }


  exams.value = exams.value
    .filter((exam) => exam.id !== id)
    .map((exam) => (exam.number > deleted.number ? { ...exam, number: exam.number - 1 } : exam))
  total.value = Math.max(0, total.value - 1)
  totalPages.value = Math.max(1, Math.ceil(total.value / PAGE_SIZE))


  if (exams.value.length === 0 && page.value > 1) {
    page.value -= 1
    await load()
  }
}


function addLocally(exam: Exam) {
  total.value += 1
  totalPages.value = Math.max(1, Math.ceil(total.value / PAGE_SIZE))
  if (page.value === totalPages.value) {
    exams.value = [...exams.value, exam]
  }
}

defineExpose({ addLocally })

onMounted(load)
</script>
