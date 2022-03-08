<template>
  <div class="home">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
      :activeIndex1="activeIndex1"
    />
    <div class="bg-banner">
      <el-row :gutter="16" style="margin-left: 0px; margin-right: 0px">
        <el-col :span="6" style="padding: 0 30px">
          <div
            style="
              height: 20rem;
              padding: '20px 0';
              background-color: rgba(100, 100, 100, 0.2);
            "
          >
            <el-row>
              <el-col
                :span="8"
                v-for="c in cates"
                :key="c"
                style="
                  color: #fff;
                  font-size: 14px;
                  padding: 20px 0;
                  text-align: center;
                "
                ><span @click="goAnchor('#anchor' + c.Cate)">{{
                  c.Catename
                }}</span></el-col
              >
            </el-row>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="line"></div>
     <el-row :gutter="16" style="margin:0;width:98%;margin:0 auto">
        <div
          style="margin:20px 20px;box-shadow: 5px 5px 5px 5px rgba(0, 0, 0, 0.12); padding: 20px"
        >
          <el-descriptions
            class="margin-top"
            title=""
            :colon="false"
            column="10"
          >
            <template slot="extra"> </template>
            <el-descriptions-item v-for="o in quikList" :key="o">
              <!-- <i class="el-icon-search"></i> -->
              <el-link :href="o.href" target="_blank">{{o.name}}</el-link>
            </el-descriptions-item>
          </el-descriptions>
        </div>
      </el-row>
    <div>
      <el-row v-for="c in cates" :key="c">
        <section :id="'anchor' + c.Cate" style="padding-left: 30px">
          <h3>{{ c.Catename }}</h3>
        </section>
        <el-col
          :span="6"
          v-for="o in navs[c.Cate]"
          :key="o"
          style="padding: 30px"
        >
          <el-card :body-style="{ padding: '0px' }">
            <div style="padding: 14px">
              <el-row :gutter="20">
                <el-col :span="5">
                  <el-avatar
                    style="background-color: #fff"
                    :src="o.Logo"
                  ></el-avatar>
                </el-col>
                <el-col :span="12">
                  <div class="linked-title">{{ o.Name }}</div>
                  <div class="bottom clearfix">
                    <time class="time">{{ o.Desc }}</time>
                  </div>
                </el-col>
                <el-col :span="7">
                  <div class="bottom clearfix">
                    <el-button
                      icon=""
                      size="small"
                      round
                      class="button"
                      @click="jumpTo(o.Url)"
                    >
                      访问</el-button
                    >
                  </div>
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <!-- 悬浮按钮 -->
    <div class="" style="position:fixed;right:40px;bottom:108px;"> 
      <router-link to="/add-link"><el-button type="primary" icon="el-icon-edit" circle></el-button></router-link>
    </div>
    <!-- 悬浮按钮结束 -->
    <FootBar />
  </div>
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
      activeIndex1: "1",
      navs: [],
      cates: [],
      quikList: [
        {
          href:'https://wennmu.github.io/',
          name:'我的博客'
        },
        {
          href:'https://www.baidu.com',
          name:'百度'
        },
        {
          href:'https://fanyi.baidu.com/',
          name:'百度翻译'
        },
        {
          href:'https://note.youdao.com/',
          name:'有道云笔记'
        },
      ],
    };
  },
  methods: {
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
  },
  created: function () {
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
}

.bottom time {
  font-size: 12px;
}

.bottom .el-button {
  font-size: 12px;
}

.button {
  padding: 0;
  float: right;
}
		.linked-title {
			overflow: hidden;
			white-space: nowrap;
		}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}

.home {
  margin: 0px;
  padding: 0px;
}
	.bg-banner {
			width: 100%;
			height: 20rem;
			background: url("https://haixin-1300602599.cos.ap-guangzhou.myqcloud.com/t0122b0e1af805d3679.jpg") no-repeat center center;
			background-color: #56CCF2;
			/* fallback for old browsers */
			background-color: -webkit-linear-gradient(to right, #2F80ED, #56CCF2);
			/* Chrome 10-25, Safari 5.1-6 */
			background-color: linear-gradient(to right, #2F80ED, #56CCF2);
			/* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
			background-size: cover;
		}
</style>