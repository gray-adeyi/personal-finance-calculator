<script lang="ts" setup>
import {main} from "../../wailsjs/go/models";
import {computed} from "vue";

const $props = defineProps<{
  result: main.FinanceResult | null;
}>();

const isCardVisible = computed(() => $props.result !== null)
</script>

<template>
  <div :class="['card', {'card--hide': !isCardVisible}]">
    <span class="card__text">Taxes: {{ $props.result?.taxes.toFixed(2)}}</span>
    <span class="card__text">Owner: {{ $props.result?.owner.toFixed(2)}}</span>
    <span class="card__text">Profit: {{ $props.result?.profit.toFixed(2)}}</span>
    <span class="card__text">OpEx: {{ $props.result?.opEx.toFixed(2)}}</span>
  </div>
</template>

<style scoped>
.card{
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 80%;
  color: black;
  background: white;
  padding: 1rem 2rem;
  box-shadow: 0 0 5px black;
  transition: opacity 500ms;

  .card__text{
    font-size: 1rem;
    font-weight: bold;
  }
}
.card--hide{
  opacity: 0;
}
</style>