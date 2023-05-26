import React from 'react';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

import renderScreen from './renderers/stack';
import screens from './screens/appointments';
import { useHeaderOptions } from './options';

const AppointmentsNavigation = () => {
  const options = useHeaderOptions();

  return (
    <Stack.Navigator screenOptions={options}>
      {screens.map(screen => renderScreen(screen))}
    </Stack.Navigator>
  );
};

const Stack = createNativeStackNavigator();

export default AppointmentsNavigation;
