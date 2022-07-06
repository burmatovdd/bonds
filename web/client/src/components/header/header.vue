<template>
  <div class="header">
    <div class="buttons">
        <button class="openMenu" @click="openMenu">add bonds</button>
        <button class="openMenu exit" @click="exit">exit</button>
    </div>
  </div>
  <div class="menu" :class="{overlay: !isActive}">
    <div class="menuContainer">
      <Overlay/>
      <button class="closeMenu" @click="openMenu">close</button>
    </div>
  </div>

</template>

<script>
import Overlay from "../overlay/overlay.vue";
import *as storage from "../../storage";

export default{
  name: "header",
  components: {
    Overlay,
  },
  data: function(){
    return{
      isActive: false
    }
  },
  methods: {
    openMenu: function() {
      this.isActive = !this.isActive;
    },
      exit: function () {
        storage.remove("token");
          this.$router.push('/');
      }
  }
}

</script>

<style>
@import "./header.css";
</style>
