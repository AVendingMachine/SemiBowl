<script setup>
import { RouterLink} from 'vue-router'
</script>
<template>
    <div class="navbar">
        <RouterLink to="/" class="title">Semibowl</RouterLink>
        <RouterLink to="/Authors">Authors</RouterLink>
    </div>
</template>
<style lang="css" scoped>
.title {
    padding: 20px;
}
</style>