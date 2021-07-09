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

  methods: {
    //get trip_id
    getTripId: async function(){
      const id = {user_id: "c8fed8a1-7e15-4516-838a-14a6bc1f703f"}
      const header = {"X-Token": "4f272392-a679-4a86-ba92-3ffd96c83ae8"}

      await axios
        .get("api/auth/user/get_trip",{params: id, headers: header})
        .then((res) => {
          (this.tripIds = res.data.trip_id)
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

