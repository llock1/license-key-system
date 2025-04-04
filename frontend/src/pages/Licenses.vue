<script setup>
import axios from "axios";
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { ref, onMounted } from "vue";
import { useToast } from "vue-toastification";
import { CheckUserIsAuthenticated } from "@/functions/index.js";
import { API_URL } from "@/config/index.js";
import DeleteLicense from "@/components/license/DeleteLicense.vue";
import CopyLicense from "@/components/license/CopyLicense.vue";
import NewLicense from "@/components/license/NewLicense.vue";

const router = useRouter()
const store = useStore()
const toast = useToast()
const token = store.getters.getToken
const licenses = ref([]);

// Get all licenses? 
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

// Remove certain license from list
// Used as an emit
const removeLicense = (keyId) => {
  licenses.value = licenses.value.filter(license => license.ID !== keyId);
};

onMounted(() => {
  CheckUserIsAuthenticated(store, router);
  getLicenses();
});
</script>

<template>
  <div class="flex justify-between mb-4">
    <h1 class="font-bold text-2xl">Licenses</h1>
    <NewLicense />
  </div>
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
          <DeleteLicense :licenseID="license.ID" @updateLicenses="removeLicense" />
          <CopyLicense :license="license.Key" />
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>

</style>