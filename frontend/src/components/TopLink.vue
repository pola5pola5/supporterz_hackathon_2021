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

h2 {
  margin-top: 40px;
  margin-bottom: 30px;
  font-size: 40px;
  font-weight: normal;
  color: #ffffff;
  letter-spacing: 0.2em;
  text-align: center;
  line-height: 1.6;
}

.TopLink {
  background: #2d2d2d;
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
  font-size: 25px;
  color: #ffffff;
  cursor: pointer;
  border: 0px;
}
</style>
