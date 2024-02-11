<template>
  <div class="flex items-center relative">
    <Icon name="ic:outline-search" class="absolute left-3 h-[16px] w-[16px]" />
    <input
      v-model="model"
      :placeholder="placeholder"
      class="rounded-full shadow-md text-sm h-[28px] pl-9 align-middle w-full bg-white"
    />
  </div>
</template>
<script setup>
defineProps(['placeholder']);
const model = defineModel();
</script>
