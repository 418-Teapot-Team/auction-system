<template>
  <ClientOnly>
    <v-select
      v-model="model"
      :options="options"
      :label="label"
      :reduce="(option) => option[`${dataKey}`]"
      :multiple="isMultiple"
      class="bg-white rounded-lg select shadow-md w-full"
      :placeholder="placeholder"
    />
  </ClientOnly>
</template>
<script setup>
defineProps(['options', 'dataKey', 'label', 'isMultiple', 'placeholder']);
const model = defineModel();
</script>
<style>
.vs__dropdown-toggle {
  height: 28px;
}
.vs__search {
  height: 20px;
  margin-bottom: 2px;
}
.vs__selected {
  height: 20px;
  margin-bottom: 2px;
}
</style>
<style scoped>
.select {
  --vs-controls-color: #222727;
  --vs-border-color: rgb(209 213 219);
  --vs-font-size: 14px;
  --vs-line-height: ;
  --vs-dropdown-bg: #eeeeee;
  --vs-dropdown-color: #eeeeee;
  --vs-dropdown-option-color: #222727;

  --vs-selected-bg: #dfdfd7;
  --vs-selected-color: #222727;

  --vs-search-input-color: #222727;
  --vs-actions-padding: 4px 6px 0 3px;

  --vs-dropdown-option--active-bg: #dfdfd7;
  --vs-dropdown-option--active-color: #222727;
}
</style>
