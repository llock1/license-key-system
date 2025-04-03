import { createRouter, createWebHistory } from 'vue-router'

// PAGES
import Login from '@/pages/Login.vue'
import Logout from '@/pages/Logout.vue'
import LicensePage from '@/pages/Licenses.vue'
import Register from '@/pages/Register.vue'

// LAYOUTS
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import BlankLout from '@/layouts/BlankLayout.vue'

const routes = [
    {
        path: '/',
        component: DashboardLayout,
        children: [
            { path: '', component: LicensePage },
        ]
    },
    {
        path: '/',
        component: BlankLout,
        children: [
            { path: 'login', component: Login },
            { path: 'logout', component: Logout },
            { path: 'register', component: Register },
        ]
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
