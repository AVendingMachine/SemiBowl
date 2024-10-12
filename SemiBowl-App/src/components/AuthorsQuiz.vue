<script setup lang="js">
import { ref } from 'vue'
import { questions } from '@/assets/questions';
const clues = questions //Object array which stores all of the author questions
const showQuiz = ref("false")
let questionOrder = [0]
let randomQuestionOrder = [0]
let currentQuestionID = 0
let currentDifficulty = 0
const rightWrong = ref("bunga")
const currentInput = ref("")
showQuiz.value = false

function randomInt(min,max) {
    min = Math.ceil(min)
    max = Math.floor(max)
    return Math.floor(Math.random()*(max-min +1)) + min
}
function randomizeArray(orderedArray) {
    let randomizedArray = [0]
    let n = 0
    while (orderedArray.length > 0) {
        let currentValue = randomInt(0,orderedArray.length-1)
        randomizedArray[n] = orderedArray[currentValue]
        orderedArray.splice(currentValue,1)
        n++
    }
    return randomizedArray
}
async function startQuiz() {
    showQuiz.value=true
    for (let i = 0;i<clues.length;i++) {
        questionOrder[i] = i;
    }
    randomQuestionOrder = randomizeArray(questionOrder)
    currentQuestionID = randomQuestionOrder[0]
}
function advanceQuestion(answer) {
    if (answer.toLowerCase() == clues[currentQuestionID].author.toLowerCase()) {
        rightWrong.value = "right"
    }
    else {
        rightWrong.value = "wrong"
    }
}
</script>
<template>
    <button v-if="!showQuiz" @click="startQuiz">Start</button>
    <div v-if="showQuiz" class="quiz">
        <h2>Question {{ clues[currentQuestionID].questionID+1 }}.</h2>
        <p>{{ clues[currentQuestionID].easyWork }}</p>
        <div v-if="showQuiz">
            <input type="text" v-model.trim="currentInput"><br>
            <input type="submit" @click="advanceQuestion(currentInput)">
        </div>
        <h4>DEBUG INFO</h4>
        <p>- currentInput: {{ currentInput }}</p>
        <p>- clues[currentQuestionId].author: {{ clues[currentQuestionID].author }}</p>
        <p>- randomQuestionOrder: {{ randomQuestionOrder }}</p>
        <p>- rightWrong: {{ rightWrong }}</p>

        <!--<p>Questions dump for debug: {{ clues }}</p>-->

    </div>
</template>
<style lang="css" scoped>
.hide {
    visibility: hidden !important;
}
</style>
