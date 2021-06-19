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
  {{ img }}
  <br>
  <div>
        <ul class="flex">
            <!-- <li v-for="file in files" :key="file.name" class="flex-col"> -->
               <li class="flex-col" v-for="(file,index) in files" :key="index" @click="deleteFile(index)">
              <div style="position: relative;">
                <span class="delete-mark">×</span>
                <img class="file_icon" src="../assets/icon.png">
              </div>
              <span>{{ file.name }}</span>
            </li>
        </ul>
        <!-- <ul>
            <li v-for="file in files" :key="file.name">{{ changeBase64(file) }}
            </li>
        </ul> -->
  </div>
  <img :src="avatar" class="image">
  <!-- <img :src="avatar" alt="Avatar" class="image"> -->
  <br>
  <div v-show="files.length">
    <button class="button" @click="changeBase64(files)">送信</button>
  </div>
  <!-- <button @click="upload()">アップロード</button> -->
  <p>{{ message }}</p>
  <p>{{ test }}</p>
  </div>
</template>

<script>
// 現状multipleだけど表示には一枚しか載せられない
// ドラッグ&ドロップだと今はダメ
// 選択して追加すると変換される
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
      img: [],
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
    dropFile: function() {
      console.log(event.dataTransfer.files)
      // console.log(event.dataTransfer.files.length)
      this.files.push(...event.dataTransfer.files)
      this.isEnter = false;
      this.files.forEach(file => {
        let form = new FormData()
        form.append('file', file)
      })
      return this.files;
    },
    deleteFile(index) {
      this.files.splice(index, 1)
    },
    // getBase64 (file) {
    //   return new Promise((resolve, reject) => {
    //     const reader = new FileReader()
    //     reader.readAsDataURL(file)
    //     reader.onload = () => resolve(reader.result)
    //     reader.onerror = error => reject(error)
    //   })
    // },
    // onImageChange (e) {
    //   const images = e.target.files || e.dataTransfer.files
    //   this.getBase64(images[0])
    //     .then(image => this.avatar = image)
    //     .catch(error => this.setError(error, '画像のアップロードに失敗しました。'))
    // },
    // upload () {
    // 	if (this.avatar) {
    //     console.log(this.avatar)
    //   	/* postで画像を送る処理をここに書く */
    //     this.message = 'アップロードしました' 
    //     this.error = ''
    //   } else {
    //   	this.error = '画像がありません'
    //   }
    // },
    changeBase64: function(files) {
      let im = null;
      console.log(files)
      for (let step = 0; step < files.length; step++) {
        console.log(files[step])
        let reader = new FileReader();
        reader.readAsDataURL(files[step]);
        console.log(reader.result)
        reader.onload = function() {
          console.log(reader.result)
          im = reader.result;
          console.log(im)
          this.img.push(im)
        }.bind(this);
      };
      axios.post("http://localhost:1323", 
      {
        user_id: "hoge",
        img: this.img
      }).then(response => {
          console.log(response.data)
        }).catch(error => {
            console.log("error")
            console.log(error)
        })
      // const jsonPostRequest = axios.post("http://localhost:1323", 
      // {
      //   user_id: "hoge",
      //   img: this.img
      // })
      // // const jsonPostRequest = await axios.post("http://localhost:1323/", {
      // //   user_id: "hoge",
      // //   img: this.img
      // // });

      // this.test = jsonPostRequest.data
    }
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