<template>
  <div class="TripMemory">
    <div id="routing">
      <div v-for="(value, index) in tripIds" :key="value.id">
        <button type="button" @click="onClickMyMap(value)">
          旅{{ index + 1 }}
        </button>
        <div>&nbsp;</div>
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
      // console.log(this.$store.getters["trip/getTripID"]);
      this.$router.push("/map");
    },
  },
};
</script>
