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
          <!--
          <template #footer>
            <CCardFooter class="px-3 py-2">
              <CLink
                  class="font-weight-bold font-xs text-muted d-flex justify-content-between"
                  href="https://coreui.io/"
                  target="_blank"
              >
                View more
                <CIcon name="cil-arrow-right" width="16"/>
              </CLink>
            </CCardFooter>
          </template>
          -->
        </CWidgetIcon>
      </CCol>
    </CRow>

    <CRow>
      <CCol sm="12">
        <CCard accent-color="info">
          <CCardHeader>
            <strong>{{ $t('message.products') }}</strong>
            <div class="card-header-actions">
              <CButton class="btn-sm btn-primary" @click="clickAddProduct">
                <CIcon name="cil-plus"/>
                {{ $t('message.add_product') }}
              </CButton>
            </div>
          </CCardHeader>
          <CCardBody>
<!--            <p v-if="flashMsg" class="alert alert-success">{{ flashMsg }}</p>-->
<!--            <CDataTable :items="appList" :fields="[-->
<!--              {key:'public',label:''},-->
<!--              {key:'created',label:$t('message.blog_tcreated')},-->
<!--              {key:'title',label:$t('message.blog_title')},-->
<!--              {key:'num_comments',label:$t('message.blog_comments')},-->
<!--              {key:'num_votes_up',label:$t('message.blog_votes')+' ↑',_style:'white-space: nowrap'},-->
<!--              {key:'num_votes_down',label:$t('message.blog_votes')+' ↓',_style:'white-space: nowrap'},-->
<!--              {key:'actions',label:$t('message.actions'),_style:'text-align: center'}-->
<!--            ]">-->
<!--              <template #public="{item}">-->
<!--                <td>-->
<!--                  <CIcon :name="`${item.is_public?'cil-check':'cil-check-alt'}`"-->
<!--                         :style="`color: ${item.is_public?'green':'grey'}`"/>-->
<!--                </td>-->
<!--              </template>-->
<!--              <template #created="{item}">-->
<!--                <td style="font-size: smaller; white-space: nowrap">{{item.t_created.substring(0,19)}} (GMT{{item.t_created.substring(26)}})</td>-->
<!--              </template>-->
<!--              <template #title="{item}">-->
<!--                <td style="font-size: smaller">{{item.title}}</td>-->
<!--              </template>-->
<!--              <template #num_comments="{item}">-->
<!--                <td style="font-size: smaller; text-align: center">{{item.num_comments}}</td>-->
<!--              </template>-->
<!--              <template #num_votes_up="{item}">-->
<!--                <td style="font-size: smaller; text-align: center">{{item.num_votes_up}}</td>-->
<!--              </template>-->
<!--              <template #num_votes_down="{item}">-->
<!--                <td style="font-size: smaller; text-align: center">{{item.num_votes_down}}</td>-->
<!--              </template>-->
<!--              <template #actions="{item}">-->
<!--                <td style="font-size: smaller; white-space: nowrap; text-align: center">-->
<!--                  <CLink @click="clickEditBlogPost(item.id)" :label="$t('message.action_edit')" class="btn btn-sm btn-primary">-->
<!--                    <CIcon name="cil-pencil"/>-->
<!--                  </CLink>-->
<!--                  &nbsp;-->
<!--                  <CLink @click="clickDeleteBlogPost(item.id)" :label="$t('message.action_delete')" class="btn btn-sm btn-danger">-->
<!--                    <CIcon name="cil-trash"/>-->
<!--                  </CLink>-->
<!--                </td>-->
<!--              </template>-->
<!--            </CDataTable>-->
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  </div>
</template>

<script>
import clientUtils from "@/utils/api_client"
// import marked from "marked"
// import DOMPurify from "dompurify"

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
      stats: {},
      prodList: [],
    }
  },
  methods: {
    clickAddProduct() {
      this.$router.push({name: "AddProduct"})
    },
    // voteValue(post) {
    //   return this.blogPostVotes[post.id]
    // },
    // avatar(post) {
    //   return "https://www.gravatar.com/avatar/" + post.owner.id.trim().toLowerCase().md5() + "?s=40"
    // },
    // displayName(post) {
    //   return post.owner.display_name
    // },
    // creationTime(post) {
    //   return post.t_created.substring(0, 19) + ' (GMT' + post.t_created.substring(26) + ')'
    // },
    // renderMarkdown(post) {
    //   return DOMPurify.sanitize(marked(post.content), {ADD_ATTR: ['target']})
    // },
    // doVote(postId, v) {
    //   const data = {vote: v}
    //   const vue = this
    //   clientUtils.apiDoPost(clientUtils.apiUserVoteForPost + "/" + postId, data,
    //       (apiRes) => {
    //         // console.log(apiRes)
    //         if (apiRes.status == 200 && apiRes.data.vote) {
    //           vue.blogPostVotes[postId] = apiRes.data.value
    //           vue.blogPostMap[postId].num_votes_up = apiRes.data.num_votes_up
    //           vue.blogPostMap[postId].num_votes_down = apiRes.data.num_votes_down
    //         }
    //       },
    //       (err) => {
    //         console.error("Error voting for post: " + err)
    //       })
    // }
  }
}
</script>
