<script setup lang="js">
import { ref } from 'vue'
import { questions } from '@/assets/questions';
const clues = questions //Object array which stores all of the author questions
const question = ref("Ready?")
const questionClue = ref("The quiz will be long, just see how many you can get! Its randomized each time.")
const showInput = ref("false")
showInput.value = false
let questionNum = 0
let questionText = "N/A"
function startQuiz() {
    question.value = "Question " + questionNum + "."
    questionClue.value = questionText
    showInput.value=true
}
</script>
<template>
    <div class="quiz">
        <h2>{{ question }}</h2>
        <p>{{ questionClue }}</p>
        <div v-if="showInput">
            <input type="text"><br><br>
        </div>
        <button @click="startQuiz">Start</button>
        <p>Questions dump for debug: {{ clues }}</p>

    </div>
</template>
<style lang="css" scoped>
.hide {
    visibility: hidden !important;
}
</style>
