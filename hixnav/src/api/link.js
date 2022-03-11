import request from '@/utils/request'

export function addArticleLink(data) {
  return request({
    url: '/api/addArticleLink',
    method: 'post',
    data
  })
}
