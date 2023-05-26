import React from 'react';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

import { NavigationScreen } from '../screens/types';

const renderScreen = (screen: NavigationScreen) => {
  return (
    <Stack.Screen
      component={screen.component}
      key={screen.name}
      name={screen.name}
      options={{
        title: screen.title,
      }}
    />
  );
};

const Stack = createNativeStackNavigator();

export default renderScreen;
