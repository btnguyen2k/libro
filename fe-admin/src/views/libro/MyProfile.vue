<template>
  <CRow>
    <CCol sm="12">
      <CAlert v-if="waitLoadUserProfile" color="info">{{ $t('message.wait') }}</CAlert>
      <CAlert v-if="errorMsg" color="danger">{{ errorMsg }}</CAlert>

      <CForm v-if="!errorMsg && !waitLoadUserProfile">
        <CCard accent-color="info">
          <CCardHeader>
            <strong>{{ $t('message.my_profile') }}</strong>
          </CCardHeader>
          <CCardBody>
            <CAlert v-if="myFlashMsgProfile" color="success" closeButton>{{ myFlashMsgProfile }}</CAlert>
            <CAlert v-if="errorMsgProfile" color="danger">{{ errorMsgProfile }}</CAlert>
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

      <CForm v-if="!errorMsg && !waitLoadUserProfile">
        <CCard accent-color="danger">
          <CCardHeader>
            <strong>{{ $t('message.user_password') }}</strong>
          </CCardHeader>
          <CCardBody>
            <CAlert v-if="myFlashMsgPwd" color="success" closeButton>{{ myFlashMsgPwd }}</CAlert>
            <CAlert v-if="errorMsgPwd" color="danger">{{ errorMsgPwd }}</CAlert>
            <CInput
                type="password"
                v-model="currentPwd"
                :label="$t('message.user_current_password')"
                :placeholder="$t('message.user_current_password_msg')"
                v-c-tooltip.hover="$t('message.user_current_password_msg')"
                horizontal
            />
            <CInput
                type="password"
                v-model="newPwd"
                :label="$t('message.user_new_password')"
                :placeholder="$t('message.user_new_password_msg')"
                v-c-tooltip.hover="$t('message.user_new_password_msg')"
                horizontal
            />
            <CInput
                type="password"
                v-model="confirmedPwd"
                :label="$t('message.user_confirmed_password')"
                :placeholder="$t('message.user_confirmed_password_msg')"
                v-c-tooltip.hover="$t('message.user_confirmed_password_msg')"
                horizontal
            />
          </CCardBody>
          <CCardFooter>
            <CButton color="danger" size="sm"  @click="doChangePassword">
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
import forge from 'node-forge'

export default {
  name: 'MyProfile',
  mounted() {
    console.log(this.$i18n)
    let vue = this
    vue.waitLoadUserProfile = true
    clientUtils.apiDoGet(clientUtils.apiInfo,
        (apiRes) => {
          if (apiRes.status != 200) {
            vue.errorMsg = "apiRes.message"
            vue.waitLoadUserProfile = false
          } else {
            vue.publicKey = forge.pki.publicKeyFromPem(apiRes.data.rsa_public_key)
            vue.loadUserProfile()
          }
        },
        (err) => {
          vue.errorMsg = err
          vue.waitLoadUserProfile = false
        })
  },
  data() {
    return {
      errorMsg: '',
      waitLoadUserProfile: false,
      userProfile: {},
      publicKey: '',

      errorMsgProfile: '',
      myFlashMsgProfile: this.flashMsg,
      errorMsgPwd: '',
      myFlashMsgPwd: this.flashMsg,
      currentPwd: '',
      newPwd: '',
      confirmedPwd: '',
    }
  },
  props: ["flashMsg"],
  methods: {
    resetMsg() {
      this.errorMsg = ''
      this.errorMsgProfile = ''
      this.errorMsgPwd = ''
      this.myFlashMsgProfile = ''
      this.myFlashMsgPwd = ''
    },
    loadUserProfile() {
      this.resetMsg()
      let session = utils.loadLoginSession()
      const vue = this
      vue.waitLoadUserProfile = true
      clientUtils.apiDoGet(clientUtils.apiUser+"/"+session.uid,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.userProfile = apiRes.data
            } else {
              vue.errorMsgProfile = apiRes.status + ": " + apiRes.message
            }
            vue.waitLoadUserProfile = false
          },
          (err) => {
            vue.errorMsgProfile = err
            vue.waitLoadUserProfile = false
          })
    },
    doUpdateMyProfile(e) {
      e.preventDefault()
      this.resetMsg()
      let vue = this
      let session = utils.loadLoginSession()
      let formData = {"display_name": vue.userProfile.name}
      clientUtils.apiDoPut(clientUtils.apiUser+"/"+session.uid, formData,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.errorMsgProfile = apiRes.status + ": " + apiRes.message
            } else {
              vue.loadUserProfile()
              vue.myFlashMsgProfile = vue.$i18n.t('message.user_profile_updated_msg', {id: session.uid})
            }
            vue.waitLoadUserProfile = false
          },
          (err) => {
            vue.errorMsgProfile = err
            vue.waitLoadUserProfile = false
          }
      )
    },
    doChangePassword(e) {
      e.preventDefault()
      this.resetMsg()
      let vue = this
      if (!vue.publicKey) {
        this.errorMsgPwd = 'Invalid public key'
        return
      }
      vue.waitLoadUserProfile = true
      let currentPwd = vue.publicKey.encrypt(vue.currentPwd)
      let newPwd = vue.publicKey.encrypt(vue.newPwd)
      let confirmedPwd = vue.publicKey.encrypt(vue.confirmedPwd)
      let formData = {"current_pwd": forge.util.encode64(currentPwd), "new_pwd": forge.util.encode64(newPwd), "confirmed_pwd": forge.util.encode64(confirmedPwd)}
      let session = utils.loadLoginSession()
      clientUtils.apiDoPut(clientUtils.apiUserPassword+"/"+session.uid, formData,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.errorMsgPwd = apiRes.status + ": " + apiRes.message
              vue.waitLoadUserProfile = false
            } else {
              vue.currentPwd = ''
              vue.newPwd = ''
              vue.confirmedPwd = ''
              vue.loadUserProfile()
              vue.myFlashMsgPwd = vue.$i18n.t('message.user_password_updated_msg', {id: session.uid})
            }
          },
          (err) => {
            vue.errorMsgPwd = err
            vue.waitLoadUserProfile = false
          }
      )
    },
  }
}
</script>
