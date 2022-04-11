<template>
  <div class="foot_bar">
    <el-backtop class="suspend">
      <div class="backup">
        <i class="el-icon-arrow-up"></i>
      </div>
    </el-backtop>
    <div class="footer">
      <el-row class="footer-section"></el-row>
      <el-row :gutter="24" class="footer-section footer-section-middle">
        <div class="copy">
          Power By
          <a
            href="https://hixnav.github.io/hixnav"
            style="text-decoration: none; color: #c0c0c0"
            >HixNav</a
          >
          ©2202 海芯导航 v1.10.2
          <a
            href="https://github.com/hixnav/hixnav/blob/gh-pages/README.md"
            style="text-decoration: none; color: #c0c0c0"
            target="_blank"
            >捐赠</a
          >
        </div>
      </el-row>
      <el-row class="footer-section"></el-row>
    </div>
    <div class="footer-copy"></div>
  </div>
</template>

<script>
export default {
  name: "FootBar",
  props: {},
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.foot_bar,
.footer-section {
  margin: 0;
  padding: 0;
}
/* .footer-copy {
  width: 100%;
  height: 4rem;
} */
.footer {
  text-align: center;
  width: 100%;
  height: 4rem;
  background: #212930;
  box-sizing: border-box;
  color: #6c7d8f;
  overflow: hidden;
  z-index: 9999;
}
.footer-section-middle {
  line-height: 2rem;
  margin: 0;
  padding: 0;
}
.footer .copy {
  text-align: center;
  font-size: 14px;
  font-weight: 700;
}

.footer .footer-section {
  margin: 0;
  padding: 0;
  height: 1rem;
}

.suspend {
  margin-bottom: 120px;
}

.backup {
  height: 100%;
  width: 100%;
  background-color: #fff;
  box-shadow: 0 0 6px rgba(0, 0, 0, 0.12);
  text-align: center;
  line-height: 40px;
  border-radius: 3px;
  color: #1989fa;
}
</style>
