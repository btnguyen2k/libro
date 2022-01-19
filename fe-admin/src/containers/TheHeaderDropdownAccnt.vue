<!-- #GovueAdmin-Customized -->
<template>
  <CDropdown inNav class="c-header-nav-items" placement="bottom-end" add-menu-classes="pt-0">
    <template #toggler>
      <CHeaderNavLink>
        <div class="c-avatar">
          <img :src="avatar" class="c-avatar-img" :title="displayName"/>
        </div>
      </CHeaderNavLink>
    </template>
    <CDropdownHeader tag="div" class="text-center" color="light">
      <strong>{{ $t('message.account') }}</strong>
    </CDropdownHeader>
    <CDropdownItem @click="doViewProfile">
      <CIcon name="cil-user"/>
      {{ $t('message.my_profile') }}
    </CDropdownItem>
    <CDropdownItem @click="doLogout">
      <CIcon name="cil-lock-locked"/>
      {{ $t('message.logout') }}
    </CDropdownItem>
  </CDropdown>
</template>

<script>
import utils from "@/utils/app_utils"
import appConfig from "@/utils/app_config"
import MD5 from "crypto-js/md5"

export default {
  name: 'TheHeaderDropdownAccnt',
  data() {
    String.prototype.md5 = function () {
      return MD5(this)
    }
    let session = utils.loadLoginSession()
    let uid = session != null ? session.uid : ""
    return {
      itemsCount: 42,
      displayName: session != null ? session.name : uid,
      avatar: "https://www.gravatar.com/avatar/" + uid.trim().toLowerCase().md5() + "?s=40",
    }
  },
  methods: {
    funcNotImplemented() {
      alert("Not implemented")
    },
    doViewProfile() {
      this.$router.push({name: "MyProfile"})
    },
    doLogout() {
      utils.localStorageSet(utils.lskeyLoginSession, null)
      utils.localStorageSet(utils.lskeyLoginSessionLastCheck, null)
      this.$router.push({name: "Login", query: {app: appConfig.APP_ID}})
    }
  }
}
</script>

<style scoped>
.c-icon {
  margin-right: 0.3rem;
}
</style>