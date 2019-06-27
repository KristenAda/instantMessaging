<template>
  <div id="app">
    <el-form
      :model="ruleForm2"
      :rules="rules2"
      status-icon
      ref="ruleForm2"
      label-position="left"
      label-width="0px"
      class="demo-ruleForm login-page"
    >
      <h3 class="title">QAQ-LOGIN</h3>
      <el-form-item prop="username">
        <el-input type="text" v-model="ruleForm2.username" auto-complete="off" placeholder="用户名"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="ruleForm2.password" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
      <el-checkbox v-model="checked" class="rememberme">记住密码</el-checkbox>
      <el-form-item style="width:100%;">
        <el-button
          type="danger"
          style="width:100%;"
          @click="handleSubmit"
          :loading="logining"
          icon="el-icon-upload"
        >登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  data() {
    return {
      logining: false,
      ruleForm2: {
        username: "tfg",
        password: "123456"
      },
      rules2: {
        username: [
          {
            required: true,
            message: "please enter your account",
            trigger: "blur"
          }
        ],
        password: [
          { required: true, message: "enter your password", trigger: "blur" }
        ]
      },
      checked: false
    };
  },
  methods: {
    handleSubmit(event) {
      this.$refs.ruleForm2.validate(valid => {
        if (valid) {
          this.logining = true;
          if (
            this.ruleForm2.username === "tfg" &&
            this.ruleForm2.password === "123456"
          ) {
            this.logining = false;
            this.$message({
              message: "登录成功,正在跳转!",
              type: "success"
            });
            sessionStorage.setItem("user", this.ruleForm2.username);
            this.$router.push({ path: "/Demo" });
          } else {
            this.logining = false;
            this.$alert("username or password wrong!", "info", {
              confirmButtonText: "ok"
            });
          }
        } else {
          console.log("error submit!");
          return false;
        }
      });
    }
  }
};
</script>
<style scoped>
.login-container {
  width: 100%;
  height: 100%;
}
.login-page {
  -webkit-border-radius: 5px;
  border-radius: 5px;
  margin: 180px auto;
  width: 350px;
  padding: 35px 35px 15px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 4px 4px 5px rgba(0, 0, 0, 0.2);
}
label.el-checkbox.rememberme {
  margin: 0px 0px 15px;
  text-align: left;
}
.bg {
  height: 100%;
  background-image: linear-gradient(to right, #7a88ff, #7affaf);
}
</style>
