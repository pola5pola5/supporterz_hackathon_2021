module.exports = {
    devServer: {
      proxy: {
        "/api": {
          target: "http://localhost:1323",
        },
      },
    },
  };
  