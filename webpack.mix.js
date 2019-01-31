let mix = require('laravel-mix');


mix.js('front/app.js', 'static/js/')
// 	.extract(['vue'])
//    .sass('resources/assets/sass/app.scss', 'public/css');
// mix.browserSync('localhost:8001');mix.js('./app.js', './build/')