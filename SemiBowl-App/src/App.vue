<script setup>
import Topbar from "./components/NavBar.vue";
import { RouterView } from "vue-router";
</script>

<template>
  <header>
    <Topbar style="position: fixed" />
    <div class="wrapper">
      <nav></nav>
    </div>
  </header>
  <body>
    <RouterView />
  </body>
</template>

<style scoped></style>
