<template>
    <div class="container">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
                <h1>Built-in Directives</h1>
                <p v-text="'Some Text'"></p>
                <p v-text="'<strong>Some Text</strong>'"></p>
            </div>
        </div>
        <hr>
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
                <h1>Customctives</h1>
                <p v-highlight:background.delay="'red'">Color this</p>
                <!-- multiple modifier && multi values-->
                <p v-local-highlight:background.delay.blink="{mainColor: 'red', secondColor: 'green', delay: 500 }">Color this</p>
                <p v-highlight="'red'">Color this</p>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        directives: {
            'local-highlight': {
                bind(el, binding, vnode) {
                    var delay = 0
                    // 3. add modifier
                    if (binding.modifiers['delay']) {
                        delay = 3000;
                    }
                    if (binding.modifiers['blink']) {
                        let mainColor = binding.value.mainColor;
                        let secondColor = binding.value.secondColor;
                        let currentColor = mainColor;
                        setTimeout(() => {
                            setInterval(() => {
                                currentColor == secondColor ? currentColor = mainColor : currentColor = secondColor
                                // 1. arguments
                                if (binding.arg == 'background') {
                                    el.style.backgroundColor = currentColor          
                                } else {
                                    // 2. value
                                    el.style.textColor = currentColor        
                                }
                            }, binding.value.delay)
                            // 1. arguments
                            if (binding.arg == 'background') {
                                el.style.backgroundColor = binding.value.mainColor        
                            } else {
                                // 2. value
                                el.style.textColor = binding.value.secondColor        
                            }
                        }, delay);                        
                    } else {
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

                }
            }
        }
    }
</script>

<style>

</style>
