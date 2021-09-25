<template>
  <div class="map" v-cloak>
    <MyMap />
  </div>
</template>

<script>
import MyMap from "@/components/MyMap.vue";
export default {
  name: "map",
  components: {
    MyMap,
  },
};
</script>
