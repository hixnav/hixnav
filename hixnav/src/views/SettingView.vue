<template>
  <div id="setting">
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
            <span slot="title">博客配置</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">数据库配置</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <span slot="title">云存储配置</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-folder"></i>
            <!-- 导出数据库 导出云存储 -->
            <span slot="title">数据迁移</span>
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
  name: "SettingView",
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