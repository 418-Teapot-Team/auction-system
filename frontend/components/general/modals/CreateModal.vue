<template>
  <div
    class="fixed top-0 left-0 right-0 bottom-0 flex justify-center items-center bg-black/40 z-10"
  >
    <div
      class="rounded-xl bg-white shadow-sm z-20 p-4 flex flex-col justify-center items-center gap-y-4 min-w-[360px]"
    >
      <div class="flex justify-between items-center w-full">
        <span class="text-cl font-semibold">Create auction</span>
        <Icon name="ic:round-close" @click="onClose" class="cursor-pointer" />
      </div>
      <AtomsInputsAppPlainInput placeholder="Title" class="w-full" />
      <textarea placeholder="Description" class="w-full"></textarea>
      <AtomsInputsAppPlainInput
        placeholder="Tags (e.g. tag1,tag2,tag3 etc)"
        class="w-full"
      />
      <AtomsInputsAppPlainInput placeholder="Minimal bid" class="w-full" />
      <AtomsButtonsGreenRoundedButton text="Add image" class="h-8 w-32" />
      <span>max 4 images</span>
      <input type="file" hidden />
      <AtomsButtonsGreenRoundedButton text="Submit" class="h-8 w-32" />
    </div>
  </div>
</template>
<script setup>
const emit = defineEmits(['onClose']);

function onClose() {
  emit('onClose');
}
</script>
