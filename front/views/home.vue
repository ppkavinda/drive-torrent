<template>
  <div class="container">
    <div class="row">
      <div v-if="user.isLoggedIn()">
        <input-torrent></input-torrent>
        <torrents></torrents>
      </div>

      <div v-else v-cloak class="valign-wrapper">
        <a href="/login" class="center-align waves-effect orange darken-1 waves-light btn-large">
          <i class="material-icons right">lock</i>Login with Google Drive
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import InputTorrent from '../components/inputTorrent'
import Torrents from '../components/torrents'

export default {
  components: {InputTorrent, Torrents},
  data() {
    return {
      user: window.User
    };
  }
};
</script>

<style>
</style>
