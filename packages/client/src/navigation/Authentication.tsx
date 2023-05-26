import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

import renderScreen from './renderers/stack';
import screens from './screens/authentication';

const AuthenticationNavigation = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator
        // initialRouteName='Authenticated'
        screenOptions={{
          headerShown: false,
        }}>
        {screens.map(screen => renderScreen(screen))}
      </Stack.Navigator>
    </NavigationContainer>
  );
};

const Stack = createNativeStackNavigator();

export default AuthenticationNavigation;
