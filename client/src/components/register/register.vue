<template>
    <div class="register">
        <Form @submit="onSubmit" class="register" v-slot="{ meta }">
            <Field name="name"
                   type="text"
                   class="register-item register-name"
                   placeholder="Name"
                   :rules="validateName"
            />
            <ErrorMessage name="name"
                          class="error-message_register"
            />
            <Field name="login"
                   type="text"
                   class="register-item register-login"
                   placeholder="login"
                   :rules="validateLogin"
            />
            <ErrorMessage name="login"
                          class="error-message_register"
            />
            <Field name="password"
                   type="password"
                   class="register-item register-password"
                   placeholder="password"
                   :rules="validatePassword"
            />
            <ErrorMessage name="password"
                          class="error-message_register"
            />
            <button class="button"
                    @click="register"
                    :disabled="!meta.valid">
                submit
            </button>
        </Form>
    </div>
</template>

<script>
import {defineComponent} from 'vue';
import {Form, Field, ErrorMessage} from 'vee-validate';
import *as httpClient from "../../httpClient";
import *as storage from "../../storage";

let token = null;

export default defineComponent({
    name: "register",
    components: {
        Form,
        Field,
        ErrorMessage,
    },
    data: function () {
        const userRegister = {
            name: null,
            login: null,
            password: null,
        }
        return {
            userRegister,
            token,
        }
    },
    methods: {
        onSubmit(values) {
            console.log("onSubmit: ", values);
        },
        validateName(value) {
            if (!value) {
                return 'This field is required';
            }
            if (value.length < 2) {
                return 'name must be at least 2 characters long'
            }
            // if the field is not a valid email
            const regex = /^[A-Za-z]/i;
            if (!regex.test(value)) {
                return 'name can contain only letters and numbers';
            }
            this.userRegister.name = value;
            // All is good
            return true;
        },
        validateLogin(value) {
            // if the field is empty
            if (!value) {
                return 'This field is required';
            }
            if (value.length < 5) {
                return 'login must be at least 5 characters long'
            }
            // if the field is not a valid email
            const regex = /^[A-Za-z0-9]/i;
            if (!regex.test(value)) {
                return 'login can contain only letters and numbers';
            }
            this.userRegister.login = value;
            // All is good
            return true;
        },
        validatePassword(value) {
            if (!value) {
                return 'This field is required';
            }
            if (value.length < 8) {
                return 'password must be at least 8 characters long'
            }
            const regex = /^[A-Za-z0-9]/i;
            if (!regex.test(value)) {
                return 'password can contain only letters and numbers';
            }
            this.userRegister.password = value;
            return true;
        },
        async register() {
            console.log("register: ", this.userRegister)
            // вынести в отдельную функцию, then.this.$router.push('/mainMenu');
            let sendUrl = "/api/register";

            let postInfo = httpClient.Post(sendUrl,this.userRegister);
            console.log("postInfo: ", postInfo);
            postInfo.then(async (data) => {
                if (data.res === false) {
                    console.log("user already exist!")
                    return 0;
                }
                storage.set("token", data.token.token);
                this.$router.push('/mainMenu');
            });

        }
    }
})
</script>

<style>
@import "register.css";
</style>
