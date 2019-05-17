<template>
<div class="admincontainer">
    <div id="main">
      <!-- START WRAPPER -->
      <div class="wrapper">
        <!-- START LEFT SIDEBAR NAV-->
        <admin-nav></admin-nav>
        <!-- END LEFT SIDEBAR NAV-->
        <!-- //////////////////////////////////////////////////////////////////////////// -->
        <!-- START CONTENT -->
        <router-view></router-view>
        <!-- END CONTENT -->
      </div>
      <!-- END WRAPPER -->
    </div>
</div>
</template>

<script>
import AdminNav from '../../components/adminNav.vue'

export default {
    components: {AdminNav},
    name: 'admin'
}
</script>
