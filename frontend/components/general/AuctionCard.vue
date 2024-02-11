<template>
  <div
    class="bg-light-milky w-full p-4 rounded-lg flex justify-between flex-row-reverse gap-4 min-h-[320px] cursor-pointer"
    @click="redirectToAuction"
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
            <span class="text-sm font-semibold">{{ auction.title }}</span>
            <div
              class="flex justify-start gap-x-2 text-xs text-darkest-grey font-semibold"
            >
              <span class="bg-light-green p-1 rounded-md shadow-md"
                >Ongoing</span
              >
              <span class="bg-white p-1 rounded-md shadow-md">Public</span>
              <!-- <span class="bg-white p-1 rounded-md shadow-md">Social</span> -->
            </div>
          </div>
        </div>
        <div class="pt-4 text-xs text-gray-600">
          {{ auction.description }}
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
          <span class="font-light">{{ auction.currentBit }}</span>
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
            <span class="font-light text-sm">{{
              auction.creator.fullName
            }}</span>
          </div>
        </div>
        <div class="flex justify-start items-center gap-2">
          <span class="font-semibold text-sm">Created at: </span>
          <div>
            <Icon name="ic:baseline-calendar-month" color="#272D2D" />
            <span class="font-light text-sm">{{
              dayjs(auction.createdAt).format('DD-MM-YYYY')
            }}</span>
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
import dayjs from 'dayjs';
defineProps(['auction']);

const router = useRouter();

function redirectToAuction() {
  router.push('/auctions/1');
}
</script>
