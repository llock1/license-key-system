<script setup>
import axios from "axios";
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ref, onMounted } from "vue";
import { useToast } from "vue-toastification";
import { CheckUserIsAuthenticated } from "@/functions/index.js";

const router = useRouter()
const store = useStore()
const toast = useToast()
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

const deleteKey = async (keyId) => {
  try {
    const response = await axios.get(`http://localhost:8090/api/keys/${keyId}/delete`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status === 200) {
      keys.value = keys.value.filter(key => key.ID !== keyId);
      toast.success("Successfully Deleted Key", {
        timeout: 2000
      });
    }
  } catch (error) {
    toast.error("Failed to Delete Key", {
      timeout: 2000
    });
  }
};


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
          ID
        </th>
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
        <th scope="col" class="px-6 py-3">
          Actions
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="key in keys" :key="key.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 border-gray-200">
        <td class="px-6 py-4">{{ key.ID }}</td>
        <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ key.Key }}</td>
        <td class="px-6 py-4">{{ key.CreatedAt }}</td>
        <td class="px-6 py-4">{{ key.UpdatedAt }}</td>
        <td class="px-6 py-4">{{ key.Hwid || 'N/A' }}</td>
        <td class="px-6 py-4">
          <button class="text-red-500" @click="deleteKey(key.ID)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6 color-red-500">
              <path stroke-linecap="round" stroke-linejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
            </svg>
          </button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>

</style>