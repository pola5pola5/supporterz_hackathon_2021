<template>
  <div class="TripMemory">
    <div id="routing">
      <div class="header"></div>
      <div class="title">Trip List</div>

      <div id="container">
        <div v-for="(value, index) in tripIds" :key="value.id" @click="onClickMyMap(value)" class="tripParent" >
          <div class="name">旅{{index + 1}}</div>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import axios from "axios";
import mapboxgl from "mapbox-gl";

export default {
  name: "memory",

  data() {
    return {
      tripIds: [],
      geojsonData: [],
      mapData: [],
    };
  },

  mounted() {
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
      
      var tripParent = document.getElementsByClassName("tripParent")[idx]
      var el = document.createElement("div")
      el.id = "map" + idx
      el.textContent = "loading"
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
        map.addLayer({
          'id': 'point',
          'source': 'point',
          'type': 'symbol',
          'layout': {
            'icon-image': 'car-15',
            'icon-rotate': ['get', 'bearing'],
            'icon-rotation-alignment': 'map',
            'icon-allow-overlap': true,
            'icon-ignore-placement': true,
            'icon-size': 2.5
          }
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
    },
  },
};
</script>

<style>
  .header{
    height: 120px;
    background-image: url("~@/assets/header.jpg");
  }

  .title{
    font-family: "Times New Roman";
    font-size: 40px;
    margin-top: 20px;
    margin-bottom: 30px;
    text-align: center;
  }

  .container{
    display: flex;
  }

  .tripParent{
    background-color: blue;
    width: 50%;
  }

  *{
    margin: 0%;
    padding: 0%;
  }
</style>

