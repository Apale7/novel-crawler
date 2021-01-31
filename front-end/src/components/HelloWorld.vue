<template>
  <div class="hello">
    <a href="https://www.xxxbiquge.com/" target="_blank"
      >找到你想要下载的小说，复制网址到下面</a
    >
    <el-form :inline="true" :model="info" class="demo-form-inline">
      <el-form-item label="网址">
        <el-input v-model="info.url" placeholder="小说的网址"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onGetList" v-loading="info.loading"
          >章节列表</el-button
        >
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDownload" v-loading="info.loading"
          >下载全本</el-button
        >
      </el-form-item>
    </el-form>
    <h2 v-if="info.success">下载成功！</h2>
    <h2 v-if="info.name">书名: {{ info.name }}</h2>
    <el-table
      :data="info.list"
      border
      style="width: 50%; height: 85%; margin: 0 auto"
      v-loading="info.loading"
    >
      <el-table-column prop="title" label="章节"> </el-table-column>
    </el-table>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  name: "HelloWorld",
  props: {},
  data() {
    return {
      info: {
        url: "",
        name: "",
        success: false,
        loading: false,
        list: [],
      },
    };
  },
  methods: {
    onGetList() {
      this.info.loading = true;
      console.log("getlist!");
      Axios.get("http://localhost:6350/list", {
        params: {
          url: this.info.url,
          headers: {
            "Content-Type": "application/json;charset=UTF-8",
          },
        },
      }).then((res) => {
        this.info.list = res.data.list;
        this.info.name = res.data.name.replace(".txt", "");
        this.info.loading = false;
      });
    },
    onDownload() {
      this.info.loading = true;
      this.info.success = false;
      console.log("download!");
      Axios.get("http://localhost:6350/novel", {
        params: {
          url: this.info.url,
          headers: {
            "Content-Type": "application/json;charset=UTF-8",
          },
        },
      }).then((res) => {
        this.info.success = true;
        this.info.loading = false;
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
