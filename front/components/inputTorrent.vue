<template>
  <div class="row">
      <form @submit.prevent="addTorrent()" class="card col s12">
        <div class="card-content">
          <div class="row">
            <div class="input-field col s12">
              <i class="material-icons prefix">get_app</i>
              <input
                type="text"
                v-model="torrent.url"
                @input="validateURL(torrent.url)"
                id="inputTorrent"
              >
              <label for="inputTorrent">Enter the magnet or URL of a torrent file</label>
              <span
                v-if="!torrent.valid"
                :class="'helper-text ' + error.color + '-text'"
                data-error="wrong"
                data-success="right"
              >{{ error.msg }}</span>
            </div>
          </div>
        </div>
      </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      torrent: {
        url: "",
        valid: true,
        type: ""
      },
      error: {
        msg: "",
        color: ""
      }
    };
  },
  methods: {
    getType() {
      if (
        this.torrent.url.match(/magnet:\?xt=urn:[a-z0-9]+:[a-z0-9]{32}/i) !==
        null
      ) {
        this.torrent.type = "magnet";
      } else {
        this.torrent.type = "url";
      }
    },
    validateURL(url) {
      this.clearError();
      try {
        new URL(url);
        this.torrent.valid = true;
        this.getType();
      } catch (_) {
        this.torrent.valid = false;
        this.error.msg = "Invalid URL";
        this.error.color = "red";
        this.torrent.type = "";
      }
    },
    clearError() {
      this.error = {
        msg: "",
        color: ""
      };
    },
    addTorrent() {
      if (!this.torrent.valid) return;

      if (this.torrent.type === "magnet") {
        let magnet = this.torrent.url;
        axios
          .post("/new/magnet", { magnet })
          .then(res => {
            this.error = { msg: "", color: "" };
            this.torrent = { url: "", valid: true, type: "" };
            this.$router.push({name: 'downloading'})
          })
          .catch(err => {
            this.error.msg = err.response.data.Message;
            this.error.color = "red";
            this.torrent.valid = false;
            if (err.response.status == 401) window.location.replace('/login')

          });
      } else if (this.torrent.type === "url") {
        let url = this.torrent.url;
        axios
          .post("/new/url", { url })
          .then(res => {
            this.error = { msg: "", color: "" };
            this.torrent = { url: "", valid: true, type: "" };
            this.$router.push({name: 'downloading'})

          })
          .catch(err => {
            this.error.msg = err.response.data.Message;
            this.error.color = "red";
            this.torrent.valid = false;
            if (err.response.status == 401) window.location.replace('/login')

          });
      }
    }
  }
};
</script>

<style>
</style>
