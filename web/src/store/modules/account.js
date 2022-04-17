import {account } from '@/api/account'

const getDefaultState = () => {
    return {}
}

const state = getDefaultState()

const mutations = {}

const actions = {
    account({ commit }, data) {
        return new Promise((resolve, reject) => {
            account(data).then(response => {
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