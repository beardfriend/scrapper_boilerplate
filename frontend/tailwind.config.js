const { nextui } = require("@nextui-org/react");

/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}", "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      spacing: {
        160: "40rem",
      },
      colors: {},
    },
  },
  darkMode: "class",
  plugins: [
    nextui({
      themes: {
        black: {
          extend: "dark", // <- inherit default values from dark theme
          colors: {
            background: "#000000",
            color: "#fffff",
            foreground: "#ffffff",
            primary: {
              DEFAULT: "#000000",
              foreground: "#ffffff",
            },
            focus: "#fffff",
          },
          layout: {
            disabledOpacity: "0.3",
            radius: {
              small: "4px",
              medium: "6px",
              large: "8px",
            },
            borderWidth: {
              small: "1px",
              medium: "2px",
              large: "3px",
            },
          },
        },
      },
    }),
  ],
};
