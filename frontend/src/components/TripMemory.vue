<template>
  <div class="TripMemory">
    <div id="routing">
      <div v-for="(value) in tripIds" :key="value.id">
        <button type="button" @click="onClickMyMap(value)">{{value}}</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "memory",  

  data(){
    return{
      tripIds: []
    }
  },

  created(){
    this.getTripId();
  },
  computed(){
    //this.onClickMyMap(value);
  },
  methods: {
    //get trip_id
    getTripId: async function(){
      const id = {user_id: "090ae791-9f53-4a84-ad61-0e84aee08634"}
      const header = {"X-Token": "0a5daf57-1fa8-452e-99d9-426718cabe8b"}

      await axios
        .get("api/auth/user/get_trip",{params: id, headers: header})
        .then((res) => {
          (this.tripIds = res.data.trip_id)
          //this.showTripMemory(this.tripIds);
        });
    },

    onClickMyMap: function(value){
      this.$store.commit("pushid", value);
      console.log(this.$store.state.tripid)
      this.$router.push("/map");
    }

  }
};
</script>

