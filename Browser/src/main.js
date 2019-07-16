import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
// import Login from './components/Login.vue'//登录页
//组件import放下面
import iView from 'iview'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'


Vue.config.productionTip = false
// var fly=require("flyio")
// Vue.use(fly)
//组件添加在这里
Vue.use(iView)
//element
Vue.use(ElementUI)

//建议不要把主页直接绑到登陆上，可以用路由跳转。
//一个路由对应到页面。不要修改这里
//如何添加功能参见router.js
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

//登录页
// new Vue({
//   router,
//   store,
//   render: h => h(Login)
// }).$mount('#app')
