<script setup>
import AuthorsQuiz from "../components/AuthorsQuiz.vue"
import { questions } from "@/assets/questions";
</script>
<template>
  <div class="content">
    <h1>{{ questions.length }} Authors You Should Probably Know for Quizbowl</h1>
    <p>This quiz contains something something</p>
    <p>and more</p>
    <AuthorsQuiz/>
  </div>
</template>

