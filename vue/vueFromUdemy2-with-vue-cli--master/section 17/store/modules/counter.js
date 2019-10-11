import * as types from '../types'

const state = {
    counter: 0
}

const  getters = { 
    [types.DOUBLE_COUNTER]: state => {
        return state.counter * 2
    },
    [types.CLICK_COUNTER]: state => {
        return state.counter + ' Clikcs';
    }
}

const mutations = {
    increment: (state, payLoad) => {
        state.counter += payLoad;
    },
    decrement: (state, payLoad) => {
        state.counter -= payLoad;
    },
}


const actions = {
    increment: ({ commit }, payLoad) => {
        commit('increment', payLoad);
    },
    decrement: ({ commit }, payLoad) => {
        commit('decrement', payLoad);
    },
    asyncIncrement: ({ commit }, payLoad) => {
        setTimeout(() => {
            setTimeout(() => {
                commit('increment', payLoad.by);
            }, payLoad.duration)
        })
    },
    asyncDecrement: ({ commit }, payLoad) => {
        setTimeout(() => {
            setTimeout(() => {
                commit('decrement', payLoad.by);
            }, payLoad.duration)
        })
    }
}

export default {
    state,
    getters,
    mutations,
    actions
}