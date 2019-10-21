<template>
    <div class="col-xs-12 col-sm-6">
        <ul class="list-group">
            <app-server 
                v-for="server in servers"
                :key="server.id"
                :server="server"></app-server>
        </ul>
    </div>
</template>

<script>
    import Server from './Server.vue'
    export default {
        props: {
            changeDetail: Function
        },
        data() {
            return {
                servers: [
                    { id: 1, status: 'normal' },
                    { id: 2, status: 'warning' },
                    { id: 3, status: 'in danger' },
                    { id: 4, status: 'inform' },
                    { id: 5, status: 'normal' },
                ]
            }
        },
        components: {
            appServer: Server
        }
    }
</script>

<style scoped>
    div {
        border: 1px solid blue;
    }
</style>
