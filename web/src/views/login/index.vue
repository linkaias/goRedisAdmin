<template>
  <div>
    <div class="login" :class="{'loading':showLogin,'active':showButton}">
      <div class="form">
        <h2>GoRedisAdmin Login</h2>
        <div class="form-field">
          <label for="login-mail"><i class="el-icon-user"></i></label>
          <input id="login-mail" v-model="user" type="text" placeholder="User" required>
        </div>
        <div class="form-field">
          <label for="login-password"><i class="el-icon-key"></i></label>
          <input id="login-password" v-model="pwd" type="password" name="password" placeholder="Password"
                 pattern=".{6,}" required>
        </div>
        <button @click="login" type="button" class="button">
          <div class="arrow-wrapper">
            <span class="arrow"></span>
          </div>
          <p class="button-text">SIGN IN</p>
        </button>
      </div>
      <div class="finished">
        <svg>
          <use href="#svg-check"/>
        </svg>
      </div>
    </div>

    <!-- //--- ## SVG SYMBOLS ############# -->
    <svg style="display:none;">
      <symbol id="svg-check" viewBox="0 0 130.2 130.2">
        <polyline points="100.2,40.2 51.5,88.8 29.8,67.5 "/>
      </symbol>
    </svg>
  </div>
</template>

<script>
import {SetToken} from "@/utils/token";

export default {
  name: "login",
  data() {
    return {
      showLogin: false,
      showButton: false,
      user: "admin",  // 部署后请设置为空：user: ""
      pwd: "123456",  // 部署后请设置为空：pwd: ""
    }
  },
  methods: {
    async login() {
      if (this.user === "" || this.pwd === "") {
        this.$message.error("用户或密码不能为空！")
        return false
      }
      //请求
      let data = {
        "user": this.user,
        "pwd": this.pwd
      }
      let res = await this.$API.dbApi.reqLogin(data)
      if (res.code === 0) {
        this.showLogin = true
        SetToken(res.data)
        setTimeout(() => {
          this.showButton = true
          location.reload()
        }, 1200)
      }
    }
  }
}
</script>

<style scoped>
@import "./css/index.css";
</style>
