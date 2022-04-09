<template>
  <div id="docs">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
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
            <span slot="title">Golang笔记</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">Linux笔记</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">开源项目</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">必会宝典</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
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
            <span slot="title">Golang笔记</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">Linux笔记</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">开源项目</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">必会宝典</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        
        <el-main>
          <!-- 操作 -->
          <!-- 内容 -->
        </el-main>
      </el-container>
    </el-container>
    <FootBar />
  </div>
</template>

<script>
import HeadBar from "@/components/HeadBar.vue";
import FootBar from "@/components/FootBar.vue";
export default {
  name: "Docs",
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
</style>