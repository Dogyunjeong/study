<template>
    <div class="container">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
                <h1>Animations</h1>
                <hr>
                <select v-model="alertAnimation" class="form-control">
                    <option value="fade">Fade</option>
                    <option value="slide">Slide</option>
                </select>
                <hr>
                <button class="btn btn-primary" @click="show = !show">Show Alert</button>
                <br><br>
                <!-- This is special tag for animation.
                    it is to show only one element, it can contain mutiple, but will show one -->

                <!-- with v-show, can work -->

                <!-- 1. To control multiple elemnt, need to set the Key
                        as vue can't recongnize without unique key
                     2. mode will define how it will work -->
                <transition :name="alertAnimation" mode="out-in">
                    <div class="alert alert-info" v-if="show" key="info">This is Some Info</div>
                    <div class="alert alert-warning" v-else key="warning">This is Some alert-warning</div>
                </transition>

                <!-- <transition name="slide">
                    <div class="alert alert-info" v-if="show">This is Some Info</div>
                </transition> -->

                <!-- vue js doesn't know  which should take between transition and animation.
                    It will take longer one.
                    So Have to declare it 
                    transition or animation-->
                <transition name="slide" type="animation">
                    <div class="alert alert-info" v-if="show">This is Some Info</div>
                </transition>

                <!-- To animate at initial -->
                <transition name="slide" type="animation" appear>
                    <div class="alert alert-info" v-if="show">This is Some Info</div>
                </transition>
                <transition 
                        enter-active-class="animated bounce"
                        leave-active-class="animated shake">
                    <div class="alert alert-info" v-if="show">This is Some Info</div>
                </transition>
                <hr>
                <hr>
                <button class="btn btn-primary" @click="load = !load">Load / Remove Element</button>
                <hr>
                <!-- If there is name in transition, vue will look up Css -->
                <!-- :css="false" mean don't look up css and learn js hooks -->
                <transition
                    @before-enter="beforeEnter"
                    @enter="enter"
                    @after-enter="afterEnter"
                    @enter-cancelled="enterCancelled"
                    
                    @before-leave="beforeLeave"
                    @leave="leave"
                    @after-leave="afterLeave"
                    @leave-cancelled="leaveCancelled"
                    :css="false">
                    <div style="width: 300px; height: 100px; background-color: lightgreen" v-if="load"></div>
                </transition>
                <hr>
                <button class="btn btn-primary"
                        @click="selectedComponent == 'app-success-alert' ? selectedComponent = 'app-danger-alert' : selectedComponent = 'app-success-alert'" >
                    Toggle Component</button>
                <br><br>
                <!-- How may element inside of transition is not matter
                    But the end there will be only one to show -->
                <transition name="fade" mode="out-in">
                    <component :is="selectedComponent"></component>                
                </transition>
                <hr>
                <button class="btn btn-primary" @click="addItem">Add Item</button>
                <br><br>
                <!-- <transition> is not redered to the DOM 
                    <trasition-group> does render a new HTML tag by default,
                    that will be a <span>, you can overwrite this by setting
                        <transition-group tag="TAG"> -->
                
                <ul class="list-group">
                    <transition-group name="slide">
                        <li
                                class="list-group-item"
                                v-for="(number, index) in numbers"
                                @click="removeItem(index)"
                                style="cursor: pointer"
                                :key="number"
                                >{{ number }}
                        </li>
                                   
                    </transition-group>
                </ul>             
            </div>
        </div>
    </div>
</template>

<script>
    import DangerAlert from './DangerAlert.vue'
    import SuccessAlert from './SuccessAlert.vue'

    export default {
        data() {
            return {
                show: false,
                alertAnimation: 'fade',
                load: true,
                elementWidth: 100,
                selectedComponent: 'app-success-alert',
                numbers: [1, 2, 3, 4, 5]
            }
        },
        methods: {
            beforeEnter(el) {
                console.log('beforeEnter');
                this.elementWidth = 100;
                el.style.width = this.elementWidth + 'px'
            },
            // done tell to this animaition is finished to vue
            // if there is css code, don't need to use done as vue know when it will be finished
            enter(el, done) {
                console.log('enter')
                let round = 1;
                const  interval = setInterval(() => {
                    el.style.width = (this.elementWidth + round * 10) + 'px';
                    round++;
                    if (round > 20) {
                        clearInterval(interval);
                        done();
                    }
                }, 20);
            },
            afterEnter(el) {
                console.log('afterEnter')
            },
            enterCancelled(el) {
                console.log('enterCancelled')
            },
            beforeLeave(el) {
                console.log('beforeLeave')
                this.elementWidth = 300;
                el.style.width = this.elementWidth + 'px';
            },
            leave(el, done) {
                console.log('leave')
                let round = 1;
                const  interval = setInterval(() => {
                    el.style.width = (this.elementWidth - round * 10) + 'px';
                    round++;
                    if (round > 20) {
                        clearInterval(interval);
                        done();
                    }
                }, 20);
            },
            afterLeave(el) {
                console.log('afterLeave')
                this.elementWidth = 100
                el.style.width = this.elementWidth + 'px';
            },
            leaveCancelled(el) {
                console.log('leaveCancelled')
            },
            addItem() {
                const pos = Math.floor(Math.random() * this.numbers.length);
                this.numbers.splice(pos, 0, this.numbers.length + 1)
            },
            removeItem(index) {
                this.numbers.splice(index, 1);
            }
        },
        components: {
            appDangerAlert: DangerAlert,
            appSuccessAlert: SuccessAlert
        }
    }
</script>

<style>
    /* it will attached only one frame at first */
    .fade-enter {
        opacity: 0;
    }
    /* this will animated the transition is for vue */
    .fade-enter-active{
        transition: opacity 1s;
    }

    .fade-leave {
        /* opacity: 1; */
    }

    .fade-leave-active {
        transition: opacity 1s;
        opacity: 0;
    }

    .slide-enter {
        opacity: 0;
        /* transform: translateY(20px) */
    }

    .slide-enter-active {
        animation: slide-in 1s ease-out forwards;
        transition: opacity .5s;
    }

    .slide-leave {
        
    }

    .slide-leave-active {
        animation: slide-out 1s ease-out forwards;
        transition: opacity 3s;
        opacity: 0;
        position: absolute;
    }

    .slide-move {
        transition: transform 1s;
    }

    @keyframes slide-in {
        from {
            transform: translateY(20px);
        }
        to {
            transform: translateY(0);
        }
    }
    @keyframes slide-out {
        from {
            transform: translateY(0);
        }
        to {
            transform: translateY(20px);
        }
    }
</style>
