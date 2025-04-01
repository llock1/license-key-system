<script setup>
import axios from "axios";
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ref, onMounted } from "vue";
import { CheckUserIsAuthenticated } from "@/functions/index.js";

const router = useRouter()
const store = useStore()
const API_URL = "127.0.0.1:8090/"
const token = store.getters.getToken
const keys = ref([]);

const getKeys = async () => {
  try {
    const response = await axios.get("http://localhost:8090/api/keys", {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
    if (response.status == 200) {
      console.log(response.data)
      keys.value = response.data
    }
  } catch (error) {
    console.log(error)
  }
}

onMounted(() => {
  getKeys();
  CheckUserIsAuthenticated(store, router);
});
</script>

<template>
  <div class="relative overflow-x-auto">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
      <tr>
        <th scope="col" class="px-6 py-3">
          Key
        </th>
        <th scope="col" class="px-6 py-3">
          Created
        </th>
        <th scope="col" class="px-6 py-3">
          Updated
        </th>
        <th scope="col" class="px-6 py-3">
          HWID
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="key in keys" :key="key.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 border-gray-200">
        <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ key.Key }}</td>
        <td class="px-6 py-4">{{ key.CreatedAt }}</td>
        <td class="px-6 py-4">{{ key.UpdatedAt }}</td>
        <td class="px-6 py-4">{{ key.Hwid || 'N/A' }}</td>  <!-- Display 'N/A' if HWID is empty -->
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>

</style>