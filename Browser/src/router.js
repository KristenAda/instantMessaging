import Vue from 'vue'
import Router from 'vue-router'
import home from './views/Home.vue'
//页面组件引入
import Login from './components/Login.vue'//登录页

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  /**
   * 在这个下面添加一个路由，对应到一个页面，首页可以直接用/访问
   * 我将登录页迁移至这里。作为例子
   */
  routes: [
    /**
     * 作为例子的登陆页，可以直接/访问
     */
    {
      path:'/',
      name:'login',
      component:Login
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path:'/home',
      name:'home',
      component:home
    }
  ]
})
