<template>
    <div class="row">
        <!-- native modifier will allow component to be clickable by vue -->
        <app-quote 
                v-for="(quote, idx) in quotes" 
                :key="idx"
                @click.native="deleteQuote(idx)">
            {{ quote }}
        </app-quote>
    </div>    
</template>

<script>
    import Quote from './Quote.vue';

    export default {
        props: ['quotes'],
        components: {
            appQuote: Quote
        },
        methods: {
            deleteQuote(idx) {
                this.$emit('quoteDeleted', idx)
            }
        }
    }
</script>

<style>

</style>
