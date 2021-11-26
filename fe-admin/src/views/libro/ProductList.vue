<template>
  <CRow>
    <CCol sm="12">
      <CCard accent-color="info">
        <CCardHeader>
          <strong>{{ $t('message.products') }}</strong>
          <div v-if="errorMsg==''" class="card-header-actions">
            <CButton class="btn-sm btn-primary" @click="clickAddProduct">
              <CIcon name="cil-library-add" class="align-top"/>
              {{ $t('message.add_product') }}
            </CButton>
          </div>
        </CCardHeader>
        <CCardBody>
          <CAlert v-if="myFlashMsg" color="success" closeButton>{{ myFlashMsg }}</CAlert>
          <CAlert v-if="waitLoadProductList" color="success">{{ $t('message.wait') }}</CAlert>
          <CAlert v-if="errorMsg" color="danger">{{ errorMsg }}</CAlert>
          <CDataTable v-if="errorMsg==''" :items="prodList" :fields="[
              {key:'is_published',label:''},
              {key:'name',label:$t('message.product_name')},
              {key:'domains',label:$t('message.product_domains')},
              {key:'actions',label:$t('message.actions'),_style:'text-align: center'}
            ]">
            <template #is_published="{item}">
              <td class="col-1">
                <CIcon :name="`${item.is_published?'cil-check':'cil-check-alt'}`"
                       :style="`color: ${item.is_published?'green':'grey'}`"/>
              </td>
            </template>
            <template #name="{item}">
              <td class="col-5">
                {{ item.name }}
                <br />
                <span style="font-size: smaller">{{ item.desc }}</span>
              </td>
            </template>
            <template #domains="{item}">
              <td class="col-4">
                {{ item.domains }}
              </td>
            </template>
            <template #actions="{item}">
              <td style="white-space: nowrap; text-align: center">
                <CLink @click="clickProductTopic(item.id)" class="btn btn-sm btn-info m-1">
                  <CIcon name="cil-list-rich" v-c-tooltip.hover="$t('message.topics')"/>
                </CLink>
                <CLink @click="clickEditProduct(item.id)" class="btn btn-sm btn-primary m-1">
                  <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                </CLink>

                <CLink @click="clickDeleteProduct(item.id)" class="btn btn-sm btn-danger m-1">
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
          <CAlert v-if="waitAddProduct" color="success">{{ $t('message.wait') }}</CAlert>
          <CAlert v-if="modalAddErr" color="danger">{{ modalAddErr }}</CAlert>
          <CTabs ref="formAddProductTabs">
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
            <CButton type="button" @click="doAddProductClick" color="primary" class="m-2" style="width: 96px">
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
  </CRow>
</template>

<script>
import clientUtils from "@/utils/api_client"

const emptyForm = {id: "", name: "", desc: "", is_published: false, domains: "", contacts: {email:"", website:"", github:"", facebook: "", linkedin: "", slack: "", twitter: ""}}

export default {
  name: 'Products',
  mounted() {
    this.loadProductList()
  },
  data() {
    return {
      modalAddShow: false,
      modalAddErr: "",
      formAdd: {...emptyForm},
      waitAddProduct: false,

      errorMsg: "",
      myFlashMsg: this.flashMsg,
      waitLoadProductList: false,
      prodList: [],
    }
  },
  props: ["flashMsg"],
  methods: {
    loadProductList() {
      const vue = this
      vue.waitLoadProductList = true
      clientUtils.apiDoGet(clientUtils.apiAdminProducts,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.prodList = apiRes.data
            } else {
              vue.errorMsg = apiRes.status + ": " + apiRes.message
            }
            vue.waitLoadProductList = false
          },
          (err) => {
            vue.errorMsg = err
            vue.waitLoadProductList = false
          })
    },
    iconize(icon) {
      return icon.startsWith("cil-")?icon.slice(4):(icon.startsWith("fa-")?icon.slice(3):icon)
    },
    clickAddProduct() {
      this.addMode = true
      this.formAdd = {...emptyForm}
      this.modalAddErr = ''
      this.modalAddShow = true
    },
    doAddProductClick() {
      // this workaround is to force switching to Product Info tab for input validation
      // before actually sending Add Product request to the backend
      const tabIndex = this.$refs.formAddProductTabs.activeTabIndex
      if (tabIndex != this.$refs.formAddProductTabInfo.index) {
        this.$refs.formAddProductTabs.changeTabTo(this.$refs.formAddProductTabInfo.index)
        setTimeout(()=>{this.$refs.btnSubmitAddProduct.click()},125)
        // this.$refs.formAddProductTabInfo.$nextTick(()=>{this.$refs.btnSubmitAddProduct.click()})
      } else {
        this.$refs.btnSubmitAddProduct.click()
      }
    },
    doAddProduct(e) {
      e.preventDefault()
      let vue = this
      let data = {...vue.formAdd}
      vue.waitAddProduct = true
      clientUtils.apiDoPost(
          clientUtils.apiAdminProducts, data,
          (apiRes) => {
            if (apiRes.status == 200 || apiRes.status == 201) {
              vue.modalAddShow = false
              vue.myFlashMsg =  apiRes.status==200?vue.$i18n.t('message.product_added_msg', {name: data.name}):apiRes.message
              vue.loadProductList()
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
    clickProductTopic(id) {
      this.$router.push({name: "ProductTopicList", params: {pid: id}})
    },
    clickEditProduct(id) {
      this.$router.push({name: "EditProduct", params: {id: id}})
    },
    clickDeleteProduct(id) {
      this.$router.push({name: "DeleteProduct", params: {id: id}})
    },
  }
}
</script>
