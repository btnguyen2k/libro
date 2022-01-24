<template>
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
                <ficon fixedWidth :icon="_iconize(item.icon)"/>
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

    <!-- pop-up dialog to confirm deleting a page -->
    <CModal color="warning" size="lg" :title="$t('message.delete_page')" :centered="true" :show.sync="modalDeleteShow" :close-on-backdrop="false">
      <p class="alert alert-warning">
        <CIcon name="cil-warning" size="lg"/>
        {{ $t('message.delete_page_msg') }}
      </p>
      <CAlert v-if="waitDeletePage" color="info">{{ $t('message.wait') }}</CAlert>
      <CAlert v-if="modalDeleteErr" color="danger">{{ modalDeleteErr }}</CAlert>
      <CInput type="text" :label="$t('message.page_icon')+' / '+$t('message.page_id')" v-model="pageToDelete.id" horizontal plaintext>
        <template #prepend>
          <CButton disabled link>
            <ficon :icon="_iconize(pageToDelete.icon)"/>
          </CButton>
        </template>
      </CInput>
      <CInput type="text" :label="$t('message.page_title')" v-model="pageToDelete.title" horizontal plaintext/>
      <CTextarea rows="2" type="text" :label="$t('message.page_summary')" v-model="pageToDelete.summary" horizontal plaintext/>
      <CTextarea rows="8" type="text" :label="$t('message.page_content')" v-model="pageToDelete.content" horizontal plaintext/>
      <template #footer>
        <CButton v-if="!waitDeletePage" type="button" color="danger" class="m-2" style="width: 96px" @click="doDeletePage">
          <CIcon name="cil-trash" class="align-top"/>
          {{ $t('message.action_delete') }}
        </CButton>
        <CButton type="button" color="secondary" style="width: 96px" @click="modalDeleteShow = false">
          <CIcon name="cil-arrow-circle-left" class="align-top"/>
          {{ $t('message.cancel') }}
        </CButton>
      </template>
    </CModal>

    <!-- pop-up form to add new page -->
    <CForm @submit.prevent="doAddPage" method="post">
      <CModal size="lg" :title="$t('message.add_page')" :centered="true" :show.sync="modalAddShow" :close-on-backdrop="false">
        <CAlert v-if="waitAddPage" color="info">{{ $t('message.wait') }}</CAlert>
        <CAlert v-if="modalAddErr" color="danger">{{ modalAddErr }}</CAlert>
        <div v-if="!waitAddPage">
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
              <CButton disabled link><ficon :icon="_iconize(formAdd.icon)"/></CButton>
            </template>
            <template #append>
              <CButton color="primary" @click="modalIconsShow = true"><ficon :icon="['fas', 'search']"/></CButton>
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
        </div>
        <template #footer>
          <CButton v-if="!waitAddPage" type="submit" color="primary" class="m-2" style="width: 96px">
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
      <CModal size="lg" :title="$t('message.edit_page')" :centered="true" :show.sync="modalEditShow" :close-on-backdrop="false">
        <CAlert v-if="waitEditPage" color="info">{{ $t('message.wait') }}</CAlert>
        <CAlert v-if="modalEditErr" color="danger">{{ modalEditErr }}</CAlert>
        <div v-if="!waitEditPage">
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
                <ficon :icon="_iconize(formEdit.icon)"/>
              </CButton>
            </template>
            <template #append>
              <CButton color="primary" @click="modalIconsShow = true">
                <ficon :icon="['fas', 'search']"/>
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
        </div>
        <template #footer>
          <CButton v-if="!waitEditPage" type="submit" color="primary" class="m-2" style="width: 96px">
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
    <CModal color="info" :title="$t('message.icons')" :centered="true" :show.sync="modalIconsShow">
      <CRow class="text-center">
        <CCol col="12" sm="12">
          <CDataTable :items="faIconList" :fields="[
                {key:'prefix', label: $t('message.icon_icon')},
                {key:'iconName', label: $t('message.icon_name')},
              ]" pagination :items-per-page="10" hover striped border small table-filter>
            <template #prefix="{item}">
              <td style="cursor: pointer" @click="clickSelectIcon(item.prefix+'-'+item.iconName)"><ficon fixedWidth :icon="[item.prefix, item.iconName]"/></td>
            </template>
            <template #iconName="{item}">
              <td style="cursor: pointer" @click="clickSelectIcon(item.prefix+'-'+item.iconName)">{{ item.prefix }}-{{ item.iconName }}</td>
            </template>
          </CDataTable>
        </CCol>
      </CRow>
      <template #footer>
        <CButton @click="modalIconsShow = false" color="secondary" style="width: 96px">
          <CIcon name="cil-x" class="align-top"/>
          {{ $t('message.close') }}
        </CButton>
      </template>
    </CModal>
  </CRow>
</template>

<script>
import clientUtils from "@/utils/api_client"
import {iconize} from './utils'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
import {markdownRender} from './utils'

export default {
  name: 'TopicPageList',
  computed: {
    previewPageContent() {
      return markdownRender(this.addMode ? this.formAdd.content : this.formEdit.content, true)
    },
    faIconList() {
      let result = []
      for (let k in fas) {
        const item = fas[k]
        result = result.concat({prefix: item.prefix, iconName: item.iconName})
      }
      for (let k in far) {
        const item = far[k]
        result = result.concat({prefix: item.prefix, iconName: item.iconName})
      }
      for (let k in fab) {
        const item = fab[k]
        result = result.concat({prefix: item.prefix, iconName: item.iconName})
      }
      return result
    },
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
      waitAddPage: false,

      modalEditShow: false,
      modalEditErr: "",
      formEdit: {id: "", icon: "", title: "", summary: "", content: ""},
      waitEditPage: false,

      modalIconsShow: false,

      modalDeleteShow: false,
      modalDeleteErr: "",
      pageToDelete: {},
      waitDeletePage: false,

      pageList: [],
      pageMap: {},

      myFlashMsg: this.flashMsg,
      errorMsg: "",
      foundStatus: -1,
    }
  },
  props: ["flashMsg"],
  methods: {
    _iconize(icon) {
      return iconize(icon)
    },
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
    // toKebabCase(str, full = false) {
    //   str = str.replace(/([a-z])([A-Z0-9])/g, '$1-$2').toLowerCase()
    //   return full ? str : str.replace(/^[a-z]+-/, '')
    // },
    clickGoback() {
      let prodId = this.$route.params.pid
      this.$router.push({name: "ProductTopicList", params: {pid: prodId}})
    },
    clickSelectIcon(iconName) {
      if (this.addMode) {
        this.formAdd.icon = iconName
        // this.formAdd.icon = this.toKebabCase(iconName, true)
      } else {
        this.formEdit.icon = iconName
        // this.formEdit.icon = this.toKebabCase(iconName, true)
      }
      this.modalIconsShow = false
    },
    clickAddPage() {
      this.addMode = true
      this.formAdd = {id: "", icon: "", title: "", summary: "", content: ""}
      this.modalAddErr = ''
      this.modalAddShow = true
      this.myFlashMsg = ''
    },
    doAddPage(e) {
      e.preventDefault()
      this.modalAddErr = ''
      let vue = this
      let data = {...vue.formAdd}
      vue.waitAddPage = true
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
            vue.waitAddPage = false
          },
          (err) => {
            vue.modalAddErr = err
            vue.waitAddPage = false
          }
      )
    },
    clickEditPage(id) {
      this.addMode = false
      this.formEdit = {...this.pageMap[id]} //shallow clone using spread syntax, alternative way: this.formEdit = Object.assign({}, this.pageMap[id])
      this.modalEditErr = ''
      this.modalEditShow = true
      this.myFlashMsg = ''
    },
    doEditPage(e) {
      e.preventDefault()
      this.modalEditErr = ''
      let vue = this
      let data = {...vue.formEdit}
      vue.waitEditPage = true
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
            vue.waitEditPage = false
          },
          (err) => {
            vue.modalEditErr = err
            vue.waitEditPage = false
          }
      )
    },
    clickDeletePage(id) {
      this.pageToDelete = this.pageMap[id]
      this.modalDeleteErr = ''
      this.modalDeleteShow = true
      this.myFlashMsg = ''
    },
    doDeletePage(e) {
      e.preventDefault()
      this.modalDeleteErr = ''
      let vue = this
      vue.waitDeletePage = true
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
            vue.waitDeletePage = false
          },
          (err) => {
            vue.modalDeleteErr = err
            vue.waitDeletePage = false
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
