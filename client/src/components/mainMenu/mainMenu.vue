<template>
    <HeaderComponent/>
    <img alt="Vue logo" src="../../assets/duck.png" />
    <InputYear/>
</template>

<script>
import InputYear from '../inputYear/inputYear.vue'
import HeaderComponent from '../header/header.vue'
export default {
  name: "mainMenu",
    components:{
      InputYear,
        HeaderComponent,
    },
}
</script>

<style scoped>

</style>
