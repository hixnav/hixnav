import request from '@/utils/request'

export function addLink(data) {
  return request({
    url: '/api/addLink',
    method: 'post',
    data
  })
}

export function editLink(data) {
  return request({
    url: '/api/editLink',
    method: 'post',
    data
  })
}
