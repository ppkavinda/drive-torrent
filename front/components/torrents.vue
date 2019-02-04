<template>
  <div>
    <h4>Torrents</h4>
    <ul class="collection with-header">
      <li v-for="(torrent, i) in torrents" :key="i" class="collection-item">
        <div>{{ torrent.Name }} : {{ torrent.Percent }}%
          <a @click="remove(torrent.InfoHash)" class="secondary-content red-text">
              Remove
            <i class="material-icons"> delete_perminantly</i>
          </a>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      msg: "",
      torrents: []
      //   wsuri: "ws://localhost:3000/sync",
      //   sock: new WebSocket('ws://localhost:3000/sync')
    };
  },
  methods: {
    send() {
      let msg = window.sock.send(this.msg);
    }
  },
  methods: {
      remove(hash) {
          axios.post("/remove", {hash})
          .then(e => console.log(e))
      },
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
