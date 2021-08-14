<template>
  <div v-if="isToken" id="nav">
    <!-- <router-link to="/home">Home</router-link> |
    <router-link to="/routing">Trip List</router-link> -->
    <!-- <router-link to="/map">Map</router-link> -->
    <!-- <router-link to="/home">logout</router-link> -->
  </div>
  <div v-else id="nav">
    <!-- <router-link to="/">Top</router-link> |
    <router-link to="/signin">Sign in</router-link> -->
  </div>
  <router-view />
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      isToken: false,
    };
  },
  computed: {
    token: function () {
      return this.$store.getters["auth/getToken"];
    },
  },
  watch: {
    token: function () {
      this.isToken = true;
    },
  },
};
</script>

<style>
/* #app {
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
} */
</style>
