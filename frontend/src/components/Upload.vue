<template>
  <div class="Upload">
    <div>
      <label>
        <!-- <input type="file" name="example" accept="image/*" multiple @change="upload"> -->
        <input ref="file" type="file" v-on:change="upload" accept="image/*" name="example" />
      </label>
    </div>
    <div>
     <p>オリジナル画像の幅：{{ widthBefore }}</p>
     <p>オリジナル画像の高さ：{{ heightBefore }}</p>
     <p> {{ exifTags }} </p>
     <p> {{ gps }} </p>
    </div>
    <div>
      <canvas id="preview-before"></canvas>
    </div>
  </div>

</template>

<script>
import ex from 'exif-js'
export default {
  name: "Upload",
  data() {
    return {
      widthBefore: 0,
      heightBefore: 0,
      exifTags: "",
      gps: ""
    };
  },
  methods: {
    upload: function(event) {
      let img = null;
      let file = event.target.files; // 複数
      console.log(file[0])
      console.log("files")
      console.log(file)
      if (file.length == 0) {
        return;
      }
      let reader = new FileReader();
      reader.readAsDataURL(file[0]);
      // ファイルを読み込んだときの処理
      reader.onload = function() {
        img = new Image();
        
        console.log("img") //base64形式のデータが表示される
        console.log(img) //base64形式のデータが表示される
        
        // 画像が読み込まれたときの処理（canvasに描画）
        img.onload = function() {
          let canvas = document.getElementById("preview-before");
          if (canvas.getContext) {
            let context = canvas.getContext("2d");
            this.widthBefore = img.width;
            this.heightBefore = img.height;
            canvas.width = this.widthBefore;
            canvas.height = this.heightBefore;
            context.drawImage(img, 0, 0);
          }
        }.bind(this);
        img.src = reader.result;

        ex.EXIF.getData(img, () => {
          this.exifTags = ex.EXIF.getTag(img,"DateTimeOriginal");
          this.gps = ex.EXIF.getAllTags(img)
          console.log(this.exifTags)
          console.log(this.gps)
        })
      }.bind(this);
    }
  }
};
// export default {
//   name: 'Upload',
//   props: {
//     value: {
//       type: String,
//       default: null
//     }
//   },
//   data() {
//     return {
//       file: null
//     }
//   },
//   methods: {
//     async upload(event) {
//       const files = event.target.files || event.dataTransfer.files
//       const file = files[0]

//       if (this.checkFile(file)) {
//         const picture = await this.getBase64(file)
//         this.$emit('input', picture)
//       }
//     },
//     getBase64(file) {
//       return new Promise((resolve, reject) => {
//         const reader = new FileReader()
//         reader.readAsDataURL(file)
//         reader.onload = () => resolve(reader.result)
//         reader.onerror = error => reject(error)
//       })
//     },
//     checkFile(file) {
//       let result = true
//       const SIZE_LIMIT = 5000000 // 5MB
//       // キャンセルしたら処理中断
//       if (!file) {
//         result = false
//       }
//       // jpeg か png 関連ファイル以外は受付けない
//       if (file.type !== 'image/jpeg' && file.type !== 'image/png') {
//         result = false
//       }
//       // 上限サイズより大きければ受付けない
//       if (file.size > SIZE_LIMIT) {
//         result = false
//       }
//       return result
//     }
//   }
// }
</script>

<style scoped>
.user-photo {
  cursor: pointer;
  outline: none;
}

.user-photo.default {
  align-items: center;
  background-color: #0074fb;
  border: 1px solid #0051b0;
  border-radius: 2px;
  box-sizing: border-box;
  display: inline-flex;
  font-weight: 600;
  justify-content: center;
  letter-spacing: 0.3px;
  color: #fff;
  height: 4rem;
  padding: 0 1.6rem;
  max-width: 177px;
}

.user-photo.default:hover {
  background-color: #4c9dfc;
}

.user-photo.default:active {
  background-color: #0051b0;
}

.user-photo-image {
  max-width: 85px;
  display: block;
}

.user-photo-image:hover {
  opacity: 0.8;
}

.file-button {
  display: none;
}

.uploaded {
  align-items: center;
  display: flex;
}
</style>