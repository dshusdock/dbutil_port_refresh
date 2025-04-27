const path = require('path');

module.exports = {
  entry: './ui/html/scripts/index.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, './ui/html/dist'),
  },
};