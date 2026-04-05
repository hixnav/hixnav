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

export function listUser(data) {
  return request({
    url: '/api/user',
    method: 'post',
    data

  }).catch((error) => {
    console.log(error)
  })
}
export function addUser(data) {
  return request({
    url: '/api/user/add',
    method: 'post',
    data

  }).catch((error) => {
    console.log(error)
  })
}
export function delUser(data) {
  return request({
    url: '/api/user/del',
    method: 'post',
    data

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