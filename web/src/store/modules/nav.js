import { addLink, editLink, delLink } from '@/api/nav'

const getDefaultState = () => {
    return {
        activeIndex: 1
    }
}

const state = getDefaultState()

const mutations = {}

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

    editLink({ commit }, data) {
        return new Promise((resolve, reject) => {
            editLink(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    delLink({ commit }, data) {
        return new Promise((resolve, reject) => {
            delLink(data).then(response => {
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