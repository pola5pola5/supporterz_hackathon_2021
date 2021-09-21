<template>
  <div class="top">
    <header class="header">
      <h1><a href="/">
        フォト旅
      </a></h1>
      <!--      <router-link to="/signin" tag="button" class="signin_btn">Sign in</router-link>-->
      <!--      <router-link to="/signup" tag="button" class="signup_btn">Sign up</router-link>-->
      <nav class="pc_nav">
        <ul>
          <li><a href="/signin">Sign In</a></li>
          <li><a href="/signup">Sign Up</a></li>
        </ul>
      </nav>
      <nav class="sp_nav" v-bind:class='{toggle:isActive}'>
        <ul>
          <li><a href="/signin">Sign In</a></li>
          <li><a href="/signup">Sign Up</a></li>
          <li class="close" v-on:click='isActive=!isActive'><span>閉じる</span></li>
        </ul>
      </nav>
      <div id="hamburger" v-on:click='isActive=!isActive'>
        <span></span>
      </div>
    </header>
    <TopLink class="toplink" msg="写真で巡るあなたの旅物語" />
    <main class="main">main</main>

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
  data(){
    return {
      isActive: false
    }
  }
};
</script>
<style scoped>
* {
  box-sizing: border-box;
}
.top{
  display: grid;
   grid-template-rows: 60px 300px 1fr
}
.header{
  grid-row-start: 1;
  grid-row-end: auto;
  font-family: serif;
  padding: 30px 4% 10px;
  background: #2d2d2d;
  display: flex;
  align-items: center;
}
h1 {
  margin: 0; padding: 0;
  font-size: 20px;
}
a {
  text-decoration: none;
  color: #ffffff;
  cursor: pointer;
}
nav {
  margin: 0 0 0 auto;
}
ul {
  list-style: none;
  margin: 0;
  display: flex;
}
li {
  margin: 0 0 0 15px;
  font-size: 14px;
}
.toplink {
  grid-row-start: 2;
  grid-row-end: auto;
  padding: 0 4%;
}
.main{
  grid-row-start: 3;
  grid-row-end: auto;
}
@media screen and (max-width: 640px) {
  .pc_nav {
    display: none;
  }
}

#hamburger {
  position: relative;
  display: block;
  width: 30px;
  height: 25px;
  margin: 0 0 0 auto;
}
#hamburger span {
  position: absolute;
  top: 50%;
  left: 0;
  display: block;
  width: 100%;
  height: 2px;
  background-color: #ffffff;
  transform: translateY(-50%);
}
#hamburger::before {
  content: '';
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #ffffff;
}
#hamburger::after {
  content: '';
  display: block;
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #ffffff;
}
.sp_nav {
  z-index: 1;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  display: block;
  background: rgba(0, 0, 0, .8);
  transition: all .2s ease-in-out;
  opacity: 0;
  transform: translateY(-100%);
}
.sp_nav ul {
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
}
.sp_nav li {
  margin: 0;
  padding: 0;
}
.sp_nav li span {
  font-size: 15px;
  color: #ffffff;
}
.sp_nav li a, .sp_nav li span {
  display: block;
  padding: 20px 0;
}
/*基準となるli要素*/
.sp_nav .close {
  position: relative;
  padding-left: 20px;
}
/*バツ印線1*/
.sp_nav .close::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  display: block;
  width: 16px;
  height: 1px;
  background: #ffffff;
  transform: rotate( 45deg );
}
/*バツ印線2*/
.sp_nav .close::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  display: block;
  width: 16px;
  height: 1px;
  background: #ffffff;
  transform: rotate( -45deg );
}
.toggle {
  transform: translateY( 0 );
  opacity: 1;
}
@media screen and (min-width: 641px){
  #hamburger {
    display: none;
  }
}
</style>
