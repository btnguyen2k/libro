/*
Client to make call to API server using Axios.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@since template-v0.1.0
@GovueAdmin
*/
import Axios from "axios"
import appConfig from "./app_config"
import utils from "@/utils/app_utils"
import router from "@/router"

const apiClient = Axios.create({
    baseURL: appConfig.APP_CONFIG.api_client.be_api_base_url,
    timeout: 10000,
});

const headerAppId = appConfig.APP_CONFIG.api_client.header_app_id
const headerAccessToken = appConfig.APP_CONFIG.api_client.header_access_token
let appId = appConfig.APP_ID

let apiInfo = "/info"
let apiLogin = "/api/login"
let apiVerifyLoginToken = "/api/verifyLoginToken"

let apiAdminStats = "/api/admin/stats"
let apiAdminProducts = "/api/admin/products"
let apiAdminProduct = "/api/admin/product"
let apiAdminDomains = "/api/admin/domains"
let apiAdminDomain = "/api/admin/domain"
let apiAdminProductTopics = "/api/admin/product/:product/topics"
let apiAdminProductTopic = "/api/admin/product/:product/topic/:topic"
let apiAdminTopicPages = "/api/admin/topic/:topic/pages"
let apiAdminTopicPage = "/api/admin/topic/:topic/page/:page"

let apiAdminUsers = "/api/admin/users"
let apiAdminUser = "/api/admin/user"

let apiSystemInfo = "/api/systemInfo"
let apiUser = "/api/user"
let apiUserPassword = "/api/userPassword"

function _apiOnSuccess(method, resp, apiUri, callbackSuccessful) {
    if (method == 'GET' && resp.hasOwnProperty("data") && resp.data.status == 403) {
        console.error("Error 403 from API [" + apiUri + "], redirecting to login page...")
        router.push({name: "Login", query: {app: appConfig.APP_ID, returnUrl: router.currentRoute.fullPath}})
        return
    }
    if (resp.hasOwnProperty("data") && resp.data.hasOwnProperty("extras") && resp.data.extras.hasOwnProperty("_access_token_")) {
        console.log("Update new access token from API [" + apiUri + "]")
        let jwt = utils.parseJwt(resp.data.extras._access_token_)
        utils.saveLoginSession({uid: jwt.payloadObj.uid, token: resp.data.extras._access_token_})
    }
    if (callbackSuccessful != null) {
        callbackSuccessful(resp.data)
    }
}

function _apiOnError(err, apiUri, callbackError) {
    console.error("Error calling api [" + apiUri + "]: " + err)
    if (callbackError != null) {
        callbackError(err)
    }
}

function apiDoGet(apiUri, callbackSuccessful, callbackError) {
    const session = utils.loadLoginSession()
    const headers = {}
    headers[headerAppId] = appId
    headers[headerAccessToken] = session != null ? session.token : ""
    return apiClient.get(apiUri, {
        headers: headers, cache: false
    }).then(res => _apiOnSuccess('GET', res, apiUri, callbackSuccessful)).catch(err => _apiOnError(err, apiUri, callbackError))
}

function apiDoPatch(apiUri, data, callbackSuccessful, callbackError) {
    const session = utils.loadLoginSession()
    const headers = {}
    headers[headerAppId] = appId
    headers[headerAccessToken] = session != null ? session.token : ""
    apiClient.patch(apiUri, data, {
        headers: headers, cache: false
    }).then(res => _apiOnSuccess('PATCH', res, apiUri, callbackSuccessful)).catch(err => _apiOnError(err, apiUri, callbackError))
}

function apiDoPost(apiUri, data, callbackSuccessful, callbackError) {
    const session = utils.loadLoginSession()
    const headers = {}
    headers[headerAppId] = appId
    headers[headerAccessToken] = session != null ? session.token : ""
    apiClient.post(apiUri, data, {
        headers: headers, cache: false
    }).then(res => _apiOnSuccess('POST', res, apiUri, callbackSuccessful)).catch(err => _apiOnError(err, apiUri, callbackError))
}

function apiDoPut(apiUri, data, callbackSuccessful, callbackError) {
    const session = utils.loadLoginSession()
    const headers = {}
    headers[headerAppId] = appId
    headers[headerAccessToken] = session != null ? session.token : ""
    apiClient.put(apiUri, data, {
        headers: headers, cache: false
    }).then(res => _apiOnSuccess('PUT', res, apiUri, callbackSuccessful)).catch(err => _apiOnError(err, apiUri, callbackError))
}

function apiDoDelete(apiUri, callbackSuccessful, callbackError) {
    const session = utils.loadLoginSession()
    const headers = {}
    headers[headerAppId] = appId
    headers[headerAccessToken] = session != null ? session.token : ""
    apiClient.delete(apiUri, {
        headers: headers, cache: false
    }).then(res => _apiOnSuccess('DELETE', res, apiUri, callbackSuccessful)).catch(err => _apiOnError(err, apiUri, callbackError))
}

export default {
    apiInfo,
    apiLogin,
    apiVerifyLoginToken,

    apiAdminStats,
    apiAdminProducts,
    apiAdminProduct,
    apiAdminDomains,
    apiAdminDomain,
    apiAdminProductTopics,
    apiAdminProductTopic,
    apiAdminTopicPages,
    apiAdminTopicPage,

    apiAdminUsers,
    apiAdminUser,

    apiSystemInfo,
    // apiGroupList,
    // apiGroup,
    // apiUserList,
    apiUser,
    apiUserPassword,

    apiDoGet,
    apiDoPatch,
    apiDoPost,
    apiDoPut,
    apiDoDelete,
}
