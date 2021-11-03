<template>
  <div id="app">
    <button type="button" @click="getLaw">Get a Hacker Law</button>
    <div>
      <h4>{{ law.name }}</h4>
      <p>{{ law.definition }}</p>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default {
  name: 'App',
  data() {
    return {
      law: {
        name: '',
        description: ''
      }
    }
  },
  methods: {
    getLaw () {
      Vue.axios.get("/law/").then(res => {
        console.log("res", res)
        this.law = res.data
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
  margin-top: 60px;
}

button {
  font-size: 18px;
  color: fuchsia;
}
</style>
