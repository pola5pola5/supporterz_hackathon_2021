<template>
  <div class="SignUp">
    <h2>Sign up</h2>
    <div>
        <input type="text" placeholder="Username" v-model="username">
    </div>
    <div>
        <!-- <input type="password" placeholder="パスワード" v-model="password" v-validate="'required|min:6|max:20'" maxlength="20" name="password"> -->
        <input type="password" placeholder="Password" v-model="password">
    </div>
    <button @click="signUp">Register</button>
    <p>Do you have an account? 
      <router-link to="/signin">sign in now!!</router-link>
    </p>
  </div>
</template>

<script>
import axios from "axios"
export default {
  name: 'SignUp',
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    signUp: function () {
      alert(this.password)
      // var resStatus;
      // var getRequest;
      axios
        .post("/api/user/create", {
          user_name: this.username,
          password: this.password,
        })
        // まだ設定途中
        .then((response) => {
          alert('Create account: ', response.data)
          // console.log('Create account: ', response.data)
          // resStatus = response.status;
          // getRequest = response.data;
          // if (resStatus == 200) {
            // console.log(resStatus);
            // console.log(getRequest);
            // this.$router.push("/home");
          // }
        })
        .catch((error) => {
          alert(error.message)
          console.log("error");
          console.log(error);
        });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
  color: #42b983;
}
.signup {
  margin-top: 20px;

  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center
}
input {
  margin: 10px 0;
  padding: 10px;
}
</style>