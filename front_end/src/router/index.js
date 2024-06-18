import { createRouter, createWebHistory } from "vue-router";
import login from "@/pages/logins/login.vue";
import index_admin from "@/pages/indexs/index_admin.vue";
import index_ordinary from "@/pages/indexs/index_ordinary.vue";
import exception_window from "@/pages/childpages/exception_window.vue";
import user_info from "@/pages/childpages/user_info.vue";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/childpages/user_info", name: "userInfo", component: user_info},
  { path: "/login", component: login},
  { path: "/index_admin", component: index_admin},
  { path: "/childpages/exception_window", component: exception_window},
  { path: "/index_ordinary", component: index_ordinary},
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
