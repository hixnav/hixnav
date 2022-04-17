import request from '@/utils/request'

export function account(data) {
    return request({
        url: '/api/account',
        method: 'post',
        data
    })
}