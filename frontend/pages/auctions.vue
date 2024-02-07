<template>
  <div>
    <GeneralAppPageTitle title="Auctions" />
    <div class="pt-2 flex flex-col justify-start gap-y-6">
      <div class="flex justify-start gap-x-4 max-w-[65vw]">
        <AtomsInputsAppSelectInput
          :options="tags"
          v-model="tag"
          label="label"
          dataKey="id"
          placeholder="Tags"
        />
        <AtomsInputsAppPlainInput
          v-model="startBetMin"
          placeholder="Minimum bet"
          type="number"
          class="w-full"
        />
        <AtomsInputsAppPlainInput
          v-model="startBetMax"
          placeholder="Maximum bet"
          type="number"
          class="w-full"
        />
        <AtomsButtonsGreenRoundedButton
          @onClick="() => console.log('Apply')"
          text="Apply filters"
          class="w-full"
        />
      </div>
      <div class="flex justify-start items-center gap-x-3">
        <AtomsTag
          v-for="item of tags"
          :key="item.id"
          :label="item.label"
          :id="item.id"
        />
      </div>
      <div class="flex justify-between items-center gap-x-6">
        <AtomsInputsAppSearchInput
          placeholder="Search For Anything"
          class="w-9/12"
        />
        <div class="w-3/12 flex justify-end items-center gap-x-6">
          <span>Sort by:</span>
          <div class="w-64">
            <AtomsInputsAppSelectInput placeholder="Sort by" />
          </div>
        </div>
      </div>
    </div>
    <div>
      <div class="mt-4 pl-1.5">
        <span class="text-sm text-gray-800">9 results found</span>
      </div>
      <div class="mt-4">
        <GeneralAuctionCard />
      </div>
    </div>
  </div>
</template>
<script setup>
const tag = ref([]);
const startBetMin = ref('');
const startBetMax = ref('');

const sorting = ref(null);

const sortOpts = ref([
  {
    id: 1,
    label: 'bet',
  },
  {
    id: 1,
    label: 'users',
  },
  {
    id: 1,
    label: 'start date',
  },
]);

const tags = ref([
  {
    id: 1,
    label: 'Games',
  },
  {
    id: 2,
    label: 'Weapons',
  },
  {
    id: 3,
    label: 'Cars',
  },
  {
    id: 4,
    label: 'Electronics',
  },
]);
</script>
