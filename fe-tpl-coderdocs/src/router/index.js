import Vue from 'vue'
import Router from 'vue-router'

const Error = () => import('@/components/Error')
const Home = () => import('@/components/Home')
const Topic = () => import('@/components/Topic')
const Page = () => import('@/components/Page')

Vue.use(Router)

let router = new Router({
    mode: 'history', // https://router.vuejs.org/api/#mode
    linkActiveClass: 'active',
    scrollBehavior: () => ({y: 0}),
    base: "/doc/",
    routes: configRoutes()
})

export default router

function configRoutes() {
    return [
        {
          path: '/_error', name: 'Error', component: Error, props: true,
        },
        {
            path: '/',
            component: {
                render(c) {
                    return c('router-view')
                }
            },
            children: [
                {
                    path: '', name: 'Home',
                    component: Home
                },
                {
                    path: '/:tid', name: 'Topic',
                    component: Topic
                },
                {
                    path: '/:tid/:pid', name: 'Page',
                    component: Page
                }
            ]
        },
    ]
}
