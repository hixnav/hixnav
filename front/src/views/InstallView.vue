<template>
  <div class="install">
    <div class="install-main">
      <a href="/" style="text-decoration: none; color: #000">
        <img class="install-logo" src="../assets/logo.png" alt="logo">
      </a>
      <h3 class="main-label">配置数据库</h3>
      <el-form ref="form" :model="config" label-width="100px">
        <el-form-item label="数据库地址">
          <el-input v-model="config.addr"></el-input>
        </el-form-item>
        <el-form-item label="数据库名称">
          <el-input v-model="config.dbname"></el-input>
        </el-form-item>
        <el-form-item label="数据库账号">
          <el-input v-model="config.account"></el-input>
        </el-form-item>
        <el-form-item label="数据库密码">
          <el-input type="password" v-model="config.password"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" size="small" @click="onSubmit(true)">保存</el-button>
          <a href="/" style="margin: 0 10px"><el-button type="plain" size="small">回到首页</el-button></a>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "InstallView",
  data: function () {
    return {
      config: {
        addr:"127.0.0.1:3306",
        dbname:"hixnav",
        account:"root",
        isMigrate: false,
      },
    }
  },
  methods: {
    onSubmit(isMigrate) {
      this.config.isMigrate = isMigrate
      this.$store
          .dispatch("config/dbMigrate", JSON.stringify(this.config))
          .then((response) => {
            console.log(response);
            if (response.code == 0){
              this.$message.success("success");
              this.$router.push("/")
            }
          })
          .catch((res) => {
            console.log(res);
          });
    }
  },
}
</script>

<style scoped>
.install {
  display: flex;
  justify-content: center;
  width: 100%;
  height: 100vh;
}

.install-logo{
  width: 200px;
  margin-top: 80px;
}

.install-main {
  min-width: 600px;
}

.main-label{
  margin: 30px 0 20px 8px;
  font-weight: 400;
  color: #1f2f3d;
  font-size: 22px;
}
</style>