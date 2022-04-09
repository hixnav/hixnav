<template>
  <div id="setting">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <div class="content">
      <el-container>
        <el-aside width="200px">
          <el-menu
            default-active="1-4-1"
            class="el-menu-vertical"
            @open="handleOpen"
            @close="handleClose"
            :collapse="isCollapse"
          >
            <el-menu-item index="1">
              <i class="el-icon-user-solid"></i>
              <span slot="title">多用户</span>
            </el-menu-item>
            <el-menu-item index="2">
              <i class="el-icon-user-solid"></i>
              <span slot="title">数据库配置</span>
            </el-menu-item>
            <el-menu-item index="3">
              <i class="el-icon-user-solid"></i>
              <span slot="title">存储配置</span>
            </el-menu-item>
            <el-menu-item index="4">
              <i class="el-icon-user-solid"></i>
              <span slot="title">数据迁移</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-container>
          <el-main>
            <!-- 主要内容 -->
            <el-button type="primary"
              ><i class="el-icon-edit"></i>添加用户</el-button
            >
            <el-table
              :data="tableData"
              style="width: 100%; margin-top: 20px"
              :row-class-name="tableRowClassName"
            >
              <el-table-column prop="uid" label="序号" width="180">
              </el-table-column>
              <el-table-column prop="account" label="账号" width="180">
              </el-table-column>

              <el-table-column prop="date" label="添加时间"> </el-table-column>
              <el-table-column fixed="right" prop="" label="操作">
                <template slot-scope="scope">
                  <el-button
                    @click.native.prevent="deleteRow(scope.$index, tableData)"
                    type="text"
                    size="small"
                  >
                    <i class="el-icon-delete"></i>
                  </el-button>
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
          account: "wangxiaohu",
          uid: "1",
        },
        {
          date: "2016-05-04",
          account: "wangxiaohu",
          uid: "2",
        },
        {
          date: "2016-05-01",
          account: "wangxiaohu",
          uid: "3",
        },
        {
          date: "2016-05-03",
          account: "wangxiaohu",
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