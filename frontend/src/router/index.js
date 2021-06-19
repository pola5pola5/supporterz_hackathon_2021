import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Upload from "@/components/Upload"
import GetPost from "@/components/GetPost"
import View from "@/components/View"

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/upload",
    name: "Upload",
    component: Upload
  },
  {
    path: '/getpost',
    name: 'GetPost',
    component: GetPost
  },
  {
    path: '/view',
    name: 'View',
    component: View
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
