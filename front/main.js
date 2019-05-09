import Vue from 'vue'
import App from './app.vue'
import router from './router'
import VueFirestore from 'vue-firestore'
import InstantSearch from 'vue-instantsearch'

import User from './models/User'
import Socket from './models/Socket'

import VueAnalytics from 'vue-analytics'

Vue.config.productionTip = false;
Vue.use(VueFirestore)
Vue.use(InstantSearch)
Vue.use(VueAnalytics, {
    id: 'UA-116031370-3',
    router
})

window.User = new User

var loc = window.location, wsUri;
loc.protocol === "https:" ? wsUri = "wss:": wsUri = "ws:"
wsUri += "//" + loc.host;
wsUri += loc.pathname + "sync";
window.sock = new Socket(wsUri);


window.axios = require('axios')
window.axios.defaults.headers.common['X-Requested-With'] = "xmlhttprequest"

new Vue({
    components:{App},
    router,
    el: "#app"
})