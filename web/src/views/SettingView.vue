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
              @select="handleSelect"
          >
            <el-menu-item index="1">
              <i class="el-icon-user"></i>
              <span slot="title">多用户</span>
            </el-menu-item>
            <el-menu-item index="2">
              <i class="el-icon-coin"></i>
              <span slot="title">数据库配置</span>
            </el-menu-item>
            <el-menu-item index="3">
              <i class="el-icon-box"></i>
              <span slot="title">存储配置</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-container>
          <el-main v-if="showMain == 1">
            <!-- 主要内容 -->
            <el-button type="primary" size="small"
            ><i class="el-icon-plus"></i>添加用户
            </el-button
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

              <el-table-column prop="date" label="添加时间"></el-table-column>
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

          <el-main v-if="showMain == 2">
            <!-- 数据库配置 -->
            <el-form ref="form" :model="form" label-width="100px">
              <el-form-item label="数据库地址" style="width: 60%;">
                <el-input v-model="form.addr"></el-input>
              </el-form-item>
              <el-form-item label="数据库账号" style="width: 60%;">
                <el-input v-model="form.account"></el-input>
              </el-form-item>
              <el-form-item label="数据库密码" style="width: 60%;">
                <el-input type="password" v-model="form.password"></el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" size="small" @click="onSubmit">保存</el-button>
              </el-form-item>
            </el-form>
          </el-main>

          <el-main v-if="showMain == 3">
            <!-- 存储配置 -->
            <el-form ref="form" :model="form" label-width="100px">
              <el-form-item label="存储地址" style="width: 60%;">
                <el-input v-model="form.addr"></el-input>
              </el-form-item>
              <el-form-item label="AccessKey" style="width: 60%;">
                <el-input v-model="form.account"></el-input>
              </el-form-item>
              <el-form-item label="SecretKey" style="width: 60%;">
                <el-input type="password" v-model="form.password"></el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" size="small" @click="onSubmit">保存</el-button>
              </el-form-item>
            </el-form>
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
      showMain:1,
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
      form:{},
    };
  },
  methods: {
    tableRowClassName({row, rowIndex}) {
      if (rowIndex === 1) {
        return "warning-row";
      } else if (rowIndex === 3) {
        return "success-row";
      }
      return "";
    },
    handleSelect: function (key, keyPath) {
      this.showMain = key;
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