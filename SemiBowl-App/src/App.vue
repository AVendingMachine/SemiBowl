<script setup>
import Topbar from './components/NavBar.vue';
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    <Topbar/>
    <div class="wrapper">
      <nav>
      </nav>
    </div>
  </header>
  <RouterView />
</template>

<style scoped>
</style>
