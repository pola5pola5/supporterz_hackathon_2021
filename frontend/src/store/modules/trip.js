export default {
  namespaced: true,
  state: {
    tripid: "",
  },
  getters: {
    getTripID: (state) => {
      return state.tripid;
    },
  },
  mutations: {
    setTripID(state, id) {
      state.tripid = id;
    },
    deleteTripID(state) {
      state.tripid = "";
    },
  },
};
