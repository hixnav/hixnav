import { addArticleLink } from '@/api/link'

const getDefaultState = () => {
  return {
  }
}

const state = getDefaultState()

const mutations = {
}

const actions = {
  addArticleLink({ commit }, data) {
    return new Promise((resolve, reject) => {
      addArticleLink(data).then(response => {
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