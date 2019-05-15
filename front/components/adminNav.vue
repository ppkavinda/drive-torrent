<template>
    <aside id="left-sidebar-nav">
        <ul id="slide-out" class="side-nav fixed leftside-navigation">
            <li class="user-details cyan darken-2">
                <div class="row">
                <div class="col col s4 m4 l4">
                    <img src="/static/img/avatar-7.png" alt="" class="circle responsive-img valign profile-image cyan">
                </div>
                <div class="col col s8 m8 l8">
                    <ul id="profile-dropdown-nav" class="dropdown-content">
                    <li>
                        <a href="#" class="grey-text text-darken-1">
                        <i class="material-icons">face</i> Profile</a>
                    </li>
                    <li>
                        <a href="#" class="grey-text text-darken-1">
                        <i class="material-icons">settings</i> Settings</a>
                    </li>
                    <li>
                        <a href="#" class="grey-text text-darken-1">
                        <i class="material-icons">live_help</i> Help</a>
                    </li>
                    <li class="divider"></li>
                    <li>
                        <a href="#" class="grey-text text-darken-1">
                        <i class="material-icons">lock_outline</i> Lock</a>
                    </li>
                    <li>
                        <a href="#" class="grey-text text-darken-1">
                        <i class="material-icons">keyboard_tab</i> Logout</a>
                    </li>
                    </ul>
                    <a class="btn-flat dropdown-button waves-effect waves-light white-text profile-btn" href="#" data-activates="profile-dropdown-nav">John Doe<i class="mdi-navigation-arrow-drop-down right"></i></a>
                    <p class="user-roal">Administrator</p>
                </div>
                </div>
            </li>
            <li class="no-padding">
                <ul class="collapsible" data-collapsible="accordion">
                <li class="bold">
                    <a href="index.html" class="waves-effect waves-cyan">
                        <i class="material-icons">pie_chart_outlined</i>
                        <span class="nav-text">Dashboard</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="cards-basic.html" class="waves-effect waves-cyan">
                        <i class="material-icons">cast</i>
                        <span class="nav-text">Cards</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="ui-basic-buttons.html" class="waves-effect waves-cyan">
                        <i class="material-icons">insert_link</i>
                        <span class="nav-text">Buttons</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="form-layouts.html" class="waves-effect waves-cyan">
                        <i class="material-icons">format_color_text</i>
                        <span class="nav-text">Forms</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="css-typography.html" class="waves-effect waves-cyan">
                        <i class="material-icons">format_size</i>
                        <span class="nav-text">Typography</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="css-color.html" class="waves-effect waves-cyan">
                        <i class="material-icons">invert_colors</i>
                        <span class="nav-text">Color</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="table-basic.html" class="waves-effect waves-cyan">
                        <i class="material-icons">border_all</i>
                        <span class="nav-text">Table</span>
                    </a>
                </li>
                <li class="bold">
                    <a href="ui-icons.html" class="waves-effect waves-cyan">
                    <i class="material-icons">lightbulb_outline</i>
                    <span class="nav-text">Icons</span>
                    </a>
                </li>
                <li>
                    <a class="btn waves-effect waves-light gradient-45deg-red-pink" href="https://pixinvent.com/materialize-material-design-admin-template/landing/" target="_blank">
                    <i class="material-icons white-text">file_upload</i>Upgrade to Pro!
                    </a>
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
