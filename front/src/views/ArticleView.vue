<template>
  <div id="article">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <div class="bg-banner">
      <el-row :gutter="16" style="margin-left: 0px; margin-right: 0px">
      </el-row>
    </div>
    <!-- 这里是内容 -->
    <div class="main">
      <el-row class="el-row" :gutter="20">
        <section style="">
          <section
            style="
              display: flex;
              flex-direction: row;
              justify-content: flex-start;
              align-items: center;
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
            <div style="font-weight: 700; font-size: 18px">文档</div>
          </section>
        </section>
        <!-- <el-col :span="24" style="padding: 20px 0">
          <el-link href="https://element.eleme.io" target="_blank"
            >Element</el-link
          >
        </el-col> -->
      </el-row>
      <el-row :gutter="20">
        <div class="main_card">
          <el-descriptions
            class="margin-top"
            title=""
            :colon="false"
            column="4"
          >
            <template slot="extra"> </template>
            <el-descriptions-item v-for="o in docLinks" :key="o">
              <el-link :href="o.Url" target="_blank">
                <el-avatar
                  style="
                    width: 20px;
                    height: 20px;
                    margin: 0 10px;
                    background-color: #fff;
                  "
                  :src="o.Logo"
                  size="small"
                ></el-avatar>
                {{ o.Name }}</el-link
              >
              <el-button-group v-if="showBtn">
                <el-button class="operator-btn" type="text" icon="el-icon-edit" size="mini" circle @click="openDrawer(o)"></el-button>
                <el-button class="operator-btn" type="text" icon="el-icon-delete" size="mini" circle @click="deleteNav(o.Id)"></el-button>
              </el-button-group>
            </el-descriptions-item>
          </el-descriptions>
        </div>
      </el-row>
      <el-row class="el-row" :gutter="20">
        <el-col :span="2" style="">
          <section
            style="
              display: flex;
              flex-direction: row;
              justify-content: flex-start;
              align-items: center;
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
            <div style="font-weight: 700; font-size: 18px">链接</div>
          </section>
        </el-col>
        <el-col :span="20" style="padding: 20px 0">
          <el-row :gutter="20">
            <el-button
              size="small"
              round
              v-for="o in linkCates"
              :key="o"
              style="display: inline; margin: 0 20px"
              >{{ o.Catename }}</el-button
            >
          </el-row>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <div class="main_card">
          <el-descriptions
            class="margin-top"
            title=""
            :colon="false"
            column="4"
          >
            <template slot="extra"> </template>
            <el-descriptions-item v-for="o in commonLinks" :key="o">
              <el-link :href="o.Url" target="_blank">
                <el-avatar
                  style="
                    width: 20px;
                    height: 20px;
                    margin: 0 10px;
                    background-color: #fff;
                  "
                  :src="o.Logo"
                  size="small"
                ></el-avatar>
                {{ o.Name }}</el-link
              >
                  <el-button-group v-if="showBtn">
                    <el-button class="operator-btn" type="text" icon="el-icon-edit" size="mini" circle @click="openDrawer(o)"></el-button>
                    <el-button class="operator-btn" type="text" icon="el-icon-delete" size="mini" circle @click="deleteLink(o.Id)"></el-button>
                  </el-button-group>
            </el-descriptions-item>
          </el-descriptions>
        </div>
      </el-row>
    </div>
    <!-- 这里内容结束 -->
    <!-- 悬浮按钮 -->
    <!-- <div class=""   >  -->
    <el-button
      style="position: fixed; right: 40px; bottom: 108px; z-index: 9999"
      type="primary"
      icon="el-icon-plus"
      circle
      @click="openAddDrawer"
    ></el-button>
    <div class="" style="position: fixed; right: 40px; bottom: 60px">
      <el-button type="primary" icon="el-icon-edit-outline" circle @click="handOffBtn"></el-button>
    </div>
    <!-- </div> -->
    <!-- 悬浮按钮结束 -->
    <!-- 这里是弹出层 -->
    <div class="drawer">
      <el-drawer
          :title="showDialog === 'add' ? '添加文链' : '修改文链'"
          :visible.sync="dialog"
          direction="rtl"
          custom-class="demo-drawer"
          ref="drawer"
      >
        <div class="demo-drawer__content" style="padding: 30px">
          <el-form :model="form">
            <el-form-item label="类别" :label-width="formLabelWidth">
              <el-select v-model="form.type" placeholder="类别">
                <el-option label="链接" value="1"></el-option>
                <el-option label="文档" value="2"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="标签" :label-width="formLabelWidth">
              <el-input
                  id="linked_label"
                  v-model="form.catename"
                  autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="名称" :label-width="formLabelWidth">
              <el-input
                  id="linked_name"
                  v-model="form.name"
                  autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="地址" :label-width="formLabelWidth">
              <el-input id="linked_url" v-model="form.url"></el-input>
            </el-form-item>
          </el-form>
          <div class="demo-drawer__footer" style="float: right">
            <el-button @click="cancelForm">取 消</el-button>
            <el-button type="primary" v-if="showDialog === 'add'" @click="linkedSubmit">添 加</el-button>
            <el-button type="primary" v-else @click="linkedSave">保 存</el-button>
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
import {getToken} from "@/utils/auth";
export default {
  name: "ArticleView",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    return {
      searchVal: "",
      activeIndex: "2",
      showBtn: false,
      showDialog: "add",
      // 链接
      linkCates: [
        // {
        //   Catename: "办公",
        // },
        // {
        //   Catename: "博文",
        // },
      ],
      docLinks: [],
      commonLinks: [],
      // 弹出层控制
      dialog: false,
      formLabelWidth: "80px",
      form: {
        name: "",
        type: "1",
        url: "",
      },
    };
  },
  methods: {
    openDrawer(row) {
      this.form.id = row.Id;
      this.form.name = row.Name;
      this.form.type = row.Type;
      this.form.url = row.Url;
      this.form.Logo = row.Logo;
      this.form.catename = row.Catename;
      this.showDialog = "edit"
      this.dialog = true;
    },
    deleteLink(id) {
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
          .then(() => {
            this.deleteLinkData(id);
          })
          .catch(() => {
          });
    },
    deleteLinkData(id) {
      self = this;
      this.$store
          .dispatch("link/delArticleLink", "id=" + id)
          .then((response) => {
            this.getLinks()
            console.log(response);
            self.$message({
              message: "成功",
              type: "success",
            });
            self.getData();
          })
          .catch((err) => {
            console.log(err);
            // this.$message.error("失败");
          });
    },
    openAddDrawer() {
      this.showDialog = "add"
      this.dialog = true;
    },
    handOffBtn() {
      const hasToken = getToken()
      if (!hasToken) {
        this.$router.push("signin")
      }
      if (this.showBtn == false) {
        this.showBtn = true;
      } else {
        this.showBtn = false;
      }
    },
    cancelForm() {
      this.dialog = false;
    },
    linkedSubmit() {
      var url = this.form.url;
      var urlArr = url.split("/");
      var hostname = urlArr[2];
      this.form.logo = urlArr[0] + "//" + hostname + "/favicon.ico";
      this.form.type = parseInt(this.form.type);
      console.log(this.form);

      self = this;
      this.$store
        .dispatch("link/addArticleLink", JSON.stringify(this.form))
        .then((response) => {
          this.getLinks()
          console.log(response);
          self.dialog = false;
          self.$notify({
            title: "成功",
            message: "手动刷新页面查看",
            type: "success",
          });
        })
        .catch((res) => {
          console.log(res);
          this.$message.error("失败");
        });
    },
    linkedSave() {
      var url = this.form.url;
      var urlArr = url.split("/");
      var hostname = urlArr[2];
      this.form.logo = urlArr[0] + "//" + hostname + "/favicon.ico";
      this.form.type = parseInt(this.form.type);
      console.log(this.form);

      self = this;
      this.$store
          .dispatch("link/editArticleLink", JSON.stringify(this.form))
          .then((response) => {
            this.getLinks()
            console.log(response);
            self.dialog = false;
            self.$notify({
              title: "成功",
              message: "手动刷新页面查看",
              type: "success",
            });
          })
          .catch((res) => {
            console.log(res);
            this.$message.error("失败");
          });
    },
    getLinks() {
      self = this;
      this.$store
          .dispatch("link/article", { Type: 2 })
          .then((response) => {
            self.docLinks = response.data.links;
          })
          .catch((res) => {
            console.log(res);
          });
      this.$store
          .dispatch("link/article", { Type: 1, Catename: "" })
          .then((response) => {
            self.commonLinks = response.data.links;
          })
          .catch((res) => {
            console.log(res);
          });
    },
  },
  created() {
    this.getLinks()
  },
};
</script>

<style scoped>
.main {
  padding: 20px 30px;
}

.main > .el-row {
  margin: 20px 0;
}
.bg-banner {
  width: 100%;
  height: 400px;
  background: url("https://picshack.net/ib/aOjphDH3MB.jpg")
    no-repeat center center;
  background-size: cover;
}

.main_card {
  box-shadow: 0px 2px 10px 0px rgba(0, 0, 0, 0.12);
  padding: 20px;
}

.operator-btn {
  color: #666666;
}
</style>
