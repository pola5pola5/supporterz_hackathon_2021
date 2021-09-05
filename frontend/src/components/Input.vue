<template>
  <div class="Input">
    <p id="error" v-show="error">{{ error }}</p>
    <p>クリックまたはドラッグ&ドロップで画像を追加してください．</p>
    <label>
      <div
        class="drop_area"
        @dragenter="dragEnter"
        @dragleave="dragLeave"
        @dragover.prevent
        @drop.prevent="dropFile()"
        :class="{ enter: isEnter }"
      >
        ファイルアップロード

        <input
          type="file"
          accept="image/jpeg"
          @change="onImageChange"
          multiple
        />
      </div>
    </label>

    <br />
    <div>
      <ul class="flex">
        <li
          class="flex-col"
          v-for="(file, index) in files"
          :key="index"
          @click="deleteFile(index)"
        >
          <!-- {{ index }} -->
          <div style="position: relative">
            <span class="delete-mark">×</span>
            <img class="file_icon" src="../assets/icon.png" />
            <!-- <img class="file_icon" :src="images[index]"> -->
          </div>
          <span>{{ file.name }}</span>
        </li>
      </ul>
    </div>

    <br />
    <div v-show="files.length">
      <button class="button" v-on:click="upload">送信</button>
    </div>

    <p>{{ message }}</p>
    <p>{{ error }}</p>
  </div>
</template>

<script>
import axios from "axios";
import loadImage from "blueimp-load-image";
export default {
  name: "Input",
  data() {
    return {
      message: "",
      error: "",
      isEnter: false,
      files: [],
      images: [],
      test: "",
    };
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
    deleteFile(index) {
      this.files.splice(index, 1);
      this.images.splice(index, 1);
    },

    checkTypeAndGPS(files) {
      return new Promise((resolve) => {
        files.forEach((file) => {
          // 拡張子のチェック
          const type = file.type;
          if (type != "image/jpeg") {
            alert(
              "エラー: " + file.name + "\n拡張子がjpegの画像を選択してください"
            );
            return;
          }
          // GPSのチェック
          loadImage.parseMetaData(file, (data) => {
            if (data.exif && data.exif.get("GPSInfo")) {
              // console.log("fileWithGPS ", fileWithGPS);
              // console.log("3: GPS取得できてる");
              this.processFile(file);
            } else {
              alert(
                "エラー: " +
                  file.name +
                  "\n位置情報を含む画像を選択してください"
              );
            }
          });
        });
        resolve();
      });
    },

    processFile(file) {
      // console.log("4: imagesを抜き出す処理")
      // console.log("5: ",file)
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => {
        const im = reader.result;
        const base64EncodedFile = im.split(",")[1];
        // console.log(base64EncodedFile); // base64にしたデータ
        this.images.push(base64EncodedFile);
        this.files.push(file);
      };
    },

    dropFile() {
      this.isEnter = false;
      const files = [...event.dataTransfer.files];
      this.checkTypeAndGPS(files);
    },
    onImageChange(e) {
      const files = e.target.files || e.dataTransfer.files;
      this.checkTypeAndGPS(files);
    },

    upload: function () {
      var resStatus;
      var getRequest;

      axios
        .post(
          "/api/auth/trip/save",
          {
            user_id: this.$store.getters["auth/getUserID"],
            imgs: this.images,
          },
          {
            headers: {
              "X-Token": this.$store.getters["auth/getToken"],
            },
          }
        )
        .then((response) => {
          resStatus = response.status;
          getRequest = response.data.trip_id;
          // console.log(getRequest);
          this.message = "アップロードしました";
          this.error = "";
          if (resStatus == 200) {
            // console.log(this.$store.state.tripid);
            // console.log(getRequest);
            this.$store.commit("trip/setTripID", getRequest);
            // console.log(this.$store.state.tripid);
            this.$router.push("/map");
          }
        });
      // .catch((error) => {
      //   // console.log("error");
      //   // console.log(error);
      // });
    },
  },
};
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
label {
  display: flex;
  justify-content: center;
  align-items: center;
}
.drop_area {
  color: #42b983;
  font-weight: bold;
  font-size: 1.2em;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 500px;
  height: 300px;
  border: 5px solid #42b983;
  border-radius: 15px;
}
.enter {
  border: 10px dotted #42b983;
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
span {
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
