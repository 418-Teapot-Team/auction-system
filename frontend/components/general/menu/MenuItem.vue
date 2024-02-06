<template>
  <li>
    <slot name="icon" />
    <span>{{ title }}</span>
  </li>
</template>
<script setup>
defineProps(['title', 'link']);
</script>
