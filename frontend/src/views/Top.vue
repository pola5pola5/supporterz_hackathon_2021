<template>
  <div class="top">
    <header class="header">
      <h1><a href="/">
        フォト旅
      </a></h1>
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
    <TopLink class="toplink" msg="写真で巡るあなたの旅物語"/>
    <div class="item01">
        <span>
          <h2 class="fade-in-top">
            旅の記録を思い出と共に残しませんか？
          </h2>
          <h3 class="fade-in-top">
            地図＋写真で振り返る新しい旅記録
          </h3>
        </span>
    </div>
    <div class="item02">
      <span>
        <h2 class="item_msg">方法は簡単<br>旅で撮った写真を<br>アップロードするだけ</h2>
      </span>
    </div>
    <div class="item03">
      <!--      <img class="photo2" src="../assets/photo2.jpg"/>-->
      <img class="phone_photo" src="../assets/phone_photo.jpg"/>
    </div>
    <div class="item04">
      <img class="map_human" src="../assets/map_human.png"/>
    </div>
    <div class="item05">
      <h2 class="item_msg">
        フォト旅が<br>あなたの旅を写真と共に<br>記録します
      </h2>
    </div>
    <div class="item06">
        <span>
        <h2 class="item_msg">
          記録された旅は<br>いつでも見返せるから<br>思い出作りに便利
        </h2>
        </span>
    </div>
    <div class="item07">
      <img class="travel" src="../assets/travel.png"/>
    </div>
    <div class="item08">
      <h2 class="msg">あなたの思い出作りにフォト旅を</h2>
      <a href="/signup" class="try_btn">Try now</a>
    </div>
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
  data() {
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

.top {
  display: grid;
  grid-template-columns: 200px 1fr 180px;
  grid-template-rows: 60px 300px 160px 220px 180px 160px 200px;
}

.header {
  grid-row-start: 1;
  grid-row-end: auto;
  grid-column: 1/4;
  font-family: serif;
  padding: 30px 4% 10px;
  background: #2d2d2d;
  display: flex;
  align-items: center;
}

h1 {
  margin: 0;
  padding: 0;
  font-size: 20px;
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
  grid-row: 2;
  grid-column: 1/4;
  padding: 0 4%;
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
  transform: rotate(45deg);
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
  transform: rotate(-45deg);
}

.toggle {
  transform: translateY(0);
  opacity: 1;
}

@media screen and (min-width: 641px) {
  #hamburger {
    display: none;
  }
}

h2 {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 18px;
  font-weight: normal;
  color: #2d2d2d;
  letter-spacing: 0.2em;
  text-align: center;
  line-height: 1.6;
}

h3 {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 16px;
  font-weight: normal;
  color: #2d2d2d;
  letter-spacing: 0.2em;
  text-align: center;
  line-height: 1.6;
}

.try_btn {
  margin-top: 20px;
  margin-bottom: 10px;
  font-size: 16px;
  font-weight: normal;
  letter-spacing: 0.2em;
  text-align: center;
  line-height: 1.6;
  color: #ffffff;
}

a {
  text-decoration: none;
  color: #ffffff;
  cursor: pointer;
}

.item01 {
  grid-row: 3;
  grid-column: 1/4;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item02 {
  grid-row: 4;
  grid-column: 1/3;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item03 {
  grid-row: 4;
  grid-column: 3;
  display: flex;
  flex-flow: column;
  justify-content: left;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item04 {
  grid-row: 5;
  grid-column: 1;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item05 {
  grid-row: 5;
  grid-column: 2/4;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item06 {
  grid-row: 6;
  grid-column: 1/3;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item07 {
  grid-row: 6;
  grid-column: 3;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  background: #ffffff;
}

.item08 {
  grid-row: 7;
  grid-column: 1/4;
  display: flex;
  flex-flow: column;
  justify-content: center;
  align-items: center;
  font-family: serif;
  padding: 0 4%;
  background: #2d2d2d;
}

.item_msg {
  text-align: left;
  font-size: 14px;
  opacity: 0;
  animation: item-msg 1.5s ease-out 0s 1 forwards;
}
@keyframes item-msg {
  0% {
    opacity: 0;
    transform: translateX(-20px);
  }
  100% {
    opacity: 1;
    transform: translateX(0);
  }
}

.msg {
  color: #f1f1f1;
}

.phone_photo {
  width: 70%;
  margin-top: 30px;
  opacity: 0;
  animation: fadein-bottom 1.5s .3s ease-out forwards;
}
@keyframes fadein-bottom {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.map_human {
  width: 90%;
  /*margin-top: 40px;*/
  /*margin-left: 20px;*/
  opacity: 0;
  animation: map-human 1.5s .3s ease-out forwards;
}

@keyframes map-human {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.travel {
  width: 90%;
  /*margin-top: 30px;*/
  /*margin-right: 20px;*/
  opacity: 0;
  animation: travel 1.5s .3s ease-out forwards;
}

@keyframes travel {
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
  animation: fadein-bottom 1.5s 0 ease-out forwards;
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
</style>
