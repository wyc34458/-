<template>
  <div class="login-container">
    <h1 class="login-title">登录</h1>
    <form @submit.prevent="login" class="login-form">
      <div class="form-group">
        <label for="name" class="form-label">用户名: </label>
        <input type="name" id="name" v-model="name" class="form-input" required autocomplete="username" />
      </div>
      <div class="form-group">
        <label for="password" class="form-label">密码: </label>
        <input type="password" id="password" v-model="password" class="form-input" required autocomplete="current-password" />
      </div>
      <div class="form-group button-container">
        <button type="submit" class="login-button">登录</button>
        <button class="register-button" @click="goToRegister">注册</button>
      </div>
    </form>
    <p v-if="loginError" class="error">{{ loginError }}</p>
  </div>
</template>
<script>
export default {
  name: 'Login',
  data() {
    return {
      name: '',
      password: '',
      loginError: ''
    };
  },
  methods: {
    login() {
  const loginData = {
    name: this.name,
    password: this.password
    
  };

  fetch('http://localhost:8080/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(loginData)
  })
    .then(response => response.json())
    .then(data => {
      if (data && data.code === 0) {
        console.log('登陆成功');
        this.$router.push('/index');
        window.sessionStorage.setItem("Authorization",data.data)
      } else {
        this.loginError = data && data.message ? data.message : '登录失败，请检查您的用户名和密码。';
        console.error('登录失败:', this.loginError);
      }
    })
    .catch(error => {
      this.loginError = '登录失败，请检查您的用户名和密码。';
      console.error('登录失败:', error);
    });
},
goToRegister() {
      this.$router.push('/create'); // 导航到注册页面
    }
  }
};
</script>

<style>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.login-title {
  font-size: 24px;
  margin-bottom: 20px;
}

.login-form {
  width: 300px;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  margin-bottom: 5px;
}

.form-input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.button-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  gap: 10px;
}

.login-button {
  padding: 8px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.register-button {
  padding: 8px 16px;
  background-color: #28a745; /* 修改为不同的颜色 */
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.error {
  color: red;
  margin-top: 10px;
}
</style>