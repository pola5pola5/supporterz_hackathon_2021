<template>
  <div class="TripMemory">
    <div id="routing"></div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "memory",  

  data(){
    return{
      tripId: [],
    }
  },

  mounted(){
    this.getTripId();
  },
  computed(){
    this.showTripMemory(this.tripId);
  },
  methods: {
    //get trip_id
    getTripId: async function(){
      const id = {user_id: "090ae791-9f53-4a84-ad61-0e84aee08634"}
      const header = {"X-Token": "0a5daf57-1fa8-452e-99d9-426718cabe8b"}

      await axios
        .get("api/auth/user/get_trip",{params: id, headers: header})
        .then((res) => {
          (this.tripId = res.data)
          this.showTripMemory(this.tripId);
        });
    },

    showTripMemory: function (tripId){
      console.log(tripId);
    },
  }
};
</script>

