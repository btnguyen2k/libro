<template>
  <CRow>
    <CCol sm="12">
      <CAlert v-if="foundStatus<0" color="info">{{ $t('message.wait') }}</CAlert>
      <CCard v-if="foundStatus==0" color="danger" text-color="white">
        <CCardHeader>
          {{ $t('message.error_product_not_found', {id: $route.params.pid}) }}
          <div class="card-header-actions">
            <CButton color="light" size="sm" @click="clickGoback">
              <CIcon name="cil-arrow-circle-left"/> {{ $t('message.action_back') }}
            </CButton>
          </div>
        </CCardHeader>
      </CCard>
      <CAlert v-if="errorMsg" color="danger" closeButton>{{ errorMsg }}</CAlert>
      <CCard accent-color="info" v-if="foundStatus>0">
        <CCardHeader>
          <strong>{{ $t('message.topics') }}</strong>
          <div class="card-header-actions">
            <CButton color="secondary" size="sm" class="m-2" @click="clickGoback">
              <CIcon name="cil-arrow-circle-left"/>
              {{ $t('message.action_back') }}
            </CButton>
            <CButton color="primary" size="sm"  @click="clickAddTopic">
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
                <CButton @click="clickTopicPages(item.id)" color="success" size="sm" class="mr-1">
                  <CIcon name="cil-notes" v-c-tooltip.hover="$t('message.pages')"/>
                </CButton>
                <CButton @click="clickEditTopic(item.id)" color="primary" size="sm" class="mr-1">
                  <CIcon name="cil-pencil" v-c-tooltip.hover="$t('message.action_edit')"/>
                </CButton>
                <CButton @click="clickDeleteTopic(item.id)" color="danger" size="sm" class="mr-3">
                  <CIcon name="cil-trash" v-c-tooltip.hover="$t('message.action_delete')"/>
                </CButton>

                <CButton @click="doMoveTopicDown(item.id)" color="info" class="mr-1" size="sm" variant="outline">
                  <CIcon name="cil-arrow-bottom" v-c-tooltip.hover="$t('message.action_move_down')"/>
                </CButton>
                <CButton @click="doMoveTopicUp(item.id)" color="info" size="sm" variant="outline">
                  <CIcon name="cil-arrow-top" v-c-tooltip.hover="$t('message.action_move_up')"/>
                </CButton>
              </td>
            </template>
          </CDataTable>
        </CCardBody>
        <CCardFooter>
          <CButton color="secondary" size="sm" class="m-2" @click="clickGoback">
            <CIcon name="cil-arrow-circle-left"/>
            {{ $t('message.action_back') }}
          </CButton>
          <CButton color="primary" size="sm"  @click="clickAddTopic">
            <CIcon name="cil-playlist-add"/>
            {{ $t('message.add_topic') }}
          </CButton>
        </CCardFooter>
      </CCard>
    </CCol>

    <!-- pop-up dialog to confirm deleting a topic -->
    <CModal color="warning" :title="$t('message.delete_topic')" :centered="true" :show.sync="modalDeleteShow" :close-on-backdrop="false">
      <p class="alert alert-warning">
        <CIcon name="cil-warning" size="lg"/>
        {{ $t('message.delete_topic_msg', {numPages: topicToDelete['num_pages']}) }}
      </p>
      <CAlert v-if="waitDeleteTopic" color="info">{{ $t('message.wait') }}</CAlert>
      <CAlert v-if="modalDeleteErr" color="danger">{{ modalDeleteErr }}</CAlert>
      <CInput type="text" :label="$t('message.topic_icon')+' / '+$t('message.topic_id')" v-model="topicToDelete.id" horizontal plaintext>
        <template #prepend>
          <CButton disabled link>
            <ficon :icon="_iconize(topicToDelete.icon)"/>
          </CButton>
        </template>
      </CInput>
      <CInput type="text" :label="$t('message.topic_title')" v-model="topicToDelete.title" horizontal plaintext/>
      <CTextarea rows="4" type="text" :label="$t('message.topic_summary')" v-model="topicToDelete.summary" horizontal plaintext/>
      <template #footer>
        <CButton v-if="!waitDeleteTopic" type="button" color="danger" class="m-2" style="width: 96px" @click="doDeleteTopic">
          <CIcon name="cil-trash" class="align-top"/>
          {{ $t('message.action_delete') }}
        </CButton>
        <CButton type="button" color="secondary" style="width: 96px" @click="modalDeleteShow = false">
          <CIcon name="cil-arrow-circle-left" class="align-top"/>
          {{ $t('message.cancel') }}
        </CButton>
      </template>
    </CModal>

    <!-- pop-up form to add new topic -->
    <CForm @submit.prevent="doAddTopic" method="post">
      <CModal :title="$t('message.add_topic')" :centered="true" :show.sync="modalAddShow" :close-on-backdrop="false">
        <CAlert v-if="waitAddTopic" color="info">{{ $t('message.wait') }}</CAlert>
        <CAlert v-if="modalAddErr" color="danger">{{ modalAddErr }}</CAlert>
        <div v-if="!waitAddTopic">
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
        </div>
        <template #footer>
          <CButton v-if="!waitAddTopic" type="submit" color="primary" class="m-2" style="width: 96px">
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

    <!-- pop-up form to edit existing topic -->
    <CForm @submit.prevent="doEditTopic" method="post">
      <CModal :title="$t('message.edit_topic')" :centered="true" :show.sync="modalEditShow" :close-on-backdrop="false">
        <CAlert v-if="waitEditTopic" color="info">{{ $t('message.wait') }}</CAlert>
        <CAlert v-if="modalEditErr" color="danger">{{ modalEditErr }}</CAlert>
        <div v-if="!waitEditTopic">
          <CInput
              type="text"
              v-model="formEdit.id"
              :label="$t('message.topic_id')"
              :placeholder="$t('message.topic_id_msg')"
              v-c-tooltip.hover="$t('message.topic_id_msg')"
              horizontal
              readonly
          />
          <CInput
              type="text"
              v-model="formEdit.icon"
              :label="$t('message.topic_icon')"
              :placeholder="$t('message.topic_icon_msg')"
              v-c-tooltip.hover="$t('message.topic_icon_msg')"
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
              v-model="formEdit.summary"
              :label="$t('message.topic_summary')"
              :placeholder="$t('message.topic_summary_msg')"
              v-c-tooltip.hover="$t('message.topic_summary_msg')"
              horizontal
              required
              was-validated
          />
        </div>
        <template #footer>
          <CButton v-if="!waitEditTopic" type="submit" color="primary" class="m-2" style="width: 96px">
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
import clientUtils from '@/utils/api_client'
import {iconize} from './utils'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'

const emptyForm = {id: '', icon: '', title: '', summary: ''}

export default {
  name: 'ProductTopicList',
  mounted() {
    this.loadProductTopicList(this.$route.params.pid)
  },
  computed:{
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
  data() {
    return {
      addMode: Boolean,

      modalAddShow: false,
      modalAddErr: '',
      formAdd: {...emptyForm},
      waitAddTopic: false,

      modalEditShow: false,
      modalEditErr: '',
      formEdit: {...emptyForm},
      waitEditTopic: false,

      modalIconsShow: false,

      modalDeleteShow: false,
      modalDeleteErr: '',
      topicToDelete: {},
      waitDeleteTopic: false,

      topicList: [],
      topicMap: {},

      myFlashMsg: this.flashMsg,
      errorMsg: '',
      foundStatus: -1,
    }
  },
  props: ["flashMsg"],
  methods: {
    _iconize(icon) {
      return iconize(icon)
    },
    loadProductTopicList(prodId) {
      this.foundStatus = -1
      const vue = this
      const apiUrl = clientUtils.apiAdminProductTopics.replaceAll(':product', prodId)
      clientUtils.apiDoGet(apiUrl,
          (apiRes) => {
            vue.foundStatus = apiRes.status == 200 ? 1 : 0
            if (vue.foundStatus == 1) {
              vue.topicList = apiRes.data
              vue.topicMap = {}
              for (let i = vue.topicList.length - 1; i >= 0; i--) {
                vue.topicMap[vue.topicList[i].id] = vue.topicList[i]
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
      this.$router.push({name: "ProductList"})
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
    clickAddTopic() {
      this.addMode = true
      this.formAdd = {...emptyForm}
      this.modalAddErr = ''
      this.modalAddShow = true
      this.myFlashMsg = ''
    },
    doAddTopic(e) {
      e.preventDefault()
      this.modalAddErr = ''
      let vue = this
      let data = {...vue.formAdd}
      vue.waitAddTopic = true
      let prodId = vue.$route.params.pid
      const apiUrl = clientUtils.apiAdminProductTopics.replaceAll(':product', prodId)
      clientUtils.apiDoPost(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalAddShow = false
              vue.myFlashMsg = vue.$i18n.t('message.topic_added_msg', {name: data.title})
              vue.loadProductTopicList(prodId)
            } else {
              vue.modalAddErr = apiRes.status + ": " + apiRes.message
            }
            vue.waitAddTopic = false
          },
          (err) => {
            vue.modalAddErr = err
            vue.waitAddTopic = false
          }
      )
    },
    clickTopicPages(topicId) {
      let prodId = this.$route.params.pid
      this.$router.push({name: "TopicPageList", params: {pid: prodId, tid: topicId}})
    },
    clickEditTopic(id) {
      this.addMode = false
      this.formEdit = {...this.topicMap[id]} //shallow clone using spread syntax, alternative way: this.formEdit = Object.assign({}, this.topicMap[id])
      this.modalEditErr = ''
      this.modalEditShow = true
      this.myFlashMsg = ''
    },
    doEditTopic(e) {
      e.preventDefault()
      this.modalEditErr = ''
      let vue = this
      let data = {...vue.formEdit}
      vue.waitEditTopic = true
      let prodId = vue.$route.params.pid
      const apiUrl = clientUtils.apiAdminProductTopic.replaceAll(':product', prodId).replaceAll(':topic', data.id)
      clientUtils.apiDoPut(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalEditShow = false
              vue.myFlashMsg = vue.$i18n.t('message.topic_updated_msg', {name: data.title})
              vue.loadProductTopicList(prodId)
            } else {
              vue.modalEditErr = apiRes.status + ": " + apiRes.message
            }
            vue.waitEditTopic = false
          },
          (err) => {
            vue.modalEditErr = err
            vue.waitEditTopic = false
          }
      )
    },
    clickDeleteTopic(id) {
      this.topicToDelete = this.topicMap[id]
      this.modalDeleteErr = ''
      this.modalDeleteShow = true
      this.myFlashMsg = ''
    },
    doDeleteTopic(e) {
      e.preventDefault()
      this.modalDeleteErr = ''
      let vue = this
      vue.waitDeleteTopic = true
      let prodId = vue.$route.params.pid
      let data = vue.topicToDelete
      const apiUrl = clientUtils.apiAdminProductTopic.replaceAll(':product', prodId).replaceAll(':topic', data.id)
      clientUtils.apiDoDelete(apiUrl,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.modalDeleteShow = false
              vue.myFlashMsg = vue.$i18n.t('message.topic_deleted_msg', {name: data.title})
              vue.loadProductTopicList(prodId)
            } else {
              vue.modalDeleteErr = apiRes.status + ": " + apiRes.message
            }
            vue.waitDeleteTopic = false
          },
          (err) => {
            vue.modalDeleteErr = err
            vue.waitDeleteTopic = false
          }
      )
    },
    _doMoveTopicUpOrDown(id, data) {
      const saveStatus = this.foundStatus
      this.foundStatus = -1
      let vue = this
      let prodId = vue.$route.params.pid
      const apiUrl = clientUtils.apiAdminProductTopic.replaceAll(':product', prodId).replaceAll(':topic', id)
      clientUtils.apiDoPatch(apiUrl, data,
          (apiRes) => {
            if (apiRes.status == 200) {
              vue.myFlashMsg = vue.$i18n.t('message.topic_updated_msg', {name: this.topicMap[id].title})
              vue.loadProductTopicList(prodId)
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
    doMoveTopicUp(id) {
      this._doMoveTopicUpOrDown(id, {action: "move_up"})
    },
    doMoveTopicDown(id) {
      this._doMoveTopicUpOrDown(id, {action: "move_down"})
    },
  }
}
</script>
