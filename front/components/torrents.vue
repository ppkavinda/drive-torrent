<template>
  <div>
    <h4>Torrents</h4>
    <ul v-if="torrents.length" class="collection with-header">
      <li v-for="(torrent, i) in torrents" :key="i" class="collection-item row">
        <div class="progress orange lighten-3">
          <div class="determinate orange darken-2" :style="progressStyle(torrent.Percent)"></div>
        </div>
        <div class="col s6">
            {{ torrent.Name }} : {{ torrent.Percent }}% <br>
             {{ Number(torrent.DownloadRate / 1024 ).toFixed(2) }} KB/s 
        </div>
        <button @click="start(torrent.InfoHash)" class="btn btn-small secondary-content green col s2">
              Start <i class="material-icons"> delete_perminantly</i>
        </button>
        <button @click="stop(torrent.InfoHash)" class="btn btn-small secondary-content orange col s2">
              Stop <i class="material-icons"> delete_perminantly</i>
        </button>
        <button @click="remove(torrent.InfoHash)" class="btn btn-small secondary-content red col s2">
              Remove <i class="material-icons"> delete_perminantly</i>
        </button>
      </li>
    </ul>
    <span v-else>No Torrents yet!</span>
  </div>
</template>

<script>
export default {
  data() {
    return {
      msg: "",
      torrents: []
    };
  },
  methods: {
    send() {
        let msg = window.sock.send(this.msg);
    },
    start (hash) {
        axios.post('/torrent/start', {hash})
        .then(e => console.log(e))
    },
    stop(hash) {
        axios.post('/torrent/stop', {hash})
        .then(e => console.log(e))
    },
    remove(hash) {
        axios.post("/torrent/remove", {hash})
        .then(e => console.log(e))
    },
    progressStyle (width) {
        return {
            width: width + '%',
        }
    }
  },
  computed: {

  },
  mounted() {

    window.sock.on("sync-torrent", e => {
      console.log(e);
      this.torrents = e;
    });
  }
};
</script>

<style>
</style>
