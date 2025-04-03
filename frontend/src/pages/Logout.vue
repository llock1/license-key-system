<script setup>
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const store = useStore()
const router = useRouter()
const token = store.getters.getToken

if (token) {
  store.commit("clearToken")
}

router.push("/login")
</script>

<template>

</template>

<style scoped>

</style>