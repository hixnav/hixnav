<template>
  <div id="add-link">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <div style="padding-top: 20px">
      <el-form
        :model="ruleForm"
        status-icon
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="标题" prop="name">
          <el-input
            type="input"
            v-model="ruleForm.name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="logo">
          <el-input
            type="input"
            v-model="ruleForm.logo"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="简述" prop="desc">
          <el-input
            type="input"
            v-model="ruleForm.desc"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="地址" prop="url">
          <el-input
            type="input"
            v-model="ruleForm.url"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="分类ID" prop="cate">
          <el-input
            type="number"
            v-model="ruleForm.cate"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="分类名称" prop="catename">
          <el-input
            type="input"
            v-model="ruleForm.catename"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')"
            >提交</el-button
          >
          <el-button type="primary" @click="submitFormWd('ruleForm')"
            >提交并查看</el-button
          >
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- <FootBar /> -->
  </div>
</template>

<script>
import HeadBar from "@/components/HeadBar.vue";
import FootBar from "@/components/FootBar.vue";
export default {
  name: "AddLink",
  components: {
    HeadBar,
    FootBar,
  },
  data: function () {
    var checkAge = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("年龄不能为空"));
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error("请输入数字值"));
        } else {
          if (value < 18) {
            callback(new Error("必须年满18岁"));
          } else {
            callback();
          }
        }
      }, 1000);
    };
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        if (this.ruleForm.checkPass !== "") {
          this.$refs.ruleForm.validateField("checkPass");
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      searchVal: "",
      activeIndex: "1s",
      ruleForm: {
        name: "",
        desc: "",
        url: "",
        logo: "",
        cate: 0,
        catename: "",
      },
      rules: {
        pass: [{ validator: validatePass, trigger: "blur" }],
        checkPass: [{ validator: validatePass2, trigger: "blur" }],
        age: [{ validator: checkAge, trigger: "blur" }],
      },
    };
  },
  methods: {
    submitForm(formName) {
      self = this;
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(this.ruleForm);
          this.ruleForm.cate = parseInt(this.ruleForm.cate);
          this.$store
            .dispatch("nav/addLink", JSON.stringify(this.ruleForm))
            .then((response) => {
              console.log(response);
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
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    submitFormWd(formName) {
      this.submitForm(formName);
      this.$router.push("/");
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>

<style scoped>
.add-link {
  /* height: 1000px; */
}
</style>