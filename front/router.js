import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/home.vue'
import Status from './views/status.vue'
import Downloading from './views/downloading.vue'
import Movie from './views/movie.vue'
import NotFound from './views/notFound.vue';
import Dashboard from './views/admin/dashboard.vue';

Vue.use(Router)

export default new Router({
    model: history,
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
        },
        {
            path: '/downloading',
            name: 'downloading',
            component: Downloading,
        },
        // {
        //     path: '/status',
        //     name: 'status',
        //     component: Status,
        // },
        {
            path: '/movie/:id',
            name: 'movie',
            component: Movie
        },
        //Admin Routes
        {
            path: '/admin/dashboard',
            name: 'dashboard',
            component: Dashboard
        },
        {
            path: '*',
            component: NotFound,
        }
    ]
})