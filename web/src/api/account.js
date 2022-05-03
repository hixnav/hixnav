import request from '@/utils/request'

export function account(data) {
    return request({
        url: '/api/account',
        method: 'post',
        data
    })
}

export function addAccount(data) {
    return request({
        url: '/api/addAccount',
        method: 'post',
        data
    })
}

export function editAccount(data) {
    return request({
        url: '/api/editAccount',
        method: 'post',
        data
    })
}

export function delAccount(data) {
    return request({
        url: '/api/delAccount',
        method: 'post',
        data
    })
}