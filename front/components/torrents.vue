<template>
  <div>
    <h4>Torrents</h4>
  </div>
</template>

<script>
export default {
  data() {
    return {
      msg: "",
      torrents: [],
    //   wsuri: "ws://localhost:3000/sync",
    //   sock: new WebSocket('ws://localhost:3000/sync')
    };
  },
  methods: {
    send() {
      let msg = window.sock.send(this.msg);
    }
  },
  mounted() {

    window.sock.onopen = function() {
      console.log("connected to " + this.wsuri);
    };
    window.sock.onclose = function(e) {
      console.log("connection closed " + e.data);
    };
    window.sock.onmessage = function(e) {
    //   console.log(JSON.stringify(e.data));
      console.log(e.data);
    };
  }
};
</script>

<style>
</style>
