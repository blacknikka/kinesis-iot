const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const webpack = require("webpack");

module.exports = {
  entry: __dirname + '/src/index.tsx',
  plugins: [
    new CleanWebpackPlugin({
      cleanAfterEveryBuildPatterns:['build']
    }),
    new HtmlWebpackPlugin({
      template: 'src/templates/index.html'
    }),
    new webpack.DefinePlugin({
      "process.env": {
        'REACT_APP_BACKEND_ENDPOINT': JSON.stringify(process.env.REACT_APP_BACKEND_ENDPOINT),
      },
    }),
  ],
  output: {
    path: __dirname + '/build',
    filename: '[name].[contenthash].js'
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js']
  },
  module: {
    rules: [
      { test: /\.tsx?$/, loader: 'ts-loader' }
    ]
  },
  devServer: {
    compress: true,
    port: 3000,
    host: '0.0.0.0',
  },
  target: 'node',
}
