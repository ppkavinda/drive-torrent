<template>
  <div v-cloak>
    <navbar :user="user"></navbar>
    <router-view></router-view>
  </div>
</template>

<script>
import User from "./models/User";
import navbar from "./components/navbar.vue";

export default {
  components: { navbar },
  props: ["initUser"],
  data() {
    return {
      user: this.initUser
    };
  },
  created: function() {
    let user = JSON.parse(this.initUser);

    this.user = new User(user.ID, user.DisplayName, user.ImageURL, user.Email);
    window.User = this.user;
    window.sock.on("sync-login", res => {
      // Vue.set(window.User, ID, res.ID);
      window.User.ID = res.ID
      console.log(res)
    });
  }
};
</script>

<style>
</style>

