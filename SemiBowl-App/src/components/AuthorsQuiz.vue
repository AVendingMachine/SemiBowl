<script setup lang="js">
import { ref } from 'vue'
import { questions } from '@/assets/questions';
const clues = questions //Object array which stores all of the author questions
const showQuiz = ref(false)
const regularQuestionOrder = ref([0])
const randomQuestionOrder = ref([0])
const rightWrong = ref("bunga")
const currentInput = ref("")
const winMessage = ref("")
const difficulty = ref(0)
let questionOrder = [0]
showQuiz.value = false
randomQuestionOrder.value = [0]

function randomInt(min,max) {
    min = Math.ceil(min)
    max = Math.floor(max)
    return Math.floor(Math.random()*(max-min +1)) + min
}
function randomizeArray(orderedArray) {
    let randomizedArray = [0]
    let copiedArray = orderedArray.slice()
    let n = 0
    while (copiedArray.length > 0) {
        let currentValue = randomInt(0,copiedArray.length-1)
        randomizedArray[n] = copiedArray[currentValue]
        copiedArray.splice(currentValue,1)
        n++
    }
    return randomizedArray
}
async function startQuiz() {
    showQuiz.value=true
    for (let i = 0;i<clues.length;i++) {
        questionOrder[i] = i;
    }
    regularQuestionOrder.value = questionOrder
    randomQuestionOrder.value = randomizeArray(questionOrder)
}
function checkIfCorrect(answer) {
  currentInput.value = ""
  if (answer.toLowerCase() === clues[randomQuestionOrder.value[0]].author.toLowerCase()) {
    regularQuestionOrder.value.shift()
    randomQuestionOrder.value.shift()
    return true
  }
  else if (randomQuestionOrder.value.length === 0) {
    difficulty.value++
  }
  else {
    regularQuestionOrder.value.push(regularQuestionOrder.value[0])
    regularQuestionOrder.value.shift()
    randomQuestionOrder.value.push(randomQuestionOrder.value[0])
    randomQuestionOrder.value.shift()
    return false
  }
}
function advanceQuestion(answer) {
    currentInput.value = ""
    if (answer.toLowerCase() === clues[randomQuestionOrder.value[0]].author.toLowerCase()) {
        regularQuestionOrder.value.shift()
        randomQuestionOrder.value.shift()
        return true
    }
    else if (randomQuestionOrder.value.length === 0) {
        difficulty.value++
    }
    else {
        regularQuestionOrder.value.push(regularQuestionOrder.value[0])
        regularQuestionOrder.value.shift()
        randomQuestionOrder.value.push(randomQuestionOrder.value[0])
        randomQuestionOrder.value.shift()
        return false
    }
}

</script>
<template>
    <button v-if="!showQuiz" @click="startQuiz">Start</button>
    <div v-if="showQuiz" class="quiz">
        <h2>Question {{ regularQuestionOrder[0]+1 }}.</h2>
        <p>{{ clues[randomQuestionOrder[0]].easyWork }}</p>
        <p>{{ winMessage }}</p>
        <div v-if="showQuiz">
            <form  @submit.prevent="checkIfCorrect(currentInput)">
                <input type="text" v-model.trim="currentInput"><br><br>
                <button>Submit</button>
            </form>
        </div>
    </div>
</template>
<style lang="css" scoped>
.hide {
    visibility: hidden !important;
}
</style>
