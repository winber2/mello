const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const BUILD_DIR = path.resolve(__dirname, './public');
const APP_DIR = path.resolve(__dirname, './src');
const STATIC_DIR = path.resolve(__dirname, '../static');

const config = {
  mode: 'development',
  entry: {
    index: `${APP_DIR}/index.jsx`
  },
  output: {
    path: BUILD_DIR,
    filename: '[name].[hash].js'
  },
  plugins: [
    new CleanWebpackPlugin([BUILD_DIR]),
    new webpack.DefinePlugin({
      'process.env': { NODE_ENV: JSON.stringify(process.env.NODE_ENV) }
    }),
    new HtmlWebpackPlugin({
      title: 'Mello',
      template: `${APP_DIR}/index.html`,
      filename: `${BUILD_DIR}/index.html`,
    }),
    new CopyWebpackPlugin([
      { from: STATIC_DIR, BUILD_DIR }
    ])
  ],
  module: {
    rules: [
      {
        include: APP_DIR,
        exclude: /(node_modules)/,
        test: /\.(js|jsx)$/,
        loader: 'babel-loader',
        query: {
          presets: [ "@babel/preset-env", "@babel/preset-react" ],
          plugins: [
            // Stage 2
            ["@babel/plugin-proposal-decorators", { "legacy": true }],
            "@babel/plugin-proposal-function-sent",
            "@babel/plugin-proposal-export-namespace-from",
            "@babel/plugin-proposal-numeric-separator",
            "@babel/plugin-proposal-throw-expressions",

            // Stage 3
            "@babel/plugin-syntax-dynamic-import",
            "@babel/plugin-syntax-import-meta",
            ["@babel/plugin-proposal-class-properties", { "loose": false }],
            "@babel/plugin-proposal-json-strings"
          ]
        }
      },
      {
        test: /\.(png|jpg|gif)$/,
        loader: 'url-loader'
      }
    ]
  },
  /* Some linux distibutions seem to have a problem without poll=true */
  watchOptions: {
    poll: true
  },
  devServer: {
    contentBase: path.join(__dirname, "./public"),
    port: '3000',
    proxy: {
      "/": {
        target: 'http://localhost:8000',
        secure: true
      }
    }
  },
  resolve: {
    extensions: [ '.js', '.jsx' ],
    modules: [
      'node_modules',
      path.resolve(__dirname, './src')
    ]
  }
};

if (process.env.NODE_ENV === 'development') {
  config.optimization = {
    minimize: true
  };
} else {
  config.devtool = 'source-map';
}

module.exports = config;
