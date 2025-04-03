<script setup>
import axios from "axios";
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ref, onMounted } from "vue";
import { useToast } from "vue-toastification";
import { CheckUserIsAuthenticated } from "@/functions/index.js";
import { API_URL } from "@/config/index.js";

const router = useRouter()
const store = useStore()
const toast = useToast()
const token = store.getters.getToken
const licenses = ref([]);

const getLicenses = async () => {
  try {
    const response = await axios.get((API_URL + "/api/licenses"), {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
    if (response.status == 200) {
      console.log(response.data)
      licenses.value = response.data
    }
  } catch (error) {
    console.log(error)
  }
}

const deleteLicense = async (keyId) => {
  try {
    const response = await axios.get((API_URL + `/api/licenses/${keyId}`), {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status === 200) {
      licenses.value = licenses.value.filter(license => license.ID !== keyId);
      toast.success("Successfully deleted license", {
        timeout: 2000
      });
    }
  } catch (error) {
    toast.error("Failed to delete license", {
      timeout: 2000
    });
  }
};

const copyLicense = (keyId) => {
    const keyArray = licenses.value;
    const license = keyArray.find(item => item.ID === keyId);
    if (license) {
      navigator.clipboard.writeText(license.Key)
          .then(() => {
            toast.success("Copied to clipboard", 2000);
          })
          .catch(err => {
            console.error('Failed to copy text to clipboard:', err);
          });
    } else {
      toast.error("License not found", 2000);
    }
}


onMounted(() => {
  CheckUserIsAuthenticated(store, router);
  getLicenses();
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
      <tr v-for="license in licenses" :key="license.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 border-gray-200">
        <td class="px-6 py-4">{{ license.ID }}</td>
        <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ license.Key }}</td>
        <td class="px-6 py-4">{{ license.CreatedAt }}</td>
        <td class="px-6 py-4">{{ license.UpdatedAt }}</td>
        <td class="px-6 py-4">{{ license.HWID || 'N/A' }}</td>
        <td class="px-6 py-4 flex flex-row items-center justify-center gap-2">
          <button class="text-red-500 cursor-pointer" @click="deleteLicense(license.ID)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6 color-red-500">
              <path stroke-linecap="round" stroke-linejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
            </svg>
          </button>
          <button class="text-white cursor-pointer" @click="copyLicense(license.ID)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.666 3.888A2.25 2.25 0 0 0 13.5 2.25h-3c-1.03 0-1.9.693-2.166 1.638m7.332 0c.055.194.084.4.084.612v0a.75.75 0 0 1-.75.75H9a.75.75 0 0 1-.75-.75v0c0-.212.03-.418.084-.612m7.332 0c.646.049 1.288.11 1.927.184 1.1.128 1.907 1.077 1.907 2.185V19.5a2.25 2.25 0 0 1-2.25 2.25H6.75A2.25 2.25 0 0 1 4.5 19.5V6.257c0-1.108.806-2.057 1.907-2.185a48.208 48.208 0 0 1 1.927-.184" />
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