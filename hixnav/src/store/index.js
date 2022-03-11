import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import user from './modules/user'
import nav from './modules/nav'
import cloud from './modules/cloud'
import link from './modules/link'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user,
    nav,
    cloud,
    link
  },
  getters
})

export default store
