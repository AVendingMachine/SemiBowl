<script setup>
import AuthorsQuiz from "../components/AuthorsQuiz.vue";
import { questions } from "@/assets/questions";
</script>
<template>
  <body>
    <div class="content">
      <h1>
        {{ questions.length }} Authors You Should Probably Know for Quizbowl
      </h1>
      <p>
        This quiz consists of {{ questions.length * 4 }} questions with answer
        choices ranging between {{ questions.length }} authors. You will have to
        guess the authors name based on 4 prompts:
      </p>
      <ol>
        <li>The title of a well-known work they wrote</li>
        <li>A slightly obscure fact about that well-known work</li>
        <li>A less-known work</li>
        <li>A fact about that obscure work</li>
      </ol>
      <p>
        The question ordered is randomized, but each section is always in the
        same order (i.e. the answer to question 1 won't always be Plath but it
        will always prompt about the title of a well-known work)
      </p>
      <AuthorsQuiz />
    </div>
  </body>
</template>
