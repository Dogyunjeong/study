import Vue from 'vue'
import * as fireBase from 'firebase'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.css'

import App from './App'
import router from './router'
import { store } from './store/index'
import DateFilter from './filters/date'
import AlertCmp from './components/Shared/Alert.vue'
import EditMeetupDetailsDialog from './components/Meetup/Edit/EditMeetupDetailsDialog.vue'
import EditMeetupDateDialog from './components/Meetup/Edit/EditMeetupDateDialog.vue'
import EditMeetupTimeDialog from './components/Meetup/Edit/EditMeetupTimeDialog.vue'
import RegisterDialog from './components/Meetup/Registration/RegisterDialog.vue'

Vue.use(Vuetify, {
  theme: {
    primary: '#3f51b5',
    secondary: '#b0bec5',
    accent: '#8c9eff',
    error: '#b71c1c'
  }
})

Vue.config.productionTip = false

Vue.filter('date', DateFilter)
Vue.component('app-alert', AlertCmp)
Vue.component('app-edit-meetup-details-dialog', EditMeetupDetailsDialog)
Vue.component('app-edit-meetup-date-dialog', EditMeetupDateDialog)
Vue.component('app-edit-meetup-time-dialog', EditMeetupTimeDialog)
Vue.component('app-meetup-register-dialog', RegisterDialog)
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App),
  created () {
    fireBase.initializeApp({
      apiKey: 'AIzaSyDvTeI8Uds_-ry7DPvBt76gl2ItjdU-ojI',
      authDomain: 'vuejs-http-6f51e.firebaseapp.com',
      databaseURL: 'https://vuejs-http-6f51e.firebaseio.com',
      projectId: 'vuejs-http-6f51e',
      storageBucket: 'gs://vuejs-http-6f51e.appspot.com'
    })
    fireBase.auth().onAuthStateChanged((user) => {
      if (user) {
        this.$store.dispatch('autoSignIn', user)
        this.$store.dispatch('fetchUserData')
      }
    })
    this.$store.dispatch('loadMeetups')
  }
})
