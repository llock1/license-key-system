<script setup>
import {onMounted, ref} from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import axios from 'axios'

const store = useStore()
const router = useRouter()

const username = ref('')
const password = ref('')

const token = store.getters.getToken

const checkTokenValidity = async () => {
  if (token) {
    try {
      const response = await axios.post("http://localhost:8090/check-token", { token })

      if (response.status === 200) {
        console.log(response.data)
        router.push("/")
      } else {
        store.dispatch('removeToken')
      }
    } catch (error) {
      console.log(error)
    }
  }
}


const handleLogin = async () => {
  try {
    const response = await axios.post("http://localhost:8090/login", {
      username: username.value,
      password: password.value,
    })

    if (response.status === 200) {
      console.log(response.data)
      store.dispatch('removeToken')
      store.dispatch('storeToken', response.data)
      router.push("/")
    } else {
      console.log("BAAD")
    }
  } catch (error) {
    console.log(error)
  }
}

onMounted(
    checkTokenValidity
)
</script>

<template>
  <section id="login">
    <div class="flex justify-center items-center h-[100dvh]">
      <div class="card border border-2 border-grey-700 bg-grey-800 p-8">
        <form @submit.prevent="handleLogin" class="flex flex-col gap-4 w-[300px]">
          <div class="flex flex-col gap-2">
            <label for="username">Username:</label>
            <input type="text" v-model="username" name="username" id="username" class="border border-2 border-black">
          </div>
          <div class="flex flex-col gap-2">
            <label for="password">Password:</label>
            <input type="password" v-model="password" name="password" id="password" class="border border-2 border-black">
          </div>
          <input type="submit" value="SUBMIT">
        </form>
      </div>
    </div>
  </section>
</template>

<style scoped>
#login {
  height: 100dvh;
}

</style>