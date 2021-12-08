<template>
  <router-view></router-view>
</template>

<style lang="css">
@import "assets/css/theme.css";
@import "assets/plugins/simplelightbox/simple-lightbox.min.css";
/*@import "//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.15.2/styles/atom-one-dark.min.css";*/
@import "//fonts.googleapis.com/css?family=Poppins:300,400,500,600,700&display=swap";
</style>

<script>
// import "./assets/plugins/popper.min.js";
// import "./assets/plugins/bootstrap/js/bootstrap.js";
// import "./assets/plugins/smoothscroll.min.js"
// // import "//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.15.8/highlight.min.js"
// import "./assets/js/highlight-custom.js"
// import "./assets/plugins/simplelightbox/simple-lightbox.min.js"
// import "./assets/plugins/gumshoe/gumshoe.polyfills.min.js"
// import "./assets/js/docs.js"

import {triggerPopstate, triggerResize} from '@/components/utils'
window.addEventListener('popstate', () => triggerPopstate())
window.addEventListener('resize', () => triggerResize())

export default {
  name: 'App'
}
</script>
