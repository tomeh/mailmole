module.exports = {
    'root': true,
      'plugins': [
          'html',
      ],
      'env': {
          'browser': true,
          'es6': true,
          'jquery': true,
      },
    'extends': [
      'airbnb-base',
    ],
    'parser': 'babel-eslint',
      'parserOptions': {
          'sourceType': 'module',
      },
      'settings': {
      'import/resolver': {
        'webpack': {
          'config': 'webpack.config.js',
        },
      }
    },
      // add your custom rules here
    'rules': {
      // don't require .vue extension when importing
      'import/extensions': ['error', 'always', {
        'js': 'never',
        'vue': 'never'
      }],
      // allow optionalDependencies
      'import/no-extraneous-dependencies': ['error', {
        'optionalDependencies': ['test/unit/index.js']
      }],
      // allow debugger during development
      'no-debugger': process.env.NODE_ENV === 'production' ? 2 : 0
    },
      'globals': {
        'Vue': true,
        'Pusher': true,
        'axios': true,
        'swal': true,
        '_': true,
      },
  };