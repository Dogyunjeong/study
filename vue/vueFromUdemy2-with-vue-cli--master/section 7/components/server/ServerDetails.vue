<template>
    <div class="col-xs-12 col-sm-6">
        <p v-if="!server">Server Details are currently not updated</p>
        <p v-else>Server #{{ server.id }}  Detail: {{ server.status }}</p>
        <button @click="resetServer">Reset Server Status</button>
    </div>

</template>

<script>
    import { ServerBus } from '../../main'
    export default {
        props: {
            server: null
        },
        methods: {
            resetServer() {
                this.server.status = 'noraml'
            }
        },
        created() {
            ServerBus.$on('serverSelected', (server) => {
                this.server = server
            })
        }
    }
</script>

<style scoped>
    div {
        border: 1px solid red;
    }
</style>
