const { VueLoaderPlugin } = require('vue-loader');
const path = require('path');

module.exports = {
  mode: 'development',
  entry: [
    './resources/js/app.js',
  ],
  module: {
    rules: [
      {
        test: /\.vue$/,
        use: 'vue-loader',
      },
      {
        test: /\.js$/,
        loader: 'babel-loader',
      },
      {
        test: /\.scss$/,
        use: [
          'vue-style-loader',
          'css-loader',
          'sass-loader',
        ],
      },
    ],
  },
  resolve: {
    extensions: ['.js', '.vue'],
    alias: {
      '@': path.resolve(__dirname, 'resources/js'),
      '@styles': path.resolve(__dirname, 'resources/sass'),
      vue$: 'vue/dist/vue.js',
    },
  },
  plugins: [
    new VueLoaderPlugin(),
  ],
};
