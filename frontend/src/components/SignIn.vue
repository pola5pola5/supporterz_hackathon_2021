<template>
  <div class="SignIn">
    <h2>Sign in</h2>
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
    </p>
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
      this.$router.push("/input");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
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
  align-items: center;
}
input {
  margin: 10px 0;
  padding: 10px;
}
</style>
