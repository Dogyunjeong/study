<template>
    <div class="row">
        <form>
            <div class="col-sm-8 col-sm-offset-2 col-x12 col-md-6 col-md-offset-3 from-group">
                <label>Quote</label>
                <textarea class="form-control" rows="3" v-model="quote"></textarea>
            </div>
            
            <div class="col-sm-8 col-sm-offset-2 col-x12 col-md-6 col-md-offset-3 from-group">
                <!-- To prevent reload and submit the form -->
                <button class="btn btn-primary" @click.prevent="createNew">Add Quote</button>
            </div>
        </form>
    </div>    
</template>

<script>
    export default {
        data() {
            return {
                quote: ''
            }
        },
        methods: {
            createNew() {
                this.$emit('quoteAdded', this.quote);
                this.quote = '';
            }
        }
    }
</script>

<style>

</style>
