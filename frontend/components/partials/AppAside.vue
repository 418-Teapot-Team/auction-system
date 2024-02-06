<template>
  <nav
    class="w-[170px] bg-dark-grey rounded-xl top-1 bottom-1 left-1 fixed flex flex-col justify-between"
  >
    <div class="px-2 pt-4">
      <div class="flex justify-start gap-x-2 items-center">
        <div class="w-10 h-10">
          <IconLogoIcon />
        </div>
        <h1 class="text-light-milky text-xl">Auctions</h1>
      </div>
      <ul class="pt-4 flex flex-col justify-start items-start gap-y-2 w-full">
        <GeneralMenuItem
          v-for="item of menuItems"
          :key="item.link"
          :title="item.title"
          :is-active="router.currentRoute.value.path === item.link"
          :link="item.link"
        >
          <template #icon>
            <Icon :name="item.icon" color="#EAEAE4" />
          </template>
        </GeneralMenuItem>
      </ul>
    </div>
    <div
      class="bg-darkest-grey h-28 w-full rounded-b-xl border-t border-t-gray-700 flex justify-center items-center"
    >
      <div class="flex justify-center items-center gap-x-2 cursor-pointer">
        <Icon
          class="rotate-180 w-7 h-7"
          name="ic:outline-log-out"
          color="#EAEAE4"
        />
        <div class="flex flex-col">
          <span class="text-light-milky">Sign out</span>
          <span class="text-xs font-thin text-gray-500">My account</span>
        </div>
      </div>
    </div>
  </nav>
</template>
<script setup>
const router = useRouter();
const menuItems = ref([
  {
    title: 'Dashboard',
    link: '/',
    icon: 'ic:baseline-dashboard',
  },
  {
    title: 'Auctions',
    link: '/auctions',
    icon: 'ic:baseline-gavel',
  },
  {
    title: 'Drafts',
    link: '/drafts',
    icon: 'ic:sharp-edit-note',
  },
  {
    title: 'Settings',
    link: '/settings',
    icon: 'ic:round-settings',
  },
]);
</script>
