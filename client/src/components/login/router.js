import {createRouter, createWebHistory} from 'vue-router';
import Login from "./login.vue";
import MainMenu from "../mainMenu/mainMenu.vue";

export default  createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Login },
        { path: '/mainMenu', component: MainMenu },
    ],
})
