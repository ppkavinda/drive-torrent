<template>
<div>
  <nav>
    <div class="nav-wrapper orange darken-1">
      <div class="container">
        <router-link class="brand-logo" to="/">Drive&nbsp;Torrent</router-link>
        <a href="" data-target="mobile-demo" class="sidenav-trigger"><i class="material-icons">menu</i></a>
        <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li v-if="user.isLoggedIn()">
            <a href="/logout">Logout</a>
          </li>
          <li v-else>
            <a href="/login">Login</a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
  <ul class="sidenav orange darken-1" id="mobile-demo">
    <li v-if="user.isLoggedIn()">
      <a class="white-text" href="/logout">Logout</a>
    </li>
    <li v-else>
      <a class="white-text" href="/login">Login</a>
    </li>
  </ul>  
</div>
</template>

<script>
export default {
  props: ["user"],
  mounted() {
    // to toggle mobile nav
    var elems = document.querySelectorAll('.sidenav');
    var instances = M.Sidenav.init(elems);
  }
};
</script>

<style scoped>
nav {
  margin-bottom: 2em;
}
</style>
