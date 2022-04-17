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
            <el-button type="primary" @click="addAccount"
            ><i class="el-icon-edit"></i>添加账号</el-button
            >
            <el-table
                :data="accountList"
                style="width: 100%; margin-top: 20px"
                :row-class-name="tableRowClassName"
            >
              <el-table-column prop="Sitename" label="站点" width="180">
                <template slot-scope="scope">
                  <a href="" target="_blank">{{ scope.row.Sitename }}</a>
                </template>
              </el-table-column>
              <el-table-column prop="Name" label="账号" width="180">
              </el-table-column>
              <el-table-column prop="Password" label="密码" width="180">
                <template slot-scope="scope">
                  <span>······</span>
                </template>
              </el-table-column>

              <el-table-column prop="Createat" label="添加时间"> </el-table-column>
              <el-table-column fixed="right" prop="" label="操作">
                <template slot-scope="scope">
                  <el-button
                      size="mini"
                      @click.native.prevent="deleteRow(scope.$index, tableData)"> <i class="el-icon-view"></i>查看密码</el-button>
                  <el-button
                      size="mini"
                      type="danger"
                      @click.native.prevent="deleteRow(scope.$index, tableData)"><i class="el-icon-delete"></i>删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-main>
        </el-container>
      </el-container>
    </div>
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
    };
  },
  methods: {
    // tableRowClassName({ row, rowIndex }) {
    //   if (rowIndex === 1) {
    //     return "warning-row";
    //   } else if (rowIndex === 3) {
    //     return "success-row";
    //   }
    //   return "";
    // },
    deleteRow(index, rows) {
      rows.splice(index, 1);
    },
    addAccount(){
      // 添加
    }
  },
  created() {
    let self = this
    this.$store.dispatch("account/account", JSON.stringify(this.ruleForm))
        .then((response) => {
          console.log(response);
          self.accountList = response.data.accounts;
        })
        .catch((res) => {
          console.log(res);
        });
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