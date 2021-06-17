<template>
  <div class="about">
    <h1>This is an about page</h1>
    <MyMap/>
  </div>
</template>

<script>
  import MyMap from '@/components/MyMap.vue'
  export default {
    name: 'about',
    components:{
      MyMap
    }
  }
</script>