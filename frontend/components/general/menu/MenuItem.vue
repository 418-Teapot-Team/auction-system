<template>
  <li
    class="flex justify-start items-center gap-x-2 rounded-full py-1 pr-2 pl-4 w-full cursor-pointer"
    :class="{
      'bg-menu-bg': isActive,
    }"
  >
    <slot name="icon" />
    <NuxtLink :to="link" class="text-light-milky font-light">{{
      title
    }}</NuxtLink>
  </li>
</template>
<script setup>
defineProps(['title', 'link', 'isActive']);
</script>
