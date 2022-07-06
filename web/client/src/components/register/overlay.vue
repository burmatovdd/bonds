<template>
<div class="container">
  <Register/>
</div>
</template>

<script>
import Register from "./register.vue"
export default {
  name: "overlay",
  components: {
    Register
  }

}
</script>

<style>
@import "register.css";
</style>