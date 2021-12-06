<template>
  <div class="m-4">
    <div class="alert alert-danger" role="alert">
      {{ myErrorMsg }}
    </div>
    <a href="javascript:;" @click="clickToHome">{{ $t('click_to_home') }}</a>
  </div>
</template>

<script>
export default {
  name: 'Error',
  props: ["errorMsg", "target"],
  computed: {
    myErrorMsg() {
      return this.errorMsg ? this.errorMsg : this.$i18n.t('error_unknown')
    },
  },
  methods: {
    clickToHome() {
      const nextPage = this.target ? this.target : 'Home'
      this.$router.push({name: nextPage})
    },
  }
}
</script>
