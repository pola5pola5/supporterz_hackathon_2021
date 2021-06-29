<template>
  <div class="home">
    <br>
    <img alt="Vue logo" src="../assets/logo.png" />

    <Top msg="写真の位置情報からマップを作成します．" />

  </div>

</template>

<script>
// @ is an alias to /src
import Top from "@/components/Top.vue";


export default {
  name: "Home",
  components: {
    Top,
  },
};
</script>
<style scoped>
p {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 40px 0 0;
  color: #42b983;
}
</style>
