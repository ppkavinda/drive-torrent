<template>
  <div class="container">
    <div class="row">
      <div v-if="user.isLoggedIn()">
        <input-torrent></input-torrent>
        <torrents></torrents>
      </div>
      <film-List></film-List>
    </div>
  </div>
</template>

<script>
import InputTorrent from '../components/inputTorrent'
import Torrents from '../components/torrents'
import filmList from '../components/filmList'

export default {
  components: {InputTorrent, Torrents,filmList},
  data() {
    return {
      user: window.User
    };
  }
};
</script>

<style>
</style>
