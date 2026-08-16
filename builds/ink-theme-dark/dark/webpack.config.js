var webpack = require('webpack')
var ExtractTextPlugin = require('extract-text-webpack-plugin')
var autoprefixer = require('autoprefixer')
var precss = require('precss')
var cssimport = require('postcss-import')

module.exports = {
  entry: './source/js/index.js',
  output: {
    path: __dirname + '/bundle/',
    filename: 'index.js'
  },
  postcss: function (webpack) {
    return [cssimport({
      addDependencyTo: webpack
    }), autoprefixer, precss]
  },
  module: {
    loaders:[
      { test: /\.css/, loader: ExtractTextPlugin.extract('style-loader', 'css!postcss'), include: __dirname },
      { test : /\.woff|\.woff2|\.svg|.eot|\.ttf/, loader : 'url-loader?limit=8192' }
    ]
  },
  plugins: [
    new webpack.optimize.UglifyJsPlugin({ minimize: true }),
    new ExtractTextPlugin('index.css')
  ],
  watch: false
}
