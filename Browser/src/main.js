import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
//组件import放下面
import iView from 'iview'

Vue.config.productionTip = false
// var fly=require("flyio")
// Vue.use(fly)
//组件添加在这里
Vue.use(iView)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
