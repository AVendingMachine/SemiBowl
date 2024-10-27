<script setup>
import { RouterLink } from "vue-router";
</script>
<template>
  <div class="navbar">
    <RouterLink to="/" class="title">Semibowl</RouterLink>
    <RouterLink to="/Authors">Authors</RouterLink>
    <a
      href="https://github.com/AVendingMachine/SemiBowl"
      target="_blank"
      rel="noopener noreferrer"
    >
      <img
        class="codeIcon"
        src="../assets/codeicon.svg"
        alt="Link to Github source"
    /></a>
  </div>
</template>
<style lang="css" scoped>
.title {
  padding: 20px;
}
</style>
