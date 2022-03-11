import request from '@/utils/request'

export function listIO(data) {
  return request({
    url: '/api/listIO',
    method: 'post',
    data
  })
}

export function uploadIO(data) {
  return request({
    url: '/api/uploadIO',
    headers: { 'Content-Type': 'multipart/form-data' },
    method: 'post',
    data
  })
}