<template>
  <CRow>
    <CCol sm="12">
      <CCard accent-color="info">
        <CCardHeader>
          <strong>{{ $t('message.users') }}</strong>
          <div v-if="errorMsg==''" class="card-header-actions">
            <CButton class="btn-sm btn-primary" @click="clickAddUser" v-if="isAdmin">
              <CIcon name="cil-user-plus" class="align-top"/>
              {{ $t('message.add_user') }}
            </CButton>
          </div>
        </CCardHeader>
        <CCardBody>
          <CAlert v-if="myFlashMsg" color="success" closeButton>{{ myFlashMsg }}</CAlert>
          <CAlert v-if="waitLoadUserList" color="info">{{ $t('message.wait') }}</CAlert>
          <CAlert v-if="errorMsg" color="danger">{{ errorMsg }}</CAlert>
          <CDataTable v-if="errorMsg==''" :items="userList" :fields="[
              {key:'is_admin', label:''},
              {key:'id', label:$t('message.user_id')},
              {key:'name', label:$t('message.user_display_name')},
              {key:'actions', label:$t('message.actions'),_style:'text-align: center'}
            ]">
            <template #is_admin="{item}">
              <td class="col-1">
                <CIcon :name="`${item.is_admin?'cil-check':'cil-check-alt'}`" :style="`color: ${item.is_admin?'green':'grey'}`"/>
              </td>
            </template>
            <template #actions="{item}">
              <td style="white-space: nowrap; text-align: center">
                <CLink v-if="isAdmin && item.id!=currentUserId" @click="clickEditUser(item.id)" class="btn btn-sm btn-primary m-1">
                  <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                </CLink><CLink v-else class="btn btn-sm btn-secondary m-1" disabled>
                  <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                </CLink>

                <CLink v-if="isAdmin && item.id!=currentUserId" @click="clickDeleteUser(item.id)" class="btn btn-sm btn-danger m-1">
                  <CIcon name="cil-trash" v-c-tooltip.hover="$t('message.action_delete')"/>
                </CLink><CLink v-else class="btn btn-sm btn-secondary m-1" disabled>
                <CIcon name="cil-trash" v-c-tooltip.hover="$t('message.action_delete')"/>
              </CLink>
              </td>
            </template>
          </CDataTable>
        </CCardBody>
      </CCard>
    </CCol>

    <!-- pop-up form to add new user -->
    <CForm @submit.prevent="doAddUser">
        <CModal size="lg" :title="$t('message.add_user')" :centered="true" :show.sync="modalAddShow" :close-on-backdrop="false">
          <CAlert v-if="waitAddUser" color="info">{{ $t('message.wait') }}</CAlert>
          <CAlert v-if="modalAddErr" color="danger">{{ modalAddErr }}</CAlert>
          <div class="form-group form-row mt-2">
            <CCol :sm="{offset:3,size:9}" class="form-inline">
              <CInputCheckbox inline :label="$t('message.user_is_admin')" :checked.sync="formAdd.is_admin"/>
              <small>({{ $t('message.user_is_admin_msg') }})</small>
            </CCol>
          </div>
          <CInput
              type="text"
              v-model="formAdd.id"
              :label="$t('message.user_id')"
              :placeholder="$t('message.user_id_msg')"
              v-c-tooltip.hover="$t('message.user_id_msg')"
              horizontal
              required
              :invalid-feedback="$t('message.error_field_mandatory')"
              was-validated
          />
          <CInput
              type="text"
              v-model="formAdd.name"
              :label="$t('message.user_display_name')"
              :placeholder="$t('message.user_display_name_msg')"
              v-c-tooltip.hover="$t('message.user_display_name_msg')"
              horizontal
          />
          <hr/>
          <CInput
              type="password"
              v-model="formAdd.new_pwd"
              :label="$t('message.user_new_password')"
              :placeholder="$t('message.user_new_password_msg')"
              v-c-tooltip.hover="$t('message.user_new_password_msg')"
              horizontal
          />
          <CInput
              type="password"
              v-model="formAdd.confirmed_pwd"
              :label="$t('message.user_confirmed_password')"
              :placeholder="$t('message.user_confirmed_password_msg')"
              v-c-tooltip.hover="$t('message.user_confirmed_password_msg')"
              horizontal
              was-validated
              :is-valid="validateNewPasswordsAddUser"
              :invalid-feedback="$t('message.error_confirmed_password_mismatch')"
          />
          <template #footer>
            <CButton v-if="!waitAddUser" type="submit" color="primary" class="m-2" style="width: 96px">
              <CIcon name="cil-save" class="align-top"/>
              {{ $t('message.action_save') }}
            </CButton>
            <CButton type="button" color="secondary" style="width: 96px" @click="modalAddShow = false">
              <CIcon name="cil-arrow-circle-left" class="align-top"/>
              {{ $t('message.cancel') }}
            </CButton>
          </template>
        </CModal>
      </CForm>

    <!-- pop-up form to update existing user -->
    <CForm @submit.prevent="doEditUser">
      <CModal size="lg" :title="$t('message.edit_user_profile')" :centered="true" :show.sync="modalEditShow" :close-on-backdrop="false">
        <CAlert v-if="waitEditUser" color="info">{{ $t('message.wait') }}</CAlert>
        <CAlert v-if="modalEditErr" color="danger">{{ modalEditErr }}</CAlert>
        <div class="form-group form-row mt-2">
          <CCol :sm="{offset:3,size:9}" class="form-inline">
            <CInputCheckbox inline :label="$t('message.user_is_admin')" :checked.sync="formEdit.is_admin"/>
            <small>({{ $t('message.user_is_admin_msg') }})</small>
          </CCol>
        </div>
        <CInput
            type="text"
            v-model="formEdit.id"
            :label="$t('message.user_id')"
            :placeholder="$t('message.user_id_msg')"
            v-c-tooltip.hover="$t('message.user_id_msg')"
            horizontal disabled
        />
        <CInput
            type="text"
            v-model="formEdit.name"
            :label="$t('message.user_display_name')"
            :placeholder="$t('message.user_display_name_msg')"
            v-c-tooltip.hover="$t('message.user_display_name_msg')"
            horizontal
        />
        <template #footer>
          <CButton v-if="!waitEditUser" type="button" @click="doEditUser" color="primary" class="m-2" style="width: 96px">
            <CIcon name="cil-save" class="align-top"/>
            {{ $t('message.action_save') }}
          </CButton>
          <CButton type="button" color="secondary" style="width: 96px" @click="modalEditShow = false">
            <CIcon name="cil-arrow-circle-left" class="align-top"/>
            {{ $t('message.cancel') }}
          </CButton>
        </template>
      </CModal>
    </CForm>

    <!-- pop-up dialog to confirm deleting an existing user -->
    <CModal size="lg" color="danger" :title="$t('message.delete_user')" :centered="true" :show.sync="modalDeleteShow" :close-on-backdrop="false">
      <CAlert color="warning">
        <CIcon name="cil-warning" size="lg" />
        {{ $t('message.delete_user_msg', {id: userToDelete.id}) }}
      </CAlert>
      <CAlert v-if="waitDeleteUser" color="info">{{ $t('message.wait') }}</CAlert>
      <CAlert v-if="modalDeleteErr" color="danger">{{ modalDeleteErr }}</CAlert>
      <CInput type="text" :label="$t('message.user_id')" v-model="userToDelete.id" horizontal plaintext />
      <CInput type="text" :label="$t('message.user_display_name')" v-model="userToDelete.name" horizontal plaintext/>
      <template #footer>
        <CButton v-if="!waitDeleteUser" type="button" color="danger" class="m-2" style="width: 96px" @click="doDeleteUser">
          <CIcon name="cil-trash" class="align-top"/>
          {{ $t('message.action_delete') }}
        </CButton>
        <CButton type="button" color="secondary" style="width: 96px" @click="modalDeleteShow = false">
          <CIcon name="cil-arrow-circle-left" class="align-top"/>
          {{ $t('message.cancel') }}
        </CButton>
      </template>
    </CModal>
  </CRow>
</template>

<script>
import clientUtils from "@/utils/api_client"
import utils from "@/utils/app_utils"
import forge from 'node-forge'

const emptyForm = {id: "", name: "", is_admin: false, new_pwd: "", confirmed_pwd: ""}

export default {
  name: 'UserList',
  mounted() {
    let vue = this
    vue.waitLoadUserList = true
    clientUtils.apiDoGet(clientUtils.apiInfo,
        (apiRes) => {
          if (apiRes.status != 200) {
            vue.errorMsg = apiRes.message
            vue.waitLoadUserList = false
          } else {
            vue.publicKey = forge.pki.publicKeyFromPem(apiRes.data.rsa_public_key)
            vue.loadUserList()
          }
        },
        (err) => {
          vue.errorMsg = err
          vue.waitLoadUserList = false
        })
  },
  data() {
    return {
      modalAddShow: false,
      modalAddErr: '',
      formAdd: {...emptyForm},
      waitAddUser: false,

      modalEditShow: false,
      modalEditErr: '',
      formEdit: {...emptyForm},
      waitEditUser: false,

      modalDeleteShow: false,
      modalDeleteErr: '',
      waitDeleteUser: false,
      userToDelete: {...emptyForm},

      errorMsg: '',
      myFlashMsg: this.flashMsg,
      waitLoadUserList: false,
      userList: [],
      userMap: {},
      isAdmin: false,
      publicKey: Object,
      currentUserId: '',
    }
  },
  props: ["flashMsg"],
  methods: {
    validateNewPasswordsAddUser(val) {
      return val == this.formAdd.new_pwd
    },
    loadUserList() {
      const vue = this
      vue.waitLoadUserList = true
      clientUtils.apiDoGet(clientUtils.apiAdminUsers,
          (apiRes) => {
            if (apiRes.status == 200) {
              let session = utils.loadLoginSession()
              vue.currentUserId = session.uid
              vue.userList = apiRes.data
              vue.userMap = {}
              for (let i = vue.userList.length - 1; i >= 0; i--) {
                vue.userMap[vue.userList[i].id] = vue.userList[i]
                if (vue.userList[i].id == session.uid && vue.userList[i].is_admin) {
                  vue.isAdmin = true
                }
              }
            } else {
              vue.errorMsg = apiRes.status + ": " + apiRes.message
            }
            vue.waitLoadUserList = false
          },
          (err) => {
            vue.errorMsg = err
            vue.waitLoadUserList = false
          })
    },
    clickAddUser() {
      this.formAdd = {...emptyForm}
      this.modalAddErr = ''
      this.modalAddShow = true
      this.myFlashMsg = ''
    },
    doAddUser(e) {
      e.preventDefault()
      this.modalAddErr = ''
      let vue = this
      let data = {...vue.formAdd}
      data.new_pwd = forge.util.encode64(vue.publicKey.encrypt(data.new_pwd))
      data.confirmed_pwd = forge.util.encode64(vue.publicKey.encrypt(data.confirmed_pwd))
      vue.waitAddUser = true
      clientUtils.apiDoPost(
          clientUtils.apiAdminUsers, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalAddShow = false
              vue.myFlashMsg =  vue.$i18n.t('message.user_added_msg', {id: data.id})
              vue.loadUserList()
            } else {
              vue.modalAddErr = apiRes.status + ": " + apiRes.message
            }
            vue.waitAddUser = false
          },
          (err) => {
            vue.modalAddErr = err
            vue.waitAddUser = false
          }
      )
    },
    clickEditUser(id) {
      this.formEdit = {...this.userMap[id]}
      this.modalEditFlash = ''
      this.modalEditErr = ''
      this.modalEditShow = true
      this.domainToMapEdit = ''
      this.waitEditUser = false
      this.myFlashMsg = ''
    },
    doEditUser(e) {
      e.preventDefault()
      this.modalEditErr = ''
      let vue = this
      let data = {...vue.formEdit}
      vue.waitEditUser = true
      clientUtils.apiDoPut(clientUtils.apiAdminUser + "/" + vue.formEdit.id, data,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.modalEditErr = apiRes.status + ": " + apiRes.message
            } else {
              vue.modalEditShow = false
              vue.loadUserList()
              vue.myFlashMsg = vue.$i18n.t('message.user_profile_updated_msg', {id: vue.formEdit.id})
            }
            vue.waitEditUser = false
          },
          (err) => {
            vue.modalEditErr = err
            vue.waitEditUser = false
          }
      )
    },
    clickDeleteUser(id) {
      this.userToDelete = this.userMap[id]
      this.modalDeleteErr = ''
      this.modalDeleteShow = true
      this.myFlashMsg = ''
    },
    doDeleteUser() {
      this.modalDeleteErr = ''
      let vue = this
      vue.waitDeleteUser = true
      clientUtils.apiDoDelete(
          clientUtils.apiAdminUser + "/" + vue.userToDelete.id,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.modalDeleteErr = apiRes.status + ": " + apiRes.message
            } else {
              vue.modalDeleteShow = false
              vue.loadUserList()
              vue.myFlashMsg = vue.$i18n.t('message.user_deleted_msg', {id: vue.userToDelete.id})
            }
            vue.waitDeleteUser = false
          },
          (err) => {
            vue.modalDeleteErr = err
            vue.waitDeleteUser = false
          }
      )
    },
  }
}
</script>
