<template>
  <div id="article">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
      :activeIndex1="activeIndex1"
    />
    <div class="bg-banner">
      <el-row :gutter="16" style="margin-left: 0px; margin-right: 0px">
      </el-row>
    </div>
    <!-- 这里是内容 -->
    <div class="main">
      <el-row class="el-row" :gutter="20">
        <section style="">
          <h3>文档</h3>
        </section>
        <el-col :span="24" style="padding: 20px 0">
          <el-link href="https://element.eleme.io" target="_blank"
            >Golang相关</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Linux相关</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Nginx</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Mysql</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >JavaScript</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Element</el-link
          >
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <div
          style="box-shadow: 5px 5px 5px 5px rgba(0, 0, 0, 0.12); padding: 20px"
        >
          <el-descriptions
            class="margin-top"
            title=""
            :colon="false"
            column="3"
          >
            <template slot="extra"> </template>
            <el-descriptions-item v-for="i in 20" :key="i">
              <i class="el-icon-search"></i>
              <el-link href="https://element.eleme.io" target="_blank"
                >百度一下</el-link
              >
            </el-descriptions-item>
          </el-descriptions>
        </div>
      </el-row>
      <el-row class="el-row" :gutter="20">
        <section style="">
          <h3>链接</h3>
        </section>
        <el-col :span="24" style="padding: 20px 0">
          <el-link href="https://element.eleme.io" target="_blank"
            >Golang相关</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Linux相关</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Nginx</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Mysql</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >JavaScript</el-link
          >
          <el-link href="https://element.eleme.io" target="_blank"
            >Element</el-link
          >
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <div
          style="box-shadow: 5px 5px 5px 5px rgba(0, 0, 0, 0.12); padding: 20px"
        >
          <el-descriptions
            class="margin-top"
            title=""
            :colon="false"
            column="3"
          >
            <template slot="extra"> </template>
            <el-descriptions-item v-for="i in 20" :key="i">
              <i class="el-icon-search"></i>
              <el-link href="https://element.eleme.io" target="_blank"
                >百度一下</el-link
              >
            </el-descriptions-item>
          </el-descriptions>
        </div>
      </el-row>
    </div>
    <!-- 这里内容结束 -->
    <!-- 这里是弹出层 -->
    <div class="drawer">
      <el-drawer
        title="添加文链"
        :visible.sync="dialog"
        direction="rtl"
        custom-class="demo-drawer"
        ref="drawer"
      >
        <div class="demo-drawer__content" style="padding: 30px">
          <el-form :model="form">
            <el-form-item label="名称" :label-width="formLabelWidth">
              <el-input
                id="linked_name"
                v-model="form.name"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="地址" :label-width="formLabelWidth">
              <el-input
                id="linked_addr"
                v-model="form.addr"
                autocomplete="off"
              ></el-input>
            </el-form-item>
          </el-form>
          <div class="demo-drawer__footer" style="float: right">
            <el-button @click="cancelForm">取 消</el-button>
            <el-button type="primary" @click="linkedSubmit">确 定</el-button>
          </div>
        </div>
      </el-drawer>
    </div>
    <!-- 这里弹出层结束 -->
    <FootBar />
  </div>
</template>

<script>
import HeadBar from "@/components/HeadBar.vue";
import FootBar from "@/components/FootBar.vue";
export default {
  name: "ArticleView",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    return {
      searchVal: "",
      activeIndex: "1",
      activeIndex2: "1",
      // 弹出层控制
      dialog: false,
      formLabelWidth: "80px",
      form: {
        name: "",
        addr: "",
      },
    };
  },
  method: {
    cancelForm() {
      this.dialog = false;
      clearTimeout(this.timer);
    },
    linkedSubmit() {
      this.axios
        .post("/addArticleLink", JSON.stringify(this.form))
        .then(function (response) {
          console.log(response);
          this.$refs.drawer.closeDrawer();
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
.main {
  padding: 20px 30px;
}

.main > .el-row {
  display: inline-block;
  margin: 20px 0;
}
.bg-banner {
  width: 100%;
  height: 20rem;
  background: url("https://haixin-1300602599.cos.ap-guangzhou.myqcloud.com/t01fce1f99b60abfba6.jpg")
    no-repeat center center;
  background-color: #56ccf2;
  /* fallback for old browsers */
  background-color: -webkit-linear-gradient(to right, #2f80ed, #56ccf2);
  /* Chrome 10-25, Safari 5.1-6 */
  background-color: linear-gradient(to right, #2f80ed, #56ccf2);
  /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
  background-size: cover;
}
</style>