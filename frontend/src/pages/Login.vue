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
<!--  <section id="login">-->
<!--    <div class="flex justify-center items-center h-[100dvh]">-->
<!--      <div class="card border border-2 border-grey-700 bg-grey-800 p-8">-->
<!--        <form @submit.prevent="handleLogin" class="flex flex-col gap-4 w-[300px]">-->
<!--          <div class="flex flex-col gap-2">-->
<!--            <label for="username">Username:</label>-->
<!--            <input type="text" v-model="username" name="username" id="username" class="border border-2 border-black">-->
<!--          </div>-->
<!--          <div class="flex flex-col gap-2">-->
<!--            <label for="password">Password:</label>-->
<!--            <input type="password" v-model="password" name="password" id="password" class="border border-2 border-black">-->
<!--          </div>-->
<!--          <input type="submit" value="SUBMIT">-->
<!--        </form>-->
<!--      </div>-->
<!--    </div>-->
<!--  </section>-->

  <section class="bg-gray-50 dark:bg-gray-900">
    <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
      <a href="#" class="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
        Licensing System
      </a>
      <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
            Sign in to your account
          </h1>
          <form class="space-y-4 md:space-y-6" @submit.prevent="handleLogin">
            <div>
              <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your username</label>
              <input type="text" v-model="username" name="username" id="username" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="username" required="">
            </div>
            <div>
              <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
              <input type="password" v-model="password" name="password" id="password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required="">
            </div>
            <div class="flex items-center justify-center">
              <a href="#" class="text-sm font-medium text-blue-600 hover:underline dark:text-primary-500">Forgot password?</a>
            </div>
            <button type="submit" class="w-full text-white bg-blue-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign in</button>
            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
              Don’t have an account yet? <a href="#" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</a>
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
#login {
  height: 100dvh;
}

</style>