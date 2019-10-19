<template>
    <div class="container">
        <div class="row">
            <div class="col-xs-12">
                <button @click="selectedComponent = 'appQuote'">Quote</button>
                <button @click="selectedComponent = 'appAuthor'">Author</button>
                <button @click="selectedComponent = 'appNew'">New</button>
                <hr>
                <p>{{ selectedComponent }}</p>
                <!-- component is for dynamic component -->
                <!-- Keep alive makes dynamic component instance kept -->
                <keep-alive>
                    <component :is="selectedComponent"></component>
                </keep-alive>
               <!-- <app-quote>
                   <h2 slot="title">{{ quoteTitle }}</h2>
                   <p>A wondeful Qoute</p>
               </app-quote> -->
            </div>
        </div>
    </div>
</template>

<script>
    import Quote from './components/Quote.vue'
    import appAuthor from './components/Author.vue'
    import appNew from './components/New.vue'

    export default {
        data() {
            return {
                quoteTitle: 'The Quote',
                selectedComponent: 'appQuote'
            }
        },
        components: {
            appQuote: Quote,
            appAuthor,
            appNew
        }
    }
</script>

<style>
</style>
