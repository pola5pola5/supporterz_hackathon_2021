<template>
    <div class='MyMap'>
        <div id='map'></div>
    </div>
</template>

<script>
    import mapboxgl from 'mapbox-gl';
    import axios from 'axios';


    export default {
        name:'MyMap',
        
        data(){
            return{
                geojsonData: [],
            }
        },
        // created:function(){
        //     this.init();
        // },
        mounted:function(){
            this.init();
        },
        computed: function(){
            this.mapCreate(this.geojsonData);
        },
        methods:{
            init: async function(){
                // json取得
                await axios.get("http://13.112.197.183:1323/api/trip/get?trip_id=" + this.$store.state.tripid)
                .then(res => {this.geojsonData = res.data, this.mapCreate(this.geojsonData)});
            },

            mapCreate:function(geojsonData){
                console.log(geojsonData);
                mapboxgl.accessToken = 'pk.eyJ1IjoidHBrdW1hIiwiYSI6ImNrb3gzbGE5aDBhZ2cyd28xb3R5cG1jZXIifQ.jI7aje2MHl9teidoNmYDPA';
                const map = new mapboxgl.Map({
                    container: 'map',
                    style: 'mapbox://styles/mapbox/streets-v11',
                    center: geojsonData.features[0].geometry.coordinates,
                    //center: [-122.486958, 37.82931],
                    zoom: 15
                });
                map.addControl(new mapboxgl.NavigationControl());
                
                map.on('load', function () {
                    map.addSource('route', {
                        'type': 'geojson',
                        'data': {
                            'type': 'Feature',
                            'properties': {},
                            'geometry': {
                                'type': 'LineString',
                                'coordinates': [
                                    geojsonData.features[0].geometry.coordinates,
                                    [-122.483696, 37.833818],
                                    [-122.483482, 37.833174],
                                    [-122.483396, 37.8327],
                                    [-122.483568, 37.832056],
                                    [-122.48404, 37.831141],
                                    [-122.48404, 37.830497],
                                    [-122.483482, 37.82992],
                                    [-122.483568, 37.829548],
                                    [-122.48507, 37.829446],
                                    [-122.4861, 37.828802],
                                    [-122.486958, 37.82931],
                                    [-122.487001, 37.830802],
                                    [-122.487516, 37.831683],
                                    [-122.488031, 37.832158],
                                    [-122.488889, 37.832971],
                                    [-122.489876, 37.832632],
                                    [-122.490434, 37.832937],
                                    [-122.49125, 37.832429],
                                    [-122.491636, 37.832564],
                                    [-122.492237, 37.833378],
                                    [-122.493782, 37.833683]
                                ]
                            }
                        }
                    });
                    map.addLayer({
                        'id': 'route',
                        'type': 'line',
                        'source': 'route',
                        'layout': {
                            'line-join': 'round',
                            'line-cap': 'round'
                        },
                        'paint': {
                            'line-color': '#888',
                            'line-width': 8
                        }
                    });
                });

                var geojson = {
                    'type': 'FeatureCollection',
                    'features': [
                        {
                            'type': 'Feature',
                            'properties': {
                            'message': 'Foo',
                            'iconSize': [60, 60]
                            },
                            'geometry': {
                            'type': 'Point',
                            'coordinates': [-122.483696, 37.833818]
                            }
                        },
                        {
                            'type': 'Feature',
                            'properties': {
                            'message': 'Bar',
                            'iconSize': [50, 50]
                            },
                            'geometry': {
                            'type': 'Point',
                            'coordinates': [-122.48404, 37.831141]
                            }
                        },
                        {
                            'type': 'Feature',
                            'properties': {
                            'message': 'Baz',
                            'iconSize': [40, 40]
                            },
                            'geometry': {
                            'type': 'Point',
                            'coordinates': [-122.487516, 37.831683]
                            }
                        }
                    ]
                };
                map.on('mouseenter', 'places', function () {
                    map.getCanvas().style.cursor = 'pointer';
                });
                geojson.features.forEach(function(marker) {
                    // create a DOM element for the marker
                    var el = document.createElement('div');
                    el.className = 'marker';
                    
                    // el.addEventListener('click', function() {
                    // window.alert(marker.properties.message);
                    // });
                    
                    new mapboxgl.Marker(el)
                    .setLngLat(marker.geometry.coordinates)
                    .addTo(map);

                    new mapboxgl.Popup(el)
                    .setLngLat(marker.geometry.coordinates)
                    .addTo(map);
                });


            }
        }
    }
</script>

<style>
    /*マップサイズ*/
    #map {
        z-index: 0;
        height: 800px;
    }
    .marker {
        background-image: url('../assets/mapbox-icon.png');
        background-size: cover;
        width: 50px;
        height: 50px;
        border-radius: 50%;
        cursor: pointer;
    }
    .mapboxgl-popup {
        max-width: 200px;
    }
</style>