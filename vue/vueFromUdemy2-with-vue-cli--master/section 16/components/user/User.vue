<template>
    <div>
        <h1>The User Page</h1>
        <hr>
        <button @click="navigateToHome" class="btn btn-primary">Go To Home</button>
        <!-- <p>Loaded ID: {{ id }} </p> -->
        <hr>
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        // data() {
        //     return {
        //         id: this.$route.params.id
        //     }
        // },
        // watch: {
        //     // can be computed for id
        //     '$route'(to, from) {
        //         this.id = to.params.id
        //     }
        // },
        methods: {
            navigateToHome() {
                // this.$router.push('/')
                // this.$router.push({ path: '/' }) 
                this.$router.push({ name: 'home' })      
            }
        }
    }
</script>