<template>

  <div class="top">
    <br>
    <b>{{ msg }}</b>
    <br>
         
    <router-link to='/view'>
      <h3>早速使ってみる</h3>
    </router-link>

    <!-- <h4>チェック用リンク</h4>
    <ul>
      <li>
        <a href="/getpost" target="_blank" rel="noopener">GetPost</a>
      </li>
      <li>
        <a href="/upload" target="_blank" rel="noopener">upload</a>
      </li>
      <li>
        <a href="/view" target="_blank" rel="noopener">view</a
        >
      </li>
    </ul> -->
    
  </div>
</template>

<script>

export default {
  name: "Top",
  props: {
    msg: String,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 40px 0 0;
  color: #42b983;
}
h3{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  background-color: #42b983;
  width: 150px;
  height: 30px;
  border-radius: 8px;

  margin-left: auto;
  margin-right: auto;

}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
