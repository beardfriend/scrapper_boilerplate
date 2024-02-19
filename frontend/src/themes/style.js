export const constantValues = {
  icon: {
    color: {
      default: "#0B132B",
      active: "#ffffff",
    },
  },
};

export const customColors = {
  navy: {
    origin: "#0B132B",
  },
};

export const brandColors = {
  main: customColors["navy"]["origin"],
};

const htmlStyles = {
  "html, body": {
    boxSizing: "border-box",
    color: "#000000",
  },
  "h1, h2, h3, h4": {
    color: "#000000",
  },

  "p, span, a": {
    color: "#000000",
  },
};

const Styles = {
  fonts: {
    heading: `'DM Sans',sans-serif`,
    body: `'Roboto', sans-serif`,
  },
  styles: {
    global: htmlStyles,
  },
  colors: customColors,
};

export default Styles;
