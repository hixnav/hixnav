<template>
  <div class="signin">
    <el-card class="signin_card">
      <div class="signin_logo">
        <img class="signin_logo" src="../assets/logo.png" alt="" />
      </div>
      <div class="signin_title">Sign in with your account</div>
      <div class="signin_tip">Please enter the correct account and password and log in

</div>
      <div>
        <el-input v-model="account" placeholder="请输入账号"></el-input>
      </div>
      <div>
        <el-input v-model="passport" placeholder="请输入密码"></el-input>
      </div>
      <div><el-button type="primary" style="width:100%">Sign In</el-button></div>
      <div class="back"><router-link to="/"> <i class="el-icon-s-home"></i>  return to home page</router-link></div>
    </el-card>
  </div>
</template>
<style>
body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  background-color: #f5f7f9;
  overflow: hidden;
}

.signin {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  display: flex;
  width: 100%;
  height: 100%;
  align-items: center;
  flex: 1 1 0%;
  flex-direction: column;
  justify-content: center;
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
  margin-bottom:10px;
}

.signin_title {
  font-size: 24px;
  font-weight: 700  ;
}

.signin_tip {
  font-size: 16px;
  font-weight: 300;
  color: darkgrey;
}
.back{
  font-size: 14px;
  color: darkgrey;
  padding-top: 20px;
}
.back a{
  color: darkgrey;
}
a {
    text-decoration: none;
}
</style>
