import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/api/info',
    method: 'get',
    params: {token}
    
  }).catch((error) => {
    console.log(error)
  })
}

export function logout(token) {
  return request({
    url: '/api/logout',
    method: 'post',
    params: {token}
  })
}