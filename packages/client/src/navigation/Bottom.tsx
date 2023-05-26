import React from 'react';
import { StyleSheet } from 'react-native';
import { Theme } from 'react-native-paper/lib/typescript/types';
import { createMaterialBottomTabNavigator } from '@react-navigation/material-bottom-tabs';
import { useTheme } from 'react-native-paper';

import renderScreen from './renderers/bottom';
import screens from './screens/authenticated';

const BottomNavigation = () => {
  const theme = useTheme();
  const styles = createStyles(theme);

  return (
    <Tabs.Navigator
      barStyle={styles.container}
      shifting={false}>
      {screens.map(screen => renderScreen(screen))}
    </Tabs.Navigator >
  );
};

const createStyles = (theme: Theme) => StyleSheet.create({
  container: {
    backgroundColor: theme.colors.primary,
  },
});

const Tabs = createMaterialBottomTabNavigator();

export default BottomNavigation;
