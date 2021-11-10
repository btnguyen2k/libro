<template>
  <div>
    <CRow>
      <CCol sm="12">
        <p v-if="foundStatus<0" class="alert alert-info">{{ $t('message.wait') }}</p>
        <p v-if="foundStatus==0" class="alert alert-danger">
          {{ $t('message.error_product_not_found', {id: $route.params.id}) }}</p>
        <CCard accent-color="info">
          <CCardHeader>
            <strong>{{ $t('message.topics') }}</strong>
            <div class="card-header-actions">
              <CButton class="btn-sm btn-primary" @click="clickAddTopic">
                <CIcon name="cil-playlist-add"/>
                {{ $t('message.add_topic') }}
              </CButton>
            </div>
          </CCardHeader>
          <CCardBody>
            <CAlert v-if="myFlashMsg" color="success" closeButton>{{ myFlashMsg }}</CAlert>
            <CDataTable :items="topicList" :fields="[
                {key:'id',label:$t('message.topic_id'),_style:'text-align: left'},
                {key:'title',label:$t('message.topic_title'),_style:'text-align: left'},
                {key:'summary',label:$t('message.topic_summary')},
                {key:'actions',label:$t('message.actions'),_style:'text-align: center'}
              ]">
              <template #id="{item}">
                <td style="white-space: nowrap;" class="col-2">
                  <CIcon :name="item.icon"/> {{ item.id }}
                </td>
              </template>
              <template #title="{item}">
                <td class="col-3">
                  {{ item.title }}
                </td>
              </template>
              <template #actions="{item}">
                <td style="white-space: nowrap; text-align: center" class="col-2">
                  <CLink @click="clickTopicPages(item.id)" class="btn btn-sm btn-success m-1">
                    <CIcon name="cil-notes" v-c-tooltip.hover="$t('message.pages')"/>
                  </CLink>
                  <CLink @click="clickEditTopic(item.id)" class="btn btn-sm btn-primary m-1">
                    <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                  </CLink>

                  <CLink @click="clickDeleteTopic(item.id)" class="btn btn-sm btn-danger m-1">
                    <CIcon name="cil-trash" v-c-tooltip.hover="$t('message.action_delete')"/>
                  </CLink>
                </td>
              </template>
            </CDataTable>
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>

    <!-- pop-up dialog to confirm deleting a topic -->
    <CModal color="warning" :title="$t('message.delete_topic')" :centered="true" :show.sync="modalDeleteShow" :close-on-backdrop="false">
      <p class="alert alert-warning"><CIcon name="cil-warning" size="lg"/> {{ $t('message.delete_topic_msg', {numPages: topicToDelete['num_pages']}) }}</p>
      <p v-if="modalDeleteErr!=''" class="alert alert-danger">{{ modalDeleteErr }}</p>
      <CInput type="text" :label="$t('message.topic_icon')+' / '+$t('message.topic_id')" v-model="topicToDelete.id" horizontal plaintext>
        <template #prepend>
          <CButton disabled link><CIcon :name="topicToDelete.icon"/></CButton>
        </template>
      </CInput>
      <CInput type="text" :label="$t('message.topic_title')" v-model="topicToDelete.title" horizontal plaintext/>
      <CTextarea rows="4" type="text" :label="$t('message.topic_summary')" v-model="topicToDelete.summary" horizontal plaintext/>
      <template #footer>
        <CButton type="button" color="danger" style="width: 96px" @click="doDeleteTopic">
          <CIcon name="cil-trash" class="align-top"/>
          {{ $t('message.action_delete') }}
        </CButton>
        <CButton type="button" color="secondary" class="ml-2" style="width: 96px" @click="modalDeleteShow = false">
          <CIcon name="cil-arrow-circle-left" class="align-top"/>
          {{ $t('message.cancel') }}
        </CButton>
      </template>
    </CModal>

    <!-- pop-up form to add new topic -->
    <CForm @submit.prevent="doAddTopic" method="post">
      <CModal :title="$t('message.add_topic')" :centered="true" :show.sync="modalAddShow" :close-on-backdrop="false">
        <p v-if="modalAddErr!=''" class="alert alert-danger">{{ modalAddErr }}</p>
        <CInput
            type="text"
            v-model="formAdd.id"
            :label="$t('message.topic_id')"
            :placeholder="$t('message.topic_id_msg')"
            v-c-tooltip.hover="$t('message.topic_id_msg')"
            horizontal
        />
        <CInput
            type="text"
            v-model="formAdd.icon"
            :label="$t('message.topic_icon')"
            :placeholder="$t('message.topic_icon_msg')"
            v-c-tooltip.hover="$t('message.topic_icon_msg')"
            horizontal
            readonly="readonly"
        >
          <template #append>
            <CButton color="primary" @click="modalIconsShow = true"><CIcon name="cil-magnifying-glass"/></CButton>
          </template>
        </CInput>
        <CInput
            type="text"
            v-model="formAdd.title"
            :label="$t('message.topic_title')"
            :placeholder="$t('message.topic_title_msg')"
            v-c-tooltip.hover="$t('message.topic_title_msg')"
            horizontal
            required
            was-validated
        />
        <CTextarea
            rows="4"
            type="text"
            v-model="formAdd.summary"
            :label="$t('message.topic_summary')"
            :placeholder="$t('message.topic_summary_msg')"
            v-c-tooltip.hover="$t('message.topic_summary_msg')"
            horizontal
            required
            was-validated
        />
        <template #footer>
          <CButton type="submit" color="primary" style="width: 96px">
            <CIcon name="cil-save" class="align-top"/>
            {{ $t('message.action_save') }}
          </CButton>
          <CButton type="button" color="secondary" class="ml-2" style="width: 96px" @click="modalAddShow = false">
            <CIcon name="cil-arrow-circle-left" class="align-top"/>
            {{ $t('message.cancel') }}
          </CButton>
        </template>
      </CModal>
    </CForm>

    <!-- pop-up dialog to pick an icon -->
    <CModal :title="$t('message.icons')" :centered="true" :show.sync="modalIconsShow">
      <CRow class="text-center">
        <template v-for="(icon, iconName) in $options.freeSet">
          <CCol class="mb-5" col="3" sm="3" :key="iconName">
            <CButton size="lg" @click="clickSelectIcon(iconName)"><CIcon size="xl" :content="icon" :title="iconName"/></CButton>
            <!--<CIcon type="button" @click="clickSelectIcon(iconName)" :height="42" :content="icon" :title="iconName"/>-->
            <div style="font-size: small">{{toKebabCase(iconName)}}</div>
          </CCol>
        </template>
      </CRow>
      <template #footer>
        <CButton @click="modalInfoShow = false" color="secondary" style="width: 96px">
          <CIcon name="cil-x" class="align-top"/>
          {{ $t('message.close') }}
        </CButton>
      </template>
    </CModal>
  </div>
</template>

<script>
import clientUtils from "@/utils/api_client";
import { freeSet } from '@coreui/icons'

export default {
  name: 'ProductTopicList',
  freeSet,
  mounted() {
    this.loadProductTopicList(this.$route.params.id)
  },
  data() {
    return {
      modalAddShow: false,
      modalAddErr: "",
      formAdd: {id: "", icon: "", title: "", summary: ""},

      modalIconsShow: false,

      modalDeleteShow: false,
      modalDeleteErr: "",
      topicToDelete: {},

      topicList: [],
      topicMap: {},

      myFlashMsg: this.flashMsg,
      errorMsg: "",
      foundStatus: -1,
    }
  },
  props: ["flashMsg"],
  methods: {
    loadProductTopicList(prodId) {
      const vue = this
      const apiUrl = clientUtils.apiAdminProductTopics.replace(':product/',prodId+'/')
      clientUtils.apiDoGet(apiUrl,
          (apiRes) => {
            vue.foundStatus = apiRes.status == 200 ? 1 : 0
            if (vue.foundStatus == 1) {
              vue.topicList = apiRes.data
              vue.topicMap = {}
              for (let i = vue.topicList.length - 1; i >= 0; i--) {
                vue.topicMap[vue.topicList[i].id] = vue.topicList[i]
              }
              console.log(vue.topicMap)
            }
          },
          (err) => {
            vue.errorMsg = err
          })
    },
    toKebabCase(str, full = false) {
      str = str.replace(/([a-z])([A-Z0-9])/g, '$1-$2').toLowerCase()
      return full ? str : str.replace(/^[a-z]+-/, '')
    },
    clickSelectIcon(iconName) {
      this.formAdd.icon = this.toKebabCase(iconName, true)
      this.modalIconsShow = false
    },
    clickAddTopic() {
      this.formAdd = {}
      this.modalAddErr = ''
      this.modalAddShow = true
    },
    doAddTopic(e) {
      e.preventDefault()
      let vue = this
      let data = vue.formAdd
      let prodId = vue.$route.params.id
      const apiUrl = clientUtils.apiAdminProductTopics.replace(':product/',prodId+'/')
      clientUtils.apiDoPost(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalAddShow = false
              vue.myFlashMsg = vue.$i18n.t('message.topic_added_msg', {name: data.title})
              vue.loadProductTopicList(prodId)
            } else {
              vue.modalAddErr = apiRes.status + ": " + apiRes.message
            }
          },
          (err) => {
            vue.modalAddErr = err
          }
      )
    },
    clickTopicPages(id) {
      console.log(id)
    },
    clickEditTopic(id) {
      console.log(id)
      // this.$router.push({name: "EditProduct", params: {id: id.toString()}})
    },
    clickDeleteTopic(id) {
      this.topicToDelete = this.topicMap[id]
      this.modalDeleteShow = true
    },
    doDeleteTopic(e) {
      e.preventDefault()
      let vue = this
      let prodId = vue.$route.params.id
      let data = vue.topicToDelete
      const apiUrl = clientUtils.apiAdminProductTopic.replace(':product/',prodId+'/')+'/'+data.id
      console.log(vue.topicToDelete)
      console.log(apiUrl)
      clientUtils.apiDoDelete(apiUrl,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalDeleteShow = false
              vue.myFlashMsg = vue.$i18n.t('message.topic_deleted_msg', {name: data.title})
              vue.loadProductTopicList(prodId)
            } else {
              vue.modalDeleteErr = apiRes.status + ": " + apiRes.message
            }
          },
          (err) => {
            vue.modalDeleteErr = err
          }
      )
    },
  }
}
</script>
