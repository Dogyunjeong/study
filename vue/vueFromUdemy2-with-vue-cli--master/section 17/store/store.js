import Vue from 'vue';
import Vuex from 'vuex';
import counter from './modules/counter'
import * as actions from './actions'
import * as getters from './getters'
import * as mutations from './mutations'

Vue.use(Vuex);

export const store = new Vuex.Store({
    state: {
        value: 0
    },
    getters,
    // Mutation always has to be synchronous
    mutations,
    // can do async task before commit
    actions,
    modules: {
        counter
    }
})

// export const store = new Vuex.Store({
//     state: {
//         counter: 0,
//         value: 0
//     },
//     getters: {
//         doubleCounter: state => {
//             return state.counter * 2
//         },
//         stringCounter: state => {
//             return state.counter + ' Clikcs';
//         },
//         value: state => {
//             return state.value
//         }
//     },
//     // Mutation always has to be synchronous
//     mutations: {
//         increment: (state, payLoad) => {
//             state.counter += payLoad;
//         },
//         decrement: (state, payLoad) => {
//             state.counter -= payLoad;
//         },
//         updateValue: (state, payLoad) => {
//             state.value = payLoad
//         }
//     },
//     // can do async task before commit
//     actions: {
//         // increment: context => {
//         //     setTimeout( , 1000)
//         //     context.commit('increment');
//         // }
//         increment: ({ commit }, payLoad) => {
//             commit('increment', payLoad);
//         },
//         decrement: ({ commit }, payLoad) => {
//             commit('decrement', payLoad);
//         },
//         asyncIncrement: ({ commit }, payLoad) => {
//             setTimeout(() => {
//                 setTimeout(() => {
//                     commit('increment', payLoad.by);
//                 }, payLoad.duration)
//             })
//         },
//         asyncDecrement: ({ commit }, payLoad) => {
//             setTimeout(() => {
//                 setTimeout(() => {
//                     commit('decrement', payLoad.by);
//                 }, payLoad.duration)
//             })
//         },
//         updateValue: ({commit}, payLoad) => {
//             commit('updateValue',)
//         }
//     }
// })