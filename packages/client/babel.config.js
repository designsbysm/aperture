module.exports = function (api) {
  api.cache(true);
  return {
    plugins: [
      [
        'module:react-native-dotenv',
        {
          allowUndefined: false,
          // safe: true,
        },
      ],
      'react-native-reanimated/plugin',
    ],
    presets: ['babel-preset-expo'],
  };
};
