import axios from "axios"

export default {
    namespaced: true,
    state: {
        userid: "",
        token: "",
    },
    getters: {
        getUserID: (state) => {
            return state.userid
        },
        getToken: (state) => {
            return state.token
        },
    },
    mutations: {
        setInfo(state, payload) {
            state.userid = payload.userid;
            state.token = payload.token;
        },
    },
    actions: {
        async signUp(context, payload) {
            const url = "/api/user/create";
            return axios.post(url, payload)
                .then(function (response) {
                    console.log("signUp_response", response.data);
                    context.commit("setInfo", response.data);
                    return response.data;
                })
                .catch((error) => {
                    console.log(error)
                })
        },
        async signIn(context, payload) {
            const url = "/api/user/login";
            return axios.post(url, payload)
                .then(function (response) {
                    console.log("signIn_response", response.data);
                    context.commit("setInfo", response.data);
                    return response.data;
                })
                .catch((error) => {
                    console.log(error)
                })
        }

    }
}