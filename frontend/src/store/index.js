// store/index.js
import { createStore } from 'vuex'

// VUEX Store
// Using this rather than normal localStorage, as it updates components when its updated
// clearAll mutation needs refactoring -> WILL DO THIS!!

const store = createStore({
    state() {
        const token = localStorage.getItem('jwt_token');
        const username = localStorage.getItem('username');
        const email = localStorage.getItem('email');
        return {
            token: token || null,
            username: username || null,
            email: email || null,
        }
    },
    mutations: {
        setToken(state, token) {
            state.token = token;
            localStorage.setItem('jwt_token', token);
        },
        clearToken(state) {
            state.token = null;
            localStorage.removeItem('jwt_token');
        },

        setEmail(state, email) {
            state.email = email;
            localStorage.setItem('email', email);
        },
        clearEmail(state) {
            state.email = null;
            localStorage.removeItem('email');
        },

        setUsername(state, username) {
            state.username = username;
            localStorage.setItem("username", username);
        },
        clearUsername(state) {
            state.username = null;
            localStorage.removeItem("username");
        },
        clearAll(state) {
            state.token = null;
            state.username = null;
            state.email = null;

            localStorage.removeItem("jwt_token");
            localStorage.removeItem("username");
            localStorage.removeItem("email");
        }
    },
    actions: {
        // Will add here if thing become more complicated
    },
    getters: {
        getToken(state) {
            return state.token;
        },
        getUsername(state) {
            return state.username;
        },
        getEmail(state) {
            return state.email;
        },
    }
})

export default store;
