const path = require('path'); // Импортируем модуль "path" для работы с путями файлов
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const SpriteLoaderPlugin = require('svg-sprite-loader/plugin');

const filePath = {
  src: {
    script: './assets/js/script.js',
    style: './assets/scss/style.scss',
    icons: './assets/icons/icons.js'
  },
  dist: './public/'
};

module.exports = {
  entry: [
    filePath.src.script,
    filePath.src.style,
    filePath.src.icons
  ], // Точка входа для сборки проекта

  output: {
    filename: './js/script.js', // Имя выходного файла сборки
    path: path.resolve(__dirname, filePath.dist) // Путь для выходного файла сборки
  },

  module: {
    rules: [
      {
        test: /\.(c|sa|sc)ss$/i, // Регулярное выражение для обработки файлов с расширением .scss, .sass, .css
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader'
        ] // Загрузчики, используемые для обработки SCSS, SASS, CSS-файлов
      },
      {
        test: /\.svg$/,
        use: [{
          loader: 'svg-sprite-loader',
          options: {
            extract: true,
            spriteFilename: './icons/sprite.svg',
            symbolId: filePath => 'icon-' + path.basename(filePath).split('.')[0]
          }
        }, {
          loader: 'svgo-loader'
        }]
      }
    ]
  },

  plugins: [
    new MiniCssExtractPlugin({
      filename: './css/style.css'
    }),
    new SpriteLoaderPlugin({
      plainSprite: true
    }),
    new CopyWebpackPlugin({
      patterns: [
        { from: './assets/images', to: 'images', noErrorOnMissing: true },
        { from: './assets/fonts', to: 'fonts', noErrorOnMissing: true }
      ]
    })
  ],

  mode: 'development' // Режим сборки
};
