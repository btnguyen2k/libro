<template>
  <div>
    <CRow>
      <CCol col="12" sm="6" lg="4">
        <CWidgetIcon
            :header="String(stats.num_products)"
            :text="$t('message.products')"
            color="gradient-primary"
            :icon-padding="false"
        >
          <CIcon name="cil-applications" class="mx-5 " width="24"/>
        </CWidgetIcon>
      </CCol>
      <CCol col="12" sm="6" lg="4">
        <CWidgetIcon
            :header="String(stats.num_topics)"
            text="Topics"
            color="gradient-info"
            :icon-padding="false"
        >
          <CIcon name="cil-list-rich" class="mx-5 " width="24"/>
        </CWidgetIcon>
      </CCol>
      <CCol col="12" sm="6" lg="4">
        <CWidgetIcon
            :header="String(stats.num_pages)"
            text="pages"
            color="gradient-success"
            :icon-padding="false"
        >
          <CIcon name="cil-notes" class="mx-5 " width="24"/>
        </CWidgetIcon>
      </CCol>
    </CRow>
  </div>
</template>

<script>
import clientUtils from "@/utils/api_client"

export default {
  name: 'Dashboard',
  mounted() {
    const vue = this
    clientUtils.apiDoGet(clientUtils.apiAdminStats,
        (apiRes) => {
          if (apiRes.status == 200) {
            vue.stats = apiRes.data
          } else {
            console.error("Getting stats was unsuccessful: " + apiRes)
          }
        },
        (err) => {
          console.error("Error getting stats: " + err)
        })
  },
  data() {
    return {
      stats: {},
    }
  },
  methods: {
  }
}
</script>
