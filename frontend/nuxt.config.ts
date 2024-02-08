// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['~/assets/css/main.scss'],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  plugins: ['~/plugins/v-select'],
  modules: ['@pinia/nuxt', 'nuxt-icon', '@sidebase/nuxt-auth'],
  auth: {
    globalAppMiddleware: false,
    baseURL: process.env.BASE_API,
    provider: {
      type: 'local',
      endpoints: {
        signIn: { path: '/auth/signin', method: 'post' },
        signOut: { path: '/auth/signout', method: 'get' },
        signUp: { path: '/auth/signup', method: 'post' },
        getSession: { path: '/users/me', method: 'get' },
      },
      pages: {
        login: '/signin',
      },
      token: {
        signInResponseTokenPointer: '/token',
      },
      sessionDataType: {
        token: '',
        id: '',
        fullName: '',
        username: '',
      },
    },
    session: {
      enableRefreshOnWindowFocus: true,
    },
    globalMiddlewareOptions: {
      allow404WithoutAuth: true,
      addDefaultCallbackUrl: false,
    },
  },
});
