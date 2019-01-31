<template>
  <nav>
    <div class="nav-wrapper orange darken-1">
      <div class="container">
        <a href="#" class="brand-logo">Drive-Torrent</a>
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
</template>

<script>
export default {
    props: ['user']

}
</script>

<style scoped>
nav {
  margin-bottom: 2em;
}
</style>
