import { createRouter, createWebHistory } from "vue-router";
import Top from "../views/Top.vue";
import Home from "../views/Home.vue";
import SignIn from '@/components/SignIn.vue'
import SignUp from '@/components/SignUp.vue'
import Input from "@/components/Input"


const routes = [
  {
    path: "/",
    name: "Top",
    component: Top,
  },
  {
    path: "/home",
    name: "Home",
    component: Home,
  },
  {
    path: '/input',
    name: 'Input',
    component: Input
  },
  {
    path: "/signin",
    name: "SignIn",
    component: SignIn
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUp

  },
  {
    path: "/map",
    name: "Map",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Map.vue"),
  },
  {
    path: "/routing",
    name: "Routing",
    component: () =>
      import("../views/Routing.vue"),
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
