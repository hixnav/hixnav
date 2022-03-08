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
        active-text-color="#909399"
      >
        <el-menu-item index="1">
          <router-link to="/">导航</router-link>
        </el-menu-item>
        <el-menu-item index="2">
          <router-link to="/article">快链</router-link>
        </el-menu-item>
        <el-menu-item index="3">
          <router-link to="/cloud-HOS">云存储</router-link>
        </el-menu-item>
        <el-menu-item index="4">
          <router-link to="/cloud-image">云图</router-link>
        </el-menu-item>
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
        <!-- <el-menu-item> 登陆 </el-menu-item> -->
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
      // console.log(keyPath);
      // if (key == 1) {
      //   location.href = "./article";
      // }
      // if (key == 2) {
      //   location.href = "./cloud-HOS";
      // }
      // if (key == 3) {
      //   location.href = "./cloud-image";
      // }
      // if (key == 4) {
      //   location.href = "./docs";
      // }
      // if (key == "2-0") {
      //   location.href = "https://wennmu.github.io";
      // }
      // if (key == "2-1") {
      //   location.href = "./add-link";
      // }
      // if (key == "2-2") {
      //   location.href = "./article?dialog";
      // }
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.head_bar {
  margin: 0;
  padding: 0;
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
