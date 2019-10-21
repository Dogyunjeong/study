<template>
    <li
        @click="serverSelected">
        Server #{{ server.id }}
    </li>
</template>

<script>
    import { ServerBus } from '../../main'
    export default {
        props: {
            server: Object
        },
        methods: {
            serverSelected() {
                ServerBus.$emit('serverSelected', this.server);
            }
        }
    }
</script>
