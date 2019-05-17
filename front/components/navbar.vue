<template>
<div class="wrapper-nav" :class="{ 'navbar-fixed' : settings.fixed }">
  <nav>
    <div class="nav-wrapper orange darken-1">
      <div class="container">
        <router-link class="brand-logo" to="/">Drive&nbsp;Torrent</router-link>
        <a href="" data-target="mobile-demo" class="sidenav-trigger"><i class="material-icons">menu</i></a>
        <ul id="nav-mobile" class="right hide-on-med-and-down">
          <div v-if="user.isLoggedIn()">
            <li>
              <router-link to="/downloading">Downloading</router-link>
            </li>
            <li>
              <a href="/logout">Logout</a>
            </li>
          </div>
          <li v-else>
            <a href="/login">Login</a>
          </li>
        </ul>
      </div>
    </div> 
     <!-- <div class="switch right">
        <label>
          <input v-model="settings.fixed" type="checkbox">
          Non Sticky
          <span class="lever"></span>
          Sticky
        </label>
      </div> -->
  </nav>
  <ul class="sidenav orange darken-1" id="mobile-demo">
      <li>
        <router-link class="white-text" to="/">Home</router-link>
      </li>
    <div v-if="user.isLoggedIn()">
      <li>
        <router-link class="white-text" to="downloading">Downloading</router-link>
      </li>
      <li>
        <a class="white-text" href="/logout">Logout</a>
      </li>
    </div>
    <li v-else>
      <a class="white-text" href="/login">Login</a>
    </li>
  </ul>  
</div>
</template>

<script>
export default {
  props: ["user"],
  data() {
    return {
      settings : {
        fixed: true
        }
    }
  },
  methods: {
    toggleFixed() {
      this.settings.fixed ? this.settings.fixed = 'navbar-fixed' : this.settings.fixed = ''
    }
  },
  mounted() {
    // to toggle mobile nav
    var elems = document.querySelectorAll('.sidenav');
    var instances = M.Sidenav.init(elems);

  }
};
</script>

<style scoped>
.wrapper-nav {
  margin-bottom: 2em;
}
</style>
