<template>

  <div class="TopLink">
    <br>
    <b>{{ msg }}</b>
    <br>
         
    <!-- <router-link to='/input'>
      <h3 class="to_input">inputに遷移</h3>
    </router-link>
    <router-link to='/home'>
      <h3 class="to_home">homeに遷移</h3>
    </router-link> -->
    <router-link to='/signin'>
      <h3 class="to_signin">ログイン</h3>
    </router-link>    
    <router-link to='/signup'>
      <h3 class="to_signup">初めての方はこちら</h3>
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
        <a href="/input" target="_blank" rel="noopener">input</a
        >
      </li>
    </ul> -->
    
  </div>
</template>

<script>

export default {
  name: "TopLink",
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
.to_input{
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
.to_home{
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
.to_signin{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  font-size: 30px;
  background-color: #42b983;
  width: 150px;
  height: 42px;
  border-radius: 8px;

  margin-left: auto;
  margin-right: auto;

}
.to_signup{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  font-size: 15px;
  background-color: #42b983;
  width: 180px;
  height: 24px;
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
