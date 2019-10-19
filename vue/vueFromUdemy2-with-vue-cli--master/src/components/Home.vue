<template>
    <div>
        <h1>Trad or Vie your Portfolio</h1>
        <h6>You may Save  & Load your data</h6>
        <h6>Click on 'End Da' to being a New Da!</h6>
        <hr>
        <p>Your funds: {{ funds }}</p>
    </div>
</template>

<script>
    export default {
        computed: {
            funds() {
                return this.$store.getters.funds;
            }
        }
    }
</script>