<template>
  <div class="header">
    <button class="openMenu" @click="openMenu">add bonds</button>
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
    }
  }
}

</script>

<style>
@import "./header.css";
</style>