<script setup lang="js">
import { nextTick, ref } from "vue";
import { questions } from "@/assets/questions.js";
import { LevDistance } from "@/assets/levDistance.js";

const clues = questions; //Object array which stores all the autfhor questions
const showQuiz = ref(false);
const showQuizContent = ref(true);
const showAnswer = ref(false);
const regularQuestionOrder = ref([0]);
const randomQuestionOrder = ref([0]);
const currentInput = ref("");
const winMessage = ref("");
const difficulty = ref(0);
const maxDifficulty = 3;
const isCorrect = ref(false);
const inputBox = ref(null);
const continueButton = ref(null);
const questionDisplay = ref("PLACEHOLDER");
const allowSummary = ref(false);
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
function randomizeQuestionOrder() {
  for (let i = 0; i < clues.length; i++) {
    questionOrder[i] = i;
  }
  regularQuestionOrder.value = questionOrder;
  randomQuestionOrder.value = randomizeArray(questionOrder);
}
function updateQuestionDisplay() {
  switch (difficulty.value) {
    case 0:
      questionDisplay.value = clues[randomQuestionOrder.value[0]].EasyWork;
      break;
    case 1:
      questionDisplay.value = clues[randomQuestionOrder.value[0]].EasyClue;
      break;
    case 2:
      questionDisplay.value = clues[randomQuestionOrder.value[0]].HardWork;
      break;
    case 3:
      questionDisplay.value = clues[randomQuestionOrder.value[0]].HardClue;
      break;
  }
}
function startQuiz() {
  showQuiz.value = true;
  randomizeQuestionOrder();
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
  updateQuestionDisplay();
}
function checkIfCorrect(answer) {
  let sanitizedAnswer = answer.toLowerCase();
  sanitizedAnswer = sanitizedAnswer.replace(/\s/g, "");
  for (let i = 0; i < clues[randomQuestionOrder.value[0]].Answers.length; i++) {
    let currentAnswer = clues[randomQuestionOrder.value[0]].Answers[i];
    if (currentAnswer === sanitizedAnswer) {
      return true;
    } else if (LevDistance(currentAnswer, sanitizedAnswer) <= 3) {
      return true;
    }
  }
  return false;
}
function checkAnswer(answer) {
  currentInput.value = "";
  isCorrect.value = checkIfCorrect(answer);
  showAnswer.value = true;
}
function resetQuiz() {
  difficulty.value = 0;
  showQuizContent.value = true;
  allowSummary.value = false;
  startQuiz();
}
function showSummary() {
  showQuizContent.value = false;
  allowSummary.value = true;
}
async function advanceQuestion() {
  if (isCorrect.value) {
    if (randomQuestionOrder.value.length > 1) {
      regularQuestionOrder.value.shift();
      randomQuestionOrder.value.shift();
    } else if (randomQuestionOrder.value.length <= 1) {
      if (difficulty.value === maxDifficulty) {
        showSummary();
      } else {
        difficulty.value++;
        randomizeQuestionOrder();
      }
    }
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
  updateQuestionDisplay();
}
</script>
<template>
  <button v-if="!showQuiz" @click="startQuiz">Start</button>
  <div v-if="showQuiz" class="quiz">
    <div v-if="showQuizContent">
      <h2>
        Question
        {{ regularQuestionOrder[0] + 1 + difficulty * clues.length }}.
      </h2>
      <p>{{ questionDisplay }}</p>
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
            {{ clues[randomQuestionOrder[0]].Author }}
          </p>
          <button @click="advanceQuestion()" ref="continueButton">
            Continue
          </button>
        </div>
      </div>
    </div>
    <div v-if="allowSummary">
      <h2>Summary</h2>
      <button @click="resetQuiz()">Restart</button>
    </div>
  </div>
</template>
