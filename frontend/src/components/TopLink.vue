<template>
  <div class="TopLink">
    <h2>{{ msg }}</h2>
    <a href="/signup">Try now</a>
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
* {
  box-sizing: border-box;
}

h2{
  margin-top: 20px;
  margin-bottom: 50px;
  font-size: 40px;
  font-weight: normal;
  color: #ffffff;
  letter-spacing : 0.2em;
  text-align: center;
}
.TopLink{
  background: #2d2d2d top center / cover no-repeat;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
}

a {
  margin-top: 20px;
  margin-bottom: 50px;
  text-decoration: none;
  font-size: 15px;
  color: #ffffff;
  cursor: pointer;
  border: 0px;
}
</style>
