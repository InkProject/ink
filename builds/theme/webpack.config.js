var MiniCssExtractPlugin = require("mini-css-extract-plugin")
var UglifyJsPlugin = require('uglifyjs-webpack-plugin')
var CssMinimizerPlugin = require('css-minimizer-webpack-plugin')

module.exports = {
  entry: './source/js/index.js',
  output: {
    path: __dirname + '/bundle/',
    filename: 'index.js'
  },
  module: {
    rules: [{
      test: /\.css$/i, use: [MiniCssExtractPlugin.loader, "css-loader", "postcss-loader"],
    }, { test: /\.woff|\.woff2|\.svg|.eot|\.ttf/, use: ['url-loader?limit=8192'] }, {
      test: /\.html$/i, use: [{ loader: 'raw-loader' },],
    },]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: "index.css",
    })
  ],
  optimization: {
    minimizer: [new UglifyJsPlugin({
      parallel: true,
      uglifyOptions: {
        output: {
          comments: false,
        },
      },
    }), new CssMinimizerPlugin({
      parallel: true,
      minimizerOptions: {
        preset: [
          'default',
          {
            discardComments: { removeAll: true },
          }
        ],
      },
    }),],
  },
  watch: false
}
