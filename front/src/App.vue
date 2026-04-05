<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: "app",
  data: function () {

  },
  created() {
    this.$store
        .dispatch("config/dbTest")
        .then((response) => {
          if (response.code != 0){
            this.$router.push("install");
          }
        })
        .catch(() => {
          this.$router.push("install");
        });
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 0;
  padding: 0;
}

html, body, #app {
  height: 100%;
  margin: 0;
  padding: 0;
}
</style>
