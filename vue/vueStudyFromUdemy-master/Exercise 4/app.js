new Vue({
  el: '#exercise',
  data: {
    effect: false,
    hello: 'hello-my-tati',
    tati: 'i-like-you',
    userClass: '',
    first: '',
    second: '',
    mystyle: {
      width: '50px',
      height: '50px',
      backgroundColor: ''
    },
    progress: false,
    progressTime: 0,
    progressStyle: {
      width: '0px',
      height: '10px',
      backgroundColor: 'red'
    }
  },
  computed: {
    highlight: function () {
      return this.effect ? 'highlight' : 'shrink'
    },
    firstClass: function () {
      return this.first
    },
    secondClass: function () {
      return this.second === 'true' ? 'small' : (this.second === 'false' || this.second === false ? 'big' : null)
    },
    customStyle: function () {
      return { backgroundColr: this.backGroundColor }
    }
    // progressStyle: function () {
    //   console.log('dogyunTest working')
    //   return { height: '10px', width: this.progressTime + 'px', backgroundColor: 'blue' }
    // }
  },
  // watch: {
  //   progress: function () {
  //     var vm = this
  //     var interval = setInterval(function () {
  //       console.log('working', vm.progressTime)
  //       vm.progressTime < 100 ? vm.progressTime += 10 : clearInterval(interval)
  //     }, 1000)
  //   }
  // },
  methods: {
    startEffect: function () {
      console.log('dogyunTest')
      var vm = this
      setInterval(function () {
        vm.effect = !vm.effect
      }, 1000)
    },
    startProgress: function () {
      // this.progress = !this.progress
      // console.log('dogyunTest', this.progress)
      var vm = this
      var width = 0
      var interval = setInterval(function () {
        width += 10
        vm.progressStyle.width = width + 'px'
        width >= 100 ? clearInterval(interval) : null
      }, 500)
    }
  }
})
