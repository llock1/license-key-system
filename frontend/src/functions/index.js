import axios from "axios";
import {API_URL} from "@/config/index.js";

// To be used in all "restricted" pages
// Redirects to login if JWT is invalid
export async function CheckUserIsAuthenticated(store, router) {
    const token = store.getters.getToken

    if (token) {
        try {
            const response = await axios.post((API_URL + "/api/check-token"), { token })

            if (response.status === 200) {
                console.log(response.data);
            } else {
                store.commit('clearAll');
                router.push("/login");
            }
        } catch (error) {
            console.log(error);
            router.push("/login");
        }
    } else {
        router.push("/login");
    }
}

// Used in Login/Register pages to check if you are already authenticated
// Very similiar to above function, but redirects to dashboard instead of login (opposite)
// Plan to merge the two
export async function CheckTokenValidity(store, router) {
    const token = store.getters.getToken;
    if (token) {
        try {
            const response = await axios.post((API_URL + "/api/check-token"), { token });

            if (response.status === 200) {
                router.push("/");
            } else {
                store.commit('clearAll');
            }
        } catch (error) {
            console.log(error);
        }
    }
}