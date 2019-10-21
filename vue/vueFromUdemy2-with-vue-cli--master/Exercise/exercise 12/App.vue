<template>
    <div class="container">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
                <h1>Directives Exercise</h1>
                <!-- Exercise -->
                <!-- Build a Custom Directive which works like v-on (Listen for Events) -->
                <p v-my-on:click="myAlert">Click Me</p>
            </div>
        </div>
        <div 
                style="width:100px; height: 100px; backgroundColor: lightgreen;"
                v-my-on:mouseenter="mouseEnter"
                v-my-on:mouseleave="mouseLeave"></div>
    </div>
</template>

<script>
    export default {
        directives: {
            'my-on': {
                bind(el, binding, vnode) {
                    const type = binding.arg
                    const listener = binding.value
                    el.addEventListener(type, listener)
                }
            }
        },
        methods: {
            myAlert() {
                return alert('Hello MyOn working')
            },
            mouseEnter() {
                console.log('mouse entered')
            },
            mouseLeave() {
                console.log('mouse leaved')
            }
        }
    }
</script>

<style>
</style>
