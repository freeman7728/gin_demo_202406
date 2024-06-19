import { createRouter, createWebHistory } from "vue-router";
import login from "@/pages/logins/login.vue";
import index_admin from "@/pages/indexs/index_admin.vue";
import index_ordinary from "@/pages/indexs/index_ordinary.vue";
import exception_window from "@/pages/childpages/exception_window.vue";
import user_info from "@/pages/childpages/user_info.vue";
import employee_info from "@/pages/childpages/employee_info.vue";
import provider from "@/pages/childpages/provider.vue";
import goods from "@/pages/childpages/goods.vue";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/childpages/user_info", name: "userInfo", component: user_info},
  { path: "/childpages/provider", name: "provider", component: provider},
  { path: "/childpages/goods", name: "goods", component: goods},
  { path: "/login", component: login},
  { path: "/index_admin", component: index_admin},
  { path: "/childpages/employee_info", name: "employeeInfo", component: employee_info},
  { path: "/childpages/exception_window", component: exception_window},
  { path: "/index_ordinary", component: index_ordinary},
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
