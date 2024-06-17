import { createRouter, createWebHistory } from "vue-router";
import login from "@/pages/logins/login.vue";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/login", component: login},
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
