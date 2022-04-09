import request from '@/utils/request'

export function addArticleLink(data) {
    return request({
        url: '/api/addArticleLink',
        method: 'post',
        data
    })
}

export function article(data) {
    return request({
        url: '/api/article',
        method: 'post',
        data
    })
}