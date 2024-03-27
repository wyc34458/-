<template>
  <div class="add-item">
    <Header />
    <h1>审核项目</h1>
    <form @submit.prevent="addItem" enctype="multipart/form-data">
      <div class="form-group">
        <label for="status">审核状态：</label>
        <select id="status" v-model="status" required>
          <option value="">请选择状态</option>
          <option value="通过">通过</option>
          <option value="不通过">不通过</option>
        </select>
      </div>
      <div class="form-group">
        <label for="because">原因：</label>
        <textarea id="because" v-model="because"></textarea>
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
  name: 'Update',
  components: {
    Header,
  },
  data() {
    return {
      status: '', // 审核状态
      postError: '',
      because:'',
      id:'',
    };
  },
  methods: {
    created() {
      console.log('Route ID:', this.$route.params.id);
      this.id = this.$route.params.id;
    },
    addItem() {
      const postData = new FormData();
      postData.append('status', this.status === '通过' ? '2' : '1'); // 映射状态值为2或1
      postData.append('because', this.because);
      this.id = this.$route.params.id;
      const token = window.sessionStorage.getItem('Authorization');
      const encodedToken = encodeURIComponent(token);
      axios.defaults.headers.common['Authorization'] = `Bearer${encodedToken}`;

      axios
        .put(`http://localhost:8080/item/admin/update/${this.id}`, postData)
        .then(response => {
          const data = response.data;
          if (data && data.code === 0) {
            console.log('审核成功');
            this.$router.push('/initialreview');
            this.items = data.data;
          } else {
            this.postError =
              data && data.message
                ? data.message
                : '审核项目失败，请稍后重试。';
            console.error('审核项目失败:', this.postError);
          }
        })
        .catch(error => {
          this.postError = '审核项目失败，请稍后重试。';
          console.error('审核项目失败:', error);
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
  
  
  
  
  
  
  