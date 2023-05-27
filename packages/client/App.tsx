import { SafeAreaView, StyleSheet, Text, View } from 'react-native';

import StartScreen from './src/screens/Start';

const App = () => {
  return (
    <SafeAreaView style={styles.container}>
      <StartScreen />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

export default App;
