//#Libro frontend, template CoderDocs
import 'core-js/stable'
import Vue from 'vue'
import App from './App'
import router from './router'
import i18n from './i18n'

import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(fab, far, fas)
Vue.component('ficon', FontAwesomeIcon)

// import { BootstrapIcon } from 'bootstrap-icons/font/bootstrap-icons'
// Vue.component('bicon', BootstrapIcon)

Vue.config.performance = true
Vue.prototype.$log = console.log.bind(console)

new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  },
  i18n,
})
