// import axios from "axios";

export default {
  namespaced: true,
  state: {
    username: "",
  },
  getters: {
    getUserName: (state) => {
      return state.username;
    },
  },
  mutations: {
    setUserName(state, name) {
      state.username = name;
    },
    deleteUserName(state) {
      state.username = "";
    },
  },
};
