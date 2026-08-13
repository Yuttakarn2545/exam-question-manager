<template>
  <section class="w-full max-w-md overflow-hidden rounded bg-white shadow">
    <header class="bg-green-500 py-2 text-center font-semibold text-white">IT 08-2</header>

    <div class="p-4">
      <form class="flex flex-col gap-2.5" @submit.prevent="onSave">
        <label class="flex items-center gap-2">
          <span class="w-[70px] shrink-0 text-sm">คำถาม</span>
          <input v-model="form.question" type="text" class="flex-1 rounded border border-gray-300 px-2 py-1.5" />
        </label>
        <label class="flex items-center gap-2">
          <span class="w-[70px] shrink-0 text-sm">คำตอบ 1</span>
          <input v-model="form.answer1" type="text" class="flex-1 rounded border border-gray-300 px-2 py-1.5" />
        </label>
        <label class="flex items-center gap-2">
          <span class="w-[70px] shrink-0 text-sm">คำตอบ 2</span>
          <input v-model="form.answer2" type="text" class="flex-1 rounded border border-gray-300 px-2 py-1.5" />
        </label>
        <label class="flex items-center gap-2">
          <span class="w-[70px] shrink-0 text-sm">คำตอบ 3</span>
          <input v-model="form.answer3" type="text" class="flex-1 rounded border border-gray-300 px-2 py-1.5" />
        </label>
        <label class="flex items-center gap-2">
          <span class="w-[70px] shrink-0 text-sm">คำตอบ 4</span>
          <input v-model="form.answer4" type="text" class="flex-1 rounded border border-gray-300 px-2 py-1.5" />
        </label>

        <p v-if="error" class="text-xs text-red-500">{{ error }}</p>

        <div class="mt-2 flex justify-center gap-2.5">
          <button
            class="rounded bg-blue-600 px-5 py-1.5 text-sm text-white hover:bg-blue-700 disabled:opacity-50"
            type="submit"
            :disabled="saving"
          >
            บันทึก
          </button>
          <button
            class="rounded bg-red-500 px-5 py-1.5 text-sm text-white hover:bg-red-600 disabled:opacity-50"
            type="button"
            :disabled="saving"
            @click="onCancel"
          >
            ยกเลิก
          </button>
        </div>
      </form>
    </div>
  </section>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { createExam } from '../services/examApi'
import type { Exam } from '../types/exam'

const emit = defineEmits<{ saved: [exam: Exam]; cancel: [] }>()

const form = reactive({
  question: '',
  answer1: '',
  answer2: '',
  answer3: '',
  answer4: '',
})

const saving = ref(false)
const error = ref('')

async function onSave() {
  if (!form.question || !form.answer1 || !form.answer2 || !form.answer3 || !form.answer4) {
    error.value = 'กรุณากรอกข้อมูลให้ครบทุกช่อง'
    return
  }

  saving.value = true
  error.value = ''
  try {
    const exam = await createExam({
      question: form.question,
      answers: [form.answer1, form.answer2, form.answer3, form.answer4],
    })
    emit('saved', exam)
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'บันทึกข้อมูลไม่สำเร็จ'
  } finally {
    saving.value = false
  }
}

function onCancel() {
  emit('cancel')
}
</script>
