// store/index.js
import { createStore } from 'vuex'

const store = createStore({
    state() {
        const token = localStorage.getItem('jwt_token')
        return {
            token: token || null,
        }
    },
    mutations: {
        setToken(state, token) {
            state.token = token
            localStorage.setItem('jwt_token', token)
        },
        clearToken(state) {
            state.token = null
            localStorage.removeItem('jwt_token')
        }
    },
    actions: {
        storeToken({ commit }, token) {
            commit('setToken', token)  // Commit the mutation to store the token
        },
        removeToken({ commit }) {
            commit('clearToken')  // Commit the mutation to clear the token
        }
    },
    getters: {
        getToken(state) {
            return state.token  // Access the token from the state
        }
    }
})

export default store
