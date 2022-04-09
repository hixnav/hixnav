import { listIO, uploadIO, downIO, delIO } from '@/api/cloud'

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

  downIO({ commit }, data) {
    return new Promise((resolve, reject) => {
      downIO(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  delIO({ commit }, data) {
    return new Promise((resolve, reject) => {
      delIO(data).then(response => {
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