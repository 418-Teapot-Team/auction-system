<template>
  <div>
    <GeneralAppPageTitle title="Auction 1" />
    <div
      class="flex flex-col justify-start gap-y-4 lg:flex-row lg:justify-between mt-10 h-full lg:gap-x-4"
    >
      <div class="flex flex-col justify-start gap-4 lg:w-1/3 w-full">
        <div class="w-full">
          <img src="/assets/images/lot.webp" class="rounded-lg" />
        </div>
        <div class="flex justify-start gap-4">
          <img
            width="70"
            height="70"
            src="/assets/images/lot.webp"
            class="rounded-lg"
          />
          <img
            width="70"
            height="70"
            src="/assets/images/lot.webp"
            class="rounded-lg"
          />
          <img
            width="70"
            height="70"
            src="/assets/images/lot.webp"
            class="rounded-lg"
          />
        </div>
      </div>
      <div class="lg:w-2/3 w-full">
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
          iure nobis? Lorem ipsum dolor sit amet consectetur adipisicing elit.
          Enim, itaque laboriosam. Ipsa, deserunt molestiae? Eius, est corrupti
          excepturi fuga, ipsa tempore nobis nihil nisi, dignissimos cumque aut.
          Deleniti, iure nobis?
        </div>
        <div class="flex justify-between w-full mt-10">
          <span class="text-lg font-semibold text-darkest-grey"
            >Current price:
          </span>
          <span class="text-lg font-semibold text-darkest-grey"
            >US $199.90</span
          >
        </div>
        <div class="flex justify-start items-center mt-5 gap-x-4 mb-5">
          <span>(USD) $</span>
          <AtomsInputsAppPlainInput type="number" />
          <AtomsButtonsGreenRoundedButton
            @on-click=""
            text="Bid now"
            class="h-6 w-20"
          />
        </div>
        <span class="text-lg font-semibold text-dark-grey">
          Realtime bids:
        </span>
        <div
          class="flex flex-col justify-start gap-y-2 p-4 max-h-[40vh] bg-light-milky rounded-lg overflow-y-auto"
        >
          <GeneralRealtimeBid />
          <GeneralRealtimeBid />
          <GeneralRealtimeBid />
          <GeneralRealtimeBid />
          <GeneralRealtimeBid />
          <GeneralRealtimeBid />
        </div>
      </div>
    </div>
  </div>
</template>
<script setup></script>
