<template>
  <CRow>
    <CCol sm="12">
      <CForm>
        <CCard accent-color="info">
          <CCardHeader>
            <strong>{{ $t('message.my_profile') }}</strong>
          </CCardHeader>
          <CCardBody>
            <CAlert v-if="myFlashMsg" color="success" closeButton>{{ myFlashMsg }}</CAlert>
            <CAlert v-if="waitLoadUserProfile" color="info">{{ $t('message.wait') }}</CAlert>
            <CAlert v-if="errorMsg" color="danger">{{ errorMsg }}</CAlert>
            <div class="form-group form-row mt-2">
              <CCol :sm="{offset:3,size:9}" class="form-inline">
                <CInputCheckbox inline :label="$t('message.user_is_admin')" :checked.sync="userProfile.is_admin" disabled/>
                <small>({{ $t('message.user_is_admin_msg') }})</small>
              </CCol>
            </div>
            <CInput
                type="text"
                v-model="userProfile.id"
                :label="$t('message.user_id')"
                :placeholder="$t('message.user_id_msg')"
                v-c-tooltip.hover="$t('message.user_id_msg')"
                horizontal readonly
            />
            <CInput
                type="text"
                v-model="userProfile.mid"
                :label="$t('message.user_mask_id')"
                :placeholder="$t('message.user_mask_id_msg')"
                v-c-tooltip.hover="$t('message.user_mask_id_msg')"
                horizontal readonly
            />
            <CInput
                type="text"
                v-model="userProfile.name"
                :label="$t('message.user_display_name')"
                :placeholder="$t('message.user_display_name_msg')"
                v-c-tooltip.hover="$t('message.user_display_name_msg')"
                horizontal
            />
          </CCardBody>
          <CCardFooter>
            <CButton color="primary" size="sm"  @click="doUpdateMyProfile">
              <CIcon name="cil-save"/>
              {{ $t('message.action_save') }}
            </CButton>
          </CCardFooter>
        </CCard>
      </CForm>
    </CCol>
  </CRow>
</template>

<script>
import clientUtils from "@/utils/api_client"
import utils from "@/utils/app_utils"

export default {
  name: 'MyProfile',
  mounted() {
    this.loadUserProfile()
  },
  data() {
    return {
      errorMsg: '',
      myFlashMsg: this.flashMsg,
      waitLoadUserProfile: false,
      userProfile: {},
    }
  },
  props: ["flashMsg"],
  methods: {
    loadUserProfile() {
      let session = utils.loadLoginSession()
      const vue = this
      vue.waitLoadUserProfile = true
      clientUtils.apiDoGet(clientUtils.apiUser+"/"+session.uid,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.userProfile = apiRes.data
            } else {
              vue.errorMsg = apiRes.status + ": " + apiRes.message
            }
            vue.waitLoadUserProfile = false
          },
          (err) => {
            vue.errorMsg = err
            vue.waitLoadUserProfile = false
          })
    },
    doUpdateMyProfile(e) {
      e.preventDefault()
      this.errorMsg = ''
      let vue = this
      let session = utils.loadLoginSession()
      let formData = {"display_name": vue.userProfile.name}
      clientUtils.apiDoPut(clientUtils.apiUser+"/"+session.uid, formData,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.errorMsg = apiRes.status + ": " + apiRes.message
            } else {
              vue.loadUserProfile()
              vue.myFlashMsg = vue.$i18n.t('message.user_profile_updated_msg', {id: session.uid})
            }
            vue.waitEditProduct = false
          },
          (err) => {
            vue.modalEditErr = err
            vue.waitEditProduct = false
          }
      )
    },
  }
}
</script>
