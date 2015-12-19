var path = require('path');
var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    entry: [
        'webpack-hot-middleware/client',
        './source/index.js',
        './source/index.html'
    ],
    output: {
        path: path.join(__dirname, 'assets'),
        filename: 'bundle.js'
    },
    module: {
        loaders: [{
                test: /\.js$/,
                loaders: ['babel'],
                include: path.join(__dirname, 'source')
            }, {
                test: /\.less$/,
                loader: 'style-loader!css!less'
                // loader: ExtractTextPlugin.extract('style-loader', 'css!less')
            }, {
                test : /\.woff|\.woff2|\.svg|.eot|\.ttf/,
                loader : 'url-loader?&limit=1&name=/styles/[name]-[hash:6].[ext]'
            }, {
                test : /\index.html$/,
                loader : 'file-loader?name=/[name].[ext]'
            }
        ]
    },
    plugins: [
        // new ExtractTextPlugin("/styles/index.css"),
        // new webpack.optimize.UglifyJsPlugin({ minimize: true }),
        new webpack.HotModuleReplacementPlugin(),
        new webpack.NoErrorsPlugin()
    ],
    devtool: 'cheap-module-eval-source-map'
};
