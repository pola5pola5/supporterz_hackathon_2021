import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";

import auth from "./modules/auth";
import trip from "./modules/trip";

const store = createStore({
  modules: {
    auth,
    trip,
  },
  plugins: [
    createPersistedState({
      key: "example",
      storage: window.sessionStorage,
    }),
  ],
});

export default store;
