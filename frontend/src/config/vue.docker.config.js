module.exports = {
  devServer: {
    proxy: {
      "/api": {
        target: "http://192.168.0.3:1323",
      },
    },
  },
};
