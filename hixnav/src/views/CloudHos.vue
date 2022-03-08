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
          default-active="1-4-1"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          :collapse="isCollapse"
        >
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">文件</span>
          </el-menu-item>
           <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">笔记</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">脑图</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">视频</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">图片</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">保险箱</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-header>
          <el-row :span="24">
            <el-menu
              :default-active="activeIndex"
              class="el-menu-demo"
              mode="horizontal"
              @select="handleSelect"
              text-color="#303133"
              active-text-color="#909399"
            >
              <el-submenu index="10">
                <template slot="title">节点</template>
                <el-menu-item index="2-1">TeCentCOS</el-menu-item>
              </el-submenu>
            </el-menu>
          </el-row>
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
    return [];
  },
  method: {},
};
</script>

<style scoped>
.cloud-hos{
  min-height: 100vh;
}
</style>