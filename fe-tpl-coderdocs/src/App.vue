<template>
  <router-view></router-view>
</template>

<style lang="css">
@import "assets/css/theme.css";
@import "//fonts.googleapis.com/css?family=Poppins:300,400,500,600,700&display=swap";
</style>

<script>
import {triggerPopstate, triggerResize} from '@/components/utils'
window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App'
}
</script>
