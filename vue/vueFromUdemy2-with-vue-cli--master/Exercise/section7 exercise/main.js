import Vue from 'vue'
import App from './App.vue'
import Home from './Home.vue'

Vue.component('app-servers', Home);

new Vue({
  el: '#app',
  // override the elment #app with App template
  render: h => h(App) // reder App.vue instead 'template'with string
})
