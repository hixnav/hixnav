<template>
  <div id="setting">
    <HeadBar title="海芯导航" :searchVal="searchVal" :activeIndex="activeIndex" />
    <div class="content">
      <el-container>
        <el-aside width="200px">
          <el-menu default-active="1-4-1" class="el-menu-vertical" @open="handleOpen" @close="handleClose"
            :collapse="isCollapse" @select="handleSelect">
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
            <el-menu-item index="4">
              <i class="el-icon-box"></i>
              <span slot="title">Banner配置</span>
            </el-menu-item>
            <el-menu-item index="5">
              <i class="el-icon-upload2"></i>
              <span slot="title">数据同步</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-container>
          <el-main v-if="showMain == 1">
            <!-- 主要内容 -->
            <el-button type="primary" size="small" @click="openMultiUserDrawer"><i class="el-icon-plus"></i>添加用户
            </el-button>
            <el-table :data="userList" style="width: 100%; margin-top: 20px" :row-class-name="tableRowClassName">
              <el-table-column prop="Id" label="序号" width="180">
              </el-table-column>
              <el-table-column prop="Account" label="账号" width="180">
              </el-table-column>

              <el-table-column prop="CreateTime" label="添加时间"></el-table-column>
              <el-table-column fixed="right" prop="" label="操作">
                <template slot-scope="scope">
                  <el-button @click.native.prevent="deleteUserRow(scope.$index, userList)" type="text" size="small">
                    <i class="el-icon-delete"></i>
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-main>

          <!-- 这里是弹出层 -->
          <div class="drawer">
            <el-drawer title="添加用户" :visible.sync="multiUserDialog" direction="rtl" custom-class="demo-drawer"
              ref="drawer">
              <div class="demo-drawer__content" style="padding: 30px">
                <el-form :model="mutilUserForm">
                  <el-form-item label="账号" :label-width="formLabelWidth">
                    <el-input id="linked_name" v-model="mutilUserForm.account" autocomplete="off"></el-input>
                  </el-form-item>
                  <el-form-item label="密码" :label-width="formLabelWidth">
                    <el-input type="password" id="linked_url" v-model="mutilUserForm.passwd"></el-input>
                  </el-form-item>
                </el-form>
                <div class="demo-drawer__footer" style="float: right">
                  <el-button @click="cancelMultiUserDrawer">取 消</el-button>
                  <el-button type="primary" @click="multiUserSubmit">确 定</el-button>
                </div>
              </div>
            </el-drawer>
          </div>
          <!-- 这里弹出层结束 -->

          <el-main v-if="showMain == 2">
            <h3>如何完成数据迁移</h3>
            <el-steps :active="active" finish-status="success" style="width:500px;margin-bottom: 30px">
              <el-step title="导出数据" description="将现在的数据导出到本地"></el-step>
              <el-step title="切换数据库" description="更换新配置数据库连接信息并且初始化数据库"></el-step>
              <el-step title="导入数据" description="将本地的数据导入新的数据库"></el-step>
            </el-steps>

            <el-tooltip class="item" effect="dark" content="将现在的数据导出到本地" placement="bottom-end">
              <el-button type="primary" size="small" @click="exportDBData">导出数据库</el-button>
            </el-tooltip>
            <el-tooltip class="item" effect="dark" content="更换新配置数据库连接信息并且初始化数据库" placement="bottom-end">
              <router-link to="/install" style="margin:0 10px"><el-button type="primary" size="small"
                  @click="exportDBData">更改数据库配置</el-button></router-link>
            </el-tooltip>

            <el-tooltip class="item" effect="dark" content="将本地的数据导入新的数据库" placement="bottom-end">
              <el-button type="primary" size="small" @click="exportDBData">导入历史数据</el-button>
            </el-tooltip>
          </el-main>

          <el-main v-if="showMain == 3">
            <!-- 存储配置 -->
            <el-form ref="form" :model="config" label-width="100px">
              <el-form-item label="存储地址" style="width: 60%;">
                <el-input v-model="config.cos" placeholder="请输入地址，https://%s-***.cos.ap-***.myqcloud.com/"></el-input>
              </el-form-item>
              <el-form-item label="AccessKey" style="width: 60%;">
                <el-input v-model="config.cosSecretID" placeholder="请输入SecretID" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="SecretKey" style="width: 60%;">
                <el-input type="password" v-model="config.cosSecretKey" placeholder="请输入SecretKey" autocomplete="off"></el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" size="small" @click="onSubmit">保存并初始化数据</el-button>
                <!--                <el-button type="plain" size="small" @click="onInitHandle">初始化</el-button>-->
              </el-form-item>
            </el-form>
          </el-main>

          <el-main v-if="showMain == 4">
            <!-- Banner配置 -->
            <el-form ref="form" :model="config" label-width="100px">
              <el-form-item label="首页 Banner" style="width: 60%;">
                <el-input v-model="config.navBanner" placeholder="请输入地址，https://..."></el-input>
              </el-form-item>
              <el-form-item label="快链 Banner" style="width: 60%;">
                <el-input v-model="config.linkBanner" placeholder="请输入地址，https://..."></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" size="small" @click="saveBanner">保 存</el-button>
              </el-form-item>
            </el-form>
          </el-main>

          <el-main v-if="showMain == 5">
            <!-- 数据同步 -->
            <div class="sync-container">
              <div class="sync-header">
                <h3><i class="el-icon-connection"></i> 数据同步</h3>
                <p class="sync-desc">将本地 SQLite 数据库与腾讯云 COS 进行双向同步，用于数据备份和多端数据一致。</p>
              </div>

              <div class="sync-cards">
                <!-- 推送到COS -->
                <div class="sync-card">
                  <div class="card-icon push-icon">
                    <i class="el-icon-upload2"></i>
                  </div>
                  <div class="card-body">
                    <h4>推送到云端</h4>
                    <p class="card-desc">将本地数据库全部表导出为 JSON 并上传至 COS 存储桶。</p>
                    <el-button type="primary" @click="handlePushToCOS" :loading="pushLoading" :disabled="pushLoading || pullLoading">
                      <i class="el-icon-upload2"></i> 同步到 COS
                    </el-button>
                  </div>
                </div>

                <!-- 从COS拉取 -->
                <div class="sync-card">
                  <div class="card-icon pull-icon">
                    <i class="el-icon-download"></i>
                  </div>
                  <div class="card-body">
                    <h4>从云端拉取</h4>
                    <p class="card-desc">从 COS 下载数据并合并到本地，按 ID 新增或更新，不删除已有数据。</p>
                    <el-button type="success" @click="handlePullFromCOS" :loading="pullLoading" :disabled="pushLoading || pullLoading">
                      <i class="el-icon-download"></i> 从 COS 拉取
                    </el-button>
                  </div>
                </div>
              </div>

              <!-- 同步状态 -->
              <div class="sync-status-card">
                <div class="status-header">
                  <h4><i class="el-icon-info"></i> 同步配置</h4>
                  <el-button type="text" size="small" @click="loadSyncStatus"><i class="el-icon-refresh"></i> 刷新</el-button>
                </div>
                <div v-if="syncStatusData" class="status-body">
                  <div class="status-item">
                    <span class="status-label">存储驱动</span>
                    <span class="status-value">{{ syncStatusData.storage_driver }}</span>
                  </div>
                  <div class="status-item">
                    <span class="status-label">SQLite 路径</span>
                    <span class="status-value">{{ syncStatusData.sqlite_path }}</span>
                  </div>
                  <div class="status-item">
                    <span class="status-label">定时同步</span>
                    <span class="status-value">
                      <el-tag :type="syncStatusData.cos_sync_enabled === 'true' ? 'success' : 'info'" size="mini">
                        {{ syncStatusData.cos_sync_enabled === 'true' ? '已开启' : '已关闭' }}
                      </el-tag>
                    </span>
                  </div>
                  <div class="status-item">
                    <span class="status-label">同步间隔</span>
                    <span class="status-value">{{ syncStatusData.cos_sync_interval }} 秒</span>
                  </div>
                  <div class="status-item">
                    <span class="status-label">COS 存储桶</span>
                    <span class="status-value">{{ syncStatusData.cos_sync_bucket }}</span>
                  </div>
                </div>
                <div v-else class="status-empty">
                  <i class="el-icon-info"></i>
                  <p>点击上方"刷新"按钮加载同步配置</p>
                </div>
              </div>
            </div>
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
      showMain: 1,
      multiUserDialog: false,
      mutilUserForm: {
        account: "",
        passwd: "",
      },
      userList: [
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
      config: {},
      pushLoading: false,
      pullLoading: false,
      syncStatusData: null,
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
    handleSelect: function (key, keyPath) {
      this.showMain = key;
    },
    deleteRow(index, rows) {
      rows.splice(index, 1);
    },
    exportDBData() {
      console.log('export db data')
    },
    onSubmit() {
      this.$store
        .dispatch("config/cosMigrate", JSON.stringify(this.config))
        .then((response) => {
          console.log(response);
          if (response.code == 0) {
            this.$message.success("success");
          }
        })
        .catch((res) => {
          console.log(res);
        });
    },

    listUser() {
      let self = this;
      this.$store
        .dispatch("user/listUser", {})
        .then((response) => {
          console.log(response);
          self.userList = response.data.users;
        })
        .catch((res) => {
          console.log(res);
        });
    },

    deleteUserRow(index, rows) {
      console.log(index, rows)
      this.$store
        .dispatch("user/delUser", "id=" + rows[index].Id)
        .then((response) => {
          if (response.code != 0) {
            this.$message({
              message: "用户删除失败",
              type: 'warning'
            })
            return
          }
          this.multiUserDialog = false;
          this.$message({
            message: "用户删除成功",
            type: 'success'
          })
          this.listUser()
        })
        .catch(() => {
          this.$message({
            message: "用户删除失败",
            type: 'warning'
          })
        });
    },
    openMultiUserDrawer() {
      this.multiUserDialog = true;
    },
    cancelMultiUserDrawer() {
      this.multiUserDialog = false;
    },
    multiUserSubmit() {
      this.$store
        .dispatch("user/addUser", JSON.stringify(this.mutilUserForm))
        .then((response) => {
          console.log(response);
          if (response.code != 0) {
            this.$message({
              message: "用户添加失败",
              type: 'warning'
            })
            return
          }
          this.multiUserDialog = false;
          this.$message({
            message: "用户添加成功",
            type: 'success'
          })
          this.listUser()
        })
        .catch(() => {
          this.$message({
            message: "用户添加失败",
            type: 'warning'
          })
        });
    },
    saveBanner(){
      this.$store.dispatch('config/saveBanner', this.config)
        .then((response) => {
          if (response.data && response.data.success) {
            this.$message.success('Banner 配置已保存');
          } else {
            this.$message.warning(response.data.message || '保存失败');
          }
        })
        .catch(() => {
          this.$message.error('保存失败');
        });
    },
    loadBanner(){
      this.$store.dispatch('config/getBanner', {})
        .then((response) => {
          if (response.data) {
            this.config.navBanner = response.data.navBanner || '';
            this.config.linkBanner = response.data.linkBanner || '';
          }
        })
        .catch(() => {});
    },
    handlePushToCOS() {
      this.$confirm('确定要将本地数据同步到COS吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.pushLoading = true;
        this.$store.dispatch('config/syncPush', {})
          .then((response) => {
            this.pushLoading = false;
            if (response.data && response.data.success) {
              this.$message.success(response.data.message || '同步成功');
            } else {
              this.$message.warning(response.data.message || '同步完成，但部分表可能失败');
            }
          })
          .catch(() => {
            this.pushLoading = false;
            this.$message.error('同步到COS失败');
          });
      }).catch(() => {});
    },
    handlePullFromCOS() {
      this.$confirm('确定要从COS拉取数据到本地吗？已有数据不会被删除，只会新增或更新。', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }).then(() => {
        this.pullLoading = true;
        this.$store.dispatch('config/syncPull', {})
          .then((response) => {
            this.pullLoading = false;
            if (response.data && response.data.success) {
              this.$message.success(response.data.message || '拉取成功');
            } else {
              this.$message.warning(response.data.message || '拉取完成，但部分表可能失败');
            }
          })
          .catch(() => {
            this.pullLoading = false;
            this.$message.error('从COS拉取数据失败');
          });
      }).catch(() => {});
    },
    loadSyncStatus() {
      this.$store.dispatch('config/syncStatus', {})
        .then((response) => {
          this.syncStatusData = response.data;
        })
        .catch(() => {
          this.$message.error('获取同步状态失败');
        });
    },
  },
  created() {
    this.listUser()
  }
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

.sync-container {
  max-width: 720px;
}

.sync-header {
  margin-bottom: 24px;
}

.sync-header h3 {
  margin: 0 0 8px 0;
  color: #303133;
  font-size: 18px;
}

.sync-header h3 i {
  margin-right: 6px;
  color: #409eff;
}

.sync-desc {
  color: #909399;
  font-size: 14px;
  margin: 0;
  line-height: 1.6;
}

/* 同步卡片 */
.sync-cards {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

.sync-card {
  flex: 1;
  background: #fff;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  padding: 24px;
  display: flex;
  align-items: flex-start;
  gap: 16px;
  transition: box-shadow 0.3s;
}

.sync-card:hover {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.card-icon {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  flex-shrink: 0;
}

.push-icon {
  background: #ecf5ff;
  color: #409eff;
}

.pull-icon {
  background: #f0f9eb;
  color: #67c23a;
}

.card-body h4 {
  margin: 0 0 6px 0;
  font-size: 15px;
  color: #303133;
}

.card-desc {
  margin: 0 0 16px 0;
  font-size: 13px;
  color: #909399;
  line-height: 1.5;
  min-height: 40px;
}

/* 同步状态卡片 */
.sync-status-card {
  background: #fff;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  padding: 20px 24px;
}

.status-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid #ebeef5;
}

.status-header h4 {
  margin: 0;
  font-size: 15px;
  color: #303133;
}

.status-header h4 i {
  margin-right: 6px;
  color: #909399;
}

.status-body {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
}

.status-item:not(:last-child) {
  border-bottom: 1px dashed #ebeef5;
}

.status-label {
  color: #606266;
  font-size: 14px;
}

.status-value {
  color: #303133;
  font-size: 14px;
  font-weight: 500;
}

.status-empty {
  text-align: center;
  padding: 24px 0;
  color: #c0c4cc;
}

.status-empty i {
  font-size: 32px;
  display: block;
  margin-bottom: 8px;
}

.status-empty p {
  margin: 0;
  font-size: 13px;
}
</style>
