import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    //聊天记录
    chatRecord:[]
  },
  mutations: {
    chatRecord(state,n){
      state.chatRecord = n;
    }
  },
  actions: {

  }
})
