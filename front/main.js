import Vue from 'vue'
import axios from 'axios'
import App from './app.vue'

import User from './models/User'

window.User = new User

window.axios = require('axios')
window.axios.defaults.headers.common['X-Requested-With'] = "xmlhttprequest"

new Vue({
    components:{App},
    el: "#app"
})