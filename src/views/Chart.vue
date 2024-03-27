<template>
  <div class="home-page">
    <Header />
    <h1>阶段图表</h1>
    <div v-if="isLoading" class="loading">加载中...</div>
    <div v-else>
      <div v-if="items && items.length === 0" class="no-items">没有项目</div>
      <table v-else>
        <thead>
          <tr>
            <th>状态</th>
            <th>数量</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.id">
            <td>{{ item.StatusName }}</td>
            <td>{{ item.Count }}</td>
          </tr>
        </tbody>
      </table>

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
  mounted() {
  this.fetchItems();
},
  methods: {
    fetchItems() {
      this.isLoading = true;
      this.error = '';

      const token = window.sessionStorage.getItem('Authorization');
      const encodedToken = encodeURIComponent(token);
      axios.defaults.headers.common['Authorization'] = `Bearer${encodedToken}`;
      axios
        .get('http://localhost:8080/item/chart')
        .then(response => {
          const data = response.data;
          if (data && data.code === 0) {
            this.items = data.data;
          } else {
            this.error = data && data.message ? data.message : '加载图表失败，请稍后重试。';
            console.error('加载图表失败:', this.error);
          }
          this.isLoading = false;
        })
        .catch(error => {
          this.error = '加载图表失败，请稍后重试。';
          console.error('加载图表失败:', error);
          this.isLoading = false;
        });
    },
  
    },
  }

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

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  padding: 10px;
  border: 1px solid #ccc;
  text-align: left;
}

.error {
  color: red;
  margin-top: 10px;
}
</style>