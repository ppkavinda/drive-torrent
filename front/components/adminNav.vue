<template>
    <aside id="left-sidebar-nav">
        <ul id="slide-out" class="side-nav fixed leftside-navigation">
            <li class="user-details cyan darken-2">
                <div class="row">
                    <div class="col col s4 m4 l4">
                        <img src="/static/img/avatar-7.png" alt="" class="circle responsive-img valign profile-image cyan">
                    </div>
                </div>
            </li>
            <li class="no-padding">
                <ul class="collapsible" data-collapsible="accordion">
                    <li class="bold">
                        <router-link :to="'/admin/dashboard'" class="waves-effect waves-cyan">
                            <i class="material-icons">pie_chart_outlined</i>
                            <span class="nav-text">Dashboard</span>
                        </router-link>
                    </li>
                    <li class="bold">
                        <router-link :to="'/admin/files'" class="waves-effect waves-cyan">
                            <i class="material-icons">cast</i>
                            <span class="nav-text">Files</span>
                        </router-link>
                    </li>
                    <li class="bold">
                        <router-link :to="'/admin/library'" class="waves-effect waves-cyan">
                            <i class="material-icons">format_color_text</i>
                            <span class="nav-text">Links Library</span>
                        </router-link>
                    </li>
                    <li class="bold">
                        <router-link :to="'/admin/users'" class="waves-effect waves-cyan">
                            <i class="material-icons">insert_link</i>
                            <span class="nav-text">Users</span>
                        </router-link>
                    </li>
                </ul>
            </li>
        </ul>
        <a href="#" data-activates="slide-out" class="sidebar-collapse btn-floating btn-medium waves-effect waves-light hide-on-large-only">
        <i class="material-icons">menu</i>
        </a>
    </aside>
</template>

<script>
export default {
    
}
</script>
