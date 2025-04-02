import axios from "axios";
import {API_URL} from "@/config/index.js";

export async function CheckUserIsAuthenticated(store, router) {
    const token = store.getters.getToken

    if (token) {
        try {
            const response = await axios.post((API_URL + "/api/check-token"), { token })

            if (response.status === 200) {
                console.log(response.data);
            } else {
                store.dispatch('removeToken');
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

