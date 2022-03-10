import request from '@/utils/request'

export function addLink(data) {
  return request({
    url: '/api/addLink',
    method: 'post',
    data
  })
}
