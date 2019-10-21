<template>
    <div>
        <h1>Filters & Mixins</h1>
        <input v-model="filterText">
        <ul>

            <!-- filter can be nice solution, But often, computed would be btter solution -->
            <!-- if put the filter in v-for, it will run when it is rerendered.
                it could huge performance deal -->
            <!-- <li v-for="(fruit, index) in fruits | fruitFilter" :key="index">{{ fruit }}</li> -->

            <!-- alternative way: use computed -->
            <li v-for="(fruit, index) in filteredFruits" :key="index">{{ fruit }}</li>
        </ul>
    </div>
</template>

<script>
    import { fruitMixin } from './fruitMixin'
    export default {
        mixins: [fruitMixin],
        created() {
            console.log('Inside List Created Hook')
        }
    }
</script>

<style>

</style>
