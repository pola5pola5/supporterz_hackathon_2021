<template>
  <div id="app">
    <tab-menu></tab-menu>
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <router-view/>
  </div>
</template>
 
<script>
import TabMenu from '@/components/TabMenu'
 
export default {
  name: 'app',
  components: {    
    TabMenu,
  },

}
</script>
 
<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin:  0 auto;
  color: #2c3e50;
  margin-top: 60px;
  width: 800px;
}
</style>