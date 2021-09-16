<template>
  <div class="Map">
    <div id="map"></div>
    <div class="overlay">
      <div class="map_header">
        <div class="title">
          <div class="headerTitle" v-on:click="onClickTitle">フォト旅</div>
        </div>
        <div class="menu">
          <div class="toHome" v-on:click="onClickTitle">Home</div>
          <div class="addtrip" v-on:click="onClickAddtrip">Add trip</div>
          <div class="userSetting" v-on:click="onClickOpenPopup">{{ username_upper }}</div>
        </div>
        <Popup :isopen='popup' :user='username' :tripnum='tripIdNum' @close="onClickClosePopup"></Popup>
      </div>
      <button id="replay">Replay</button>
    </div>
  </div>
</template>

<script>
import mapboxgl from "mapbox-gl";
import axios from "axios";
import Popup from "@/components/UserPopup.vue";
//import turf from "@turf/turf"

export default {
  components: {Popup} ,
  name: "Mymap",

  data() {
    return {
      geojsonData: [],
      mapData: [],
      tripIdNum: "",
      popup: false,
      username: "undefined user",
      username_upper: "A"
    };
  },

  mounted: function () {
    this.getGeojson();
  },

  computed: function () {
    this.mapCreate(this.mapData, this.geojsonData);
    this.getMapApi();
  },

  methods: {
    onClickTitle: function () {
      this.$router.push("/home");
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

    //get json
    getGeojson: async function () {
      this.tripIdNum = this.$store.getters["trip/getNumTripID"];
      this.username = this.$store.getters["user/getUserName"];
      this.username_upper = this.username.slice(0,1).toUpperCase()
      const id = { trip_id: this.$store.getters["trip/getTripID"] };
      const header = { "X-Token": this.$store.getters["auth/getToken"] };

      await axios
        .get("/api/auth/trip/get", { params: id, headers: header })
        .then((res) => {
          this.geojsonData = res.data;
          this.getMapApi();
        });
    },

    getMapApi: async function () {
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
          this.mapCreate(this.mapData, this.geojsonData);
        });
    },

    mapCreate: function (mapData, geojsonData) {
      const data = mapData.routes[0];
      var route = data.geometry.coordinates;

      //cretate map
      mapboxgl.accessToken = process.env.VUE_APP_MAPBOX_API_KEY;
      const map = new mapboxgl.Map({
        container: "map",
        style: "mapbox://styles/tpkuma/ckr1c20cv1c4f18qcbsrr2gmm",
        center: route[0],
        zoom: 15,
      });

      var point = {
        type: "FeatureCollection",
        features: [
          {
            type: "Feature",
            properties: {},
            geometry: {
              type: "Point",
              coordinates: route[0],
            },
          },
        ],
      };

      var counter = 0;
      var steps = route.length;

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
          data: point,
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
          id: "point",
          source: "point",
          type: "symbol",
          layout: {
            "icon-image": "car-15",
            "icon-rotate": ["get", "bearing"],
            "icon-rotation-alignment": "map",
            "icon-allow-overlap": true,
            "icon-ignore-placement": true,
            "icon-size": 2.5,
          },
        });
        function animate() {
          var start = route[counter >= steps ? counter - 1 : counter];
          var end = route[counter >= steps ? counter : counter + 1];
          if (!start || !end) return;

          point.features[0].geometry.coordinates = route[counter];

          // Calculate the bearing to ensure the icon is rotated to match the route arc
          // The bearing is calculated between the current point and the next point, except
          // at the end of the arc, which uses the previous point and the current point
          // point.features[0].properties.bearing = turf.bearing(
          //   turf.point(start),
          //   turf.point(end)
          // );

          map.getSource("point").setData(point);

          if (counter < steps) {
            requestAnimationFrame(animate);
          }

          counter = counter + 1;
        }
        document
          .getElementById("replay")
          .addEventListener("click", function () {
            // Set the coordinates of the original point back to origin
            point.features[0].geometry.coordinates = origin;

            // Update the source layer
            map.getSource("point").setData(point);

            // Reset the counter
            counter = 0;

            // Restart the animation
            animate(counter);
          });

        animate(counter);
      });

      map.on("mouseenter", "places", function () {
        map.getCanvas().style.cursor = "pointer";
      });

      //make
      geojsonData.features.forEach(function (marker) {
        // create a DOM element for the marker
        var el = document.createElement("div");
        el.className = "map_marker";

        var pop = document.createElement("div");
        pop.className = "img";
        pop.style.background = "url(" + marker.properties.img_url + ")";
        pop.style.width = 300 + "px";
        pop.style.height = 300 + "px";
        pop.style.backgroundSize = "cover";

        const popup = new mapboxgl.Popup({
          offset: 25,
          maxWidth: 1000,
        }).setDOMContent(pop);

        new mapboxgl.Marker({
          element: el,
          anchor: "bottom",
        })
          .setLngLat(marker.geometry.coordinates)
          .setPopup(popup)
          .addTo(map);
      });
    },
  },
};
</script>

<style scoped>
/*マップサイズ*/
  #map {
    z-index: 0;
    height: 800px;
  }

  .map_header{
    height: 50px;
    width: 100%;
    background-color: black;
    background-size: cover;
    background-position: center center;
    display: flex;
    align-items: center;
    opacity: 0.67;
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

  .title{
    margin-right: auto;
  }

  .menu{
    display: flex;
    justify-content: space-around;
    align-items: center;
    margin-left: 50%;
  }

  .toHome{
    font-family: serif;
    color: white;
    cursor: pointer;
    font-size: 20px;
    text-align: center;
  }

  .addtrip{
    font-family: serif;
    color: white;
    margin-left: 40px;
    cursor: pointer;
    font-size: 20px;
    /* margin-left: auto; */
    border-color: white;
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

  .overlay {
    position: absolute;
    top: 0px;
    left: 0px;
    width: 100%;
  }
  .overlay button {
    font: 600 12px/20px 'Helvetica Neue', Arial, Helvetica, sans-serif;
    background-color: #F45252;
    color: #fff;
    display: inline-block;
    margin-left: 30px;
    margin-top: 20px;
    padding: 10px 20px;
    border: none;
    cursor: pointer;
    border-radius: 3px;
    z-index: 1;
  }
  
  .overlay button:hover {
    background-color: #e94040;
  }
</style>

<style>
  .map_marker {
    background-image: url("../assets/marker.jpg");
    background-size: cover;
    width: 50px;
    height: 50px;
    border-radius: 50%;
    cursor: pointer;
  }
</style>
