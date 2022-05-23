<template>
<div class="container">
  <AddBonds/>
</div>
</template>

<script>
import AddBonds from '../addBonds/addBonds.vue'
export default {
  name: "overlay",
  components: {
    AddBonds
  }

}
</script>

<style>
@import "overlay.css";
</style>
