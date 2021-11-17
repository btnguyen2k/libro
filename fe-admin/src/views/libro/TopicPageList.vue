<template>
  <div>
    <CRow>
      <CCol sm="12">
        <CAlert v-if="foundStatus<0" color="info">{{ $t('message.wait') }}</CAlert>
        <CCard v-if="foundStatus==0" color="danger" text-color="white">
          <CCardHeader>
            {{ $t('message.error_topic_not_found', {id: $route.params.tid}) }}
            <div class="card-header-actions">
              <CButton color="light" size="sm" @click="clickGoback">
                <CIcon name="cil-arrow-circle-left"/>
                {{ $t('message.action_back') }}
              </CButton>
            </div>
          </CCardHeader>
        </CCard>
        <CAlert v-if="errorMsg" color="danger" closeButton>{{ errorMsg }}</CAlert>
        <CCard accent-color="info" v-if="foundStatus>0">
          <CCardHeader>
            <strong>{{ $t('message.pages') }}</strong>
            <div class="card-header-actions">
              <CButton color="secondary" size="sm" class="m-2" @click="clickGoback">
                <CIcon name="cil-arrow-circle-left"/>
                {{ $t('message.action_back') }}
              </CButton>
              <CButton color="primary" size="sm" @click="clickAddPage">
                <CIcon name="cil-note-add"/>
                {{ $t('message.add_page') }}
              </CButton>
            </div>
          </CCardHeader>
          <CCardBody>
            <CAlert v-if="myFlashMsg" color="success" closeButton>{{ myFlashMsg }}</CAlert>
            <CDataTable :items="pageList" :fields="[
                {key:'id',label:$t('message.page_id'),_style:'text-align: left'},
                {key:'title',label:$t('message.page_title'),_style:'text-align: left'},
                {key:'summary',label:$t('message.page_summary')},
                {key:'actions',label:$t('message.actions'),_style:'text-align: center'}
              ]">
              <template #id="{item}">
                <td style="white-space: nowrap; font-size: small" class="col-2">
                  <CIcon :name="item.icon"/>
                  {{ item.id }}
                </td>
              </template>
              <template #title="{item}">
                <td class="col-3">
                  {{ item.title }}
                </td>
              </template>
              <template #actions="{item}">
                <td style="white-space: nowrap; text-align: center">
                  <CButton @click="clickEditPage(item.id)" color="primary" size="sm" class="mr-1">
                    <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                  </CButton>
                  <CButton @click="clickDeletePage(item.id)" color="danger" size="sm" class="mr-3">
                    <CIcon name="cil-trash" v-c-tooltip.hover="$t('message.action_delete')"/>
                  </CButton>

                  <CButton @click="doMovePageDown(item.id)" color="info" class="mr-1" size="sm" variant="outline">
                    <CIcon name="cil-arrow-bottom" v-c-tooltip.hover="$t('message.action_move_down')"/>
                  </CButton>
                  <CButton @click="doMovePageUp(item.id)" color="info" size="sm" variant="outline">
                    <CIcon name="cil-arrow-top" v-c-tooltip.hover="$t('message.action_move_up')"/>
                  </CButton>
                </td>
              </template>
            </CDataTable>
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>

    <!--    &lt;!&ndash; pop-up dialog to confirm deleting a topic &ndash;&gt;-->
    <!--    <CModal color="warning" :title="$t('message.delete_topic')" :centered="true" :show.sync="modalDeleteShow"-->
    <!--            :close-on-backdrop="false">-->
    <!--      <p class="alert alert-warning">-->
    <!--        <CIcon name="cil-warning" size="lg"/>-->
    <!--        {{ $t('message.delete_topic_msg', {numPages: topicToDelete['num_pages']}) }}-->
    <!--      </p>-->
    <!--      <p v-if="modalDeleteErr!=''" class="alert alert-danger">{{ modalDeleteErr }}</p>-->
    <!--      <CInput type="text" :label="$t('message.topic_icon')+' / '+$t('message.topic_id')" v-model="topicToDelete.id"-->
    <!--              horizontal plaintext>-->
    <!--        <template #prepend>-->
    <!--          <CButton disabled link>-->
    <!--            <CIcon :name="topicToDelete.icon"/>-->
    <!--          </CButton>-->
    <!--        </template>-->
    <!--      </CInput>-->
    <!--      <CInput type="text" :label="$t('message.topic_title')" v-model="topicToDelete.title" horizontal plaintext/>-->
    <!--      <CTextarea rows="4" type="text" :label="$t('message.topic_summary')" v-model="topicToDelete.summary" horizontal-->
    <!--                 plaintext/>-->
    <!--      <template #footer>-->
    <!--        <CButton type="button" color="danger" class="m-2" style="width: 96px" @click="doDeleteTopic">-->
    <!--          <CIcon name="cil-trash" class="align-top"/>-->
    <!--          {{ $t('message.action_delete') }}-->
    <!--        </CButton>-->
    <!--        <CButton type="button" color="secondary" style="width: 96px" @click="modalDeleteShow = false">-->
    <!--          <CIcon name="cil-arrow-circle-left" class="align-top"/>-->
    <!--          {{ $t('message.cancel') }}-->
    <!--        </CButton>-->
    <!--      </template>-->
    <!--    </CModal>-->

    <!-- pop-up form to add new page -->
    <CForm @submit.prevent="doAddPage" method="post">
      <CModal size="lg" :title="$t('message.add_page')" :centered="true" :show.sync="modalAddShow"
              :close-on-backdrop="false">
        <p v-if="modalAddErr!=''" class="alert alert-danger">{{ modalAddErr }}</p>
        <CInput
            type="text"
            v-model="formAdd.id"
            :label="$t('message.page_id')"
            :placeholder="$t('message.page_id_msg')"
            v-c-tooltip.hover="$t('message.page_id_msg')"
            horizontal
        />
        <CInput
            type="text"
            v-model="formAdd.icon"
            :label="$t('message.page_icon')"
            :placeholder="$t('message.page_icon_msg')"
            v-c-tooltip.hover="$t('message.page_icon_msg')"
            horizontal
            readonly
        >
          <template #prepend>
            <CButton disabled link>
              <CIcon :name="formAdd.icon"/>
            </CButton>
          </template>
          <template #append>
            <CButton color="primary" @click="modalIconsShow = true">
              <CIcon name="cil-magnifying-glass"/>
            </CButton>
          </template>
        </CInput>
        <CInput
            type="text"
            v-model="formAdd.title"
            :label="$t('message.page_title')"
            :placeholder="$t('message.page_title_msg')"
            v-c-tooltip.hover="$t('message.page_title_msg')"
            horizontal
            required
            was-validated
        />
        <CTextarea
            rows="2"
            type="text"
            v-model="formAdd.summary"
            :label="$t('message.page_summary')"
            :placeholder="$t('message.page_summary_msg')"
            v-c-tooltip.hover="$t('message.page_summary_msg')"
            horizontal
            required
            was-validated
        />
        <CTabs>
          <CTab active>
            <template slot="title">
              <CIcon name="cib-markdown"/>
              {{ $t('message.content_editor') }}
            </template>
            <CTextarea
                rows="8"
                type="text"
                v-model="formAdd.content"
                :label="$t('message.page_content')"
                :placeholder="$t('message.page_content_msg')"
                horizontal
                required
                was-validated
            />
          </CTab>
          <CTab>
            <template slot="title">
              <CIcon name="cil-calculator"/>
              {{ $t('message.content_preview') }}
            </template>
            <div v-html="previewPageContent"></div>
          </CTab>
        </CTabs>
        <template #footer>
          <CButton type="submit" color="primary" class="m-2" style="width: 96px">
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

    <!-- pop-up form to edit existing page -->
    <CForm @submit.prevent="doEditPage" method="post">
      <CModal size="lg" :title="$t('message.edit_page')" :centered="true" :show.sync="modalEditShow"
              :close-on-backdrop="false">
        <p v-if="modalEditErr!=''" class="alert alert-danger">{{ modalEditErr }}</p>
        <CInput
            type="text"
            v-model="formEdit.id"
            :label="$t('message.page_id')"
            :placeholder="$t('message.page_id_msg')"
            v-c-tooltip.hover="$t('message.page_id_msg')"
            horizontal
            readonly
        />
        <CInput
            type="text"
            v-model="formEdit.icon"
            :label="$t('message.page_icon')"
            :placeholder="$t('message.page_icon_msg')"
            v-c-tooltip.hover="$t('message.page_icon_msg')"
            horizontal
            readonly
        >
          <template #prepend>
            <CButton disabled link>
              <CIcon :name="formEdit.icon"/>
            </CButton>
          </template>
          <template #append>
            <CButton color="primary" @click="modalIconsShow = true">
              <CIcon name="cil-magnifying-glass"/>
            </CButton>
          </template>
        </CInput>
        <CInput
            type="text"
            v-model="formEdit.title"
            :label="$t('message.page_title')"
            :placeholder="$t('message.page_title_msg')"
            v-c-tooltip.hover="$t('message.page_title_msg')"
            horizontal
            required
            was-validated
        />
        <CTextarea
            rows="2"
            type="text"
            v-model="formEdit.summary"
            :label="$t('message.page_summary')"
            :placeholder="$t('message.page_summary_msg')"
            v-c-tooltip.hover="$t('message.page_summary_msg')"
            horizontal
            required
            was-validated
        />
        <CTabs>
          <CTab active>
            <template slot="title">
              <CIcon name="cib-markdown"/>
              {{ $t('message.content_editor') }}
            </template>
            <CTextarea
                rows="8"
                type="text"
                v-model="formEdit.content"
                :label="$t('message.page_content')"
                :placeholder="$t('message.page_content_msg')"
                horizontal
                required
                was-validated
            />
          </CTab>
          <CTab>
            <template slot="title">
              <CIcon name="cil-calculator"/>
              {{ $t('message.content_preview') }}
            </template>
            <div v-html="previewPageContent"></div>
          </CTab>
        </CTabs>
        <template #footer>
          <CButton type="submit" color="primary" class="m-2" style="width: 96px">
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

    <!-- pop-up dialog to pick an icon -->
    <CModal :title="$t('message.icons')" :centered="true" :show.sync="modalIconsShow">
      <CRow class="text-center">
        <template v-for="(icon, iconName) in $options.freeSet">
          <CCol class="mb-5" col="3" sm="3" :key="iconName">
            <CButton size="lg" @click="clickSelectIcon(iconName)">
              <CIcon size="xl" :content="icon" :title="iconName"/>
            </CButton>
            <!--<CIcon type="button" @click="clickSelectIcon(iconName)" :height="42" :content="icon" :title="iconName"/>-->
            <div style="font-size: small">{{ toKebabCase(iconName) }}</div>
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
import clientUtils from "@/utils/api_client"
import {freeSet} from '@coreui/icons'
import {markdownRender} from './utils'

export default {
  name: 'TopicPageList',
  freeSet,
  computed: {
    previewPageContent() {
      return markdownRender(this.addMode ? this.formAdd.content : this.formEdit.content, true)
    }
  },
  mounted() {
    this.loadTopicPageList(this.$route.params.tid)
  },
  data() {
    return {
      addMode: Boolean,

      modalAddShow: false,
      modalAddErr: "",
      formAdd: {id: "", icon: "", title: "", summary: "", content: ""},

      modalEditShow: false,
      modalEditErr: "",
      formEdit: {id: "", icon: "", title: "", summary: "", content: ""},

      modalIconsShow: false,

      modalDeleteShow: false,
      modalDeleteErr: "",
      pageToDelete: {},

      pageList: [],
      pageMap: {},

      myFlashMsg: this.flashMsg,
      errorMsg: "",
      foundStatus: -1,
    }
  },
  props: ["flashMsg"],
  methods: {
    loadTopicPageList(topicId) {
      this.foundStatus = -1
      const vue = this
      const apiUrl = clientUtils.apiAdminTopicPages.replaceAll(':topic', topicId)
      clientUtils.apiDoGet(apiUrl,
          (apiRes) => {
            vue.foundStatus = apiRes.status == 200 ? 1 : 0
            if (vue.foundStatus == 1) {
              vue.pageList = apiRes.data
              vue.pageMap = {}
              for (let i = vue.pageList.length - 1; i >= 0; i--) {
                vue.pageMap[vue.pageList[i].id] = vue.pageList[i]
              }
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
    clickGoback() {
      let prodId = this.$route.params.pid
      this.$router.push({name: "ProductTopicList", params: {pid: prodId}})
    },
    clickSelectIcon(iconName) {
      if (this.addMode) {
        this.formAdd.icon = this.toKebabCase(iconName, true)
      } else {
        this.formEdit.icon = this.toKebabCase(iconName, true)
      }
      this.modalIconsShow = false
    },
    clickAddPage() {
      this.addMode = true
      this.formAdd = {id: "", icon: "", title: "", summary: "", content: ""}
      this.modalAddErr = ''
      this.modalAddShow = true
    },
    doAddPage(e) {
      e.preventDefault()
      let vue = this
      let data = vue.formAdd
      let topicId = vue.$route.params.tid
      const apiUrl = clientUtils.apiAdminTopicPages.replaceAll(':topic', topicId)
      clientUtils.apiDoPost(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalAddShow = false
              vue.myFlashMsg = vue.$i18n.t('message.page_added_msg', {name: data.title})
              vue.loadTopicPageList(topicId)
            } else {
              vue.modalAddErr = apiRes.status + ": " + apiRes.message
            }
          },
          (err) => {
            vue.modalAddErr = err
          }
      )
    },
    clickEditPage(id) {
      this.addMode = false
      this.formEdit = {...this.pageMap[id]} //shallow clone using spread syntax, alternative way: this.formEdit = Object.assign({}, this.pageMap[id])
      this.modalEditErr = ''
      this.modalEditShow = true
    },
    doEditPage(e) {
      e.preventDefault()
      let vue = this
      let data = vue.formEdit
      let topicId = vue.$route.params.tid
      const apiUrl = clientUtils.apiAdminTopicPage.replaceAll(':topic', topicId).replaceAll(':page', data.id)
      clientUtils.apiDoPut(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalEditShow = false
              vue.myFlashMsg = vue.$i18n.t('message.page_updated_msg', {name: data.title})
              vue.loadTopicPageList(topicId)
            } else {
              vue.modalEditErr = apiRes.status + ": " + apiRes.message
            }
          },
          (err) => {
            vue.modalEditErr = err
          }
      )
    },
    clickDeletePage(id) {
      this.pageToDelete = this.pageMap[id]
      this.modalDeleteShow = true
    },
    doDeletePage(e) {
      e.preventDefault()
      let vue = this
      let topicId = vue.$route.params.tid
      let data = vue.pageToDelete
      const apiUrl = clientUtils.apiAdminTopicPage.replaceAll(':topic', topicId).replaceAll(':page', data.id)
      clientUtils.apiDoDelete(apiUrl,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalDeleteShow = false
              vue.myFlashMsg = vue.$i18n.t('message.page_deleted_msg', {name: data.title})
              vue.loadTopicPageList(topicId)
            } else {
              vue.modalDeleteErr = apiRes.status + ": " + apiRes.message
            }
          },
          (err) => {
            vue.modalDeleteErr = err
          }
      )
    },
    _doMovePageUpOrDown(id, data) {
      const saveStatus = this.foundStatus
      this.foundStatus = -1
      let vue = this
      let topicId = vue.$route.params.tid
      const apiUrl = clientUtils.apiAdminTopicPage.replaceAll(':topic', topicId).replaceAll(':page', id)
      clientUtils.apiDoPatch(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.myFlashMsg = vue.$i18n.t('message.page_updated_msg', {name: this.pageMap[id].title})
              vue.loadTopicPageList(topicId)
            } else {
              vue.errorMsg = apiRes.status + ": " + apiRes.message
              vue.foundStatus = saveStatus
            }
          },
          (err) => {
            vue.errorMsg = err
            vue.foundStatus = saveStatus
          }
      )
    },
    doMovePageUp(id) {
      this._doMovePageUpOrDown(id, {action: "move_up"})
    },
    doMovePageDown(id) {
      this._doMovePageUpOrDown(id, {action: "move_down"})
    },
  }
}
</script>
