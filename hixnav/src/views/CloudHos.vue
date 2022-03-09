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
        <el-menu :default-active="param.active" class="el-menu-vertical-demo">
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
            class="upload-file"
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
          <el-card class="box-card" shadow="never">
            <template>
              <el-table :data="fileLists" style="width: 100%">
                <el-table-column prop="Key" label="文件名" width="240">
                </el-table-column>
                <el-table-column prop="Size" label="文件大小" width="180">
                </el-table-column>
                <el-table-column fixed="right" label="操作" width="100">
                  <template slot-scope="">
                    <el-button type="text" size="small">下载</el-button>
                    <el-button type="text" size="small">删除</el-button>
                  </template>
                </el-table-column>
              </el-table>
            </template>
          </el-card>
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
      searchVal: "",
      activeIndex: "1",
      activeIndex1: "1",
      param: {
        active: "1", //默认上传文件类型
      },

      fileLists: [],
    };
  },
  methods: {
    listFile: function () {
      let self = this;
      this.axios
        .post("/api/listIO", this.param)
        .then(function (response) {
          console.log(response);
          self.fileLists = response.data.files;
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
  created: function () {
    this.listFile();
  },
};
</script>

<style scoped>
.cloud-hos {
  min-height: 100vh;
}

.upload-file {
  line-height: 60px;
}

.text {
  font-size: 14px;
}

.item {
  padding: 10px 0;
}

.box-card {
  width: 98%;
}
</style>