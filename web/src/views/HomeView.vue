<template>
  <div class="home">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <div class="bg-banner"></div>
    <div class="line"></div>
    <el-row
      :gutter="16"
      style="
        margin: 0;
        padding-right: 3rem;
        width: 98%;
        margin: 0 auto;
        margin-top: 2rem;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
      "
    >
      <!-- 搜索 -->
      <div class="main-search">
        <el-row style="min-width: 400px; padding: 14px 50px 20px">
          <el-input
            placeholder="搜索关键词"
            :value="searchVal"
            class="input-with-select"
            suffix-icon="el-icon-search"
            size="large"
            style="border: none"
          />
          <div class="history">
            <div v-for="o in quikList" :key="o">
              <!-- <i class="el-icon-search"></i> -->
              <el-link :href="o.href" target="_blank" style="font-size: 10px"
                >{{ o.name }}
              </el-link>
            </div>
          </div>
        </el-row>
      </div>
      <!-- 标签 -->
      <div style="flex: 1; border-left: 1px solid #ccc; padding-left: 20px">
        <h3 style="margin-left: 10px">
          <svg
            t="1650003512498"
            class="icon"
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            p-id="2554"
            width="20"
            height="20"
            style="position: relative; top: 4px"
          >
            <path
              d="M194.889274 0v1024l317.082033-224.864391 317.139419 224.864391V0z"
              p-id="2555"
              fill="#1296db"
            ></path>
          </svg>
          标签
        </h3>
        <el-tag
          size="small"
          v-for="c in cates"
          :key="c"
          @click="goAnchor('#anchor' + c.Cate)"
          style="margin: 0.4rem 0.6rem; color: #fff; font-weight: 700"
          :style="{
            'background-color': colorlists[Math.floor(Math.random() * 10)],
          }"
          >{{ c.Catename }}
        </el-tag>
      </div>
    </el-row>
    <div>
      <el-row v-for="c in cates" :key="c">
        <section
          :id="'anchor' + c.Cate"
          style="
            display: flex;
            flex-direction: row;
            justify-content: flex-start;
            align-items: center;
            padding-left: 30px;
            padding-top: 2rem;
          "
        >
          <img
            src="../assets/tag-blue.png"
            alt=""
            style="
              display: inline-block;
              width: 10px;
              height: 20px;
              padding: 0 10px 0;
              margin-top: 2px;
            "
          />
          <div style="font-weight: 700; font-size: 18px">{{ c.Catename }}</div>
        </section>
        <el-row
          style="
            width: 100%;
            display: flex;
            flex-direction: row;
            flex-wrap: wrap;
            justify-content: flex-start;
            align-items: center;
            padding: 0 30px;
          "
        >
          <div v-for="o in navs[c.Cate]" :key="o" style="margin: 30px 20px">
            <el-card :body-style="{ padding: '0px', width: '20rem' }">
              <div style="padding: 14px">
                <el-row :gutter="24">
                  <el-col :span="4">
                    <el-avatar
                      style="background-color: #fff"
                      :src="o.Logo"
                      fit="cover"
                      shape="square"
                    ></el-avatar>
                  </el-col>
                  <el-col :span="12">
                    <div class="linked-title">{{ o.Name }}</div>
                    <div class="bottom">
                      <time class="time">{{ o.Desc }}</time>
                    </div>
                  </el-col>
                  <el-col
                    :span="8"
                    style="padding: 10px 0 0; text-align: center"
                  >
                    <div class="bottom clearfix">
                      <el-button
                        v-if="showBtn != true"
                        icon=""
                        size="mini"
                        circle
                        class="button"
                        @click="jumpTo(o.Url)"
                        icon="el-icon-arrow-right"
                      >
                      </el-button>
                      <el-button-group v-else>
                        <el-button
                          size="mini"
                          circle
                          class="button"
                          @click="jumpTo(o.Url)"
                          icon="el-icon-arrow-right"
                        >
                        </el-button>
                        <el-button
                          icon="el-icon-edit"
                          size="mini"
                          circle
                          @click="openDrawer(o)"
                        ></el-button>
                        <el-button
                          icon="el-icon-delete"
                          size="mini"
                          circle
                          @click="deleteNav(o.Id)"
                        ></el-button>
                      </el-button-group>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </el-card>
          </div>
        </el-row>
      </el-row>
    </div>
    <!-- 悬浮按钮 -->
    <div class="" style="position: fixed; right: 40px; bottom: 108px">
      <el-button-group>
        <router-link to="/add-link">
          <el-button type="primary" icon="el-icon-plus" circle></el-button>
        </router-link>
      </el-button-group>
    </div>
    <div class="" style="position: fixed; right: 40px; bottom: 60px">
      <el-button
        type="primary"
        icon="el-icon-edit-outline"
        circle
        @click="handOffBtn"
      ></el-button>
    </div>
    <!-- 悬浮按钮结束 -->
    <!-- 这里是弹出层 -->
    <div class="drawer">
      <el-drawer
        title="修改导航"
        :visible.sync="dialog"
        direction="rtl"
        custom-class="demo-drawer"
        ref="drawer"
      >
        <div class="demo-drawer__content" style="padding: 30px">
          <el-form
            :model="ruleForm"
            status-icon
            :rules="rules"
            ref="ruleForm"
            label-width="100px"
            class="demo-ruleForm"
          >
            <el-form-item label="标题" prop="name">
              <el-input
                type="input"
                v-model="ruleForm.name"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="图标" prop="logo">
              <el-input
                type="input"
                v-model="ruleForm.logo"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="简述" prop="desc">
              <el-input
                type="input"
                v-model="ruleForm.desc"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="地址" prop="url">
              <el-input
                type="input"
                v-model="ruleForm.url"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="分类ID" prop="cate">
              <el-input
                type="number"
                v-model="ruleForm.cate"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="分类名称" prop="catename">
              <el-input
                type="input"
                v-model="ruleForm.catename"
                autocomplete="off"
              ></el-input>
            </el-form-item>
          </el-form>
          <div class="demo-drawer__footer" style="float: right">
            <el-button @click="cancelForm">取 消</el-button>
            <el-button type="primary" @click="submitForm('ruleForm')"
              >确 定</el-button
            >
          </div>
        </div>
      </el-drawer>
    </div>
    <!-- 这里弹出层结束 -->
    <FootBar />
  </div>
  <!-- <div class="loadding" v-else>
    <div  v-loading="loading"></div>
  </div> -->
</template>

<script>
// @ is an alias to /src
import HeadBar from "@/components/HeadBar.vue";
import FootBar from "@/components/FootBar.vue";

export default {
  name: "HomeView",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    return {
      searchVal: "",
      activeIndex: "1",
      showBtn: false,
      // 弹出层控制
      dialog: false,
      ruleForm: {
        id: 0,
        name: "",
        desc: "",
        url: "",
        logo: "",
        cate: 0,
        catename: "",
      },
      navs: [],
      cates: [],
      colorlists: [
        "rgb(229, 0, 19)",
        "rgb(206,194,28)",
        "rgb(0,161,233)",
        "rgb(109,185,45)",
        "rgb(166,0,130)",
        "rgb(237,108,0)",
        "rgb(240, 28, 131)",
        "rgb(84, 21, 226)",
        "rgb( 0,128,0)",
        "rgb( 255,69,0)",
        "rgb( 255,165,0)",
        "rgb( 178,34,34)",
        "rgb( 255,0,255)",
        "rgb(65,105,225)",
        "blueviolet",
      ],
      quikList: [
        {
          href: "https://wennmu.github.io/",
          name: "我的博客",
        },
        {
          href: "https://www.baidu.com",
          name: "百度",
        },
        {
          href: "https://cn.bing.com/",
          name: "必应",
        },
        {
          href: "https://fanyi.baidu.com/",
          name: "百度翻译",
        },
        {
          href: "https://note.youdao.com/",
          name: "有道云笔记",
        },
        {
          href: "https://mail.google.com/",
          name: "Gmail",
        },
      ],
    };
  },
  methods: {
    openDrawer(navInfo) {
      this.ruleForm.id = navInfo.Id;
      this.ruleForm.name = navInfo.Name;
      this.ruleForm.desc = navInfo.Desc;
      this.ruleForm.url = navInfo.Url;
      this.ruleForm.logo = navInfo.Logo;
      this.ruleForm.cate = navInfo.Cate;
      this.ruleForm.catename = navInfo.Catename;
      this.dialog = true;
    },
    cancelForm() {
      this.dialog = false;
    },
    submitForm(formName) {
      self = this;
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(this.ruleForm);
          this.ruleForm.cate = parseInt(this.ruleForm.cate);
          this.$store
            .dispatch("nav/editLink", JSON.stringify(this.ruleForm))
            .then((response) => {
              console.log(response);
              self.$notify({
                title: "成功",
                message: "手动刷新页面查看",
                type: "success",
              });
              self.cancelForm();
              self.getData();
            })
            .catch((res) => {
              console.log(res);
              this.$message.error("失败");
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    deleteNav(id) {
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.deleteNavData(id);
        })
        .catch(() => {});
    },
    deleteNavData(id) {
      self = this;
      this.$store
        .dispatch("nav/delLink", "id=" + id)
        .then((response) => {
          console.log(response);
          self.$notify({
            title: "成功",
            message: "",
            type: "success",
          });
          self.getData();
        })
        .catch((res) => {
          console.log(res);
          this.$message.error("失败");
        });
    },
    handOffBtn() {
      console.log("打开操作" + this.showBtn);
      if (this.showBtn == false) {
        this.showBtn = true;
      } else {
        this.showBtn = false;
      }
    },
    jumpTo(url) {
      window.open(url);
    },
    goAnchor(selector) {
      document.querySelector(selector).scrollIntoView({
        behavior: "smooth",
        block: "center",
        inline: "nearest",
      });
    },
    getData() {
      let self = this;
      this.axios
        .post("/api/home")
        .then(function (response) {
          // console.log(response);
          self.navs = response.data.navs;
        })
        .catch(function (error) {
          console.log(error);
        });
      this.axios
        .post("/api/cates")
        .then(function (response) {
          // console.log(response);
          self.cates = response.data.cates;
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
  created: function () {
    let self = this;
    this.getData();
  },
  mounted: function () {
    this.getData();
    // this.$nextTick(function () {
    //   // 仅在整个视图都被渲染之后才会运行的代码
    // });
  },
};
</script>

<style scoped>
.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 10px;
  line-height: 12px;
  white-space: nowrap;
  overflow: hidden;
  padding-left: 6px;
}

.bottom time {
  font-size: 12px;
}

.bottom .el-button {
  font-size: 12px;
}

.button {
}

.linked-title {
  font-weight: 600;
  padding-left: 6px;
  overflow: hidden;
  white-space: nowrap;
}

.image {
  width: 100%;
  display: block;
}

.home {
  margin: 0px;
  padding: 0px;
}

.bg-banner {
  width: 100%;
  height: 400px;
  background: url("https://cdn.jsdelivr.net/gh/hututu-tech/IMG-gongfeng@main/2022/04/11/6253fd6c6f522.jpg")
    no-repeat center center;
  background-size: cover;
}

.main-search {
  width: 100%;
  height: 150px;
  margin: 20px 20px;
  flex: 3;
  /* background: url("http://pic.5tu.cn/uploads/allimg/2111/pic_5tu_big_6320187_7f6b7f0c6f9205bbc7fa6a1616b67864.jpg")
    no-repeat left bottom;
  background-size: cover; */
  background-color: #fff;
  border-radius: 5px;
}
.history {
  display: flex;
  flex-direction: row;
  align-items: left;
}
.history > div {
  /* width: 80px; */
  margin: 10px 20px 0 0;
}
</style>