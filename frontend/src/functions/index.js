import axios from "axios";

export async function CheckUserIsAuthenticated(store, router) {
    const token = store.getters.getToken

    if (token) {
        try {
            const response = await axios.post("http://localhost:8090/api/check-token", { token })

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

