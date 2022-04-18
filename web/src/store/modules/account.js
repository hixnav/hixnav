import {account,addAccount,delAccount } from '@/api/account'

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

    addAccount({ commit }, data) {
        return new Promise((resolve, reject) => {
            addAccount(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    delAccount({ commit }, data) {
        return new Promise((resolve, reject) => {
            delAccount(data).then(response => {
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