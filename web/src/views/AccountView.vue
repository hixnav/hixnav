<template>
  <div id="setting">
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
            <el-button type="primary"
            ><i class="el-icon-edit"></i>添加账号</el-button
            >
            <el-table
                :data="tableData"
                style="width: 100%; margin-top: 20px"
                :row-class-name="tableRowClassName"
            >
              <el-table-column prop="site" label="站点" width="180">
              </el-table-column>
              <el-table-column prop="account" label="账号" width="180">
              </el-table-column>
              <el-table-column prop="password" label="密码" width="180">
              </el-table-column>

              <el-table-column prop="date" label="添加时间"> </el-table-column>
              <el-table-column fixed="right" prop="" label="操作">
                <template slot-scope="scope">
                  <el-button
                      size="mini"
                      @click.native.prevent="deleteRow(scope.$index, tableData)""> <i class="el-icon-view"></i>查看密码</el-button>
                  <el-button
                      size="mini"
                      type="danger"
                      @click.native.prevent="deleteRow(scope.$index, tableData)""><i class="el-icon-delete"></i>删除</el-button>
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
  name: "SettingView",
  components: {
    HeadBar,
  },
  data: function () {
    return {
      searchVal: "",
      activeIndex: "99",
      tableData: [
        {
          date: "2016-05-02",
          site: "百度一下",
          account: "wangxiaohu",
          password:"······",
          uid: "1",
        },
        {
          date: "2016-05-04",
          site: "百度一下",
          account: "wangxiaohu",
          password:"······",
          uid: "2",
        },
        {
          date: "2016-05-01",
          site: "百度一下",
          account: "wangxiaohu",
          password:"······",
          uid: "3",
        },
        {
          date: "2016-05-03",
          site: "百度一下",
          account: "wangxiaohu",
          password:"······",
          uid: "4",
        },
      ],
    };
  },
  methods: {
    tableRowClassName({ row, rowIndex }) {
      if (rowIndex === 1) {
        return "warning-row";
      } else if (rowIndex === 3) {
        return "success-row";
      }
      return "";
    },
    deleteRow(index, rows) {
      rows.splice(index, 1);
    },
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