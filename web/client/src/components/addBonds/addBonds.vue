<template>
<div class="container">
  <div class="add">
    <input id="inputBond" type="text" class="input" placeholder="enter bond">
    <input id="inputBondCount" type="text" class="input" placeholder="enter count">
    <button class="addBonds" @click="sendBond">Add</button>
  </div>
</div>
</template>

<script>
import { defineComponent } from 'vue';
import *as storage from "../../storage";
import { useRouter, useRoute } from 'vue-router'
import *as httpClient from "../../httpClient";


export default defineComponent({
  name: "addBonds",
  setup(){
      const router = useRouter();
    let bonds = {
        bond: {
            name: null,
        },
      count: null,
    }
    let token = storage.get("token");
    if (token == null){
        router.push("/");
    }
    async function sendBond() {
      bonds.bond.name = document.getElementById("inputBond").value;
      bonds.count = +document.getElementById("inputBondCount").value;// + - int

      let sendUrl = "/api/bonds";

      let postInfo = httpClient.Post(sendUrl,bonds,token);

    }
    return {
      bonds,
      sendBond
    }
  }
})
</script>

<style>
@import "./addBonds.css";
</style>
