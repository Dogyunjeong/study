import Vue from 'vue'
import VueResource from 'vue-resource'
import App from './App.vue'

// Register vueResource to vue.
// Every instance will be implemented vue-resource
Vue.use(VueResource);

// Access from global vue object, $ is not needed
// Vue.http.options.root = 'https://vuejs-http-6f51e.firebaseio.com/data.json';
Vue.http.options.root = 'https://vuejs-http-6f51e.firebaseio.com/';
// Vue.http.options.headers and etc

// Can trim the any request. Like filters
Vue.http.interceptors.push((request, next) => {
  console.log(request);
  if (request.method == 'POST') {
    request.method = 'PUT';
  }
  next(response => {
    response.json = () => { return { message: response.body }}
  })
})

new Vue({
  el: '#app',
  render: h => h(App)
})
