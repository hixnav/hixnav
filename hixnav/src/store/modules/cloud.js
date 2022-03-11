import { listIO, uploadIO } from '@/api/cloud'

const getDefaultState = () => {
  return {
  }
}

const state = getDefaultState()

const mutations = {
}

const actions = {
  listIO({ commit }, data) {
    return new Promise((resolve, reject) => {
      listIO(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  uploadIO({ commit }, data) {
    return new Promise((resolve, reject) => {
      uploadIO(data).then(response => {
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