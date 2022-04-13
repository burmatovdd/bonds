<template>
<div class="container">
  <div class="login">
    <input id="login" type="text" class="input-login" placeholder="login">
    <input id="password" type="password" class="input-password" placeholder="password">
    <button class="button" @click="login">login</button>
    <button class="button button-register" @click="register">register</button>
  </div>
  <div class="menu-register" :class="{overlay: !isActive}">
    <div class="menu-register_container">
      <Overlay/>
      <button class="button" @click="register">close</button>
    </div>
  </div>
</div>
</template>

<script>
import { defineComponent } from 'vue';
import Overlay from "../register/overlay.vue";
export default defineComponent ({
  name: "login",
  components:{
    Overlay
  },
  data: function(){
    let user = {
      login: null,
      password: null
    }
    return{
      isActive: false,
      user
    }
  },
  methods: {
    register: function() {
      this.isActive = !this.isActive;
    },
    login: async function(){
      this.user.login = document.getElementById("login").value
      this.user.password = document.getElementById("password").value
      let sendUrl = "http://localhost:8080/login";

      await fetch(sendUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(this.user)
      })
    }
  }
})
</script>

<style>
@import "./login.css";
</style>