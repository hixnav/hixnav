import request from '@/utils/request'

export function search(data) {
    return request({
        url: '/api/search',
        method: 'post',
        data
    })
}
