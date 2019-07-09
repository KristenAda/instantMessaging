<template>
  <div>
    <!--<Input v-model="value" placeholder="Enter something..." style="width: 300px" />-->
    <el-input type="text"  v-model="value" placeholder="Enter something..." style="width:200px"/>
    <!--<Button type="success" long @click="request">SUBMIT</Button>-->
    <el-button type="button" long @click="request">SUBMIT</el-button>
    <ul>
      <li v-for="chat in this.$store.state.chatRecord" :key="chat" :value="chat">{{chat}}</li>
    </ul>
  </div>
</template>

<script>
import fly from 'flyio'
export default {
  name: 'HelloWorld',
  data(){
    return{
      value:'',
      ws:{}
    }
  },
  created(){
    let that = this;
    that.ws = new WebSocket('ws://localhost:9999/ws');
    let w = that.ws;
    let c = 0;
    w.onopen = function () {
      console.log('ws onopen');
      w.send('from client: hello');
      for(let i = 0; i < 10; i ++){
        w.send('想发什么发什么' + i);
      }
    };
    w.onmessage = function (e) {
      console.log('ws onmessage');
      console.log('from server: ' + e.data);
      let arr = that.$store.state.chatRecord;
      arr.push(e.data);
      that.$store.commit('chatRecord',arr);
      c++;
      if(c<10){
        w.send('再次回信'+c+'哈哈哈');
      }
    };
    w.onclose = function(e){
      console.log('ws onclose');
      console.log('close con:' + e.reason);
    }
    w.onerror = function(e){
      console.log('ws onerror');
      console.log('err con:' + e);
    }
  },
  methods:{
    request(){
      fly.post('http://localhost:8888/user',{
        username:'zhangxiaojiao',
        password:'zxj'
      }).then((result) => {
        console.log(result.data);
      }).catch((err) => {
        console.log(err);
      });
      this.ws.close();
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
