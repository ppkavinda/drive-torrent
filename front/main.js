import Vue from 'vue'
import App from './app.vue'
import router from './router'

import User from './models/User'

window.User = new User

window.axios = require('axios')
window.axios.defaults.headers.common['X-Requested-With'] = "xmlhttprequest"

new Vue({
    components:{App},
    router,
    el: "#app"
})