<template>
  <div class="home">
    <br />
    <img alt="Vue logo" src="../assets/logo.png" />

    <HomeLink />
    <!-- <HomeLink msg="こんにちは，ゲストさん" /> -->

    <h3 @click="signOut">sign out</h3>
  </div>
</template>

<script>
// @ is an alias to /src
import HomeLink from "@/components/HomeLink.vue";

export default {
  name: "Home",
  components: {
    HomeLink,
  },
  methods: {
    signOut: function () {
      this.$store.commit("auth/deleteInfo");
      this.$router.push("/");
    }
  }
};
</script>
<style scoped>
p {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 40px 0 0;
  color: #42b983;
}
</style>
