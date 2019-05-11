import Vue from "vue";
import Router from "vue-router";

let NProgress = require("nprogress");

import "nprogress/nprogress.css";

Vue.use(Router);
const router = new Router({
    base: "/",
    mode: "hash",
    linkActiveClass: "active",
    routes: [
        {
            path: "*",
            meta: {
                public: true
            },
            redirect: {
                path: "/404"
            }
        },
        {
            path: "/404",
            meta: {
                public: true
            },
            name: "NotFound",
            component: () => import(`@/views/NotFound.vue`)
        },
        {
            path: "/403",
            meta: {
                public: true
            },
            name: "AccessDenied",
            component: () => import(  `@/views/Deny.vue`)
        },
        {
            path: "/500",
            meta: {
                public: true
            },
            name: "ServerError",
            component: () => import(`@/views/Error.vue`)
        },
        {
            path: "/",
            meta: {},
            name: "Root",
            redirect: {
                name: "registry"
            }
        },
        {
            path: "/call",
            meta: {breadcrumb: true},
            name: "call",
            component: () => import(`@/views/call/Call.vue`)
        },
        {
            path: "/cli",
            meta: {breadcrumb: true},
            name: "cli",
            component: () => import(`@/views/cli/Cli.vue`)
        },
        {
            path: "/home",
            meta: {breadcrumb: true},
            name: "home",
            component: () => import(`@/views/home/Home.vue`)
        },
        {
            path: "/registry",
            meta: {breadcrumb: true},
            name: "registry",
            component: () => import(`@/views/registry/RegistryPage.vue`)
        },
        {
            path: "/stats/api",
            meta: {breadcrumb: true},
            name: "apiStatistics",
            component: () => import(`@/views/statistics/APIStatistics.vue`),
        },
        {
            path: "/stats/service",
            meta: {breadcrumb: true},
            name: "serviceStatistics",
            component: () => import(`@/views/statistics/ServiceStatistics.vue`),
        }
    ]
});
// router gards
router.beforeEach((to, from, next) => {
    NProgress.start();
    next();
});

router.afterEach((to, from) => {
    // ...
    NProgress.done();
});

export default router;
