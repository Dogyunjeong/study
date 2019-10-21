import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import { routes }from './routes';

// add Vue plugin
Vue.use(VueRouter)

const router = new VueRouter({
  routes,
  mode: 'history',
  scrollBehavior(to, from, savedPosition) {
    // If user come back to previous page it will scroll where user saw
    if (savedPosition) {
      return savedPosition;
    }
    if (to.hash) {
      return { selector: to.hash }
    }
    // scroll
    return { x: 0, y: 0 }
  }
});

// before every router action, it wil execute
router.beforeEach((to, from, next) => {
  console.log('global beforeEach');
  next();
  // stay at there
  // next(false)
  // To redirection
  // next({})
})

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
