import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import Map from "@/views/Map.vue"
import Routing from "@/views/Routing.vue"
import SignIn from '@/components/SignIn.vue'
import SignUp from '@/components/SignUp.vue'
import View from "@/components/View"

import Store from "@/store/index.js";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      isPublic: true
    }
  },
  {
    path: "/signin",
    name: "SignIn",
    component: SignIn,
    meta: {
      isPublic: true
    }
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUp,
    meta: {
      isPublic: true
    }
  },
  {
    path: '/view',
    name: 'View',
    component: View,
  },
  {
    path: "/map",
    name: "Map",
    component: Map,
  },
  {
    path: "/routing",
    name: "Routing",
    component: Routing
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

// ログインが必要なページかどうかを判定
router.beforeEach((to, from, next) => {
  if (to.matched.some(page => page.meta.isPublic) || Store.state.auth.token) {
    next()
  } else {
    next('/signin')
  }
})

export default router;
