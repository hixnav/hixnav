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
      <div style="margin: 20px 20px; flex: 3">
        <el-row style="min-width: 400px; padding: 14px 50px 20px">
          <el-input
            placeholder="搜索关键词"
            :value="searchVal"
            class="input-with-select"
            suffix-icon="el-icon-search"
            size="large"
          />
          <el-descriptions
            class="margin-top"
            title=""
            :colon="false"
            column="10"
          >
            <template slot="extra"></template>
            <el-descriptions-item v-for="o in quikList" :key="o">
              <!-- <i class="el-icon-search"></i> -->
              <el-link :href="o.href" target="_blank" style="font-size: 10px"
                >{{ o.name }}
              </el-link>
            </el-descriptions-item>
          </el-descriptions>
        </el-row>
      </div>
      <div style="flex: 1; border-left: 1px solid #ccc; padding-left: 20px">
        <h3 style="margin-left: 10px">标签</h3>
        <el-tag
          size="mini"
          v-for="c in cates"
          :key="c"
          @click="goAnchor('#anchor' + c.Cate)"
          style="margin: 0.4rem 0.6rem"
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
                        ></el-button>
                        <el-button
                          icon="el-icon-delete"
                          size="mini"
                          circle
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
      navs: [],
      cates: [],
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
    handOffBtn() {
      console.log("打开操作" + this.showBtn);
      if (this.showBtn == false) {
        this.showBtn = true;
      } else {
        this.showBtn = false;
      }
    },
    jumpTo(url) {
      //this.$router.push({path:url});
      location.href = url;
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

/* .loadding{
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.loadding > div{
  width: 100%;
  height: 200px;
  flex-direction: column;
  align-items: center;
} */
</style>