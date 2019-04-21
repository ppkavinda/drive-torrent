<template>
  <div class="container">
    <div class="row">
      <div v-if="user.isLoggedIn()">
        <input-torrent></input-torrent>
        <torrents></torrents>
      </div>
      <search></search>
      <!-- <film-List></film-List> -->
    </div>
  </div>
</template>

<script>
import InputTorrent from '../components/inputTorrent'
import Torrents from '../components/torrents'
// import filmList from '../components/filmList'
import search from '../components/search'

export default {
  components: {InputTorrent, Torrents, search},
  data() {
    return {
      user: window.User
    };
  }
};
</script>

<style>
</style>
