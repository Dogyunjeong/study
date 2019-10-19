import Vue from 'vue'
import App from './App.vue'

Vue.directive('highlight', {
  bind(el, binding, vnode) {
    // el.style.backgroundColor = 'green'
    // el.style.backgroundColor = binding.value
    var delay = 0
    // 3. add modifier
    if (binding.modifiers['delay']) {
      delay = 3000;
    }
    setTimeout(() => {
      // 1. arguments
      if (binding.arg == 'background') {
        el.style.backgroundColor = binding.value          
      } else {
        // 2. value
        el.style.textColor = binding.value        
      }
    }, delay)

  }
})

new Vue({
  el: '#app',
  render: h => h(App)
})
