<template>
  <div v-if="tripIds.length == 0">
    ぜひ新しい旅を記録しましょう！
    <router-link to="/input">
      <h3 class="to_input">新しい旅を記録する</h3>
    </router-link>
  </div>
  <div v-else class="TripMemory">
    <div id="routing">
      <div v-for="(value, index) in tripIds" :key="value.id">
        <button type="button" @click="onClickMyMap(value)">旅{{ index + 1 }}</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "memory",

  data() {
    return {
      tripIds: [],
    };
  },

  created() {
    this.getTripId();
  },
  methods: {
    //get trip_id
    getTripId: async function () {
      const id = { user_id: this.$store.getters["auth/getUserID"] };
      const header = { "X-Token": this.$store.getters["auth/getToken"] };

      await axios
        .get("api/auth/user/get_trip", { params: id, headers: header })
        .then((res) => {
          this.tripIds = res.data.trip_id;
          //this.showTripMemory(this.tripIds);
        });
    },

    onClickMyMap: function (value) {
      this.$store.commit("trip/setTripID", value);
      console.log(this.$store.getters["trip/getTripID"]);
      this.$router.push("/map");
    },
  },
};
</script>
<style scoped>
h1 {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 40px 0 0;
  color: #42b983;
}
.to_input {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  background-color: #42b983;
  width: 200px;
  height: 40px;
  border-radius: 8px;
  line-height: 40px;
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