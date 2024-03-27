<template>
  <div class="home-page">
    <Header />
    <h1>初次审核通过的项目</h1>
    <div v-if="isLoading" class="loading">加载中...</div>
    <div v-else>
      <div v-if="items && items.length === 0" class="no-items">没有项目</div>
      <ul v-else>
        <li v-for="item in items" :key="item.id" class="item">
          <h2>{{ item.name }}</h2>
          <p>简单描述：{{ item.description }}</p>
          <p>发布人：{{ item.publisher }}</p>
          <p>原因：{{ item.because }}</p>
          <router-link :to="'/item/details/' + item.id" class="view-details">查看详情</router-link>|
          <router-link :to="'/item/review2/' + item.id" class="view-details">最终审批</router-link>
        </li>
      </ul>
    </div>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<script>
import Header from '../store/header.vue';
import axios from 'axios';

export default {
  name: 'HomePage',
  components: {
    Header,
  },
  data() {
    return {
      items: [],
      isLoading: false,
      error: '',
    };
  },
  created() {
    this.fetchItems();
  },
  methods: {
    fetchItems() {
      this.isLoading = true;
      this.error = '';

      const token = window.sessionStorage.getItem('Authorization');
      const encodedToken = encodeURIComponent(token);
      axios.defaults.headers.common['Authorization'] = `Bearer${encodedToken}`;
      axios.get('http://localhost:8080/item/second')
        .then(response => {
          const data = response.data;
          if (data && data.code === 0) {
            this.items = data.data;
          } else {
            this.error = data && data.message ? data.message : '加载项目失败，请稍后重试。';
            console.error('加载项目失败:', this.error);
          }
          this.isLoading = false;
        })
        .catch(error => {
          this.error = '加载项目失败，请稍后重试。';
          console.error('加载项目失败:', error);
          this.isLoading = false;
        });
    },
  },
};
</script>

<style>
.home-page {
  max-width: 800px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  margin: 20px 0;
}

.no-items {
  text-align: center;
  margin: 20px 0;
}

.item {
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.item h2 {
  margin: 0;
}

.error {
  color: red;
  margin-top: 10px;
}
</style>