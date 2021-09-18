<template>
  <div>
    <CRow>
      <CCol sm="12">
        <CCard>
          <CCardHeader><h5>{{ $t('message.add_product') }}</h5></CCardHeader>
          <CForm @submit.prevent="doSubmit" method="post">
            <CCardBody>
              <p v-if="errorMsg!=''" class="alert alert-danger">{{ errorMsg }}</p>
              <div class="form-group form-row">
                <CCol :sm="{offset:3,size:9}" class="form-inline">
                  <CInputCheckbox inline :label="$t('message.product_is_published')" :checked.sync="form.isPublished"/>
                  <small>({{ $t('message.product_is_published_msg') }})</small>
                </CCol>
              </div>
              <CInput
                  type="text"
                  v-model="form.name"
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
                  v-model="form.desc"
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
                  v-model="form.domains"
                  :label="$t('message.product_domains')"
                  :placeholder="$t('message.product_domains_msg')"
                  v-c-tooltip.hover="$t('message.product_domains_msg')"
                  horizontal
                  required
                  was-validated
              />
            </CCardBody>
            <CCardFooter>
              <CButton type="submit" color="primary" style="width: 96px">
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
  </div>
</template>

<script>
import router from "@/router"
import clientUtils from "@/utils/api_client"

export default {
  name: 'AddProduct',
  data() {
    return {
      form: {name: "", desc: "", isPublished: false, domains: ""},
      errorMsg: "",
    }
  },
  methods: {
    doCancel() {
      router.push({name: "ProductList"})
    },
    doSubmit(e) {
      e.preventDefault()
      let vue = this
      let data = {is_published: vue.form.isPublished, name: vue.form.name, description: vue.form.desc, domains: vue.form.domains}
      clientUtils.apiDoPost(
          clientUtils.apiAdminProducts, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.$router.push({
                name: "ProductList",
                params: {flashMsg: vue.$i18n.t('message.product_added_msg', {name: vue.form.name})},
              })
            } else if (apiRes.status == 201) {
              vue.$router.push({
                name: "ProductList",
                params: {flashMsg: apiRes.message},
              })
            } else {
              vue.errorMsg = apiRes.status + ": " + apiRes.message
            }
          },
          (err) => {
            vue.errorMsg = err
          }
      )
    },
  }
}
</script>
