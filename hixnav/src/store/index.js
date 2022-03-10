import Vue from 'vue'
import Vuex from 'vuex'
import user from './modules/user'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  getters: {
    token: state => state.user.token,
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user
  }
})
