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
        <section id="content">
          <!--start container-->
          <div class="admincontainer">
            <!--card stats start-->
            <div id="card-stats">
              <div class="row mt-1">
                <div class="col s12 m6 l3">
                  <div class="card gradient-45deg-light-blue-cyan gradient-shadow min-height-100 white-text">
                    <div class="padding-4">
                      <div class="col s7 m7">
                        <i class="material-icons background-round mt-5">add_shopping_cart</i>
                        <p>Orders</p>
                      </div>
                      <div class="col s5 m5 right-align">
                        <h5 class="mb-0">690</h5>
                        <p class="no-margin">New</p>
                        <p>6,00,00</p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col s12 m6 l3">
                  <div class="card gradient-45deg-red-pink gradient-shadow min-height-100 white-text">
                    <div class="padding-4">
                      <div class="col s7 m7">
                        <i class="material-icons background-round mt-5">perm_identity</i>
                        <p>Clients</p>
                      </div>
                      <div class="col s5 m5 right-align">
                        <h5 class="mb-0">1885</h5>
                        <p class="no-margin">New</p>
                        <p>1,12,900</p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col s12 m6 l3">
                  <div class="card gradient-45deg-amber-amber gradient-shadow min-height-100 white-text">
                    <div class="padding-4">
                      <div class="col s7 m7">
                        <i class="material-icons background-round mt-5">timeline</i>
                        <p>Sales</p>
                      </div>
                      <div class="col s5 m5 right-align">
                        <h5 class="mb-0">80%</h5>
                        <p class="no-margin">Growth</p>
                        <p>3,42,230</p>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col s12 m6 l3">
                  <div class="card gradient-45deg-green-teal gradient-shadow min-height-100 white-text">
                    <div class="padding-4">
                      <div class="col s7 m7">
                        <i class="material-icons background-round mt-5">attach_money</i>
                        <p>Profit</p>
                      </div>
                      <div class="col s5 m5 right-align">
                        <h5 class="mb-0">$890</h5>
                        <p class="no-margin">Today</p>
                        <p>$25,000</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div id="card-stats">
             <div class="row">
               <div class="col s12 m6 l3">
                 <div class="card">
                   <div class="card-content cyan white-text">
                     <p class="card-stats-title">
                       <i class="material-icons">person_outline</i> New Clients</p>
                     <h4 class="card-stats-number">566</h4>
                     <p class="card-stats-compare">
                       <i class="material-icons">keyboard_arrow_up</i> 15%
                       <span class="cyan text text-lighten-5">from yesterday</span>
                     </p>
                   </div>
                 </div>
               </div>
               <div class="col s12 m6 l3">
                 <div class="card">
                   <div class="card-content red accent-2 white-text">
                     <p class="card-stats-title">
                       <i class="material-icons">attach_money</i>Total Sales</p>
                     <h4 class="card-stats-number">$8990.63</h4>
                     <p class="card-stats-compare">
                       <i class="material-icons">keyboard_arrow_up</i> 70%
                       <span class="red-text text-lighten-5">last month</span>
                     </p>
                   </div>
                 </div>
               </div>
               <div class="col s12 m6 l3">
                 <div class="card">
                   <div class="card-content teal accent-4 white-text">
                     <p class="card-stats-title">
                       <i class="material-icons">trending_up</i> Today Profit</p>
                     <h4 class="card-stats-number">$806.52</h4>
                     <p class="card-stats-compare">
                       <i class="material-icons">keyboard_arrow_up</i> 80%
                       <span class="teal-text text-lighten-5">from yesterday</span>
                     </p>
                   </div>
                    
                 </div>
               </div>
               <div class="col s12 m6 l3">
                 <div class="card">
                   <div class="card-content deep-orange accent-2 white-text">
                     <p class="card-stats-title">
                       <i class="material-icons">content_copy</i> New Invoice</p>
                     <h4 class="card-stats-number">1806</h4>
                     <p class="card-stats-compare">
                       <i class="material-icons">keyboard_arrow_down</i> 3%
                       <span class="deep-orange-text text-lighten-5">from last month</span>
                     </p>
                   </div>
                 </div>
               </div>
             </div>
            </div>
            <!--card stats end-->
          </div>
          <!--end container-->
        </section>
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
    data(){
        return{}
    },
    methods:{},
    mounted(){}
}
</script>

<style>
    @import url('/static/css/materialize.css');
    @import url('/static/css/style.css');
    @import url('/static/css/custom.css');
    @import url('/static/css/flag-icon.min.css');
    @import url('/static/css/perfect-scrollbar.css');
</style>
