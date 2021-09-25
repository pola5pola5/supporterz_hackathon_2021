<template>
  <div class="TopLink">
    <h2 class="fade-in-top">{{ msg }}</h2>
    <a href="/signup" class="fade-in-bottom">Try now</a>
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
  color: #f1f1f1;
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

.fade-in-top {
  opacity: 0;
  animation: fadein-top 1.5s .3s ease-out forwards;
}

@keyframes fadein-top {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.fade-in-bottom {
  opacity: 0;
  animation: fadein-bottom 1.5s ease-out forwards;
}

@keyframes fadein-bottom {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
