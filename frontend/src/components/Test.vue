<template>
    {{ answer }}
    <img :src="img_base64_content" alt="" />
</template>

<script>
import axios from "axios";

export default {
    data() {
        return {
            img_base64_content: "",
            answer: "まだ送られてないです",
        };
    },
    created: function() {
        this.getImgData();
    },
    
    methods: {
        
        getImgData: async function() {
            this.answer = "送受信中"
            await axios({
                    method: 'get',
                    url: 'http://13.231.129.55/api/auth/trip/get_img', 
                    params: { "img_path": "img1.jpg" },
                    headers: { "X-Token": "4f272392-a679-4a86-ba92-3ffd96c83ae8"},
                })
                .then(response => {
                    this.img_base64_content = "data:image/jpg;base64," + response.data;//string
                    this.answer = "帰ってきました"
                })
                .catch(error => {
                    this.answer = "error!" + error
                });
        },
    }
    
}

</script>


