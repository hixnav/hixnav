<template>
  <div id="cloud-image">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
      :activeIndex1="activeIndex1"
    />
    <!-- 云图上传部分 -->
    <section class="ci-main">
      <el-row>
        <el-col :span="8" :offset="8">
          <el-upload
            class="upload-demo"
            drag
            :auto-upload="false"
            :file-list="fileList"
            list-type="picture"
            action="/api/upload"
            multiple
            :limit="3"
            ref="upload"
            :on-success="onSuccess"
          >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
              将文件拖到此处，或<em>点击上传</em>
            </div>

            <div class="el-upload__tip" slot="tip">
              <span>只能上传jpg/png文件，且不超过500kb</span>
              <el-button
                style="margin-left: 10px"
                size="small"
                type="primary"
                @click="submitUpload"
              >
                开始上传
              </el-button>
            </div>
          </el-upload>
          <div style="margin-top: 15px">
            <el-input
              placeholder="http://"
              v-for="(item, index) in cloudImages"
              :key="item"
              v-model="item.file"
              ref="demoInput"
            >
              <el-button
                slot="append"
                icon="el-icon-document-copy"
                @click="copyClicked(index)"
              ></el-button>
            </el-input>
          </div>
        </el-col>
      </el-row>
    </section>
    <FootBar />
  </div>
</template>

<script>
import HeadBar from "@/components/HeadBar.vue";
import FootBar from "@/components/FootBar.vue";
export default {
  name: "CloudImage",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    return {
      activeIndex: "1",
      searchVal:"",
      activeIndex1:"",
      fileList: [],
      cloudImages: [],
    };
  },
  methods: {
    submitUpload() {
      this.$refs.upload.submit();
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      console.log(file);
    },
    onSuccess(res, file) {
      console.log(res, file);
      this.cloudImages.push(res);
    },
    copyClicked(index) {
      this.$refs.demoInput[index].select();
      document.execCommand("copy");
    },
  },
};
</script>

<style scoped>
.ci-main {
  margin: 4rem 0;
  text-align: center;
}
</style>