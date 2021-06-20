<template>
  <div class="routing">
      <h1>Routing Test Page</h1>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Routing',
  //created: functionはなくて良い
  created:function(){
    this.init();
  },
  methods:{
    // 全体的な流れ
    // 1. axios.postでデータ送信
    // 2. 帰ってきたjsonからtrip_idを抜き取ると同時にstatusコード確認
    // 3. trip_idをstoreに格納
    // 4. statusコード確認からページ遷移
    init: async function(){
      let resStatus
      let getRequest
      //json取得 **画像送信のPOSTに変更**
      await axios.get("http://13.112.197.183:1323/api/trip/get?trip_id=" + this.$store.state.tripid)
      .then(res => {
        resStatus = res.status, 
        getRequest = res.data.features[0].properties.trip_id});

      if(resStatus === 200){
        console.log(this.$store.state.tripid);
        console.log(getRequest);
        this.$store.commit('pushid', getRequest);
        console.log(this.$store.state.tripid);
        this.$router.push("/about");
      }
    }
  }
}
</script>