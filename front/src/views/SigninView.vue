<template>
  <div class="signin">
    <el-card class="signin_card">
      <div class="signin_logo">
        <router-link to="/"
          ><img class="signin_logo" src="../assets/logo.png" alt=""
        /></router-link>
      </div>
      <div class="signin_title">Sign in with your account</div>
      <div class="signin_tip">
        Please enter the correct account and password and log in
      </div>
      <div>
        <el-form
          ref="loginForm"
          :model="loginForm"
          :rules="loginRules"
          class="login-form"
          auto-complete="on"
          label-position="left"
        >
          <el-form-item prop="username">
            <el-input
              ref="username"
              type="text"
              name="username"
              v-model="loginForm.username"
              placeholder="请输入账号"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :key="passwordType"
              ref="password"
              name="password"
              :type="passwordType"
              v-model="loginForm.password"
              @keyup.enter.native="handleLogin"
              placeholder="请输入密码"
            >
              <template slot="suffix"
                ><i class="el-icon-view" @click="showPwd"></i></template
            ></el-input>
          </el-form-item>
        </el-form>
      </div>
      <div style="padding-top: 12px">
        <el-button
          :loading="loading"
          type="primary"
          style="width: 100%"
          @click.native.prevent="handleLogin"
          >Sign In</el-button
        >
      </div>
      <div class="back">
        <router-link to="/">
          <i class="el-icon-s-home"></i> return to home page</router-link
        >
      </div>
    </el-card>
  </div>
</template>
<script>
export default {
  name: "signin",
  data() {
    const validateUsername = (rule, value, callback) => {
      if (value.length < 1) {
        callback(new Error("Please enter the correct user name"));
      } else {
        callback();
      }
    };
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error("The password can not be less than 6 digits"));
      } else {
        callback();
      }
    };
    return {
      loginForm: {
        username: "",
        password: "",
      },
      loginRules: {
        username: [
          { required: true, trigger: "blur", validator: validateUsername },
        ],
        password: [
          { required: true, trigger: "blur", validator: validatePassword },
        ],
      },
      loading: false,
      passwordType: "password",
      redirect: undefined,
    };
  },
  methods: {
    showPwd() {
      if (this.passwordType === "password") {
        this.passwordType = "";
      } else {
        this.passwordType = "password";
      }
      this.$nextTick(() => {
        this.$refs.password.focus();
      });
    },
    handleLogin() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          this.$store
            .dispatch("user/login", this.loginForm)
            .then((res) => {
              console.log(this.redirect);
              this.$router.push({ path: this.redirect || "/" });
              this.loading = false;
              // this.$message({
              //   message: "登陆成功",
              //   type: "success",
              // });
            })
            .catch((res) => {
              console.log(res);
              // this.$message.error("登陆失败");
              this.loading = false;
            });
        } else {
          // console.log("error submit!!");
          this.$message.error("登陆失败");
          return false;
        }
      });
    },
  },
  watch: {
    $route: {
      handler: function (route) {
        this.redirect = route.query && route.query.redirect;
      },
      immediate: true,
    },
  },
};
</script>

<style>
body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
}

.signin {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  overflow: hidden;

  display: flex;
  width: 100%;
  height: 100%;
  align-items: center;
  flex: 1 1 0%;
  flex-direction: column;
  justify-content: center;
  background-color: #f5f7f9;
}

.signin_card {
  background-color: rgb(255, 255, 255);
  border-radius: 6px;
  margin: 16px;
  padding: 40px;
  width: 300px;
  text-align: center;
}

.el-card__body {
  padding: 0 0 20px;
}

.signin_card div {
  width: 100%;
  margin-top: 10px;
}

.signin_logo {
  width: 200px;
  margin: 0 auto;
  margin-bottom: 10px;
}

.signin_title {
  font-size: 24px;
  font-weight: 700;
}

.signin_tip {
  font-size: 16px;
  font-weight: 300;
  color: darkgrey;
}
.back {
  font-size: 14px;
  color: darkgrey;
  padding-top: 20px;
}
.back a {
  color: darkgrey;
}
a {
  text-decoration: none;
}
</style>
