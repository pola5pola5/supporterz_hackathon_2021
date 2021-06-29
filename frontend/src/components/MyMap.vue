<template>
  <div class="MyMap">
    <div id="map"></div>
  </div>
</template>

<script>
import mapboxgl from "mapbox-gl";
import axios from "axios";

export default {
  name: "MyMap",

  data() {
    return {
      geojsonData: [],
      geocoordinate: [],
    };
  },
  mounted: function () {
    this.init();
  },
  computed: function () {
    this.mapCreate(this.geojsonData, this.geocoordinate);
  },
  methods: {
    //get json
    init: async function () {
      await axios
        .get("/api/trip/get?trip_id=" + this.$store.state.tripid)
        .then((res) => {
          (this.geojsonData = res.data),
            this.mapCreate(this.geojsonData, this.geocoordinate);
        });
    },

    //create map
    mapCreate: function (geojsonData, geocoordinate) {
      console.log(geojsonData);
      mapboxgl.accessToken =
        "pk.eyJ1IjoidHBrdW1hIiwiYSI6ImNrb3gzbGE5aDBhZ2cyd28xb3R5cG1jZXIifQ.jI7aje2MHl9teidoNmYDPA";
      //initialize
      const map = new mapboxgl.Map({
        container: "map",
        style: "mapbox://styles/mapbox/streets-v11",
        center: geojsonData.features[0].geometry.coordinates,
        zoom: 15,
      });
      map.addControl(new mapboxgl.NavigationControl());

      //make line coordinate
      geojsonData.features.forEach(function (json) {
        geocoordinate.push(json.geometry.coordinates);
      });

      //make line
      map.on("load", function () {
        map.addSource("route", {
          type: "geojson",
          data: {
            type: "Feature",
            properties: {},
            geometry: {
              type: "LineString",
              coordinates: geocoordinate,
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
