<template>
  <header class="banner">
    <nav>
      <router-link to="/index">首页</router-link> |
      <router-link to="/">登录</router-link> |
      <router-link to="/create">注册</router-link> |
      <router-link to="/post">申报项目</router-link> |
      <router-link to="/initialreview">初次审批</router-link> |
      <router-link to="/finalreview">最终审核</router-link> |
      <router-link to="/user">我的</router-link> |
      <router-link to="/chart">图表</router-link> |
    </nav>
  </header>
</template>

<script>
export default {
  name: 'Header',
  data() {
    return {
      isLoggedIn: false,
      showDropdown: false,
    };
  },
  methods: {
    logout() {
      // 执行退出登录的逻辑
      // 将 isLoggedIn 设置为 false
      this.isLoggedIn = false;
    },
  },
};
</script>

<style>
.banner {
  background-color: #f1f1f1;
  padding: 20px;
  text-align: left;
}

nav {
  display: flex;
  align-items: center;
}

.avatar-wrapper {
  margin-left: auto;
  position: relative;
}

.dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  background-color: #fff;
  padding: 10px;
  border: 1px solid #ccc;
  display: none;
}

.avatar-wrapper:hover .dropdown {
  display: block;
}

nav img.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
}
</style>