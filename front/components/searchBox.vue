<template>
  <div>
    <h3>Recent YTS</h3>
    <form @submit.prevent="addTorrent()" class="col s12">
      <div class="card-content">
        <div class="row">
          <div class="input-field col">
            <i class="material-icons prefix">search</i>
            <input type="text" v-model="query.text" @keyup="searchTorrents(query.text)" id="searchYts">
            <label for="searchYts">Search YTS torrents</label>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import {db} from '../config/firebase';

export default {
  data() {
    return {
      query: {}
    };
  },
  methods: {
      searchTorrents (query) {
          console.log(query);
          
          db.collection('films').where('title_long', '>=', query).where('title_long', '<=', query).onSnapshot(snap => {
              console.log(snap);
              
          })
      }
  }
};
</script>

<style>
</style>
