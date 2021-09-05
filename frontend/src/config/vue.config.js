module.exports = {
  pages: {
    index: {
      entry: "src/main.js",
      title: "フォトたび",
    },
  },
  devServer: {
    proxy: {
      "/api": {
        // target: "http://3.112.234.249:1323",
        target: "http://3.112.130.126:1323",
      },
    },
  },
};
