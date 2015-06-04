var path = require('path'),
    gulp = require('gulp'),
    connect = require('gulp-connect'),
    clean = require('gulp-clean'),
    coffee = require('gulp-coffee'),
    less = require('gulp-less');

var compile = function(file, type) {
    var method = coffee;
    var outpath = 'js';
    if (type === 'less') {
        method = less;
        outpath = 'css';
    }
    gulp.src(file)
        .pipe(method().on('error', function(err) {
            console.error(err.stack);
        }))
        .pipe(gulp.dest(outpath))
        .pipe(connect.reload());
};

gulp.task('watch', function() {
    gulp.watch(['coffee/**/*.coffee', 'less/**/*.less'], function(data) {
        console.info(data.type + ': ' + data.path);
        compile(data.path, data.type);
    });
});

gulp.task('default', function() {
    compile('coffee/**/*.coffee', 'coffee');
    compile('less/**/*.less', 'less');
});
