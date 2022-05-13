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


export default defineComponent({
  name: "addBonds",
  setup(){
      const router = useRouter();
    let bonds = {
      bond: null,
      count: null,
    }
    let token = storage.get("token");
    if (token == null){
        router.push("/");
    }
    async function sendBond() {
      bonds.bond = document.getElementById("inputBond").value;
      bonds.count = document.getElementById("inputBondCount").value;

      let sendUrl = "http://localhost:8080/bonds";

      await fetch(sendUrl, {
        method: 'POST',
        headers: {
            'Authorization': token,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(bonds)
      })
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
