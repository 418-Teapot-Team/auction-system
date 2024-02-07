/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './components/**/*.{js,vue,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.{js,ts}',
    './app.vue',
    './error.vue',
  ],
  theme: {
    extend: {
      fontFamily: {
        poppins: ['Poppins', 'Helvetica', 'Arial', 'sans-serif'],
      },
      colors: {
        'dark-grey': '#272D2D',
        'darkest-grey': '#222727',
        milky: '#DFDFD7',
        'light-milky': '#EAEAE4',
        white: '#F1F1EE',
        grey: '#585855',
        'light-green': '#77FF61',
        'menu-bg': '#383E3E',
      },
    },
  },
  plugins: [],
};
