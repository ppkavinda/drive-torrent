<template>
  <div>
    <h4>Torrents</h4>
    <ul v-if="torrents.length" class="collection">
      <li style="margin-bottom:1em;" v-for="(torrent, i) in torrents" :key="i" :class="'collection-item row lighten-5 ' + getStatus(torrent)">
        <div class="progress orange lighten-3">
          <div class="determinate orange darken-2" :style="progressStyle(torrent.Percent)"></div>
        </div>
        <div class="col s6">
          <span v-if="torrentDownloading(torrent)"><i class="material-icons">file_download</i> Downloading</span>
          <span v-if="torrentUploading(torrent)"><i class="material-icons">file_upload</i> Uploading</span>
          <br>
            {{ torrent.Name }} : <strong>{{ torrent.Percent }}%</strong> <br>
             {{ Number(torrent.DownloadRate / 1024 ).toFixed(2) }} KB/s 
        </div>
        <div class="right">
        <a v-if="!torrent.Started" @click="start(torrent.InfoHash)" class="btn btn-small  waves-effect waves-light green">
          Start <i class="material-icons right"> play_arrow</i>
        </a>
        <a v-if="torrent.Started" @click="stop(torrent.InfoHash)" class="btn waves-effect waves-light btn-small orange">
          Stop <i class="material-icons right"> stop</i>
        </a>
        <a @click="remove(torrent.InfoHash)" class="waves-effect waves-light btn red">
          <i class="material-icons right">delete</i>Remove
        </a>
        </div>
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
    getStatus(torrent) {
      if (torrent.Started && !torrent.Finished) {
        return 'orange'
      } else if (torrent.Started && torrent.Finished) {
        return 'green'
      }
    },
    torrentDownloading (torrent) {
      if (torrent.Started && !torrent.Finished) {
        return false
      }
    },
    torrentUploading (torrent) {
      if (torrent.Started && torrent.Finished) {
        return 'file_upload'
      }
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
/* li {
  margin-bottom: 1em;
} */
</style>
