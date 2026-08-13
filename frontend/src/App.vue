<template>
  <div class="min-h-screen flex justify-center bg-gray-100 p-8 font-sans text-sm text-gray-800">
    <ExamList v-show="view === 'list'" ref="examListRef" @add="view = 'form'" />
    <ExamForm v-if="view === 'form'" @saved="onSaved" @cancel="view = 'list'" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ExamList from './components/ExamList.vue'
import ExamForm from './components/ExamForm.vue'
import type { Exam } from './types/exam'

const view = ref<'list' | 'form'>('list')
const examListRef = ref<InstanceType<typeof ExamList> | null>(null)

function onSaved(exam: Exam) {
  examListRef.value?.addLocally(exam)
  view.value = 'list'
}
</script>
