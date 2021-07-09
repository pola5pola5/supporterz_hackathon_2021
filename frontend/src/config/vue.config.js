module.exports = {
  devServer: {
    proxy: {
      "/api": {
        target: "http://3.112.234.249:1323",
      },
    },
  },
};
