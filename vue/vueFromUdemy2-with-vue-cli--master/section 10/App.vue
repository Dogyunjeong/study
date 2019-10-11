<template>
    <div class="container">
        <app-header
            :quoteCount="quotes.length"
            :maxQuotes="maxQuotes"></app-header>
        <app-new-quote @quoteAdded="newQuote"></app-new-quote>
        <app-quote-grid 
            :quotes="quotes"
            @quoteDeleted="deleteQuote"></app-quote-grid>
        <div class="row">
            <div class="col-sm-12 text-center">
                <div class="alert alert-info">Info: Click on a Quote to delete it!</div>
            </div>
        </div>
    </div>
</template>

<script>
    import QuoteGrid from './component/QuoteGrid.vue'
    import NewQuote from './component/NewQuote.vue';
    import Header from './component/Header.vue';

    export default {
        data() {
            return {
                maxQuotes: 10,
                quotes: [
                    'Just a Qoute to see something'
                ]
            }
        },
        components: {
            appQuoteGrid: QuoteGrid,
            appNewQuote: NewQuote,
            appHeader: Header,
        },
        methods: {
            newQuote(newQuote) {
                if (this.quotes.length >= this.maxQuotes) {
                    return alert('Please delete Quotes first!')
                }
                this.quotes.push(newQuote);
            },
            deleteQuote(index) {
                this.quotes.splice(index, 1)
            }
        }
    }
</script>

<style>
</style>
