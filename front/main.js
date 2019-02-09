import Vue from 'vue'
import App from './app.vue'
import router from './router'

import User from './models/User'
import Socket from './models/Socket'

window.User = new User
// window.sock = new Socket('ws://10.22.167.160:3000/sync');
window.sock = new Socket('ws://localhost:3000/sync');
// window.sock = new WebSocket('ws://localhost:3000/sync')


window.axios = require('axios')
window.axios.defaults.headers.common['X-Requested-With'] = "xmlhttprequest"

new Vue({
    components:{App},
    router,
    el: "#app"
})