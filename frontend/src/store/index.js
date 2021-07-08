import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";

import auth from "./modules/auth";

const store = createStore({
  state() {
    return {
      tripid: "",
    };
  },
  mutations: {
    pushid(state, id) {
      state.tripid = id;
    },
  },
  modules: {
    auth,
  },
  plugins: [
    createPersistedState({
      key: "example",
      storage: window.sessionStorage,
    }),
  ],
});

export default store;
