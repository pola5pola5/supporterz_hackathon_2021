export default {
  namespaced: true,
  state: {
    tripid: "",
    tripid_num: "",
  },
  getters: {
    getTripID: (state) => {
      return state.tripid;
    },
    getNumTripID: (state) => {
      return state.tripid_num;
    },
  },
  mutations: {
    setTripID(state, id) {
      state.tripid = id;
    },
    deleteTripID(state) {
      state.tripid = "";
    },
    setNumTripID(state, id_num) {
      state.tripid_num = id_num;
    },
  },
};
