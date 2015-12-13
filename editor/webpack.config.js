var path = require('path');
var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    entry: [
        'webpack-hot-middleware/client',
        './source/app.js'
    ],
    output: {
      path: path.join(__dirname, 'dist'),
      filename: 'bundle.js',
      publicPath: '/static/'
    },
    module: {
        loaders: [{
                test: /\.js$/,
                loaders: ['babel'],
                include: path.join(__dirname, 'source')
            },{
                test: /\.less$/,
                loader: "style!css!less"
            },
            {test : /\.woff|\.woff2|\.svg|.eot|\.ttf/, loader : 'url-loader?limit=8192'}
        ]
    },
    plugins: [
        // new webpack.optimize.UglifyJsPlugin({ minimize: true }),
        new webpack.HotModuleReplacementPlugin(),
        new webpack.NoErrorsPlugin()
    ],
    devtool: 'cheap-module-eval-source-map'
};
