<template>
  <div id="account">
    <HeadBar
      title="海芯导航"
      :searchVal="searchVal"
      :activeIndex="activeIndex"
    />
    <div>开发中...</div>
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
      tableData: [
        {
          site: "2016-05-02",
          account: "wangxiaohu",
          uid: "1",
        },
      ],
    };
  },
  methods: {
    deleteRow(index, rows) {
      rows.splice(index, 1);
    },
  },
};
</script>

<style scoped>
</style>