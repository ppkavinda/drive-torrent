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
  }
};
</script>

<style>
</style>

