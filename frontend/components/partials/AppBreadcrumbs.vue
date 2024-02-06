<template>
  <nav id="breadcrumbs" aria-label="Breadcrumb">
    <ul class="flex flex-row gap-x-1 text-sm font-normal">
      <li>
        <NuxtLink
          to="/"
          :aria-current="ariaCurrent(-1)"
          :class="{
            'text-gray-700': getBreadcrumbs.length > 0,
          }"
          >Dashboard</NuxtLink
        >
      </li>
      <li
        v-for="(breadcrumb, index) in getBreadcrumbs"
        :key="index"
        class="before:pr-1 before:content-['/']"
      >
        <NuxtLink
          :to="breadcrumb.path"
          :aria-current="ariaCurrent(index)"
          class="capitalize"
          :class="{
            'text-gray-700': !(breadcrumb.path === route.path),
          }"
        >
          {{ breadcrumb.name }}
        </NuxtLink>
      </li>
    </ul>
  </nav>
</template>

<script setup>
const route = useRoute();
const router = useRouter();

const getBreadcrumbs = computed(() => {
  const fullPath = route.fullPath;
  const requestPath = fullPath;

  const crumbs = requestPath.split('/');
  const breadcrumbs = [];
  let path = '';

  crumbs.forEach((crumb) => {
    if (crumb) {
      path = `${path}/${crumb}`;
      const breadcrumb = router.getRoutes().find((r) => r.path === path);

      if (breadcrumb) {
        breadcrumbs.push(breadcrumb);
      }
    }
  });

  return breadcrumbs;
});
const ariaCurrent = (index) =>
  index === getBreadcrumbs.value.length - 1 ? 'page' : 'false';
</script>
