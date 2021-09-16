<template>
  <div>
    <div id="overlay" v-show="isopen" @click.self="closePopup">
      <div id="content">
        <div class="userIcon">{{user.slice(0,1).toUpperCase()}}</div>
        <div class="name">{{user}}</div>
        <div class="line"></div>
        <div class="name">{{tripnum}} Trip</div>
        <div class="line"></div>
        <div class="btn" @click="onClickLogout">Log out</div>
        <div><button class="btn" @click="closePopup">close</button></div>
      </div>
    </div>
  </div>
</template>

<script>
export default{
  props:{
    isopen: Boolean,
    user: String,
    tripnum: Number,
  },

  data(){
    return{
      popup: this.isopen,
    }
  },

  methods:{
    closePopup(){
      this.popup = false
      this.$emit("close", this.popup)
    },

    onClickLogout(){
      this.$store.commit("auth/deleteInfo");
      this.$router.push("/");
    }
  }
}
</script>

<style scoped>
  #overlay{
    z-index:1;
    position: fixed;
    top:0;
    left:0;
    width:100%;
    height:100%;
  }

  #content{
    z-index:2;
    width:200px;
    height: 250px;
    float: right;
    background:#2F2F2F;
    margin-top: 70px;
    margin-right: 30px;
    border-radius: 10px;
    text-align: center;
    color: white;
  }

  .userIcon{
    color: white;
    background-color: #C850BC;
    margin-top: 30px;
    margin-bottom: 10px;
    margin-left: auto;
    margin-right: auto;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    text-align: center;
    line-height: 40px;
  }

  .line{
    border-bottom: 1px solid white;
    margin-top: 10px;
    margin-bottom: 10px;
  }

  .name{
    margin-top: 10px;
    margin-bottom: 10px;
  }

  .btn{
    margin-top: 10px;
    margin-bottom: 10px;
    cursor: pointer;
  }
</style>