<template>
  <div id="cloud-hos">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <el-container>
      <el-aside width="200px">
        <el-menu
          :default-active="param.active"
          class="el-menu-vertical-demo"
          @select="handleSelect"
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
            class="upload-file"
            :name="file"
            action="/api/uploadIO"
            :data="param"
            :limit="1"
            show-file-list="false"
            :auto-upload="true"
            :http-request="uploadIO"
            :with-credentials="true"
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
                <el-table-column prop="key" label="文件名" width="280">
                </el-table-column>
                <el-table-column prop="size" label="文件大小">
                </el-table-column>
                <el-table-column fixed="right" label="操作" width="120">
                  <template slot-scope="scope">
                    <el-button
                      type="text"
                      size="small"
                      @click="downFile(scope.row)"
                      >下载</el-button
                    >
                    <el-button
                      type="text"
                      size="small"
                      @click="delFile(scope.row)"
                      >删除</el-button
                    >
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
let Base64 = require("js-base64").Base64;
export default {
  name: "CloundHos",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    return {
      searchVal: "",
      activeIndex: "3",
      file: "file",
      param: {
        active: "1", //默认上传文件类型
      },

      fileLists: [],
    };
  },
  methods: {
    downFile(row) {
      row.active = this.param.active;
      console.log(row);
      this.$store
        .dispatch("cloud/downIO", row)
        .then((response) => {
          console.log(response);
          // let blob = Base64.decode(response.data.file);
          // let bdata = [];
          // bdata.push(blob);

          let raw = window.atob(response.data.file);
          let rawLength = response.data.contentLength;
          let uInt8Array = new Uint8Array(rawLength);

          for (let i = 0; i < rawLength; ++i) {
            uInt8Array[i] = raw.charCodeAt(i);
          }
          let objUrl = window.URL.createObjectURL(
            new Blob([uInt8Array], { type: response.data.contentType })
          );

          let a = document.createElement("a");
          a.href = objUrl;
          a.download = row.key;
          a.style.display = "none";
          a.click();
          a.remove();
        })
        .catch((res) => {
          console.log(res);
        });
    },
    base64ToBlob: function (code) {
      let parts = code.split(";base64,");
      let contentType = parts[0].split(":")[1];
      let raw = window.atob(parts[1]);
      let rawLength = raw.length;

      let uInt8Array = new Uint8Array(rawLength);

      for (let i = 0; i < rawLength; ++i) {
        uInt8Array[i] = raw.charCodeAt(i);
      }
      return new Blob([uInt8Array], { type: contentType });
    },
    delFile(row) {
      row.active = this.param.active;
      console.log(row);
      this.$store
        .dispatch("cloud/delIO", row)
        .then((response) => {
          console.log(response);
          this.$notify({
            title: "成功",
            message: "操作成功",
            type: "success",
          });
          location.reload();
        })
        .catch((res) => {
          console.log(res);
          this.$message.error("操作失败");
        });
    },
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
      this.param.active = key;
      this.listFile();
    },
    listFile: function () {
      let self = this;
      this.$store
        .dispatch("cloud/listIO", this.param)
        .then((response) => {
          console.log(response);
          self.fileLists = response.data.files;
        })
        .catch((res) => {
          console.log(res);
        });
    },
    uploadIO: function (params) {
      let self = this,
        file = params.file,
        formData = new window.FormData();
      // fileType = file.type,
      // isImage = fileType.indexOf('image') != -1,

      formData.append(self.file, file);
      formData.append("active", this.param.active);
      this.$store
        .dispatch("cloud/uploadIO", formData)
        .then((response) => {
          console.log(response);
          self.$notify({
            title: "成功",
            message: "上传成功",
            type: "success",
          });
          location.reload();
        })
        .catch((res) => {
          console.log(res);
          this.$message.error("上传失败");
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