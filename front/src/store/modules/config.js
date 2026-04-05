import {dbTest, dbMigrate, cosMigrate, syncPush, syncPull, syncStatus, saveBanner, getBanner} from '@/api/config'

const getDefaultState = () => {
    return {}
}

const state = getDefaultState()

const mutations = {}

const actions = {
    dbTest({ commit }, data) {
        return new Promise((resolve, reject) => {
            dbTest(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    dbMigrate({ commit }, data) {
        return new Promise((resolve, reject) => {
            dbMigrate(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    cosMigrate({ commit }, data) {
        return new Promise((resolve, reject) => {
            cosMigrate(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    syncPush({ commit }, data) {
        return new Promise((resolve, reject) => {
            syncPush(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    syncPull({ commit }, data) {
        return new Promise((resolve, reject) => {
            syncPull(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    syncStatus({ commit }, data) {
        return new Promise((resolve, reject) => {
            syncStatus(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    saveBanner({ commit }, data) {
        return new Promise((resolve, reject) => {
            saveBanner(data).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
    },

    getBanner({ commit }, data) {
        return new Promise((resolve, reject) => {
            getBanner(data).then(response => {
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