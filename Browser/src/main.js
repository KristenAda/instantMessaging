import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Login from './components/Login.vue'//登录页
import Demo from './components/Demo.vue'//登录页
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
// new Vue({
//   router,
//   store,
//   render: h => h(App)
// }).$mount('#app')

//登录页
new Vue({
  router,
  store,
  render: h => h(Demo)
}).$mount('#app')
