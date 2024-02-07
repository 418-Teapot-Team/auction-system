<template>
  <input
    class="h-[28px] rounded-md bg-[#F1F1EE] p-2 text-darkest-grey border border-solid border-gray-300 shadow-md text-sm"
    v-model="model"
    :type="type"
    :placeholder="placeholder"
  />
</template>
<script setup>
defineProps(['type', 'placeholder']);
const model = defineModel();
</script>
