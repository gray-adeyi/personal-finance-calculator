<script lang="ts" setup></script>

<template>
  <button>Calculate</button>
</template>

<style scoped>
button{
  display: block;
  width: 90%;
  text-align: center;
  padding: .5rem;
  color: white;
  background: var(--app-accent-color);
  border: none;
  border-radius: 7px;
  font-weight: bold;
  font-size: 1rem;
  cursor: pointer;
  transition: background 500ms;

  &:hover{
    background: #598bf6;
  }

  &:active{
    background: var(--app-accent-color);
  }
}
</style>