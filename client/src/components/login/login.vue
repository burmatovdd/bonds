<template>
    <div class="container container-login">
        <h1>Welcome to Amenophis</h1>
        <div class="login">
            <Form @submit="onSubmit" class="form" v-slot="{ meta }">
                <Field name="login"
                       type="text"
                       class="input-login"
                       placeholder="login"
                       :rules="validateLogin"/>
                <ErrorMessage name="login"
                              class="error-message"/>
                <Field name="password"
                       type="password"
                       class="input-password"
                       placeholder="password"
                       :rules="validatePassword"/>
                <ErrorMessage name="password"
                              class="error-message"/>
                <button class="button"
                        :disabled="!meta.valid"
                        @click="login">
                    login
                </button>
            </Form>
            <button class="button button-register"
                    @click="register">
                register
            </button>
        </div>
        <div class="menu-register" :class="{overlay: !isActive}">
            <div class="menu-register_container">
                <Overlay/>
                <button class="button" @click="register">close</button>
            </div>
        </div>
        .<footer>
        <p>
            Amenophis is a learning project that aims to teach you how to start your own project from scratch. In the future, the project will develop.</p>
    </footer>
    </div>

</template>

<script>
import {defineComponent} from 'vue';
import {Form, Field, ErrorMessage} from 'vee-validate';
import Overlay from "../register/overlay.vue";
import *as storage from "../../storage";
import *as httpClient from "../../httpClient";


export default defineComponent({
    name: "login",
    components: {
        Overlay,
        Form,
        Field,
        ErrorMessage,
    },
    data: function () {
        const user = {
            login: null,
            password: null
        }
        return {
            isActive: false,
            user,
            currentPath: window.location.hash,
        }
    },
    methods: {
        register: function () {
            this.isActive = !this.isActive;
        },
        onSubmit(values) {
            this.user.login = values.login;
            this.user.password = values.password;
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
            this.user.login = value;
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
            this.user.password = value;
            return true;
        },
        login: async function () {
            console.log("login: ", this.user)

            let sendUrl = "http://localhost:8080/login";

            let postInfo = httpClient.Post(sendUrl,this.user);
            console.log("postInfo: ", postInfo);
            postInfo.then((data) => {
                if (data.response === true) {
                    console.log("user find")
                    storage.set("token",data.token.token);
                    this.$router.push('/mainMenu');
                } else {
                    console.log("invalid username or password")
                }
            });

        },
    }
})
</script>

<style>
@import "./login.css";
</style>
