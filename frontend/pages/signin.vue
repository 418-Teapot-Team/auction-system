<template>
  <div
    class="min-w-screen min-h-screen bg-milky flex justify-center items-center"
  >
    <form
      class="bg-light-milky shadow-md rounded-lg flex flex-col gap-4 p-4"
      @submit.prevent="onSubmit"
    >
      <h1 class="text-lg font-semibold text-darkest-grey text-center">
        Sign In
      </h1>
      <div class="flex justify-between items-center gap-2">
        <span>Username:</span>
        <AtomsInputsAppPlainInput type="text" v-model="username" />
      </div>
      <div class="flex justify-between items-center gap-2">
        <span>Password:</span>
        <AtomsInputsAppPlainInput type="password" v-model="password" />
      </div>
      <AtomsButtonsGreenRoundedButton
        type="submit"
        text="Submit"
        class="h-[28px]"
      />
      <div class="flex justify-center">
        <NuxtLink to="/signup" class="text-sm text-grey cursor-pointer"
          >Sign Up</NuxtLink
        >
      </div>
    </form>
  </div>
</template>
<script setup>
definePageMeta({
  layout: false,
});

const auth = useAuth();
const router = useRouter();

const username = ref('');
const password = ref('');

function onSubmit() {
  if (!username.value.length || !password.value.length) {
    return;
  }
  auth.signIn(
    { username: username.value, password: password.value },
    { external: false, callbackUrl: '/' }
  );
}
</script>
