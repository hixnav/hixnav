<template>
  <div id="account">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <div class="content">
      <el-container>
        <el-container>
          <el-main>
            <!-- 主要内容 -->
            <el-button type="primary" size="mini" @click="openAddDrawer"
              ><i class="el-icon-edit"></i>添加账号
            </el-button>
            <el-table
              :data="accountList"
              style="width: 100%; margin-top: 20px"
              :row-class-name="tableRowClassName"
            >
              <el-table-column prop="Sitename" label="站点" width="140">
                <template slot-scope="scope">
                  <a :href="scope.row.Siteurl" target="_blank">{{
                    scope.row.Sitename
                  }}</a>
                </template>
              </el-table-column>
              <el-table-column prop="Name" label="账号" width="160">
              </el-table-column>
              <el-table-column prop="Password" label="密码" width="140">
                <template slot-scope="scope">
                  <span>******</span>
                </template>
              </el-table-column>

              <el-table-column
                prop="CreateTime"
                label="添加时间"
              ></el-table-column>
              <el-table-column fixed="right" prop="" label="操作">
                <template slot-scope="scope">
                  <el-button
                    size="mini"
                    @click.native.prevent="viewCode(scope.$index, accountList)"
                    ><i class="el-icon-view"></i>查看密码
                  </el-button>
                  <el-button
                    size="mini"
                    @click.native.prevent="openEditDrawer(scope.$index, accountList)"
                    ><i class="el-icon-edit"></i>编辑
                  </el-button>
                  <el-button
                    size="mini"
                    type="danger"
                    @click.native.prevent="deleteRow(scope.$index, accountList)"
                    ><i class="el-icon-delete"></i>删除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-main>
        </el-container>
      </el-container>
    </div>

    <!-- 这里是弹出层 -->
    <div class="drawer">
      <el-drawer
        :title="addShow ? '添加账号':'编辑账号'"
        :visible.sync="dialog"
        direction="rtl"
        custom-class="demo-drawer"
        ref="drawer"
      >
        <div class="demo-drawer__content" style="padding: 30px">
          <el-form :model="form">
            <el-form-item label="名称" :label-width="formLabelWidth">
              <el-input
                id="linked_label"
                v-model="form.sitename"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="地址" :label-width="formLabelWidth">
              <el-input
                id="linked_label"
                v-model="form.siteurl"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="账号" :label-width="formLabelWidth">
              <el-input
                id="linked_name"
                v-model="form.name"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item label="秘钥" :label-width="formLabelWidth">
              <el-input
                id="linked_url"
                type="password"
                v-model="form.password"
              ></el-input>
            </el-form-item>
          </el-form>
          <div class="demo-drawer__footer" style="float: right">
            <el-button @click="cancelForm">取 消</el-button>
            <el-button type="primary" v-show="addShow" @click="addAccount">确 定</el-button>
            <el-button type="primary" v-show="editShow" @click="editAccount">确 定</el-button>
          </div>
        </div>
      </el-drawer>
    </div>
    <!-- 这里弹出层结束 -->
  </div>
</template>

<script>
import HeadBar from "@/components/HeadBar.vue";

export default {
  name: "AccountView",
  components: {
    HeadBar,
  },
  data: function () {
    return {
      searchVal: "",
      activeIndex: "5",
      accountList: [],
      // 弹出层控制
      dialog: false,
      addShow: true,
      editShow:false,
      form: {
        id: 0,
        sitename: "",
        siteurl: "",
        name: "",
        password: "",
      },
    };
  },
  methods: {
    openAddDrawer() {
      console.log("drawer");
      this.addShow = true
      this.editShow = false
      this.dialog = true;
      this.form = []
    },
    openEditDrawer(index, rows) {
      console.log(rows[index]);
      this.addShow = false
      this.editShow = true
      this.dialog = true;
      this.form.id = rows[index].ID
      this.form.sitename = rows[index].Sitename
      this.form.siteurl = rows[index].Siteurl
      this.form.name = rows[index].Name
    },
    cancelForm() {
      this.dialog = false;
    },
    deleteRow(index, rows) {
      console.log(rows[index]);
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.deleteAccountData(rows[index].ID);
          rows.splice(index, 1);
        })
        .catch(() => {});
    },
    editAccount() {
      console.log(this.form);
      let self = this;
      this.$store
          .dispatch("account/editAccount", JSON.stringify(this.form))
          .then((response) => {
            console.log(response);
            self.dialog = false;
            self.$notify({
              title: "成功",
              message: "",
              type: "success",
            });
            this.getData();
          })
          .catch((res) => {
            console.log(res);
            this.$message.error("失败");
          });
    },
    addAccount() {
      // 添加
      console.log(this.form);
      let self = this;
      this.$store
        .dispatch("account/addAccount", JSON.stringify(this.form))
        .then((response) => {
          console.log(response);
          self.dialog = false;
          self.$notify({
            title: "成功",
            message: "",
            type: "success",
          });
          this.getData();
        })
        .catch((res) => {
          console.log(res);
          this.$message.error("失败");
        });
    },
    deleteAccountData(id) {
      let self = this;
      this.$store
        .dispatch("account/delAccount", "id=" + id)
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
    getData() {
      let self = this;
      this.$store
        .dispatch("account/account", JSON.stringify(this.ruleForm))
        .then((response) => {
          console.log(response);
          self.accountList = response.data.accounts;
        })
        .catch((res) => {
          console.log(res);
        });
    },
  },
  created() {
    this.getData();
  },
};
</script>

<style scoped>
.el-menu-vertical {
  border: 0;
  padding-right: 20px;
}

.setting {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.content {
  width: 1024px;
  display: block;
  padding-top: 20px;
  margin: 0 auto;
}

.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>