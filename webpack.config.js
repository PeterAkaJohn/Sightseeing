const path = require('path');
const webpack = require('webpack');

module.exports = {
  context: path.resolve(__dirname, './app'),
  entry: './index.js',
  output: {
    path: __dirname,
    filename: './bundle.js'
  },
  resolve: {
    modules: [path.resolve(__dirname, './app'), 'node_modules', 'bower_components'],
  },
  devServer: {
    historyApiFallback: true,
    contentBase: path.resolve(__dirname, './app')
  },
  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        loader: 'babel-loader',
        exclude: /(node_modules|bower_components)/,
        query: {
          presets: ['es2017', 'react', 'stage-3'],
          plugins: ['transform-decorators-legacy'],
        }
      }
    ]
  }
}
