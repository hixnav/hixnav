import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import user from './modules/user'
import nav from './modules/nav'
import cloud from './modules/cloud'
import link from './modules/link'
import account from './modules/account'
import createPersistedState from "vuex-persistedstate"

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
    link,
    account,
  },
  getters,
  plugins: [createPersistedState({
    // paths:  ['user.signin','nav.activeIndex']
    storage: window.localStorage,
    reducer (val) {
      return {
        // 只保存module user内部所有变量持久化
        user: val.user,
        // nav: val.nav
      }
    }
  })],
})

export default store
