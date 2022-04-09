module.exports = {
  mode: 'jit',
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        lightGray: '#F3F3F3',
        primaryBlue: '#0070F3',
      },
      fontFamily: {
        primary: "'Inter', sans-serif",
      }
    },
  },
  plugins: [],
}
