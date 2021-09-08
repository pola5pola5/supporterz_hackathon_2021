<template>
  <div class="HomeLink">
    <!-- <br />
    <b>こんにちは、{{ getUserName }}さん</b>
    <br /> -->

    <!-- <router-link to="/">
      <h3 class="to_top">Homeに戻る</h3>
    </router-link> -->
    <!-- <router-link to="/input">
      <h3 class="to_input">旅の追加</h3>
    </router-link>
    <router-link to="/routing">
      <h3 class="to_top">旅の一覧</h3>
    </router-link> -->

    <!-- <router-link to='/signin'>
      <h3 class="to_signin">早速使ってみる</h3>
    </router-link>    
    <router-link to='/signup'>
      <h3 class="to_signup">初めての方はこちら</h3>
    </router-link> -->

    <!-- <h4>チェック用リンク</h4>
    <ul>
      <li>
        <a href="/getpost" target="_blank" rel="noopener">GetPost</a>
      </li>
      <li>
        <a href="/upload" target="_blank" rel="noopener">upload</a>
      </li>
      <li>
        <a href="/input" target="_blank" rel="noopener">input</a
        >
      </li>
    </ul> -->
    <div class="header">
      <div class="headerTitle" v-on:click="onClickTitle()">フォト旅</div>
      <div class="border"></div>
      <div class="username">Hello {{username}}!</div>
      <div class="addtrip" v-on:click="onClickAddtrip()">Add trip</div>
      <div class="userSetting" v-on:click="onClickOpenPopup()">{{username.slice(0,1).toUpperCase()}}</div>
      <Popup :isopen='popup' :user='this.username' :tripnum='this.tripIds.length' @close="onClickClosePopup"></Popup>
    </div>

    <div class="title">Trip List</div>
    <div class="container">
      <div v-for="(value, index) in tripIds" :key="value.id" v-on:click="onClickMyMap(value)" class="tripParent" >
        <div class="name_bg">
          <div class="name">旅{{index + 1}}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// export default {
//   name: "HomeLink",
//   props: {
//     // msg: String,
//   },
//   computed: {
//     getUserName: function () {
//       return this.$store.getters["user/getUserName"];
//     },
//   },
// };
import axios from "axios";
import mapboxgl from "mapbox-gl";
import Popup from "@/components/UserPopup.vue";

export default {
  components: { Popup },
  name: "memory",

  data() {
    return {
      tripIds: [],
      geojsonData: [],
      mapData: [],
      username: "hoge",
      popup: false,
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

    onClickOpenPopup: function () {
      this.popup = true
    },

    onClickClosePopup: function () {
      this.popup = false
    },

    onClickOpenModal: function () {
      this.modal = true;
    },

    getTripData: async function (tripID, idx) {
      const id = { trip_id: tripID };
      const header = { "X-Token": this.$store.getters["auth/getToken"] };

      await axios
        .get("/api/auth/trip/get", { params: id, headers: header })
        .then((res) => {
          this.geojsonData[idx] = res.data;
          this.getMap(idx);
        });
    },

    getMap: async function (idx) {
      var jsonCoordinates = [];
      this.geojsonData[idx].features.forEach(function (jsonData, index, array) {
        jsonCoordinates =
          jsonCoordinates +
          jsonData.geometry.coordinates[0] +
          "," +
          jsonData.geometry.coordinates[1];
        if (index < array.length - 1) {
          jsonCoordinates = jsonCoordinates + ";";
        }
        if (index === array.length - 1) {
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
          this.mapData[idx] = res.data;
          this.createMap(this.mapData[idx], this.geojsonData[idx], idx);
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
        photoel.src = photo
        photoel.style.width = 5 + "rem"
        photoel.style.height = 5 + "rem"
        tripParent.appendChild(photoel)
      }
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* h1 {
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
  width: 150px;
  height: 30px;
  border-radius: 8px;

  margin-left: auto;
  margin-right: auto;
}
.to_top {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  background-color: #42b983;
  width: 150px;
  height: 30px;
  border-radius: 8px;

  margin-left: auto;
  margin-right: auto;
}
.to_signin {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  background-color: #42b983;
  width: 150px;
  height: 30px;
  border-radius: 8px;

  margin-left: auto;
  margin-right: auto;
}
.to_signup {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: relative;
  color: white;
  background-color: #42b983;
  width: 180px;
  height: 30px;
  border-radius: 8px;

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
} */
.header{
    height: 100px;
    width: 100%;
    background-image: url("~@/assets/header.jpg");
    background-size: cover;
    background-position: center center;
    display: flex;
    align-items: center;
  }

  .headerTitle{
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

  .userSetting{
    color: white;
    margin-left: 40px;
    margin-right: 20px;
    cursor: pointer;
    background-color: #C850BC;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    text-align: center;
    line-height: 40px;
  }

  .container{
    display: flex;
    flex-wrap: wrap;
  }

  .name{
    margin-left: 5px;
  }

  .name_bg{
    background-color: #060B38;
  }

  .tripParent{
    color: white;
    width: calc(50% - 2em);
    height: 303px;
    margin: 1em;
    cursor: pointer;
    box-shadow: 5px 2.5px 2.5px gray;
  }

  *{
    margin: 0%;
    padding: 0%;
  }
</style>
