import Vue from 'vue'
import Router from 'vue-router'

const Home = () => import('@/components/Home')

Vue.use(Router)

let router = new Router({
    mode: 'history', // https://router.vuejs.org/api/#mode
    linkActiveClass: 'active',
    scrollBehavior: () => ({y: 0}),
    base: "/doc/",
    routes: configRoutes()
})

// import appConfig from "@/utils/app_config"
// import utils from "@/utils/app_utils"
// import clientUtils from "@/utils/api_client"

// router.beforeEach((to, from, next) => {
//     if (!to.matched.some(record => record.meta.allowGuest)) {
//         let session = utils.loadLoginSession()
//         if (session == null) {
//             //redirect to login page if not logged in
//             return next({name: "Login", query: {returnUrl: router.resolve(to, from).href}})
//         }
//         let lastUserTokenCheck = utils.localStorageGetAsInt(utils.lskeyLoginSessionLastCheck)
//         if (lastUserTokenCheck + 60 < utils.getUnixTimestamp()) {
//             lastUserTokenCheck = utils.getUnixTimestamp()
//             let token = session.token
//             clientUtils.apiDoPost(clientUtils.apiVerifyLoginToken, {app: appConfig.APP_ID, token: token},
//                 (apiRes) => {
//                     if (apiRes.status != 200) {
//                         //redirect to login page if session verification failed
//                         console.error("Session verification failed: " + JSON.stringify(apiRes))
//                         return next({name: "Login", query: {returnUrl: router.resolve(to, from).href}})
//                     } else {
//                         utils.localStorageSet(utils.lskeyLoginSessionLastCheck, lastUserTokenCheck)
//                         next()
//                     }
//                 },
//                 (err) => {
//                     console.error("Session verification error: " + err)
//                     //redirect to login page if cannot verify session
//                     return next({name: "Login", query: {returnUrl: router.resolve(to, from).href}})
//                 })
//         } else {
//             next()
//         }
//     } else {
//         next()
//     }
// })

export default router

import i18n from '../i18n'

function configRoutes() {
    return [
        {
            path: '/',
            component: {
                render(c) {
                    return c('router-view')
                }
            },
            children: [
                {
                    path: '', name: 'Home', meta: {label: i18n.t('message.products')},
                    component: Home
                }

                // {
                //     path: 'dashboard',
                //     name: 'Dashboard',
                //     meta: {label: i18n.t('message.dashboard')},
                //     component: Dashboard
                // },
                // {
                //     path: 'products', meta: {label: i18n.t('message.products')},
                //     component: {
                //         render(c) {
                //             return c('router-view')
                //         }
                //     },
                //     children: [
                //         {
                //             path: '',
                //             meta: {label: i18n.t('message.products')},
                //             name: 'ProductList',
                //             component: ProductList,
                //             props: true, //[props=true] to pass flashMsg
                //         },
                //         {
                //             path: '_add',
                //             meta: {label: i18n.t('message.add_product')},
                //             name: 'AddProduct',
                //             component: AddProduct,
                //         },
                //         {
                //             path: '_edit/:id',
                //             meta: {label: i18n.t('message.edit_product')},
                //             name: 'EditProduct',
                //             component: EditProduct,
                //         },
                //         {
                //             path: '_delete/:id',
                //             meta: {label: i18n.t('message.delete_product')},
                //             name: 'DeleteProduct',
                //             component: DeleteProduct,
                //         },
                //         {
                //             path: '_topics/:pid',
                //             meta: {label: i18n.t('message.topics')},
                //             name: 'ProductTopicList',
                //             component: ProductTopicList,
                //         },
                //         {
                //             path: '_pages/:pid/:tid',
                //             meta: {label: i18n.t('message.pages')},
                //             name: 'TopicPageList',
                //             component: TopicPageList,
                //         },
                //     ]
                // },
            ]
        },
    ]
}
