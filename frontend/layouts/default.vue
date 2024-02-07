<template>
  <main
    class="bg-milky min-h-screen min-w-screen p-1 flex justify-start font-poppins"
  >
    <PartialsAppAside />
    <div class="ml-[170px] w-full px-8">
      <div class="pt-4 flex justify-between mb-4">
        <PartialsAppBreadcrumbs />
        <AtomsInputsAppSearchInput placeholder="Search For Anything" />
      </div>
      <slot />
    </div>
  </main>
</template>
<script setup></script>
