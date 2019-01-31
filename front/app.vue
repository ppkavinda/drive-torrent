<template>
  <div v-cloak>
    <navbar :user="user"></navbar>
    <div class="container">
      <div class="row">
        <div v-if="user.isLoggedIn()">loggedin</div>

        <div v-else v-cloak class="valign-wrapper">
          <a href="/login" class="center-align waves-effect orange darken-1 waves-light btn-large">
            <i class="material-icons right">lock</i>Login with Google Drive
          </a>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script>
import User from "./models/User";
import navbar from "./components/navbar.vue";

export default {
  components: { navbar },
  props: ['initUser'],
  data() {
    return {
      user: this.initUser,
    };
  },
  created: function() {
    let user = JSON.parse(this.initUser)

    this.user = new User(
        user.ID,
        user.DisplayName,
        user.ImageURL,
        user.Email
    )
  }
};
</script>

<style>

</style>

