<template>
  <div v-if="foundStatus<0" class="alert alert-info m-4" role="alert">{{ $t('wait') }}</div>
  <div v-else-if="foundStatus==0" class="alert alert-danger m-4" role="alert">
    {{ $t('error_product_not_found', {domain: currentHost}) }}
  </div>
  <div v-else-if="targetTopic==null" class="alert alert-danger m-4" role="alert">
    {{ $t('error_product_not_found', {domain: currentHost}) }}
  </div>
  <div v-else>
    <header class="header fixed-top">
      <div class="branding docs-branding">
        <div class="container-fluid position-relative py-2">
          <div class="docs-logo-wrapper">
            <button ref="docs-sidebar-toggler" id="docs-sidebar-toggler" class="docs-sidebar-toggler docs-sidebar-visible me-2 d-xl-none" type="button">
              <span></span>
              <span></span>
              <span></span>
            </button>
            <div class="site-logo">
              <a class="navbar-brand" href="javascript:;" @click="goHome">
                <img class="logo-icon me-2" src="images/coderdocs-logo.svg" alt="logo">
                <span class="logo-text">{{ prodNameFirst }}<span class="text-alt">{{ prodNameLast }}</span></span>
              </a>
            </div>
          </div>
          <div class="docs-top-utilities d-flex justify-content-end align-items-center">
            <div class="top-search-box d-none d-lg-flex">
              <form class="search-form" @submit.prevent="popup('not implemented yet')">
                <input type="text" placeholder="Search the docs..." name="search" class="form-control search-input">
                <button type="submit" class="btn search-btn" value="Search">
                  <ficon :icon="['fas', 'search']"/>
                </button>
              </form>
            </div>
            <ul class="social-list list-inline mx-md-3 mx-lg-5 mb-0 d-none d-lg-flex">
              <li v-if="product.contacts.website" class="list-inline-item">
                <a :href="product.contacts.website" title="Website"><ficon fixedWidth :icon="['fas', 'globe']"/></a>
              </li>
              <li v-if="product.contacts.github" class="list-inline-item">
                <a :href="product.contacts.github" title="GitHub"><ficon fixedWidth :icon="['fab', 'github']"/></a>
              </li>
              <li v-if="product.contacts.facebook" class="list-inline-item">
                <a :href="product.contacts.facebook" title="Facebook"><ficon fixedWidth :icon="['fab', 'facebook']"/></a>
              </li>
              <li v-if="product.contacts.linkedin" class="list-inline-item">
                <a :href="product.contacts.linkedin" title="LinkedIn"><ficon fixedWidth :icon="['fab', 'linkedin']"/></a>
              </li>
              <li v-if="product.contacts.slack" class="list-inline-item">
                <a :href="product.contacts.slack" title="Slack"><ficon fixedWidth :icon="['fab', 'slack']"/></a>
              </li>
              <li v-if="product.contacts.twitter" class="list-inline-item">
                <a :href="product.contacts.twitter" title="Twitter"><ficon fixedWidth :icon="['fab', 'twitter']"/></a>
              </li>
            </ul>
            <!--            <a href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderdocs-free-bootstrap-5-documentation-template-for-software-projects/" class="btn btn-primary d-none d-lg-flex">Download</a>-->
          </div>
        </div>
      </div>
    </header>

    <div class="docs-wrapper">
      <div ref="docs-sidebar" id="docs-sidebar" class="docs-sidebar">
        <div class="top-search-box d-lg-none p-3">
          <form class="search-form">
            <input type="text" placeholder="Search the docs..." name="search" class="form-control search-input">
            <button type="submit" class="btn search-btn" value="Search"><i class="fas fa-search"></i></button>
          </form>
        </div>
        <nav id="docs-nav" class="docs-nav navbar">
          <ul class="section-items list-unstyled nav flex-column pb-3">
            <li class="nav-item section-title"><a class="nav-link scrollto active" href="#section-1"><span class="theme-icon-holder me-2"><i class="fas fa-map-signs"></i></span>Introduction</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-1-1">Section Item 1.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-1-2">Section Item 1.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-1-3">Section Item 1.3</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-1-4">Section Item 1.4</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-1-5">Section Item 1.5</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-1-6">Section Item 1.6</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-2"><span class="theme-icon-holder me-2"><i class="fas fa-arrow-down"></i></span>Installation</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-2-1">Section Item 2.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-2-2">Section Item 2.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-2-3">Section Item 2.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-3"><span class="theme-icon-holder me-2"><i class="fas fa-box"></i></span>APIs</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-3-1">Section Item 3.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-3-2">Section Item 3.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-3-3">Section Item 3.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-4"><span class="theme-icon-holder me-2"><i class="fas fa-cogs"></i></span>Integrations</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-4-1">Section Item 4.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-4-2">Section Item 4.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-4-3">Section Item 4.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-5"><span class="theme-icon-holder me-2"><i class="fas fa-tools"></i></span>Utilities</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-5-1">Section Item 5.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-5-2">Section Item 5.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-5-3">Section Item 5.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-6"><span class="theme-icon-holder me-2"><i class="fas fa-laptop-code"></i></span>Web</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-6-1">Section Item 6.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-6-2">Section Item 6.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-6-3">Section Item 6.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-7"><span class="theme-icon-holder me-2"><i class="fas fa-tablet-alt"></i></span>Mobile</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-7-1">Section Item 7.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-7-2">Section Item 7.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-7-3">Section Item 7.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-8"><span class="theme-icon-holder me-2"><i class="fas fa-book-reader"></i></span>Resources</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-8-1">Section Item 8.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-8-2">Section Item 8.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-8-3">Section Item 8.3</a></li>
            <li class="nav-item section-title mt-3"><a class="nav-link scrollto" href="#section-9"><span class="theme-icon-holder me-2"><i class="fas fa-lightbulb"></i></span>FAQs</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-9-1">Section Item 9.1</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-9-2">Section Item 9.2</a></li>
            <li class="nav-item"><a class="nav-link scrollto" href="#item-9-3">Section Item 9.3</a></li>
          </ul>
        </nav>
      </div>

    </div>
  </div>
</template>

<script>
import clientUtils from "@/utils/api_client"
import {iconize} from "@/components/utils"

export default {
  name: 'Topic',
  unmounted() {
    window.removeEventListener('resize', this.responsiveSidebar)
  },
  mounted() {
    this.foundStatus = -1
    const vue = this
    const apiUrl = clientUtils.apiProduct.replaceAll(':domain', vue.currentHost)
    clientUtils.apiDoGet(apiUrl,
        (apiRes) => {
          vue.foundStatus = apiRes.status == 200 ? 1 : 0
          if (vue.foundStatus == 1) {
            console.log(apiRes)
            vue.product = apiRes.data
            vue.topicList = vue.product.topics
            for (let i = vue.topicList.length-1; i >= 0; i--) {
              const id = vue.topicList[i].id
              vue.topicMap[id] = vue.topicList[i]
            }
            vue.foundStatus = vue.product.is_published ? 1 : 0

            const tid = vue.$route.params.tid
            if (vue.foundStatus && !vue.topicMap[tid]) {
              console.log(vue.$i18n.t('error_topic_not_found', {topic: tid}))
              vue.$router.push({
                name: "Error",
                params: {errorMsg: vue.$i18n.t('error_topic_not_found', {topic: tid}), target: 'Home'}
              })
            } else {
              // CoderDocs
              window.addEventListener('resize', this.responsiveSidebar)
              vue.$nextTick(()=>vue.responsiveSidebar())
            }
          }
        },
        (err) => {
          vue.errorMsg = err
        })
  },
  computed: {
    prodNameFirst() {
      return this.product.name ? this.product.name.slice(0, 2) : ""
    },
    prodNameLast() {
      return this.product.name ? this.product.name.slice(2) : ""
    },
    currentHost() {
      return window.location.host
    },
    targetTopic() {
      const tid = this.$route.params.pid
      return this.topicMap[tid] ? this.topicMap[tid] : null
    },
  },
  methods: {
    _iconize(icon) {
      return iconize(icon)
    },
    goHome() {
      this.$router.push({name: "Home"})
    },
    popup(msg) {
      alert(msg)
    },
    responsiveSidebar() { // CoderDocs
      const w = window.innerWidth
      const sidebar = this.$refs['docs-sidebar']
      if (sidebar) {
        if (w >= 1024) {
          sidebar.classList.remove('sidebar-hidden')
          sidebar.classList.add('sidebar-visible')
        } else {
          sidebar.classList.remove('sidebar-visible')
          sidebar.classList.add('sidebar-hidden')
        }
      }
    }
  },
  data() {
    return {
      product: {},
      topicList: [],
      topicMap: {},
      topic: undefined,
      foundStatus: -1,
      errorMsg: '',
    }
  },
}
</script>
