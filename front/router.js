import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/home.vue'
import Downloading from './views/downloading.vue'
import Movie from './views/movie.vue'
import NotFound from './views/notFound.vue';
import AdminRoutes from './views/admin/adminRouter'

Vue.use(Router)

const baseRoutes = [
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
    {
        path: '/movie/:id',
        name: 'movie',
        component: Movie
    },
    {
        path: '*',
        component: NotFound,
    }
]

const routes = baseRoutes.concat(AdminRoutes)

export default new Router({
    model: history,
    base: process.env.BASE_URL,
    routes:routes
})