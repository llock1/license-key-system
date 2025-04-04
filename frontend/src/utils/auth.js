import axios from "axios";
import {API_URL} from "@/config/index.js";

// To be used in all "restricted" pages
// Redirects to login if JWT is invalid
export async function checkUserIsAuthenticated(
    store, router
) {
    const loginPath = "/login"
    const token = store.getters.getToken

    if (token) {
        try {
            const response = await axios.post((API_URL + "/api/check-token"), { token })

            if (response.status === 200) {
                console.log(response.data);
            } else {
                store.commit('clearAll');
                router.push(loginPath);
            }
        } catch (error) {
            console.log(error);
            router.push(loginPath);
        }
    } else {
        router.push(loginPath);
    }
}

// Used in Login/Register pages to check if you are already authenticated
// Very similiar to above function, but redirects to dashboard instead of login (opposite)
// Plan to merge the two
export async function checkTokenValidity(store, router) {
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
