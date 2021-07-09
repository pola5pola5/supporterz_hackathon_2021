<template>
  <div class="Map">
    <div id="map"></div>
  </div>
</template>

<script>
import mapboxgl from "mapbox-gl";
import axios from "axios";

export default {
  data() {
    return {
      geojsonData: [],
      mapData: [],
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
    //get json
    getGeojson: async function () {
      const id = { trip_id: this.$store.state.tripid };
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
      });
      await axios
        .get(
          "https://api.mapbox.com/directions/v5/mapbox/driving/" +
            jsonCoordinates +
            "?access_token=pk.eyJ1IjoidHBrdW1hIiwiYSI6ImNrb3gzbGE5aDBhZ2cyd28xb3R5cG1jZXIifQ.jI7aje2MHl9teidoNmYDPA&depart_at=2019-05-02T15:00&overview=full&geometries=geojson"
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
      mapboxgl.accessToken =
        "pk.eyJ1IjoidHBrdW1hIiwiYSI6ImNrb3gzbGE5aDBhZ2cyd28xb3R5cG1jZXIifQ.jI7aje2MHl9teidoNmYDPA";
      const map = new mapboxgl.Map({
        container: "map",
        style: "mapbox://styles/mapbox/streets-v11",
        center: route[0],
        zoom: 15,
      });

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

      map.on("mouseenter", "places", function () {
        map.getCanvas().style.cursor = "pointer";
      });

      //make
      geojsonData.features.forEach(function (marker) {
        // create a DOM element for the marker
        var el = document.createElement("div");
        el.className = "marker";

        var pop = document.createElement("div");
        pop.className = "img";
        pop.style.background = "url(" + marker.properties.img_url + ")";
        pop.style.width = 300 + "px";
        pop.style.height = 300 + "px";
        pop.style.backgroundSize = "cover";
        console.log(marker.properties.img_url);

        const popup = new mapboxgl.Popup({
          offset: 25,
          maxWidth: 1000,
        }).setDOMContent(pop);

        new mapboxgl.Marker(el)
          .setLngLat(marker.geometry.coordinates)
          .setPopup(popup)
          .addTo(map);
      });
    },
  },
};
</script>

<style>
/*マップサイズ*/
#map {
  z-index: 0;
  height: 800px;
}
.marker {
  background-image: url("../assets/mapbox-icon.png");
  background-size: cover;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  cursor: pointer;
}
</style>
