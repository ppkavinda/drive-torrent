<template>
  <div class="container">
    <div class="row">
        <input-torrent></input-torrent>
        <search-box></search-box>
      <film-List></film-List>
    </div>
  </div>
</template>

<script>
import InputTorrent from "../components/inputTorrent";
import Torrents from "../components/torrents";
import filmList from "../components/filmList";
import SearchBox from '../components/searchBox';

export default {
  components: { InputTorrent, Torrents, filmList, SearchBox },
  data() {
    return {
      user: window.User
    };
  }
};
</script>

<style>
</style>
