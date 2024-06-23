import { createRouter, createWebHistory } from "vue-router";
import login from "@/pages/logins/login.vue";
import index_admin from "@/pages/indexs/index_admin.vue";
import index_ordinary from "@/pages/indexs/index_ordinary.vue";
import user_info from "@/pages/childpages/user_info.vue";
import user_info_o from "@/pages/childpages/user_info_o.vue";
import employee_info from "@/pages/childpages/employee_info.vue";
import provider from "@/pages/childpages/provider.vue";
import goods from "@/pages/childpages/goods.vue";
import list from "@/pages/childpages/list.vue";
import order from "@/pages/childpages/order.vue";
import employee_info_o from "@/pages/childpages/employee_info_o.vue";
import provider_o from "@/pages/childpages/provider_o.vue";
import goods_o from "@/pages/childpages/goods_o.vue";
import list_o from "@/pages/childpages/list_o.vue";
import order_o from "@/pages/childpages/order_o.vue";
import addEmployee from "@/components/addEmployee.vue";
import addClient from "@/components/addClient.vue";
import addGoods from "@/components/addGoods.vue";
import addList from "@/components/addList.vue";
import addOrder from "@/components/addOrder.vue";
import updateEmployee from "@/components/updateEmployee.vue";
import updateGoods from "@/components/updateGoods.vue";
import updateClient from "@/components/updateClient.vue";
import updateList from "@/components/updateList.vue";
import updateOrder from "@/components/updateOrder.vue";
import delEmployee from "@/components/delEmployee.vue";
import delClient from "@/components/delClient.vue";
import delGoods from "@/components/delGoods.vue";
import delList from "@/components/delList.vue";
import delOrder from "@/components/delOrder.vue";
import searchEmployee from "@/components/searchEmployee.vue";
import searchClient from "@/components/searchClient.vue";
import searchList from "@/components/searchList.vue";
import searchGoods from "@/components/searchGoods.vue";
import searchOrder from "@/components/searchOrder.vue";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/childpages/user_info", name: "userInfo", component: user_info},
  { path: "/childpages/user_info_o", name: "user_info_o", component: user_info_o},
  { path: "/childpages/provider", name: "provider", component: provider},
  { path: "/childpages/order", name: "order", component: order},
  { path: "/childpages/list", name: "list", component: list},
  { path: "/childpages/goods", name: "goods", component: goods},
  { path: "/childpages/provider_o", name: "provider_o", component: provider_o},
  { path: "/childpages/order_o", name: "order_o", component: order_o},
  { path: "/childpages/list_o", name: "list_o", component: list_o},
  { path: "/childpages/goods_o", name: "goods_o", component: goods_o},
  { path: "/login", component: login},
  { path: "/addEmployee", name: "addEmployee", component: addEmployee},
  { path: "/addClient", name: "addClient", component: addClient},
  { path: "/addGoods", name: "addGoods", component: addGoods},
  { path: "/addList", name: "addList", component: addList},
  { path: "/addOrder", name: "addOrder", component: addOrder},
  { path: "/index_admin", component: index_admin},
  { path: "/childpages/employee_info", name: "employeeInfo", component: employee_info},
  { path: "/childpages/employee_info_o", name: "employeeInfo_o", component: employee_info_o},
  { path: "/updateEmployee", name:"updateEmployee",component: updateEmployee},
  { path: "/updateGoods", name:"updateGoods",component: updateGoods},
  { path: "/updateClient", name:"updateClient",component: updateClient},
  { path: "/updateList", name:"updateList",component: updateList},
  { path: "/updateOrder", name:"updateOrder",component: updateOrder},
  { path: "/delEmployee", name:"delEmployee" ,component: delEmployee},
  { path: "/delClient", name:"delClient" ,component: delClient},
  { path: "/delGoods", name:"delGoods" ,component: delGoods},
  { path: "/delList", name:"delList" ,component: delList},
  { path: "/delOrder", name:"delOrder" ,component: delOrder},
  { path: "/searchEmployee", name:"searchEmployee" ,component: searchEmployee},
  { path: "/searchClient", name: "searchClient", component: searchClient},
  { path: "/searchGoods", name:"searchGoods" ,component: searchGoods},
  { path: "/searchList", name: "searchList", component: searchList},
  { path: "/searchOrder", name: "searchOrder", component: searchOrder},
  { path: "/index_ordinary", component: index_ordinary},
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
