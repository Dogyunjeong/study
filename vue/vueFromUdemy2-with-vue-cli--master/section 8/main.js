import Vue from 'vue'
import App from './App.vue'

export const eventBus = new Vue({
  methods: {
    changeAge(age) {
      this.$emit('ageWasEdited', age);
    }
  }
}); // It need to be before running vue instance

new Vue({
  el: '#app',
  render: h => h(App)
})

