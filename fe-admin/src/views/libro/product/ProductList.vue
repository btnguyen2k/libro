<template>
  <CRow>
    <CCol sm="12">
      <CCard accent-color="info">
        <CCardHeader>
          <strong>{{ $t('message.products') }}</strong>
          <div class="card-header-actions">
            <CButton class="btn-sm btn-primary" @click="clickAddProduct">
              <CIcon name="cil-image-plus"/>
              {{ $t('message.add_product') }}
            </CButton>
          </div>
        </CCardHeader>
        <CCardBody>
          <p v-if="flashMsg" class="alert alert-success">{{ flashMsg }}</p>
          <CDataTable :items="prodList" :fields="[
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
  </CRow>
</template>

<script>
import clientUtils from "@/utils/api_client";

export default {
  name: 'Products',
  mounted() {
    const vue = this
    clientUtils.apiDoGet(clientUtils.apiAdminProducts,
        (apiRes) => {
          if (apiRes.status == 200) {
            vue.prodList = apiRes.data
          } else {
            console.error("Getting product list was unsuccessful: " + apiRes)
          }
        },
        (err) => {
          console.error("Error getting product list: " + err)
        })
  },
  data() {
    return {
      prodList: [],
    }
  },
  props: ["flashMsg"],
  methods: {
    clickAddProduct() {
      this.$router.push({name: "AddProduct"})
    },
    clickEditProduct(id) {
      this.$router.push({name: "EditProduct", params: {id: id.toString()}})
    },
    clickDeleteProduct(id) {
      this.$router.push({name: "DeleteProduct", params: {id: id.toString()}})
    },
  }
}
</script>
