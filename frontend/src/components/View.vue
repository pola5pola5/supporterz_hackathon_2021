<template>
  <div class="View">
    <!-- <input type="file" name="example" ref="preview" accept="image/*" multiple required> -->
  <p id="error" v-show="error">{{ error }}</p>
    <p>クリックまたはドラッグ&ドロップで画像を追加してください．</p>
  <label>
    <!-- <img :src="avatar" alt="Avatar" class="image"> -->
    <div 
      class="drop_area" 
      @dragenter="dragEnter"
      @dragleave="dragLeave"
      @dragover.prevent
      @drop.prevent="dropFile()"
      :class="{enter: isEnter}"
      @change="onImageChange"

      >

      ファイルアップロード
    <!-- <div> -->
    <input
           type="file"
           id="avatar_name"
           accept="image/jpeg, image/png"
           @change="onImageChange"
           multiple
           />
    </div>
  </label>
  {{ files }}
  {{ files.length }}
  {{ images }}
  <br>
  <div>
        <ul class="flex">
          <li class="flex-col" v-for="(file,index) in files" :key="index" @click="deleteFile(index)">
            {{ index }}
              <div style="position: relative;">
                <span class="delete-mark">×</span>
                <img class="file_icon" :src="images[index]">
              </div>
              <span>{{ file.name }}</span>
          </li>
        </ul>
  
  </div>
  <!-- <img :src="avatar" class="image"> -->
  <!-- <img :src="avatar" alt="Avatar" class="image"> -->
  <br>
  <div v-show="files.length">
    <button class="button" v-on:click="upload">送信</button>
    <!-- <button class="button" @click="changeBase64(images)">送信</button> -->
  </div>

  <p>{{ message }}</p>
  <p>{{ error }}</p>
  </div>
</template>

<script>

import axios from "axios"
export default {
  name: "View",
  data () {
    return {
   		avatar: '',
      message: '',
      error: '',
      isEnter: false,
      files: [],
      images: [],
      min_imgs: [],
      test: ''
    }
  },
  methods: {
    // setError (error, text) {
    //   this.error = (error.response && error.response.data && error.response.data.error) || text
    // },
    dragEnter() {
      this.isEnter = true;
    },
    dragLeave() {
      this.isEnter = false;
    },
    getBase64 (file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = () => resolve(reader.result)
        reader.onerror = error => reject(error)
      })
    },
    dropFile: function() {
      console.log(event.dataTransfer.files)
      // console.log(event.dataTransfer.files.length)
      this.files.push(...event.dataTransfer.files)
      this.isEnter = false;

      // それぞれのファイルに対して変換処理
      this.files.forEach(file => {
        console.log(file)
        let im = null

        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = function() {
          im = reader.result;
          console.log(im)
          this.images.push(im)

          // 小さくする処理いれる
          const resizedCanvas = this.createResizedCanvasElement(im)
          const resizedBase64 = resizedCanvas.toDataURL(file.type)
          this.min_imgs.push(resizedBase64)
        }.bind(this);
      })
      return this.files, this.images;
    },
    deleteFile(index) {
      this.files.splice(index, 1)
      this.images.splice(index, 1)
    },


    createResizedCanvasElement (originalImg) {
      const originalImgWidth = originalImg.width
      const orifinalImgHeight = originalImg.height

      // resizeWidthAndHeight関数については下記参照
      const [resizedWidth, resizedHeight] = this.resizeWidthAndHeight(originalImgWidth, orifinalImgHeight)
      const canvas = document.createElement('canvas')
      const ctx = canvas.getContext('2d')
      canvas.width = resizedWidth
      canvas.height = resizedHeight

      // drawImage関数の仕様はcanvasAPIのドキュメントを参照下さい
      ctx.drawImage(originalImg, 0, 0, resizedWidth, resizedHeight)
      return canvas
    },
    // 縦横の比率を変えず、定めた大きさを超えないWidthとHeightの値を割り出す関数
    resizeWidthAndHeight (width, height) {

      // 今回は400x400のサイズにしましたが、ここはプロジェクトによって柔軟に変更してよいと思います
      const MAX_WIDTH = 400
      const MAX_HEIGHT = 400

      // 縦と横の比率を保つ
      if (width > height) {
        if (width > MAX_WIDTH) {
          height *= MAX_WIDTH / width
          width = MAX_WIDTH
        }
      } else {
        if (height > MAX_HEIGHT) {
          width *= MAX_HEIGHT / height
          height = MAX_HEIGHT
        }
      }
      return [width, height]
    },

    upload: function() {
      console.log("test")
      console.log(axios.get("http://13.112.197.183:1323/api/setting/get").then(response => {
        console.log("test")
        console.log(response.data)
        }))
      console.log("unchi: ", this.images) // データは入ってた

      // パターン1(ダメだった)
      let params = new URLSearchParams();
      params.append('user_id', 'c8fed8a1-7e15-4516-838a-14a6bc1f703f');
      params.append('imgs', this.images);
      console.log(params.values())

      axios.post("http://13.112.197.183:1323/api/trip/save", 
        params
      ).then(response => {
          console.log(response.data)
          this.message = "アップロードしました"
          this.error = ""
        }).catch(error => {
            console.log("error")
            console.log(error)
        })

      // パターン2(昨日と同じやり方)
      // パターン1(ダメだった)
      axios.post("http://13.112.197.183:1323/api/trip/save", {
        user_id: "c8fed8a1-7e15-4516-838a-14a6bc1f703f",
        imgs: this.images
      }).then(response => {
          console.log(response.data)
          this.message = "アップロードしました"
          this.error = ""
        }).catch(error => {
            console.log("error")
            console.log(error)
        })

    },

    // changeBase64: function(images) {
    //   // let im = null;
    //   // console.log(files)
    //   console.log(images)
    //   // for (let step = 0; step < files.length; step++) {
    //   //   // console.log(files[step])
    //   //   const reader = new FileReader();
    //   //   reader.readAsDataURL(files[step]);
    //   //   // console.log(reader.result)
    //   //   reader.onload = function() {
    //   //     // console.log(reader.result)
    //   //     im = reader.result;
    //   //     // console.log(im)
    //   //     this.images.push(im)
    //   //     console.log(this.images)
    //   //   }.bind(this);

    //   axios.post("http://13.112.197.183:1323/api/trip/save", {
    //     user_id: "c8fed8a1-7e15-4516-838a-14a6bc1f703f",
    //     imgs: this.images
    //   }
    //   ).then(response => {
    //     console.log(response.data)
    //       this.message = "アップロードしました"
    //       this.error = ""
    //     }).catch(error => {
    //       console.log("error")
    //         console.log(error)
    //     });
    //       // }
    // }
  }

}
</script>
<style scoped>
input {
  display: none;
}

img:hover {
  opacity: 0.7;
}

#error {
  color: red;
}
label{
  display: flex;
  justify-content: center;
  align-items: center;
}
.drop_area {
  color: gray;
  font-weight: bold;
  font-size: 1.2em;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 500px;
  height: 300px;
  border: 5px solid gray;
  border-radius: 15px;
    }
.enter {
    border: 10px dotted powderblue;
}
ul {
    margin: 0;
    padding: 0;
    list-style-type: none;
}
.flex {
    display: flex;
    align-items: center;
}

.flex-col {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 0.5em;
    font-size: 10px;
}

.delete-mark {
    position: absolute;
    top: -14px;
    right: -10px;
    font-size: 20px;
}
span{
    max-width: 100px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.button {
    padding: 0.5em 1.5em;
    background-color: #0070a7;
    color: white;
    font-size: 14px;
    font-weight: bold;
    border-radius: 5px;
    border-color: #0070a7;
}
</style>