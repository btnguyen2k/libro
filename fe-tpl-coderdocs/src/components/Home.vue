<template>
  <div v-if="foundStatus<0" class="alert alert-info" role="alert">{{ $t('message.wait') }}</div>
  <div v-else-if="foundStatus==0" class="alert alert-danger" role="alert">
    {{ $t('message.error_product_not_found', {domain: currnetHost}) }}
  </div>
  <div v-else>
    <header class="header fixed-top">
      <div class="branding docs-branding">
        <div class="container-fluid position-relative py-2">
          <div class="docs-logo-wrapper">
            <div class="site-logo">
              <a class="navbar-brand" href="index.html"><img class="logo-icon me-2" src="images/coderdocs-logo.svg"
                                                             alt="logo">
                <span class="logo-text">{{ prodNameFirst }}<span class="text-alt">{{ prodNameLast }}</span></span>
              </a>
            </div>
          </div>
          <div class="docs-top-utilities d-flex justify-content-end align-items-center">
            <ul class="social-list list-inline mx-md-3 mx-lg-5 mb-0 d-none d-lg-flex">
              <li class="list-inline-item"><a href="#">
                <ficon fixedWidth :icon="['fab', 'github']"/>
              </a></li>
              <li class="list-inline-item"><a href="#">
                <ficon fixedWidth :icon="['fab', 'twitter']"/>
              </a></li>
              <li class="list-inline-item"><a href="#">
                <ficon fixedWidth :icon="['fab', 'slack']"/>
              </a></li>
              <li class="list-inline-item"><a href="#">
                <ficon fixedWidth :icon="['fab', 'product-hunt']"/>
              </a></li>
            </ul>
            <a href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderdocs-free-bootstrap-5-documentation-template-for-software-projects/"
               class="btn btn-primary d-none d-lg-flex">Download</a>
          </div>
        </div>
      </div>
    </header>

    <div class="page-header theme-bg-dark py-5 text-center position-relative">
      <div class="theme-bg-shapes-right"></div>
      <div class="theme-bg-shapes-left"></div>
      <div class="container">
        <h1 class="page-heading single-col-max mx-auto">Documentation</h1>
        <div class="page-intro single-col-max mx-auto">Everything you need to get your software documentation online.
        </div>
        <div class="main-search-box pt-3 d-block mx-auto">
          <form class="search-form w-100">
            <input type="text" placeholder="Search the docs..." name="search" class="form-control search-input">
            <button type="submit" class="btn search-btn" value="Search">
              <ficon :icon="['fas', 'search']"/>
            </button>
          </form>
        </div>
      </div>
    </div>

    <div class="page-content">
      <div class="container">
        <div class="docs-overview py-5">
          <div class="row justify-content-center">
            <div v-for="topic in topicList" class="col-12 col-lg-4 py-3">
              <div class="card shadow-sm">
                <div class="card-body">
                  <h5 class="card-title mb-3">
                    <span class="theme-icon-holder card-icon-holder me-2"><ficon :icon="['fas', iconize(topic.icon)]"/></span>
                    <span class="card-title-text">{{ topic.title }}</span>
                  </h5>
                  <div class="card-text">{{ topic.summary }}</div>
                  <a class="card-link-mask" href="docs-page.html#section-1"></a>
                </div>
              </div>
            </div>

            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-arrow-down"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">Installation</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-2"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-box fa-fw"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">APIs</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-3"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-cogs fa-fw"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">Integrations</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-4"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-tools"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">Utilities</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-5"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-laptop-code"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">Web</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-6"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-tablet-alt"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">Mobile</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-7"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-book-reader"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">Resources</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-8"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
            <!--            <div class="col-12 col-lg-4 py-3">-->
            <!--              <div class="card shadow-sm">-->
            <!--                <div class="card-body">-->
            <!--                  <h5 class="card-title mb-3">-->
            <!--								    <span class="theme-icon-holder card-icon-holder me-2">-->
            <!--								        <i class="fas fa-lightbulb"></i>-->
            <!--							        </span>&lt;!&ndash;//card-icon-holder&ndash;&gt;-->
            <!--                    <span class="card-title-text">FAQs</span>-->
            <!--                  </h5>-->
            <!--                  <div class="card-text">-->
            <!--                    Section overview goes here. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.-->
            <!--                  </div>-->
            <!--                  <a class="card-link-mask" href="docs-page.html#section-9"></a>-->
            <!--                </div>&lt;!&ndash;//card-body&ndash;&gt;-->
            <!--              </div>&lt;!&ndash;//card&ndash;&gt;-->
            <!--            </div>&lt;!&ndash;//col&ndash;&gt;-->
          </div>
        </div>
      </div>
    </div>

    <section class="cta-section text-center py-5 theme-bg-dark position-relative">
      <div class="theme-bg-shapes-right"></div>
      <div class="theme-bg-shapes-left"></div>
      <div class="container">
        <h3 class="mb-2 text-white mb-3">Launch Your Software Project Like A Pro</h3>
        <div class="section-intro text-white mb-3 single-col-max mx-auto">Want to launch your software project and start
          getting traction from your target users? Check out our premium <a class="text-white"
                                                                            href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderpro-bootstrap-5-startup-template-for-software-projects/">Bootstrap
            5 startup template CoderPro</a>! It has everything you need to promote your product.
        </div>
        <div class="pt-3 text-center">
          <a class="btn btn-light"
             href="https://themes.3rdwavemedia.com/bootstrap-templates/startup/coderpro-bootstrap-5-startup-template-for-software-projects/">
            Get CoderPro
            <ficon class="ms-1" :icon="['fas', 'arrow-alt-circle-right']"/>
          </a>
        </div>
      </div>
    </section>

    <footer class="footer">
      <div class="footer-bottom text-center py-5">
        <ul class="social-list list-unstyled pb-4 mb-0">
          <li class="list-inline-item"><a href="#">
            <ficon fixedWidth :icon="['fab', 'github']"/>
          </a></li>
          <li class="list-inline-item"><a href="#">
            <ficon fixedWidth :icon="['fab', 'twitter']"/>
          </a></li>
          <li class="list-inline-item"><a href="#">
            <ficon fixedWidth :icon="['fab', 'slack']"/>
          </a></li>
          <li class="list-inline-item"><a href="#">
            <ficon fixedWidth :icon="['fab', 'product-hunt']"/>
          </a></li>
          <li class="list-inline-item"><a href="#">
            <ficon fixedWidth :icon="['fab', 'facebook-f']"/>
          </a></li>
          <li class="list-inline-item"><a href="#">
            <ficon fixedWidth :icon="['fab', 'instagram']"/>
          </a></li>
        </ul>

        <!--/* This template is free as long as you keep the footer attribution link. If you'd like to use the template without the attribution link, you can buy the commercial license via our website: themes.3rdwavemedia.com Thank you for your support. :) */-->
        <small class="copyright">Designed with
          <ficon style="color: #fb866a;" :icon="['fas', 'heart']"/>
          by <a class="theme-link" href="http://themes.3rdwavemedia.com" target="_blank">Xiaoying Riley</a> for
          developers</small>
      </div>
    </footer>
  </div>
</template>

<script>
import clientUtils from "@/utils/api_client"

export default {
  name: 'Home',
  mounted() {
    this.foundStatus = -1
    const vue = this
    const apiUrl = clientUtils.apiProduct.replaceAll(':domain', vue.currnetHost)
    clientUtils.apiDoGet(apiUrl,
        (apiRes) => {
          vue.foundStatus = apiRes.status == 200 ? 1 : 0
          if (vue.foundStatus == 1) {
            vue.product = apiRes.data
            vue.topicList = vue.product.topics
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
    currnetHost() {
      return window.location.host
    },
  },
  methods: {
    iconize(icon) {
      if (icon.startsWith("cil-")) {
        return icon.slice(4)
      }
      if (icon.startsWith("fa-")) {
        return icon.slice(3)
      }
      return icon
    },
  },
  data() {
    return {
      product: {},
      topicList: [],
      foundStatus: -1,
      errorMsg: "",
    }
  },
}
</script>
