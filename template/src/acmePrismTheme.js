// An acme-inspired theme for CodeSurfer / mdx-deck
// Ref: https://github.com/plan9-for-vimspace/acme-colors/blob/master/colors/acme.vim

var theme = {
  plain: {
    color: "#000000",
    backgroundColor: "#ffffea"
  },
  styles: [
    {
      types: ["comment"],
      style: {
        color: "#000000"
      }
    },
    {
      types: ["string", "number", "builtin", "variable"],
      style: {
        color: "#000000"
      }
    },
    {
      types: ["class-name", "function", "tag", "attr-name"],
      style: {
        color: "#000000"
      }
    }
  ]
};

module.exports = theme;
