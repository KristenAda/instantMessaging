<template>
  <div id="app">
    <el-container style="height: 500px; border: 1px solid #eee">
  <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
    <el-menu :default-openeds="['1', '3']">
      <el-submenu index="1">
        <template slot="title"><i class="el-icon-message"></i>好友</template>
        <el-menu-item-group>
          <template slot="title">分组一</template>
          <el-menu-item index="1-1">tfg</el-menu-item>
          <el-menu-item index="1-2">zxj</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="分组2">
          <el-menu-item index="1-3">sc</el-menu-item>
        </el-menu-item-group>
        <el-submenu index="1-4">
          <template slot="title">选项4</template>
          <el-menu-item index="1-4-1">选项4-1</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-submenu index="2">
        <template slot="title"><i class="el-icon-menu"></i>导航</template>
        <el-menu-item-group>
          <template slot="title">分组一</template>
          <el-menu-item index="2-1">选项1</el-menu-item>
          <el-menu-item index="2-2">选项2</el-menu-item>
        </el-menu-item-group>
        
      </el-submenu>
      <el-submenu index="3">
        <template slot="title"><i class="el-icon-setting"></i>设置</template>
        
      </el-submenu>
    </el-menu>
  </el-aside>
  
  <el-container>
    <el-header style="text-align: right; font-size: 12px;color:#fff">
      <el-dropdown>
        <i class="el-icon-setting" style="margin-right: 15px;color:#fff"></i>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>查看</el-dropdown-item>
          <el-dropdown-item>新增</el-dropdown-item>
          <el-dropdown-item>删除</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <span>tong</span>
    </el-header>
    
    <el-main>
      
    </el-main>
  </el-container>
</el-container>

  </div>
</template>

<style>
  .el-header {
    background-color: #f56c6c;
    color: #333;
    line-height: 60px;
  }
  .el-aside {
    color: #333;
  }
  /* .el-container{
    
  } */
  .el-menu{
    height: 100%;
  }
  #app{
    position: relative;
    height: 500px;
    top:100px;
    width: 70%;
    margin:0 auto;
    box-shadow: 2px 2px 10px rgba(0, 0, 0, 0.2)
  }
</style>

<script>
  export default {
    data() {
      const item = {
        date: '2019-06-27',
        name: '童',
        address: '成都'
      };
      return {
        tableData: Array(5).fill(item)
      }
    }
  };
</script>