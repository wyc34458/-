<template>
  <div class="add-item">
    <Header />
    <h1>申报项目</h1>
    <form @submit.prevent="addItem" enctype="multipart/form-data">
      <div class="form-group">
        <label for="name">项目名称：</label>
        <input type="text" id="name" v-model="name" required />
      </div>
      <div class="form-group">
        <label for="description">项目描述：</label>
        <textarea id="description" v-model="description"></textarea>
      </div>
      <div class="form-group">
        <label for="budget">预算：</label>
        <input type="number" id="budget" v-model="budget" required />
      </div>
      <div class="form-group">
        <label for="publisher">发布人：</label>
        <input type="text" id="publisher" v-model="publisher" required />
      </div>
      <div class="form-group">
        <label for="file">上传文件：</label>
        <input type="file" id="file" ref="file" required />
      </div>
      <div class="form-group">
        <button type="submit">提交</button>
      </div>
    </form>
    <p v-if="postError" class="error">{{ postError }}</p>
  </div>
</template>

<script>
import Header from '../store/header.vue';
import axios from 'axios';

export default {
  name: 'Post',
  components: {
    Header,
  },
  data() {
    return {
      name: '',
      description: '',
      budget: '',
      publisher: '',
      postError: '',
    };
  },
  methods: {
    addItem() {
      const postData = new FormData();
      postData.append('name', this.name);
      postData.append('description', this.description);
      postData.append('budget', this.budget);
      postData.append('publisher', this.publisher);
      postData.append('file', this.$refs.file.files[0]);

      const token = window.sessionStorage.getItem('Authorization');
      const encodedToken = encodeURIComponent(token);
      axios.defaults.headers.common['Authorization'] = `Bearer${encodedToken}`;

      axios
        .post('http://localhost:8080/item/add', postData)
        .then(response => {
          const data = response.data;
          if (data && data.code === 0) {
            console.log('上传成功');
        this.$router.push('/index');
            this.items = data.data;
          } else {
            this.postError = data && data.message ? data.message : '上传项目失败，请稍后重试。';
            console.error('上传项目失败:', this.postError);
          }
        })
        .catch(error => {
          this.postError = '上传项目失败，请稍后重试。';
          console.error('上传项目失败:', error);
        });
    },
  },
};
</script>

<style>
.add-item {
  max-width: 400px;
  margin: 0 auto;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input[type="text"],
input[type="number"],
textarea {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  padding: 8px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.error {
  color: red;
  margin-top: 10px;
}

.success {
  color: green;
  margin-top: 10px;
}
</style>






