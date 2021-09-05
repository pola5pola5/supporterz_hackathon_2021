import { createRouter, createWebHistory } from "vue-router";

import Top from "@/views/Top.vue";
import Home from "@/views/Home.vue";
import Map from "@/views/Map.vue";
import Routing from "@/views/Routing.vue";

import SignIn from "@/components/SignIn.vue";
import SignUp from "@/components/SignUp.vue";
import Input from "@/components/Input";

import Test from "@/components/Test";
import Store from "@/store/index.js";

const routes = [
  {
    //初めのページ
    path: "/",
    name: "Top",
    component: Top,
    meta: {
      isPublic: true,
    },
  },
  {
    //ログイン後に現れるページ
    path: "/home",
    name: "Home",
    component: Home,
  },
  {
    path: "/test",
    name: "Test",
    component: Test,
    meta: {
      isPublic: true,
    },
  },
  {
    path: "/signin",
    name: "SignIn",
    component: SignIn,
    meta: {
      isPublic: true,
    },
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUp,
    meta: {
      isPublic: true,
    },
  },
  {
    //元々viewだったページ
    path: "/input",
    name: "Input",
    component: Input,
  },
  {
    path: "/map",
    name: "Map",
    component: Map,
  },
  {
    path: "/routing",
    name: "Routing",
    component: Routing,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

// ログインが必要なページかどうかを判定
router.beforeEach((to, from, next) => {
  if (to.matched.some((page) => page.meta.isPublic) || Store.state.auth.token) {
    next();
  } else {
    next("/signin");
  }
});

export default router;
