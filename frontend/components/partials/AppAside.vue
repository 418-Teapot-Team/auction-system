<template>
  <nav
    class="w-[170px] bg-dark-grey rounded-xl top-1 bottom-1 left-1 fixed p-4"
  >
    <div class="flex justify-start gap-x-2 items-center">
      <div class="w-10 h-10">
        <IconLogoIcon />
      </div>
      <h1 class="text-light-milky text-xl">Auctions</h1>
    </div>
    <ul>
      <GeneralMenuMenuItem
        v-for="item of menuItems"
        :key="item.link"
        :title="item.title"
      >
        <Icon :name="item.icon" />
      </GeneralMenuMenuItem>
    </ul>
  </nav>
</template>
<script setup>
const menuItems = ref([
  {
    title: 'Dashboard',
    link: '/',
    icon: '',
  },
]);
</script>
