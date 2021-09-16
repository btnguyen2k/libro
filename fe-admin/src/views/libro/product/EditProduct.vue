<template>
  <div>
    <CRow>
      <CCol sm="12">
        <CCard>
          <CCardHeader><h5>{{ $t('message.edit_product') }}</h5></CCardHeader>
          <CForm @submit.prevent="doSubmit" method="post">
            <CCardBody>
              <p v-if="foundStatus<0" class="alert alert-info">{{ $t('message.wait') }}</p>
              <p v-if="foundStatus==0" class="alert alert-danger">
                {{ $t('message.error_product_not_found', {id: this.$route.params.id}) }}</p>
              <p v-if="errorMsg!=''" class="alert alert-danger">{{ errorMsg }}</p>
              <div class="form-group form-row" v-if="foundStatus>0">
                <CCol :sm="{offset:3,size:9}" class="form-inline">
                  <CInputCheckbox inline :label="$t('message.product_is_published')"
                                  :checked.sync="product.is_published"/>
                  <small>({{ $t('message.product_is_published_msg') }})</small>
                </CCol>
              </div>
              <CInput v-if="foundStatus>0"
                      type="text"
                      :value="product.name"
                      :label="$t('message.product_name')"
                      :placeholder="$t('message.product_name_msg')"
                      v-c-tooltip.hover="$t('message.product_name_msg')"
                      horizontal
                      required
                      was-validated
              />
              <CTextarea v-if="foundStatus>0"
                         rows="2"
                         type="text"
                         :value="product.desc"
                         :label="$t('message.product_desc')"
                         :placeholder="$t('message.product_desc_msg')"
                         v-c-tooltip.hover="$t('message.product_desc_msg')"
                         horizontal
                         required
                         was-validated
              />
              <div class="form-group form-row" v-if="foundStatus>0">
                <CCol sm="3">
                  {{ $t('message.product_domains') }}
                </CCol>
                <CCol sm="9">
                  <CInput type="text" v-c-tooltip.hover="$t('message.product_map_domain_msg')">
                    <template #append>
                      <CButton color="primary">{{ $t('message.product_map_domain') }}</CButton>
                    </template>
                  </CInput>
                  <CDropdown v-for="domain in product.domains"
                             size="sm"
                             :toggler-text="domain"
                             color="info"
                             class="m-0 d-block"
                             :label="$t('message.product_domains')"
                  >
                    <CDropdownItem @click="clickUnmapDomain(domain)">{{
                        $t('message.product_unmap_domain')
                      }}
                    </CDropdownItem>
                  </CDropdown>
                </CCol>
              </div>
            </CCardBody>
            <CCardFooter>
              <CButton v-if="foundStatus>0" type="submit" color="primary" style="width: 96px">
                <CIcon name="cil-save"/>
                {{ $t('message.action_save') }}
              </CButton>
              <CButton type="button" color="info" class="ml-2" style="width: 96px" @click="doCancel">
                <CIcon name="cil-arrow-circle-left"/>
                {{ $t('message.action_back') }}
              </CButton>
            </CCardFooter>
          </CForm>
        </CCard>
      </CCol>
    </CRow>

    <CModal :title="$t('message.product_unmap_domain')" color="danger" :centered="true" :show.sync="modalUnmap">
      {{ modalUnmapMessage }}
      <template #footer>
        <CButton @click="modalUnmap = false" color="danger" style="width: 96px">
          <CIcon name="cil-trash"/>
          {{ $t('message.ok') }}
        </CButton>
        <CButton @click="modalUnmap = false" color="secondary" class="ml-2" style="width: 96px">
          <CIcon name="cil-arrow-circle-left"/>
          {{ $t('message.cancel') }}
        </CButton>
      </template>
    </CModal>
  </div>
</template>

<script>
import router from "@/router"
import clientUtils from "@/utils/api_client"

export default {
  name: 'EditProduct',
  mounted() {
    const vue = this
    clientUtils.apiDoGet(clientUtils.apiAdminProduct + "/" + vue.$route.params.id,
        (apiRes) => {
          vue.foundStatus = apiRes.status == 200 ? 1 : 0
          if (vue.foundStatus == 1) {
            vue.product = apiRes.data
          }
        },
        (err) => {
          vue.errorMsg = err
        })
  },
  data() {
    return {
      modalUnmap: false,
      modalUnmapMessage: "",
      modalUnmapData: "",
      product: {},
      errorMsg: "",
      foundStatus: -1,
    }
  },
  methods: {
    doCancel() {
      router.push({name: "ProductList"})
    },
    clickUnmapDomain(domain) {
      this.modalUnmapData = domain
      this.modalUnmap = true
      this.modalUnmapMessage = this.$i18n.t('message.product_unmap_domain_msg', {domain: domain})
    },
    clickMapDomain(domain) {

    },
    doSubmit(e) {
      e.preventDefault()
      let data = {is_public: this.post.is_public, title: this.post.title, content: this.post.content}
      clientUtils.apiDoPut(
          clientUtils.apiPost + "/" + this.$route.params.id, data,
          (apiRes) => {
            if (apiRes.status != 200) {
              this.errorMsg = apiRes.status + ": " + apiRes.message
            } else {
              this.$router.push({
                name: "ProductList",
                params: {flashMsg: this.$i18n.t('message.product_updated_msg', {title: this.product.title})},
              })
            }
          },
          (err) => {
            this.errorMsg = err
          }
      )
    },
  }
}
</script>
