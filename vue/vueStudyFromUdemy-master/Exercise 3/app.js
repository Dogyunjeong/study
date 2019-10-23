new Vue({
        el: '#exercise',
        data: {
            value: 0
        },
        computed: {
            result: function () {
                return this.value > 39 ? 'Done' : 'not there yet'
            }
        },
        watch: {
            result: function() {
                var vm = this
                setTimeout(function () {
                    vm.value = 0
                },5000)
            }
        }
    });