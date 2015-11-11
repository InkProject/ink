var ExtractTextPlugin = require('extract-text-webpack-plugin');
var webpack = require('webpack');

module.exports = {
    entry: './source/js/index.coffee',
    output: {
        path: __dirname + '/bundle/',
        filename: 'index.js'
    },
    module: {
        loaders:[
            { test: /\.coffee$/, loader: 'coffee-loader' },
            { test: /\.less$/,  loader: ExtractTextPlugin.extract('style', 'css!less') },
            { test : /\.woff|\.woff2|\.svg|.eot|\.ttf/, loader : 'url-loader?limit=8192' }
        ]
    },
    plugins: [
        new ExtractTextPlugin('index.css'),
        new webpack.optimize.UglifyJsPlugin({ minimize: true })
    ],
    watch: true
};
