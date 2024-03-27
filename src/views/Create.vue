<template>
  <div class="register">
    <h1>注册</h1>
    <form @submit.prevent="register" class="register-form">
      <div class="form-group">
        <label for="name" class="form-label">用户名: </label>
        <input type="name" id="name" v-model="name" class="form-input" required />
      </div>
      <div class="form-group">
        <label for="password" class="form-label">密码: </label>
        <input type="password" id="password" v-model="password" class="form-input" required />
      </div>
      <div class="form-group">
        <label for="password" class="form-label">确认密码: </label>
        <input type="password" id="password" v-model="password2" class="form-input" required />
      </div>
      <div class="form-group">
        <label for="phone" class="form-label">手机号: </label>
        <input type="phone" id="phone" v-model="phone" class="form-input" required />
      </div>
      <div class="form-group">
        <button type="submit" class="register-button">注册</button>
        <button class="login-button" @click="goToLogin">登录</button>
      </div>
    </form>
    <p v-if="registrationError" class="error">{{ registrationError }}</p>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data() {
    return {
      name: '',
      password: '',
      password2: '',
      phone: '',
      registrationError: ''
    };
  },
  methods: {
    goToLogin() {
      this.$router.push('/'); // 导航到登录页面
    },
    register() {
      const registerData = {
        name: this.name,
        password: this.password,
        password2: this.password2,
        phone: this.phone
      };

      fetch('http://localhost:8080/user/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(registerData)
      })
        .then(response => response.json())
        .then(data => {
  if (data && data.code === 0) {
    console.log('注册成功');
    this.$router.push('/'); // 导航到登录页面
  } else {
    this.registrationError = data && data.message ? data.message : '注册失败，请检查您输入的信息。';
    console.error('注册失败:', this.registrationError);
  }
})
        .catch(error => {
          this.registrationError = '注册失败，请检查您的用户名和密码。';
          console.error('注册失败:', error);
        });
    }
  }
};
</script>

<style>
.register {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.form-title {
  font-size: 24px;
  margin-bottom: 20px;
}

.register-form {
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

.button-group {
  display: flex;
  justify-content: center;
  align-items: center;
}

.button-gap {
  flex: 1;
  min-width: 10px;
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

.login-button {
  padding: 8px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>