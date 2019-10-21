<template>
    <div class="component">
        <h3>You may view the User Details here</h3>
        <p>Many Details</p>
        <p>User Name: {{ switchName() }}</p>
        <p>User Age: {{ userAge }}</p>
        <button @click="resetName">Reset Name</button>
        <button @click="resetFn">Reset Name</button>
    </div>
</template>

<script>
    import { eventBus } from '../main.js'
    export default {
        // props: ['myName'], //props can be used as normal property in data object
        props: {
            // myName: String  // validating props
            // myname: [String, Array]
            myName: {
                type: String,
                // required: true,  if there is default, There is value anyway
                default: 'Max'
            },
            // for object
            // myName: {
            //     type: Object,
            //     default() {
            //         return {
            //             name: 'Max'
            //         }
            //     }
            // }
            resetFn: Function,
            userAge: Number
        },
        methods: {
            switchName() {
                return this.myName.split("").reverse().join("");
            },
            resetName() {
                this.myName = 'Max';
                this.$emit('nameWasReset', this.myName);   // emit custom event
            }
        },
        created() {
            eventBus.$on('ageWasEdited', (age) => {
                this.userAge = age;
            })
        }
    }
</script>

<style scoped>
    div {
        background-color: lightcoral;
    }
</style>
