<template>
  <div class="SignIn">
    <div class="content">
      <div class="bg"></div>
      <div class="input_area">
        <div class="center_area">
          <h2>Sign in</h2>
          <div>
            <div>username</div>
            <input class="input_user" type="text" placeholder="Username" v-model="info.user_name" />
          </div>
          <div>
            <div>password</div>
            <input class="input_password" type="password" placeholder="Password" v-model="info.password" />
          </div>
          <button class="sign_btn" @click="signIn">Signin</button>
          <div class="text_center">
            <p>You don't have an account?</p>
            <router-link to="/signup">create account now!!</router-link>
          </div>
        </div>
      </div>
    </div>
    <!-- <h2>Sign in</h2>
    <div>
      <input type="text" placeholder="Username" v-model="info.user_name" />
    </div>
    <div>
      <input type="password" placeholder="Password" v-model="info.password" />
    </div>
    <button @click="signIn">Signin</button>
    <p>
      You don't have an account?
      <router-link to="/signup">create account now!!</router-link>
    </p> -->
  </div>
</template>

<script>
export default {
  name: "SignIn",
  data() {
    return {
      info: {
        user_name: "",
        password: "",
      },
    };
  },
  methods: {
    signIn: function () {
      this.$store.dispatch("auth/signIn", this.info);
    },
  },
  computed: {
    token: function () {
      return this.$store.getters["auth/getToken"];
    },
  },
  watch: {
    token: function () {
      this.$store.commit("user/setUserName", this.info.user_name);
      this.$router.push("/home");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1,
  h2 {
    font-size: 35px;
    margin-bottom: 20px;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #CCB800;
  }
  input {
    margin: 10px 0;
    padding: 10px;
  }
  .input_user{
    margin-top: 0;
    width: 100%;
    margin-bottom: 30px;
  }
  .input_password{
    margin-top: 0;
    width: 100%;
  }
  .center_area{
    width: 80%;
    font-family: serif;
    color: #060B38;
  }
 .bg{
   background-color: #060B38;
   width: 50%;
   height: 100vh;
   margin: 0%;
 }
 .sign_btn{
   margin-top: 20px;
   margin-bottom: 50px;
   margin-left: calc(50% - 50px);
   color: white;
   cursor: pointer;
   font-family: serif;
   font-size: 15px;
   background-color: #060B38;
   width: 100px;
   height: 30px;
   border: 0px;
 }
 .text_center{
   text-align: center;
 }
 .input_area{
   display: flex;
   flex-direction: column;
   align-items: center;
   justify-content: center;
   width: 50%;
   height: 100vh;
 }
 .content{
   display: flex;
   justify-content: space-around;
 }
 *{
    margin: 0%;
    padding: 0%;
  }
</style>
