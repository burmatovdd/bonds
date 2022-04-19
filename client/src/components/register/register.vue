<template>
      <div class="register">
<!--        <input id="name" type="text" class="register-item register-name" placeholder="Name">-->
<!--        <input id="regLogin" type="text" class="register-item register-login" placeholder="login">-->
<!--        <input id="regPassword" type="password" class="register-item register-password" placeholder="password">-->
<!--        <input type="password" class="register-item register-password" placeholder="password">-->
        <Form  class="register">
          <Field name="name" type="text" class="register-item register-name" placeholder="Name" />
          <ErrorMessage name="login" class="error-message"/>
          <Field name="login" type="text" class="register-item register-login" placeholder="login" />
          <ErrorMessage name="login" class="error-message"/>
          <Field name="password" type="password" class="register-item register-password" placeholder="password" />
          <ErrorMessage name="password" class="error-message"/>
          <Field name="password" type="password" class="register-item register-password" placeholder="password" />
          <ErrorMessage name="password" class="error-message"/>
          <button class="button" @click="register">submit</button>
        </Form>
      </div>
</template>

<script>
import { defineComponent } from 'vue';
import { Form, Field,ErrorMessage } from 'vee-validate';
export default  defineComponent({
  name: "register",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
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