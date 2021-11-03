import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false


const client = axios.create({
  baseURL: '/api/v1',
})

Vue.use(VueAxios, client)
new Vue({
  render: h => h(App),
}).$mount('#app')
