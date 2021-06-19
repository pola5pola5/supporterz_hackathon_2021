import { createStore } from "vuex"

const store = createStore({
    state() {
        return{
            tripid: "ade01599-4acb-40a8-870c-4b9eea358d22"
        }
    },
    mutations: {
        pushid(state, id){
            state.tripid = id;
        }
    }
});

export default store;
