<template>
  <div v-if="isToken" id="nav">
    <router-link @click="checkToken" to="/home">Home</router-link> |
    <router-link @click="checkToken" to="/routing">Trip List</router-link> |
    <router-link @click="checkToken" to="/map">Map</router-link> |         
    <router-link @click="checkToken" to="/home">logout</router-link>
  </div>
  <div v-else id="nav">
    <router-link @click="checkToken" to="/">Top</router-link> |
    <router-link @click="checkToken" to="/signin">Sign in</router-link>
  </div>
  <router-view />
</template>

<script>
import Store from "@/store/index.js";
export default {
  name: "Header",
  data() {
    return {
      isToken: false,
    }
  },
  methods: {
    checkToken: function() {
      if (Store.state.auth.token) {
        this.isToken = true
      } else {
        this.isToken = false
      }
    }
  }

}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
