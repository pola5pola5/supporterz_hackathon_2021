<template>
  <div class="top">
    <br />
    <img alt="Vue logo" src="../assets/logo.png" />

    <TopLink msg="写真の位置情報からマップを作成します．" />
  </div>
</template>

<script>
// @ is an alias to /src
import TopLink from "@/components/TopLink.vue";

export default {
  name: "Top",
  components: {
    TopLink,
  },
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
