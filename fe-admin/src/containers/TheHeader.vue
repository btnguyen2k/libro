<!-- #GovueAdmin-Customized -->
<template>
  <CHeader fixed with-subheader light>
    <CToggler
        in-header
        class="ml-3 d-lg-none"
        @click="$store.commit('toggleSidebarMobile')"
    />
    <CToggler
        in-header
        class="ml-3 d-md-down-none"
        @click="$store.commit('toggleSidebarDesktop')"
    />
    <a href="/" class="c-header-brand mx-auto d-lg-none"
       style="font-weight: bolder; font-size: x-large">{{ appName }}</a>
    <!--    <CHeaderBrand class="mx-auto d-lg-none" to="/">-->
    <!--      <CIcon name="logo" height="48" alt="Logo"/>-->
    <!--    </CHeaderBrand>-->
    <CHeaderNav class="d-md-down-none mr-auto">
      <CHeaderNavItem class="px-3">
        <CHeaderNavLink :to="{name:'Dashboard'}">
          {{ $t('message.dashboard') }}
        </CHeaderNavLink>
      </CHeaderNavItem>
      <CHeaderNavItem class="px-3">
        <CHeaderNavLink :to="{name:'ProductList'}">
          {{ $t('message.products') }}
        </CHeaderNavLink>
      </CHeaderNavItem>
      <CHeaderNavItem class="px-3">
        <CHeaderNavLink :to="{name:'UserList'}">
          {{ $t('message.users') }}
        </CHeaderNavLink>
      </CHeaderNavItem>
    </CHeaderNav>
    <CHeaderNav class="mr-4">
      <CDropdown inNav class="c-header-nav-items" placement="bottom-end" add-menu-classes="pt-0">
        <template #toggler>
          <CHeaderNavLink>
            <CIcon name="cil-flag-alt"/>
          </CHeaderNavLink>
        </template>
        <CDropdownItem v-for="(locale, _) in $i18n.availableLocales" @click="doSwitchLanguage(locale)">
          <CIcon :name="$i18n.messages[locale]._flag"/>
          <span class="px-2">{{ $i18n.messages[locale]._name }}</span>
        </CDropdownItem>
      </CDropdown>
      <TheHeaderDropdownAccnt/>
    </CHeaderNav>
    <CSubheader class="px-3">
      <CBreadcrumbRouter class="border-0 mb-0"/>
    </CSubheader>
  </CHeader>
</template>

<script>
import TheHeaderDropdownAccnt from './TheHeaderDropdownAccnt'
import cfg from '@/utils/app_config'

export default {
  name: 'TheHeader',
  computed: {
    appName() {
      return cfg.APP_CONFIG.app.name
    },
  },
  components: {
    TheHeaderDropdownAccnt
  },
  methods: {
    doSwitchLanguage(locale) {
      this.$i18n.locale = locale
    },
  }
}
</script>
