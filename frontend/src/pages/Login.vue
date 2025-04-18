<script setup>
// Imports
import {onMounted, ref} from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { useToast } from "vue-toastification";
import {checkTokenValidity} from "@/utils/auth.js";
import {API_URL} from "@/config/index.js";
import axios from 'axios'

// Setup Variables
const store = useStore()
const router = useRouter()
const toast = useToast()

const username = ref('')
const password = ref('')

const token = store.getters.getToken

// Functions
const handleLogin = async () => {
  try {
    const response = await axios.post((API_URL + "/api/login"), {
      user: username.value,
      password: password.value,
    })

    if (response.status === 200) {
      store.commit('clearToken')
      store.commit('setToken', response.data['token'])
      store.commit('setUsername', response.data['user']['username'])
      store.commit('setEmail', response.data['user']['email'])
      router.push("/")
      toast.success("Successfully logged in!", {
        timeout: 2000
      });
    } else {
      toast.error("Sorry, there was an error", {
        timeout: 2000
      });
    }
  } catch (error) {
    toast.error("Your username/email or password is incorrect!", {
      timeout: 2000
    });
  }
}

// On page load
onMounted(() => {
  checkTokenValidity(store, router);
});
</script>

<template>
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
          <form class="space-y-4 md:space-y-6" @submit.prevent="handleLogin" autocomplete="off">
            <div>
              <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your Username or Email</label>
              <input type="text" v-model="username" name="username" id="username" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="username/email" required="">
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
              Don’t have an account yet? <router-link to="/register" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</router-link>
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