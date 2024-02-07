<template>
  <button
    class="bg-light-green rounded-full text-darkest-grey shadow-md text-sm"
    @click="click"
    :type="type ? type : 'button'"
  >
    {{ text }}
  </button>
</template>
<script setup>
defineProps(['text', 'type']);
const emit = defineEmits(['onClick']);

function click() {
  emit('onClick');
}
</script>
