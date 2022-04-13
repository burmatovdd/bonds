<template>
      <div class="register">
        <input id="name" type="text" class="register-item register-name" placeholder="Name">
        <input id="regLogin" type="text" class="register-item register-login" placeholder="login">
        <input id="regPassword" type="password" class="register-item register-password" placeholder="password">
        <input type="password" class="register-item register-password" placeholder="password">
        <button class="button" @click="register">submit</button>
      </div>
</template>

<script>
import { defineComponent } from 'vue';
export default  defineComponent({
  name: "register",
  setup(){
    let user = {
      name: null,
      login: null,
      password: null
    }
    async function register(){
      user.name = document.getElementById("name").value;
      user.login = document.getElementById("regLogin").value;
      user.password = document.getElementById("regPassword").value;
      let sendUrl = "http://localhost:8080/register";

      await fetch(sendUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(user)
      })
    }
    return{
      user,
      register
    }
  }
})
</script>

<style>
@import "register.css";
</style>