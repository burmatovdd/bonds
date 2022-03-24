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
export default defineComponent({
  name: "addBonds",
  setup(){
    let bonds = {
      bond: null,
      count: null,
    }
    async function sendBond() {
      bonds.bond = document.getElementById("inputBond").value;
      bonds.count = document.getElementById("inputBondCount").value;

      let sendUrl = "http://localhost:8080/bonds";

      await fetch(sendUrl, {
        method: 'POST',
        headers: {
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