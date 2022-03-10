import { addLink } from '@/api/nav'

const getDefaultState = () => {
  return {
  }
}

const state = getDefaultState()

const mutations = {
}

const actions = {
  addLink({ commit }, data) {
    return new Promise((resolve, reject) => {
        addLink(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}