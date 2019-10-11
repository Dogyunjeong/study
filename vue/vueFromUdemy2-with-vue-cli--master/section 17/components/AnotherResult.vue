<template>
    <div>        
        <p>Counter is: {{ doubleCounter }}</p>
        <p>Number of Clicks: {{ stringCounter }}</p>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';
    import * as types from '../store/types';
    export default {
        computed: {
            // ES6 syntax. it will sprate all the properties of object
            // ...mapGetters([
            //     'doubleCounter',
            //     'stringCounter'
            // ]),
            ...mapGetters({
                doubleCounter: types.DOUBLE_COUNTER,
                stringCounter: types.CLICK_COUNTER
            }),
            myComputed() {

            }
        }
     }

    //  export default {
    //      computed: {  
    //          // counter() {
    //         //     return this.$store.getters.doubleCounter;
    //         // },
    //         // clicks() {
    //         //     return this.$store.getters.stringCounter;
    //         // }}
    //  }
</script>