<template>
  <div>
    <h4>Torrents</h4>
    <ul v-if="torrents.length" class="collection">
      <li style="margin-bottom:1em;" v-for="(torrent, i) in torrents" :key="i" class="">
        <div class="collection-item lighten-5 orange row" v-if="torrentDownloading(torrent)">
          <div class="progress orange lighten-3">
            <div class="determinate orange darken-2" :style="progressStyle(torrent.Percent)"></div>
          </div>
          <div class="col s6">
            <span><i class="material-icons">file_download</i> Downloading</span>
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
        </div>
        <div class="collection-item lighten-5 green row" v-if="torrentUploading(torrent)">
          <div class="progress green lighten-3">
            <div class="determinate green darken-2" :style="progressStyle(getUploadPercentage(torrent))"></div>
          </div>
          <div class="col s6">
            <span><i class="material-icons">file_upload</i> Uploading</span>
            <br>
              {{ torrent.Name }} : <strong>{{ getUploadPercentage(torrent) }}%</strong> <br>
              {{ torrent.UploadRate }}
          </div>
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
    getUploadPercentage(torrent) {
      let percentage = (torrent.UploadedTotal / torrent.Size) * 100
      return percentage ? (percentage).toFixed(2) : 0
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
        return true
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
    },
    syncTorrents(e) {
      console.log(e);
      this.torrents = e;
    }
  },
  computed: {

  },
  mounted() {
    window.sock.on("sync-torrent", this.syncTorrents);
  },
  beforeDestroy() {
    window.sock.off('sync-torrent', this.syncTorrents);
  }
};
</script>

<style>
/* li {
  margin-bottom: 1em;
} */
</style>
