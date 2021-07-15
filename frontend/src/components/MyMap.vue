<template>
  <div class="Map">
    <div id="map"></div>
    <div class="overlay">
       <button id="replay">Replay</button>
    </div>
  </div>
</template>

<script>
import mapboxgl from "mapbox-gl";
import axios from "axios";
//import turf from "@turf/turf"

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
        if(idx === array.length - 1){
          jsonCoordinates = jsonCoordinates + ";" 
          + jsonData.geometry.coordinates[0] + "," 
          + jsonData.geometry.coordinates[1];
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
        function animate() {
          var start =
            route[
              counter >= steps ? counter - 1 : counter
            ];
          var end =
            route[
              counter >= steps ? counter : counter + 1
            ];
          if (!start || !end) return;
          
          point.features[0].geometry.coordinates =
          route[counter];
          
          // Calculate the bearing to ensure the icon is rotated to match the route arc
          // The bearing is calculated between the current point and the next point, except
          // at the end of the arc, which uses the previous point and the current point
          // point.features[0].properties.bearing = turf.bearing(
          //   turf.point(start),
          //   turf.point(end)
          // );
          
          map.getSource('point').setData(point);
          
          if (counter < steps) {
            requestAnimationFrame(animate);
          }
          
          counter = counter + 1;
        }
        document
          .getElementById('replay')
          .addEventListener('click', function () {
            // Set the coordinates of the original point back to origin
            point.features[0].geometry.coordinates = origin;
            
            // Update the source layer
            map.getSource('point').setData(point);
            
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
        el.className = "marker";

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
            anchor: 'bottom'
          })
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
  background-image: url("../assets/marker.jpg");
  background-size: cover;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  cursor: pointer;
}
.overlay {
  position: absolute;
  top: 100px;
  left: 30px;
}
.overlay button {
  font: 600 12px/20px 'Helvetica Neue', Arial, Helvetica, sans-serif;
  background-color: #3386c0;
  color: #fff;
  display: inline-block;
  margin: 0;
  padding: 10px 20px;
  border: none;
  cursor: pointer;
  border-radius: 3px;
}
 
.overlay button:hover {
  background-color: #4ea0da;
}
</style>
