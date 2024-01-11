<script lang="ts" setup>
const model = defineModel({required: true, default: 0})
</script>

<template>
  <input type="text" v-model.number="model" />
</template>

<style scoped>
input {
  display: block;
  width: 100%;
  outline: none;
  border: 2px solid var(--app-accent-color);
  background: #5d5d5dff;
  padding: .5rem;
  text-align: center;
  color: white;
}
</style>
