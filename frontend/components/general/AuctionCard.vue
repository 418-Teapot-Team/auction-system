<template>
  <div
    class="bg-light-milky w-full p-4 rounded-lg flex justify-between flex-row-reverse gap-4 min-h-[320px] cursor-pointer"
  >
    <div class="w-1/2 flex flex-col justify-between items-start">
      <div>
        <div class="flex justify-start gap-x-2">
          <div
            class="w-14 h-14 rounded-full bg-dark-grey flex justify-center items-center"
          >
            <Icon name="mdi:scale-unbalanced" class="w-8 h-8" color="#77FF61" />
          </div>
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
        <div class="pt-4 text-xs text-gray-600">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Enim, itaque
          laboriosam. Ipsa, deserunt molestiae? Eius, est corrupti excepturi
          fuga, ipsa tempore nobis nihil nisi, dignissimos cumque aut. Deleniti,
          iure nobis?
        </div>
      </div>
      <div class="flex justify-start items-center 5">
        <span class="font-semibold">Current bid:</span>
        <div class="flex items-center justify-center">
          <Icon
            name="ic:baseline-attach-money"
            class="w-5 h-5"
            color="#272D2D"
          />
          <span class="font-light">100</span>
        </div>
      </div>
      <div>
        <div class="flex justify-start items-center gap-2">
          <span class="font-semibold text-sm">Author: </span>
          <div>
            <Icon
              name="ic:baseline-account-circle"
              color="#272D2D"
              class="mr-0.5"
            />
            <span class="font-light text-sm">mekola</span>
          </div>
        </div>
        <div class="flex justify-start items-center gap-2">
          <span class="font-semibold text-sm">Created at: </span>
          <div>
            <Icon name="ic:baseline-calendar-month" color="#272D2D" />
            <span class="font-light text-sm">11-05-2024</span>
          </div>
        </div>
      </div>
    </div>

    <div class="bg-white flex justify-start gap-4 p-4 rounded-lg w-1/2">
      <img
        src="/assets/images/lot.webp"
        alt="lot image"
        class="rounded-lg w-full object-contain"
      />
    </div>
  </div>
</template>
<script setup>
defineProps(['auction']);
</script>
