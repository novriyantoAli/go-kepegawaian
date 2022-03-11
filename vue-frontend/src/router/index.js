/* eslint-disable */
import Vue from 'vue';
import Router from 'vue-router';

import { authenticationService } from '@/_services';
import DashboardPage from '@/components/dashboard/DashboardPage';
import LoginPage from '@/components/login/LoginPage';

Vue.use(Router);

export const router = new Router({
    mode: 'history',
    routes: [
        { 
            path: '/', 
            name: 'dashboard',
            component: DashboardPage, 
            meta: { authorize: [] } 
        },
        { 
            path: '/login', 
            name: 'login',
            component: LoginPage 
        },
    ]
});

router.beforeEach((to, from, next) => {
    // redirect to login page if not logged in and trying to access a restricted page
    const { authorize } = to.meta;
    const currentUser = authenticationService.currentUserValue;

    if (authorize) {
        if (!currentUser) {
            // not logged in so redirect to login page with the return url
            return next({ path: '/login', query: { returnUrl: to.path } });
        }

        // check if route is restricted by role
        if (authorize.length && !authorize.includes(currentUser.level)) {
            // role not authorised so redirect to home page
            return next({ path: '/' });
        }
    }

    next();
})

export default router;