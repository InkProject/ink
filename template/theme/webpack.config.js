var MiniCssExtractPlugin = require("mini-css-extract-plugin")

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
    new MiniCssExtractPlugin()
  ],
  watch: false
}
