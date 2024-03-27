<template>
  <div class="item-detail-page">
    <Header />
    <h1>项目详情</h1>
    <div v-if="isLoading" class="loading">加载中...</div>
    <div v-else>
      <div v-if="item" class="item">
        <h2>{{ item.name }}</h2>
        <p>描述：{{ item.description }}</p>
        <p>预算：{{ item.budget }}</p>
        <p>发布人：{{ item.publisher }}</p>
        <p>文件：<a :href="item.file_url" target="_blank">{{ item.file_url }}</a></p>
        <p>审核状态：{{ item.statusName }}</p>
        <p>原因：{{ item.because }}</p>
      </div>
      <div v-else class="no-item">项目不存在或加载失败。</div>
    </div>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<script>

import Header from '../store/header.vue';
import axios from 'axios';

export default {
  name: 'DetailsView',
  components: {
    Header,
  },
  data() {
    return {
      item: null,
      isLoading: false,
      error: '',
      id : null,
    };
  },
  created() {
    this.fetchItem();
  },
  methods: {
    fetchItem() {
      this.isLoading = true;
      this.error = '';

      const id = this.$route.params.id;
      const token = window.sessionStorage.getItem('Authorization');
      const encodedToken = encodeURIComponent(token);
      axios.defaults.headers.common['Authorization'] = `Bearer${encodedToken}`;
      axios
        .get(`http://localhost:8080/item/details/${id}`)
        .then(response => {
          const data = response.data;
          if (data && data.code === 0) {
            this.item = data.data;
          } else {
            this.error =
              data && data.message
                ? data.message
                : '加载项目详情失败，请稍后重试。';
            console.error('加载项目详情失败:', this.error);
          }
        })
        .catch(error => {
          this.error = '加载项目详情失败，请稍后重试。';
          console.error('加载项目详情失败:', error);
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
};
</script>

<style>
.item-detail-page {
  max-width: 800px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  margin: 20px 0;
}

.no-item {
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