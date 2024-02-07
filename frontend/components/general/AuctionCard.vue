<template>
  <div class="bg-light-milky w-1/2 p-4 rounded-lg">
    <div class="flex justify-between">
      <div class="flex justify-start gap-x-2">
        <img
          class="w-14 h-14 rounded-full"
          src="/public/images/fotor-03d1a91a0cec4542927f53c87e0599f6.webp"
        />
        <div class="flex flex-col justify-center gap-y-1">
          <span class="text-sm font-semibold">summitlg</span>
          <div
            class="flex justify-start gap-x-2 text-xs text-darkest-grey font-semibold"
          >
            <span class="bg-light-green p-1 rounded-md shadow-md">Games</span>
            <span class="bg-white p-1 rounded-md shadow-md">Science</span>
            <span class="bg-white p-1 rounded-md shadow-md">Social</span>
          </div>
        </div>
      </div>
      <div>
        <div>save</div>
      </div>
    </div>
    <div class="pt-4 text-xs text-gray-600">
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Enim, itaque
      laboriosam. Ipsa, deserunt molestiae? Eius, est corrupti excepturi fuga,
      ipsa tempore nobis nihil nisi, dignissimos cumque aut. Deleniti, iure
      nobis?
    </div>
    <div>Card inner</div>
    <div>Card footer</div>
  </div>
</template>
<script setup>
defineProps(['auction']);
</script>
