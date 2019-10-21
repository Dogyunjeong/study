<template>
    <div class="container">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
                <h1>Filters & Mixins</h1>
                <!-- pipe works throw from left to right -->
                <p>{{ text | toUppercase | toLowercase}}</p>
                <hr>
                <button @click="fruits.push('Berries')">Add New Item</button>
                <input v-model="filterText">
                <ul>

                    <!-- filter can be nice solution, But often, computed would be btter solution -->
                    <!-- if put the filter in v-for, it will run when it is rerendered.
                        it could huge performance deal -->
                    <!-- <li v-for="(fruit, index) in fruits | fruitFilter" :key="index">{{ fruit }}</li> -->

                    <!-- alternative way: use computed -->
                    <li v-for="(fruit, index) in filteredFruits" :key="index">{{ fruit }}</li>
                </ul>
                <hr>
                <app-list></app-list>
            </div>
        </div>
    </div>
</template>

<script>
    import List from './List.vue'
    import { fruitMixin } from './fruitMixin'
    export default {
        mixins: [fruitMixin],
        data() {
            return {
                text: 'Hello there!',
                // fruits: ['Apple', 'Banana', 'Mango', 'Melon'],
                // filterText: ''
            }
        },
        filters: {
            toUppercase(value) {
                return value.toUpperCase();
            }
        },
        // computed: {
        //     filteredFruits() {
        //         return this.fruits.filter((element) => {
        //             return element.match(this.filterText)
        //         })
        //     }
        // },
        components: {
            appList: List
        },
        created() {
            console.log('Inside App created Hook')
        }
    }
</script>

<style>

</style>
