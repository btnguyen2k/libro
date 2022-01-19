//#GovueAdmin-Customized
import Vue from 'vue'
import Router from 'vue-router'

const TheContainer = () => import('@/containers/TheContainer')
const Login = () => import('@/views/gva/pages/Login')
const Dashboard = () => import('@/views/libro/Dashboard')

const MyProfile = () => import('@/views/libro/MyProfile')

const ProductList = () => import('@/views/libro/ProductList')
const ProductTopicList = () => import('@/views/libro/ProductTopicList')
const TopicPageList = () => import('@/views/libro/TopicPageList')

const UserList = () => import('@/views/libro/UserList')

Vue.use(Router)

let router = new Router({
    mode: 'hash', // https://router.vuejs.org/api/#mode
    linkActiveClass: 'active',
    scrollBehavior: () => ({y: 0}),
    base: "/admin/",
    routes: configRoutes()
})

import appConfig from "@/utils/app_config"
import utils from "@/utils/app_utils"
import clientUtils from "@/utils/api_client"

router.beforeEach((to, from, next) => {
    if (!to.matched.some(record => record.meta.allowGuest)) {
        let session = utils.loadLoginSession()
        if (session == null) {
            //redirect to login page if not logged in
            return next({name: "Login", query: {returnUrl: router.resolve(to, from).href}})
        }
        let lastUserTokenCheck = utils.localStorageGetAsInt(utils.lskeyLoginSessionLastCheck)
        if (lastUserTokenCheck + 60 < utils.getUnixTimestamp()) {
            lastUserTokenCheck = utils.getUnixTimestamp()
            let token = session.token
            clientUtils.apiDoPost(clientUtils.apiVerifyLoginToken, {app: appConfig.APP_ID, token: token},
                (apiRes) => {
                    if (apiRes.status != 200) {
                        //redirect to login page if session verification failed
                        console.error("Session verification failed: " + JSON.stringify(apiRes))
                        return next({name: "Login", query: {returnUrl: router.resolve(to, from).href}})
                    } else {
                        utils.localStorageSet(utils.lskeyLoginSessionLastCheck, lastUserTokenCheck)
                        next()
                    }
                },
                (err) => {
                    console.error("Session verification error: " + err)
                    //redirect to login page if cannot verify session
                    return next({name: "Login", query: {returnUrl: router.resolve(to, from).href}})
                })
        } else {
            next()
        }
    } else {
        next()
    }
})

export default router

import i18n from '../i18n'

function configRoutes() {
    return [
        {
            path: '/',
            redirect: {name: "Dashboard"},
            name: 'Home', meta: {label: i18n.t('message.home')},
            component: TheContainer,
            children: [
                {
                    path: 'dashboard',
                    name: 'Dashboard',
                    meta: {label: i18n.t('message.dashboard')},
                    component: Dashboard
                },
                {
                    path: 'profile',
                    name: 'MyProfile',
                    meta: {label: i18n.t('message.my_profile')},
                    component: MyProfile
                },
                {
                    path: 'products', meta: {label: i18n.t('message.products')},
                    component: {
                        render(c) {
                            return c('router-view')
                        }
                    },
                    children: [
                        {
                            path: '',
                            meta: {label: i18n.t('message.products')},
                            name: 'ProductList',
                            component: ProductList,
                            props: true, //[props=true] to pass flashMsg
                        },
                        {
                            path: '_topics/:pid',
                            meta: {label: i18n.t('message.topics')},
                            name: 'ProductTopicList',
                            component: ProductTopicList,
                            props: true, //[props=true] to pass flashMsg
                        },
                        {
                            path: '_pages/:pid/:tid',
                            meta: {label: i18n.t('message.pages')},
                            name: 'TopicPageList',
                            component: TopicPageList,
                            props: true, //[props=true] to pass flashMsg
                        },
                    ]
                },
                {
                    path: 'users', meta: {label: i18n.t('message.users')},
                    name: 'UserList',
                    component: UserList,
                    props: true, //[props=true] to pass flashMsg
                    // component: {
                    //     render(c) {
                    //         return c('router-view')
                    //     }
                    // },
                    // children: [
                    //     {
                    //         path: '',
                    //         meta: {label: i18n.t('message.users')},
                    //         name: 'UserList',
                    //         component: ProductList,
                    //         props: true, //[props=true] to pass flashMsg
                    //     },
                    // ]
                },
            ]
        },
        {
            path: '/pages', redirect: {name: "Page404"}, name: 'Pages',
            component: {
                render(c) {
                    return c('router-view')
                }
            },
            meta: {
                allowGuest: true //do not required login to view
            },
            children: [
                {
                    path: 'login', name: 'Login', component: Login,
                    props: (route) => ({returnUrl: route.query.returnUrl}),
                    params: (route) => ({returnUrl: route.query.returnUrl}),
                },
            ]
        }
    ]
}
