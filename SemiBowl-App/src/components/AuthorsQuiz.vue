<script setup lang="js">
import { nextTick, ref } from "vue";
import { questions } from "@/assets/questions";
const clues = questions; //Object array which stores all the author questions
const showQuiz = ref(false);
const showAnswer = ref(false);
const regularQuestionOrder = ref([0]);
const randomQuestionOrder = ref([0]);
const currentInput = ref("");
const winMessage = ref("");
const difficulty = ref(0);
const isCorrect = ref(false);
const inputBox = ref(null);
const continueButton = ref(null);
let questionOrder = [0];
showQuiz.value = false;
randomQuestionOrder.value = [0];

function randomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}
function randomizeArray(orderedArray) {
  let randomizedArray = [0];
  let copiedArray = orderedArray.slice();
  let n = 0;
  while (copiedArray.length > 0) {
    let currentValue = randomInt(0, copiedArray.length - 1);
    randomizedArray[n] = copiedArray[currentValue];
    copiedArray.splice(currentValue, 1);
    n++;
  }
  return randomizedArray;
}
function startQuiz() {
  showQuiz.value = true;
  for (let i = 0; i < clues.length; i++) {
    questionOrder[i] = i;
  }
  regularQuestionOrder.value = questionOrder;
  randomQuestionOrder.value = randomizeArray(questionOrder);
  nextTick(function () {
    inputBox.value.focus();
  });
  window.addEventListener("keyup", (event) => {
    if (event.key === "Enter") {
      if (showAnswer.value) {
        advanceQuestion();
      } else if (showAnswer.value === false) {
        checkAnswer(currentInput.value);
      }
    }
  });
}
function checkIfCorrect(answer) {
  return (
    answer.toLowerCase() ===
    clues[randomQuestionOrder.value[0]].author.toLowerCase()
  );
}
function checkAnswer(answer) {
  currentInput.value = "";
  isCorrect.value = checkIfCorrect(answer);
  showAnswer.value = true;
}
function advanceQuestion() {
  if (isCorrect.value) {
    regularQuestionOrder.value.shift();
    randomQuestionOrder.value.shift();
  } else if (randomQuestionOrder.value.length <= 0) {
    difficulty.value++;
    randomQuestionOrder.value = randomizeArray(questionOrder);
  } else if (!isCorrect.value) {
    regularQuestionOrder.value.push(regularQuestionOrder.value[0]);
    regularQuestionOrder.value.shift();
    randomQuestionOrder.value.push(randomQuestionOrder.value[0]);
    randomQuestionOrder.value.shift();
  }
  showAnswer.value = false;
  nextTick(function () {
    inputBox.value.focus();
  });
}
</script>
<template>
  <button v-if="!showQuiz" @click="startQuiz">Start</button>
  <div v-if="showQuiz" class="quiz">
    <h2>
      Question
      {{
        regularQuestionOrder[0] + 1 + difficulty * regularQuestionOrder.length
      }}.
    </h2>
    <p>{{ clues[randomQuestionOrder[0]].easyWork }}</p>
    <p>{{ winMessage }}</p>
    <div v-if="showQuiz">
      <input
        type="text"
        v-model.trim="currentInput"
        :disabled="showAnswer"
        ref="inputBox"
      /><br />
      <div v-if="!showAnswer">
        <br />
        <button @click="checkAnswer(currentInput)">Submit</button>
      </div>
      <div v-if="showAnswer">
        <p v-if="isCorrect">Correct!</p>
        <p v-if="!isCorrect">
          Incorrect, the correct answer was
          {{ clues[randomQuestionOrder[0]].author }}
        </p>
        <button @click="advanceQuestion()" ref="continueButton">
          Continue
        </button>
      </div>
    </div>
  </div>
</template>
