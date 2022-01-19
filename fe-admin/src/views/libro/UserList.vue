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
                <CLink v-if="isAdmin" @click="clickEditProduct(item.id)" class="btn btn-sm btn-primary m-1">
                  <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                </CLink><CLink v-else class="btn btn-sm btn-secondary m-1" disabled>
                  <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                </CLink>

                <CLink v-if="isAdmin" @click="clickDeleteProduct(item.id)" class="btn btn-sm btn-danger m-1">
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

    <!-- pop-up form to add new product -->
    <CForm @submit.prevent="doAddProduct">
        <CModal size="lg" :title="$t('message.add_product')" :centered="true" :show.sync="modalAddShow" :close-on-backdrop="false">
          <CAlert v-if="waitAddProduct" color="info">{{ $t('message.wait') }}</CAlert>
          <CAlert v-if="modalAddErr" color="danger">{{ modalAddErr }}</CAlert>
          <CTabs ref="formAddProductTabs" v-if="!waitAddProduct">
            <CTab ref="formAddProductTabInfo" active>
              <template slot="title">
                {{ $t('message.product_info') }}
              </template>
              <div class="form-group form-row mt-2">
                <CCol :sm="{offset:3,size:9}" class="form-inline">
                  <CInputCheckbox inline :label="$t('message.product_is_published')" :checked.sync="formAdd.is_published"/>
                  <small>({{ $t('message.product_is_published_msg') }})</small>
                </CCol>
              </div>
              <CInput
                  type="text"
                  v-model="formAdd.id"
                  :label="$t('message.product_id')"
                  :placeholder="$t('message.product_id_msg')"
                  v-c-tooltip.hover="$t('message.product_id_msg')"
                  horizontal
              />
              <CInput
                  type="text"
                  v-model="formAdd.name"
                  :label="$t('message.product_name')"
                  :placeholder="$t('message.product_name_msg')"
                  v-c-tooltip.hover="$t('message.product_name_msg')"
                  horizontal
                  required
                  was-validated
              />
              <CTextarea
                  rows="2"
                  type="text"
                  v-model="formAdd.desc"
                  :label="$t('message.product_desc')"
                  :placeholder="$t('message.product_desc_msg')"
                  v-c-tooltip.hover="$t('message.product_desc_msg')"
                  horizontal
                  required
                  was-validated
              />
              <CTextarea
                  rows="4"
                  type="text"
                  v-model="formAdd.domains"
                  :label="$t('message.product_domains')"
                  :placeholder="$t('message.product_domains_msg')"
                  v-c-tooltip.hover="$t('message.product_domains_msg')"
                  horizontal
                  required
                  was-validated
              />
            </CTab>
            <CTab>
              <template slot="title">
                {{ $t('message.product_contacts') }}
              </template>
              <CInput class="mt-2" type="text" v-model="formAdd.contacts.email" :placeholder="$t('message.product_email')" v-c-tooltip.hover="$t('message.product_email_msg')">
                <template #prepend-content><CIcon name="cil-envelope-closed"/></template>
              </CInput>
              <CInput type="text" v-model="formAdd.contacts.website" :placeholder="$t('message.product_website')" v-c-tooltip.hover="$t('message.product_website_msg')">
                <template #prepend-content><CIcon name="cil-globe-alt"/></template>
              </CInput>
              <CInput type="text" v-model="formAdd.contacts.github" :placeholder="$t('message.product_github')" v-c-tooltip.hover="$t('message.product_github_msg')">
                <template #prepend-content><CIcon name="cib-github"/></template>
              </CInput>
              <CInput type="text" v-model="formAdd.contacts.facebook" :placeholder="$t('message.product_facebook')" v-c-tooltip.hover="$t('message.product_facebook_msg')">
                <template #prepend-content><CIcon name="cib-facebook"/></template>
              </CInput>
              <CInput type="text" v-model="formAdd.contacts.linkedin" :placeholder="$t('message.product_linkedin')" v-c-tooltip.hover="$t('message.product_linkedin_msg')">
                <template #prepend-content><CIcon name="cib-linkedin"/></template>
              </CInput>
              <CInput type="text" v-model="formAdd.contacts.slack" :placeholder="$t('message.product_slack')" v-c-tooltip.hover="$t('message.product_slack_msg')">
                <template #prepend-content><CIcon name="cib-slack"/></template>
              </CInput>
              <CInput type="text" v-model="formAdd.contacts.twitter" :placeholder="$t('message.product_twitter')" v-c-tooltip.hover="$t('message.product_twitter_msg')">
                <template #prepend-content><CIcon name="cib-twitter"/></template>
              </CInput>
            </CTab>
          </CTabs>
          <template #footer>
            <button type="submit" ref="btnSubmitAddProduct" style="display:none;" />
            <CButton v-if="!waitAddProduct" type="button" @click="doAddProductClick" color="primary" class="m-2" style="width: 96px">
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

    <!-- pop-up form to update existing product -->
    <CForm @submit.prevent="doEditProduct">
      <CModal size="lg" :title="$t('message.edit_product')" :centered="true" :show.sync="modalEditShow" :close-on-backdrop="false">
        <CAlert v-if="waitEditProduct" color="info">{{ $t('message.wait') }}</CAlert>
        <CAlert v-if="modalEditErr" color="danger">{{ modalEditErr }}</CAlert>
        <CAlert v-if="modalEditFlash" color="success" closeButton>{{ modalEditFlash }}</CAlert>
        <CTabs ref="formEditProductTabs" v-if="!waitEditProduct">
          <CTab ref="formEditProductTabInfo" active>
            <template slot="title">
              {{ $t('message.product_info') }}
            </template>
            <div class="form-group form-row mt-2">
              <CCol :sm="{offset:3,size:9}" class="form-inline">
                <CInputCheckbox inline :label="$t('message.product_is_published')" :checked.sync="formEdit.is_published"/>
                <small>({{ $t('message.product_is_published_msg') }})</small>
              </CCol>
            </div>
            <CInput
                type="text"
                v-model="formEdit.id"
                :label="$t('message.product_id')"
                :placeholder="$t('message.product_id_msg')"
                v-c-tooltip.hover="$t('message.product_id_msg')"
                horizontal
                disabled
            />
            <CInput
                type="text"
                v-model="formEdit.name"
                :label="$t('message.product_name')"
                :placeholder="$t('message.product_name_msg')"
                v-c-tooltip.hover="$t('message.product_name_msg')"
                horizontal
                required
                was-validated
            />
            <CTextarea
                rows="2"
                type="text"
                v-model="formEdit.desc"
                :label="$t('message.product_desc')"
                :placeholder="$t('message.product_desc_msg')"
                v-c-tooltip.hover="$t('message.product_desc_msg')"
                horizontal
                required
                was-validated
            />
            <div class="form-group form-row">
              <CCol sm="3">
                {{ $t('message.product_domains') }}
              </CCol>
              <CCol sm="9">
                <CInput type="text" v-c-tooltip.hover="$t('message.product_map_domain_msg')" v-model="domainToMapEdit">
                  <template #append>
                    <CButton type="button" color="primary" @click="doMapDomain"><ficon :title="$t('message.product_map_domain')" :icon="['fas','link']"/></CButton>
                  </template>
                </CInput>
                <CDropdown v-for="(domain, _) in formEdit.domains" size="sm"
                           :toggler-text="domain" color="info" class="mb-1 mr-1 d-inline"
                           :label="$t('message.product_domains')">
                  <CDropdownItem @click="clickUnmapDomain(domain)"><ficon fixedWidth :icon="['fas', 'unlink']"/> {{ $t('message.product_unmap_domain') }}</CDropdownItem>
                </CDropdown>
              </CCol>
            </div>
          </CTab>
          <CTab>
            <template slot="title">
              {{ $t('message.product_contacts') }}
            </template>
            <CInput class="mt-2" type="text" v-model="formEdit.contacts.email" :placeholder="$t('message.product_email')" v-c-tooltip.hover="$t('message.product_email_msg')">
              <template #prepend-content><CIcon name="cil-envelope-closed"/></template>
            </CInput>
            <CInput type="text" v-model="formEdit.contacts.website" :placeholder="$t('message.product_website')" v-c-tooltip.hover="$t('message.product_website_msg')">
              <template #prepend-content><CIcon name="cil-globe-alt"/></template>
            </CInput>
            <CInput type="text" v-model="formEdit.contacts.github" :placeholder="$t('message.product_github')" v-c-tooltip.hover="$t('message.product_github_msg')">
              <template #prepend-content><CIcon name="cib-github"/></template>
            </CInput>
            <CInput type="text" v-model="formEdit.contacts.facebook" :placeholder="$t('message.product_facebook')" v-c-tooltip.hover="$t('message.product_facebook_msg')">
              <template #prepend-content><CIcon name="cib-facebook"/></template>
            </CInput>
            <CInput type="text" v-model="formEdit.contacts.linkedin" :placeholder="$t('message.product_linkedin')" v-c-tooltip.hover="$t('message.product_linkedin_msg')">
              <template #prepend-content><CIcon name="cib-linkedin"/></template>
            </CInput>
            <CInput type="text" v-model="formEdit.contacts.slack" :placeholder="$t('message.product_slack')" v-c-tooltip.hover="$t('message.product_slack_msg')">
              <template #prepend-content><CIcon name="cib-slack"/></template>
            </CInput>
            <CInput type="text" v-model="formEdit.contacts.twitter" :placeholder="$t('message.product_twitter')" v-c-tooltip.hover="$t('message.product_twitter_msg')">
              <template #prepend-content><CIcon name="cib-twitter"/></template>
            </CInput>
          </CTab>
        </CTabs>
        <template #footer>
          <button type="submit" ref="btnSubmitEditProduct" style="display:none;" />
          <CButton v-if="!waitEditProduct" type="button" @click="doEditProductClick" color="primary" class="m-2" style="width: 96px">
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

    <!-- pop-up form to confirm unmapping domain -->
    <CModal :title="$t('message.product_unmap_domain')" color="danger" :centered="true" :show.sync="modalUnmapShow">
      {{ modalUnmapMessage }}
      <template #footer>
        <CButton @click="doUnmapDomain(modalUnmapData)" color="danger" style="width: 96px">
          <CIcon name="cil-trash" class="align-top"/>
          {{ $t('message.ok') }}
        </CButton>
        <CButton @click="modalUnmapShow = false" color="secondary" class="ml-2" style="width: 96px">
          <CIcon name="cil-arrow-circle-left" class="align-top"/>
          {{ $t('message.cancel') }}
        </CButton>
      </template>
    </CModal>

    <!-- pop-up dialog to confirm deleting an existing product -->
    <CModal color="danger" :title="$t('message.delete_product')" :centered="true" :show.sync="modalDeleteShow" :close-on-backdrop="false">
      <CAlert color="warning">
        <CIcon name="cil-warning" size="lg" />
        {{ $t('message.delete_product_msg', {numTopics: prodToDelete['num_topics']}) }}
      </CAlert>
      <CAlert v-if="waitDeleteProduct" color="info">{{ $t('message.wait') }}</CAlert>
      <CAlert v-if="modalDeleteErr" color="danger">{{ modalDeleteErr }}</CAlert>
      <CInput type="text" :label="$t('message.product_id')" v-model="prodToDelete.id" horizontal plaintext />
      <CInput type="text" :label="$t('message.product_name')" v-model="prodToDelete.name" horizontal plaintext/>
      <CTextarea rows="2" type="text" :label="$t('message.product_desc')" v-model="prodToDelete.desc" horizontal plaintext/>
      <template #footer>
        <CButton v-if="!waitDeleteProduct" type="button" color="danger" class="m-2" style="width: 96px" @click="doDeleteProduct">
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

const emptyForm = {id: "", name: "", desc: "", is_published: false, domains: "", contacts: {email:"", website:"", github:"", facebook: "", linkedin: "", slack: "", twitter: ""}}

export default {
  name: 'UserList',
  mounted() {
    this.loadUserList()
  },
  data() {
    return {
      modalAddShow: false,
      modalAddErr: '',
      formAdd: {...emptyForm},
      waitAddProduct: false,

      modalEditShow: false,
      modalEditErr: '',
      modalEditFlash: '',
      formEdit: {...emptyForm},
      waitEditProduct: false,
      domainToMapEdit: '',

      modalUnmapShow: false,
      modalUnmapMessage: '',
      modalUnmapData: '',

      modalDeleteShow: false,
      modalDeleteErr: '',
      waitDeleteProduct: false,
      prodToDelete: {...emptyForm},

      errorMsg: '',
      myFlashMsg: this.flashMsg,
      waitLoadUserList: false,
      userList: [],
      userMap: {},
      isAdmin: false,
    }
  },
  props: ["flashMsg"],
  methods: {
    loadUserList() {
      const vue = this
      vue.waitLoadUserList = true
      clientUtils.apiDoGet(clientUtils.apiAdminUsers,
          (apiRes) => {
            if (apiRes.status == 200) {
              let session = utils.loadLoginSession()
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
    doAddProductClick() {
      // this workaround is to force switching to Product Info tab for input validation
      // before actually sending Add Product request to the backend
      const tabIndex = this.$refs.formAddProductTabs.activeTabIndex
      if (tabIndex != this.$refs.formAddProductTabInfo.index) {
        this.$refs.formAddProductTabs.changeTabTo(this.$refs.formAddProductTabInfo.index)
        setTimeout(()=>{this.$refs.btnSubmitAddProduct.click()},125)
      } else {
        this.$refs.btnSubmitAddProduct.click()
      }
    },
    doAddProduct(e) {
      e.preventDefault()
      this.modalAddErr = ''
      let vue = this
      let data = {...vue.formAdd}
      vue.waitAddProduct = true
      clientUtils.apiDoPost(
          clientUtils.apiAdminProducts, data,
          (apiRes) => {
            if (apiRes.status == 200 || apiRes.status == 201) {
              vue.modalAddShow = false
              vue.myFlashMsg =  apiRes.status==200?vue.$i18n.t('message.product_added_msg', {name: data.name}):apiRes.message
              vue.loadUserList()
            } else {
              vue.modalAddErr = apiRes.status + ": " + apiRes.message
            }
            vue.waitAddProduct = false
          },
          (err) => {
            vue.modalAddErr = err
            vue.waitAddProduct = false
          }
      )
    },
    clickEditProduct(id) {
      this.formEdit = {...this.userMap[id]}
      this.modalEditFlash = ''
      this.modalEditErr = ''
      this.modalEditShow = true
      this.domainToMapEdit = ''
      this.waitEditProduct = false
      this.myFlashMsg = ''
    },
    doEditProductClick() {
      // this workaround is to force switching to Product Info tab for input validation
      // before actually sending Edit Product request to the backend
      const tabIndex = this.$refs.formEditProductTabs.activeTabIndex
      if (tabIndex != this.$refs.formEditProductTabInfo.index) {
        this.$refs.formEditProductTabs.changeTabTo(this.$refs.formEditProductTabInfo.index)
        setTimeout(()=>{this.$refs.btnSubmitEditProduct.click()},125)
      } else {
        this.$refs.btnSubmitEditProduct.click()
      }
    },
    doEditProduct(e) {
      e.preventDefault()
      this.modalEditFlash = ''
      this.modalEditErr = ''
      if (this.domainToMapEdit.trim() != '') {
        this.doMapDomain()
        return
      }
      let vue = this
      let data = {...vue.formEdit}
      vue.waitEditProduct = true
      clientUtils.apiDoPut(clientUtils.apiAdminProduct + "/" + vue.formEdit.id, data,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.modalEditErr = apiRes.status + ": " + apiRes.message
            } else {
              vue.modalEditShow = false
              vue.loadUserList()
              vue.myFlashMsg = vue.$i18n.t('message.product_updated_msg', {name: vue.formEdit.name})
            }
            vue.waitEditProduct = false
          },
          (err) => {
            vue.modalEditErr = err
            vue.waitEditProduct = false
          }
      )
    },
    doMapDomain() {
      this.modalEditFlash = ''
      this.modalEditErr = ''
      let vue = this
      const domainName = vue.domainToMapEdit.trim()
      if (domainName != '') {
        vue.waitEditProduct = true
        let data = {pid: vue.formEdit.id, domain: domainName}
        clientUtils.apiDoPost(
            clientUtils.apiAdminDomains, data,
            (apiRes) => {
              if (apiRes.status != 200) {
                vue.modalEditErr = apiRes.status + ": " + apiRes.message
              } else {
                vue.modalEditFlash = vue.$t('message.product_domain_mapped_msg', {domain: domainName})
                vue.domainToMapEdit = ''
                vue.formEdit.domains = apiRes.data
              }
              vue.waitEditProduct = false
              vue.loadUserList()
            },
            (err) => {
              vue.waitEditProduct = false
              vue.modalEditErr = err
            }
        )
      }
    },
    clickUnmapDomain(domain) {
      this.modalUnmapShow = true
      this.modalUnmapData = domain.trim()
      this.modalUnmapMessage = this.$i18n.t('message.product_unmap_domain_msg', {domain: domain})
    },
    doUnmapDomain(domain) {
      this.modalEditFlash = ''
      this.modalEditErr = ''
      let vue = this
      vue.modalUnmapShow = false
      vue.waitEditProduct = true
      clientUtils.apiDoDelete(
          clientUtils.apiAdminDomain+"/"+domain+"/"+vue.formEdit.id,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.modalEditErr = apiRes.status + ": " + apiRes.message
            } else {
              vue.modalEditFlash = vue.$t('message.product_domain_unmapped_msg', {domain: domain})
              vue.formEdit.domains = apiRes.data
            }
            vue.waitEditProduct = false
            vue.loadUserList()
          },
          (err) => {
            vue.waitEditProduct = false
            vue.modalEditErr = err
          }
      )
    },
    clickDeleteProduct(id) {
      this.prodToDelete = this.userMap[id]
      this.modalDeleteErr = ''
      this.modalDeleteShow = true
      this.myFlashMsg = ''
    },
    doDeleteProduct() {
      this.modalDeleteErr = ''
      let vue = this
      vue.waitDeleteProduct = true
      clientUtils.apiDoDelete(
          clientUtils.apiAdminProduct + "/" + vue.prodToDelete.id,
          (apiRes) => {
            if (apiRes.status != 200) {
              vue.modalDeleteErr = apiRes.status + ": " + apiRes.message
            } else {
              vue.modalDeleteShow = false
              vue.loadUserList()
              vue.myFlashMsg = vue.$i18n.t('message.product_deleted_msg', {name: vue.prodToDelete.name})
            }
            vue.waitDeleteProduct = false
          },
          (err) => {
            vue.modalDeleteErr = err
            vue.waitDeleteProduct = false
          }
      )
    },
  }
}
</script>
