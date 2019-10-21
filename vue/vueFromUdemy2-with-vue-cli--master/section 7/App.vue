<template>
    <div class="container">
        <appHeader></appHeader>
        <!--<app-Header></app-Header> It is work as vue will give hypernated version from camel cases -->
        <hr>
        <div class="row">
            <app-servers :changeDetail="changeDetail"></app-servers>
            <app-server-details :server="server"></app-server-details>
        </div>
        <hr>
        <app-footer></app-footer>
    </div>
</template>

<script>
    import Header from './components/shares/Header.vue';
    import Footer from './components/shares/Footer.vue';
    import Servers from './components/server/Servers.vue';
    import ServerDetails from './components/server/ServerDetails.vue';

    export default {
        data() {
            return {
                server: {}
            }
        },
        components: {
            // It is not dorm, so it is case sensitive
            appHeader: Header,
            'app-servers': Servers,
            'app-server-details': ServerDetails,
            'app-footer': Footer
        },
        methods: {
            changeDetail(server) {
                this.server = server
            }
        }
    }
</script>

<style>

</style>
