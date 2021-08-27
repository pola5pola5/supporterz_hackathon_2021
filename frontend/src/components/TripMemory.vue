<template>
  <div class="TripMemory">
    <div id="routing">
      <div class="header">
        <div class="header_title" @click="onClickTitle()">フォト旅</div>
        <div class="border"></div>
        <div class="username">Hello {{username}}!</div>
        <div class="addtrip" @click="onClickAddtrip()">Add trip</div>
        <div class="user_setting">hoge</div>

        <div class="example-modal-window">
          <p>ボタンを押すとモーダルウィンドウが開きます</p>
          <button @click="openModal">開く</button>

          <!-- コンポーネント MyModal -->
          <MyModal @close="closeModal" v-if="modal">
            <!-- default スロットコンテンツ -->
            <p>Vue.js Modal Window!</p>
            <div><input v-model="message"></div>
            <!-- /default -->
          </MyModal>
        </div>
      </div>

      <div class="title">Trip List</div>
      <div class="container">
        <div v-for="(value, index) in tripIds" :key="value.id" @click="onClickMyMap(value)" class="tripParent" >
          <div class="name_bg">
            <div class="name">旅{{index + 1}}</div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import axios from "axios";
import mapboxgl from "mapbox-gl";
import MyModal from './MyModal.vue'

export default {
  components: { MyModal },
  name: "memory",

  data() {
    return {
      tripIds: [],
      geojsonData: [],
      mapData: [],
      username: [],
      modal: false,
      message: ''
    };
  },

  mounted() {
    this.getTripId();
  },

  methods: {

    //get trip_id
    getTripId: async function () {
      this.username = this.$store.getters["user/getUserName"];
      const id = { user_id: this.$store.getters["auth/getUserID"] };
      const header = { "X-Token": this.$store.getters["auth/getToken"] };

      await axios
        .get("api/auth/user/get_trip", { params: id, headers: header })
        .then((res) => {
          this.tripIds = res.data.trip_id;
          this.makeEachMap()
        });
    },

    makeEachMap: function (){
      let self = this
      for(const [idx, tripID] of this.tripIds.entries()){
        self.getTripData(tripID, idx)
      }
    },

    onClickMyMap: function (value) {
      this.$store.commit("trip/setTripID", value);
      this.$router.push("/map");
    },

    onClickTitle: function () {
      this.$router.push("/routing");
    },

    onClickAddtrip: function () {
      this.$router.push("/input");
    },

    openModal() {
      this.modal = true
    },
    closeModal() {
      this.modal = false
    },
    doSend() {
      if (this.message.length > 0) {
        alert(this.message)
        this.message = ''
        this.closeModal()
      } else {
        alert('メッセージを入力してください')
      }
    },

    getTripData: async function (tripID, idx) {
      const id = { trip_id: tripID };
      const header = { "X-Token": this.$store.getters["auth/getToken"] };

      await axios
        .get("/api/auth/trip/get", { params: id, headers: header })
        .then((res) => {
          this.geojsonData = res.data;
          this.getMap(idx);
        });
    },

    getMap: async function (idx) {
      var jsonCoordinates = [];
      this.geojsonData.features.forEach(function (jsonData, idx, array) {
        jsonCoordinates =
          jsonCoordinates +
          jsonData.geometry.coordinates[0] +
          "," +
          jsonData.geometry.coordinates[1];
        if (idx < array.length - 1) {
          jsonCoordinates = jsonCoordinates + ";";
        }
        if (idx === array.length - 1) {
          jsonCoordinates =
            jsonCoordinates +
            ";" +
            jsonData.geometry.coordinates[0] +
            "," +
            jsonData.geometry.coordinates[1];
        }
      });
      await axios
        .get(
          "https://api.mapbox.com/directions/v5/mapbox/driving/" +
            jsonCoordinates +
            "?access_token=" +
            process.env.VUE_APP_MAPBOX_API_KEY +
            "&depart_at=2019-05-02T15:00&overview=full&geometries=geojson"
        )
        .then((res) => {
          this.mapData = res.data;
          this.createMap(this.mapData, this.geojsonData, idx);
        });
    },

    createMap: function (mapData, geojsonData, idx) {
      const data = mapData.routes[0];
      var route = data.geometry.coordinates;
      const photos = geojsonData.features

      var tripParent = document.getElementsByClassName("tripParent")[idx]
      var el = document.createElement("div")
      el.id = "map" + idx
      el.textContent = "loading"
      el.style.width= 100 + "%";
      el.style.height = 200 + "px"
      tripParent.appendChild(el)

      //cretate map
      mapboxgl.accessToken = process.env.VUE_APP_MAPBOX_API_KEY;
      const map = new mapboxgl.Map({
        container: el.id,
        style: "mapbox://styles/tpkuma/ckr1c20cv1c4f18qcbsrr2gmm",
        center: route[0],
        zoom: 15,
      });

      var point = {
        'type': 'FeatureCollection',
        'features': [
          {
            'type': 'Feature',
            'properties': {},
            'geometry': {
              'type': 'Point',
              'coordinates': route[0]
            }
          }
        ]
      };

      map.on("load", function () {
        map.addSource("route", {
          type: "geojson",
          data: {
            type: "Feature",
            properties: {},
            geometry: {
              type: "LineString",
              coordinates: route,
            },
          },
        });
        map.addSource("point", {
          type: "geojson",
          data: point
        });
        map.addLayer({
          id: "route",
          type: "line",
          source: "route",
          layout: {
            "line-join": "round",
            "line-cap": "round",
          },
          paint: {
            "line-color": "#888",
            "line-width": 8,
          },
        });
      });

      geojsonData.features.forEach(function (marker) {
        var el = document.createElement("div");
        el.className = "marker";

        new mapboxgl.Marker({
            element: el,
            anchor: 'bottom'
          })
          .setLngLat(marker.geometry.coordinates)
          .addTo(map);
      });

      const len = photos.length
      for(let i = 0; i < len; i++){
        if(i > 5) break
        const photo = photos[i].properties.img_url
        var photoel = document.createElement("img")
        photoel.src =  photo
        photoel.style.width = 80 + "px"
        photoel.style.height = 80 + "px"
        tripParent.appendChild(photoel)
      }
    },
  },
};
</script>

<style>
  .header{
    height: 100px;
    width: 100%;
    background-image: url("~@/assets/header.jpg");
    background-size: cover;
    background-position: center center;
    display: flex;
    align-items: center;
  }

  .header_title{
    width: 170px;
    font-family: serif;
    font-size: 30px;
    cursor: pointer;
    color: white;
    display:table-cell;
    vertical-align:middle;
    text-align: center;
  }

  .border{
    border-left: solid #C4C4C4;
    padding-left: 10px;
    height: 70px;
  }

  .username{
    font-family: serif;
    color: white;
    font-size: 25px;
  }

  .title{
    font-family: "Times New Roman";
    font-size: 40px;
    margin-top: 20px;
    margin-bottom: 20px;
    text-align: center;
    color: #060B38;
  }

  .addtrip{
    font-family: serif;
    color: white;
    cursor: pointer;
    font-size: 20px;
    margin-left: auto;
    background-color: #52A7F4;
    padding: 8px;
    border-radius: 10px;
  }

  .user_setting{
    color: white;
    margin-left: 40px;
    margin-right: 20px;
  }

  #overlay{
    z-index:1;
    position:fixed;
    top:0;
    left:0;
    width:100%;
    height:100%;
    background-color:rgba(0,0,0,0.5);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  #content{
    z-index:2;
    width:50%;
    padding: 1em;
    background:#fff;
  }

  .container{
    display: flex;
  }

  .name{
    margin-left: 5px;
  }

  .name_bg{
    background-color: #060B38;
  }

  .tripParent{
    color: white;
    width: 50%;
    height: 303px;
    margin: 1em;
    cursor: pointer;
    box-shadow: 5px 2.5px 2.5px gray;
    /* background-color: #f1f0f0; */
  }

  *{
    margin: 0%;
    padding: 0%;
  }
</style>

