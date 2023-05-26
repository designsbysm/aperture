import 'react-native-gesture-handler';

import React from 'react';
import { Colors, DefaultTheme, Provider as PaperProvider } from 'react-native-paper';
import { Provider as StoreProvider } from 'react-redux';
import { StatusBar } from 'expo-status-bar';
import { SafeAreaView, StyleSheet, View } from 'react-native';
import { Theme } from 'react-native-paper/lib/typescript/types';

import AuthenticationNavigation from './src/navigation/Authentication';
import { store } from './src/store';

const App: React.FC = () => {
  const theme = {
    ...DefaultTheme,
    colors: {
      ...DefaultTheme.colors,
      accent: '#eb942a',
      background: '#f2f2f2',
      error: Colors.red600,
      primary: '#0075be',
    },
    roundness: 4,
  };

  const styles = createStyles(theme);

  return (
    <SafeAreaView style={styles.container}>
      <StatusBar style='light' />
      <StoreProvider store={store}>
        <PaperProvider theme={theme}>
          <AuthenticationNavigation />
        </PaperProvider>
      </StoreProvider>
    </SafeAreaView>
  );
};

const createStyles = (theme: Theme) => StyleSheet.create({
  container: {
    backgroundColor: theme.colors.primary,
    flex: 1,
  },
});

export default App;
