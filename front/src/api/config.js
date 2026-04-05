import request from '@/utils/request'

export function dbTest(data) {
    return request({
        url: '/api/db/test',
        method: 'post',
        data
    })
}

export function dbMigrate(data) {
    return request({
        url: '/api/migrate/db',
        method: 'post',
        data
    })
}

export function cosMigrate(data) {
    return request({
        url: '/api/config/oss',
        method: 'post',
        data
    })
}

export function syncPush(data) {
    return request({
        url: '/api/sync/push',
        method: 'post',
        data
    })
}

export function syncPull(data) {
    return request({
        url: '/api/sync/pull',
        method: 'post',
        data
    })
}

export function syncStatus(data) {
    return request({
        url: '/api/sync/status',
        method: 'post',
        data
    })
}

export function saveBanner(data) {
    return request({
        url: '/api/config/banner/save',
        method: 'post',
        data
    })
}

export function getBanner(data) {
    return request({
        url: '/api/config/banner/get',
        method: 'post',
        data
    })
}
