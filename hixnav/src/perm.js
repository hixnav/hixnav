import { getToken } from '@/utils/auth' 
import router from './router'
import store from './store'
import NProgress from 'nprogress'

const whiteList = ['/signin','/']

router.beforeEach(async(to, from, next) => {
    NProgress.start()
    const hasToken = getToken()
    if (hasToken) {
      if (to.path === '/signin') {
        next({ path: '/' })
        NProgress.done()
      } else {
        const userInfo = store.getters.name
        if (userInfo) {
          next()
        } else {
          try {
            await store.dispatch('user/getInfo')
            next()
          } catch (error) {
            await store.dispatch('user/resetToken')
            Message.error(error || 'Has Error')
            next(`/signin?redirect=${to.path}`)
            NProgress.done()
          }
        }
      }
    } else {
      if (whiteList.indexOf(to.path) !== -1) {
        next()
      } else {
        next(`/signin?redirect=${to.path}`)
        NProgress.done()
      }
    }
  })
  
  router.afterEach(() => {
    NProgress.done()
  })