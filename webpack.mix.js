let mix = require('laravel-mix');


mix.disableSuccessNotifications();

mix.js('front/main.js', 'static/js/app.js')
// 	.extract(['vue'])
//    .sass('resources/assets/sass/app.scss', 'public/css');
// mix.browserSync('localhost:8001');mix.js('./app.js', './build/')