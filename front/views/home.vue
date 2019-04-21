<template>
  <div class="container">
    <div class="row">
        <input-torrent></input-torrent>
<<<<<<< HEAD
        <torrents></torrents>
      </div>
      <search></search>
      <!-- <film-List></film-List> -->
=======
        <search-box></search-box>
      <film-List></film-List>
>>>>>>> d8a6dbba662b565c9d45529fa7edd5767fdc10b8
    </div>
  </div>
</template>

<script>
<<<<<<< HEAD
import InputTorrent from '../components/inputTorrent'
import Torrents from '../components/torrents'
// import filmList from '../components/filmList'
import search from '../components/search'

export default {
  components: {InputTorrent, Torrents, search},
=======
import InputTorrent from "../components/inputTorrent";
import Torrents from "../components/torrents";
import filmList from "../components/filmList";
import SearchBox from '../components/searchBox';

export default {
  components: { InputTorrent, Torrents, filmList, SearchBox },
>>>>>>> d8a6dbba662b565c9d45529fa7edd5767fdc10b8
  data() {
    return {
      user: window.User
    };
  }
};
</script>

<style>
</style>
