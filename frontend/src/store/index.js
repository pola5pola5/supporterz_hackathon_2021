import { createStore } from "vuex"

const store = createStore({
    state() {
        return{
            tripid: "b0add9e8-cdba-42ee-85d5-be89bd9875f0",
        }
    },
    mutations: {
        pushid(state, id){
            state.tripid = id;
        },
    }
});

export default store;
