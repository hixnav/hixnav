<template>
  <div id="cloud-hos">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
      :activeIndex1="activeIndex1"
    />
    <el-container>
      <el-aside width="200px">
        <el-menu
          :default-active="param.active"
          class="el-menu-vertical-demo"
        >
          <el-menu-item index="1">
            <i class="el-icon-folder"></i>
            <span slot="title">文件</span>
          </el-menu-item>
          <el-menu-item index="2">
            <i class="el-icon-folder"></i>
            <span slot="title">笔记</span>
          </el-menu-item>
          <el-menu-item index="3">
            <i class="el-icon-folder"></i>
            <span slot="title">脑图</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">视频</span>
          </el-menu-item>
          <el-menu-item index="5">
            <i class="el-icon-folder"></i>
            <span slot="title">图片</span>
          </el-menu-item>
          <!-- <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">保险箱</span>
          </el-menu-item> -->
        </el-menu>
      </el-aside>
      <el-container>
        <el-header>
          <el-upload
            class="upload-demo"
            action="/api/uploadIO"
            :data="param"
            show-file-list="false"
          >
            <el-button size="small" type="primary">点击上传</el-button>
          </el-upload>
        </el-header>
        <el-main>
          <!-- 操作 -->
          <!-- 内容 -->
        </el-main>
      </el-container>
    </el-container>
    <!-- <FootBar /> -->
  </div>
</template>

<script>
import HeadBar from "@/components/HeadBar.vue";
import FootBar from "@/components/FootBar.vue";
export default {
  name: "CloundHos",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    return {
      param:{
        active:"1", //默认上传文件类型
      }
    };
  },
  method: {},
};
</script>

<style scoped>
.cloud-hos {
  min-height: 100vh;
}
</style>