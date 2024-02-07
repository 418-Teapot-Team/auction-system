<template>
  <div
    class="bg-light-milky h-8 flex justify-center items-center gap-x-2 p-2 rounded-full text-sm"
  >
    <span>{{ label }}</span>
    <Icon
      class="cursor-pointer"
      name="ic:round-close"
      color="#60605D"
      @click="onRemove"
    />
  </div>
</template>
<script setup>
const props = defineProps(['label', 'id']);

const emit = defineEmits(['onRemove']);

function onRemove() {
  emit('onRemove', props.id);
}
</script>
