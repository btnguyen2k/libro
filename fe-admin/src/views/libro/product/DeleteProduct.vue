<template>
  <div>
    <CRow>
      <CCol sm="12">
        <CCard>
          <CCardHeader><h5>{{ $t('message.delete_product') }}</h5></CCardHeader>
          <CForm @submit.prevent="doSubmit" method="post">
            <CCardBody>
              <p v-if="foundStatus<0" class="alert alert-info">{{ $t('message.wait') }}</p>
              <p v-if="foundStatus==0" class="alert alert-danger">{{ $t('message.error_product_not_found', {id: this.$route.params.id}) }}</p>
              <p v-if="errorMsg!=''" class="alert alert-danger">{{ errorMsg }}</p>
              <div class="form-group form-row" v-if="foundStatus>0">
                <CCol :sm="{offset:3,size:9}" class="form-inline">
                  <CInputCheckbox inline :label="$t('message.product_is_published')" :checked.sync="product.is_published" disabled="disabled"/>
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
                  readonly="readonly"
              />
              <CTextarea v-if="foundStatus>0"
                  rows="2"
                  type="text"
                  :value="product.desc"
                  :label="$t('message.product_desc')"
                  :placeholder="$t('message.product_desc_msg')"
                  v-c-tooltip.hover="$t('message.product_desc_msg')"
                  horizontal
                  readonly="readonly"
              />
              <CTextarea v-if="foundStatus>0"
                  rows="2"
                  type="text"
                  :value="JSON.stringify(product.domains)"
                  :label="$t('message.product_domains')"
                  :placeholder="$t('message.product_domains_msg')"
                  v-c-tooltip.hover="$t('message.product_domains_msg')"
                  horizontal
                  readonly="readonly"
              />
            </CCardBody>
            <CCardFooter>
              <CButton v-if="foundStatus>0" type="submit" color="danger" style="width: 96px">
                <CIcon name="cil-trash"/>
                {{ $t('message.action_delete') }}
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
  </div>
</template>

<script>
import router from "@/router"
import clientUtils from "@/utils/api_client"

export default {
  name: 'DeleteProduct',
  mounted() {
    const vue = this
    clientUtils.apiDoGet(clientUtils.apiAdminProduct + "/" + vue.$route.params.id,
        (apiRes) => {
          vue.foundStatus = apiRes.status==200 ? 1 : 0
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
      product: {},
      errorMsg: "",
      foundStatus: -1,
    }
  },
  methods: {
    doCancel() {
      router.push({name: "ProductList"})
    },
    doSubmit(e) {
      e.preventDefault()
      clientUtils.apiDoDelete(
          clientUtils.apiAdminProduct + "/" + this.$route.params.id,
          (apiRes) => {
            if (apiRes.status != 200) {
              this.errorMsg = apiRes.status + ": " + apiRes.message
            } else {
              this.$router.push({
                name: "ProductList",
                params: {flashMsg: this.$i18n.t('message.product_deleted_msg', {name: this.product.name})},
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
