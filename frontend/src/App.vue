<script lang="ts" setup>
import AppInput from "./components/AppInput.vue";
import AppButton from "./components/AppButton.vue";
import ResultCard from "./components/ResultCard.vue";
import {ref} from 'vue';
import {calculatePersonalFinance} from "./backendService";
import {main} from "../wailsjs/go/models";
import {ShowErrorDialog} from "../wailsjs/go/main/App";

const _income = ref(0)
const _result = ref<main.FinanceResult|null>(null)

async function calculate(){
  try {
    _result.value = await calculatePersonalFinance(_income.value)
  } catch (e) {
    await ShowErrorDialog((e as Error).toString())
  }
}
</script>

<template>
  <div class="app">
    <h1 class="app__header">Enter Total Income</h1>
    <AppInput class="app__input" v-model.number="_income" @keyup.enter="calculate" />
    <AppButton class="app__button" @click="calculate" />
    <ResultCard class="app__result-card" :result="_result" />
  </div>
</template>

<style scoped>
.app{
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 2rem;

  .app__header{
    margin-block-start: 3rem;
    margin-block-end: 2rem;
  }

  .app__input{
    margin-block-end: 2rem;
  }

  .app__button{
    margin-block-end: 3rem;
  }
}
</style>
