new Vue({
        el: '#exercise',
        data: {
            value: ''
        },
        methods: {
            buttonClickAlert: function () {
                alert('Button is clicked')
            },
            changeValueWhenKeydown: function(event) {
                return this.value = event.target.value
            },
            changeValueWhenEnter: function(event) {
                console.log('dogyunTest event: ', event)
                return this.value = event.target.value
            }
        }
    });