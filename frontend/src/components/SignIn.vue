<template>
  <div class="SignIn">
    <h2>Sign in</h2>
    <div>
        <input type="text" placeholder="Username" v-model="username">
    </div>
    <div>
        <input type="password" placeholder="Password" v-model="password">
    </div>
    <button @click="signIn">Signin</button>
    <p>You don't have an account? 
      <router-link to="/signup">create account now!!</router-link>
    </p>
  </div>
</template>

<script>
import axios from "axios"
export default {
  name: 'Signin',
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    signIn: function () {
      var resStatus;
      var getRequest;
      axios
        .post("/user/login", {
          user_name: this.username,
          password: this.password,
        })
        // まだ設定途中
        .then((response) => {
          resStatus = response.status;
          getRequest = response.data;
          console.log(getRequest);
          if (resStatus == 200) {
            console.log(getRequest);
            this.$router.push("/view");
          }
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
.signin {
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