<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
      <a @click.prevent="run">Click Me</a>
    </div>
    <router-view/>
  </div>
</template>

<script>
import authClient from '../../pb/js/auth_grpc_web_pb'
import authObjs from '../../pb/js/auth_pb'


export default {
  name: "app",
  methods: {
    run() {
      const client = new authClient.V1AuthClient("http://localhost:8080")
      const request = new authObjs.LoginRequest()

      request.setEmail("foo@bar.com")
      request.setPassword("1234")

      client.login(request, {'custom-header-1': 'value1'}, (err, resp) => {

        alert(resp.getToken())

      })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
