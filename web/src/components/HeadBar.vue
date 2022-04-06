<template>
  <div class="head_bar">
    <el-row :span="24">
      <el-col :span="6">
        <a href="/" style="text-decoration: none; color: #000">
          <div style="padding: 0 10px 0 30px; float: left; align: center">
            <div class="logo"></div>
          </div>
          <div
            style="
              line-height: 60px;
              font-weight: 800;
              font-size: 20px;
              padding: 0 60px;
              color: #333333;
            "
          >
            <!-- <span>{{ title }}</span> -->
          </div>
        </a>
      </el-col>
      <el-col :span="6" style="padding: 14px 10px 0 30px; float: left">
        <!-- <el-input
          placeholder="搜索关键词"
          :value="searchVal"
          class="input-with-select"
        />
        <el-button slot="append" icon="el-icon-search"></el-button> -->
      </el-col>
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        style="float: right; text-decoration: none"
        text-color="#303133"
        active-text-color="#409EFF"
        @select="handleSelect"
      >
        <el-menu-item index="1"> 导航 </el-menu-item>
        <el-menu-item index="2"> 快链 </el-menu-item>
        <el-menu-item index="3"> 云存储 </el-menu-item>
        <el-menu-item index="4"> 云图 </el-menu-item>
        <el-menu-item index="5"> 云账号 </el-menu-item>
        <!-- <el-menu-item index="4">
          <router-link to="/docs">云笔记</router-link>
        </el-menu-item> -->
        <!-- <el-submenu index="10"> -->
        <!-- <template slot="title">我的</template> -->
        <!-- <el-menu-item index="2-0">博客</el-menu-item> -->
        <!-- <el-menu-item index="2-1"
            ><router-link to="/add-link">添加导航</router-link></el-menu-item
          > -->
        <!-- <el-menu-item index="2-2"
            ><router-link to="/settting"
              >设置中心</router-link
            ></el-menu-item
          > -->
        <!-- </el-submenu> -->
        <el-menu-item v-if="!this.$store.state.user.signin" index="9"
          >登陆</el-menu-item
        >
        <el-menu-item v-else index="99">
          <el-submenu>
            <template slot="title">我的</template>
            <el-menu-item index="2-0"
              ><router-link to="/setting">设置中心</router-link></el-menu-item
            >
            <el-menu-item index="10">退出</el-menu-item>
          </el-submenu>
        </el-menu-item>
      </el-menu>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "HeadBar",
  props: {
    title: String,
    searchVal: String,
    activeIndex: String,
    activeIndex2: String,
  },
  methods: {
    handleSelect: function (key, keyPath) {
      this.activeIndex = key;
      // console.log(keyPath);
      if (key == 1) {
        this.$router.push("/");
      }
      if (key == 2) {
        this.$router.push("/article");
      }
      if (key == 3) {
        this.$router.push("/cloud-HOS");
      }
      if (key == 4) {
        this.$router.push("/cloud-image");
      }
      if (key == 5) {
        this.$router.push("/cloud-account");
      }
      // if (key == "2-0") {
      //   location.href = "https://wennmu.github.io";
      // }
      // if (key == "2-1") {
      //   location.href = "./add-link";
      // }
      // if (key == "2-2") {
      //   location.href = "./article?dialog";
      // }
      if (key == 9) {
        this.$router.push("/signin");
      }
      if (key == 10) {
        this.logout();
      }
    },
    logout: function () {
      this.$store
        .dispatch("user/logout")
        .then((response) => {
          this.$router.go(0);
        })
        .catch((res) => {
          console.log(res);
        });
    },
  },
  created: function () {
    console.log(this.$store.state.nav.activeIndex);
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.head_bar {
  margin: 0;
  padding: 0;
  border-bottom: solid 1px #e6e6e6;
}
.logo {
  width: 132px;
  height: 30px;
  padding-top: 12px;
  margin-top: 10px;
  background-image: url("../assets/logo.png");
  background-size: cover;
}
a {
  text-decoration: none;
}
</style>
