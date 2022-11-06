/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      boxShadow: {
        go: "0px 4px 6px 0px rgba(0,0,0,0.1)",
      },
      backgroundImage: {
        "red-gradient":
          "linear-gradient(270deg, rgba(216,216,216,0) 0%, #FFE0E0 100%)",
      },
    },
  },
  plugins: [],
};
