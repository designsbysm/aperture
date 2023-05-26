import React from 'react';
import { createDrawerNavigator } from '@react-navigation/drawer';

import renderScreen from './renderers/drawer';
import screens from './screens/authenticated';

const DrawerNavigation = () => {
  return (
    <Drawer.Navigator>
      {screens.map(screen => renderScreen(screen))}
    </Drawer.Navigator>
  );
};

const Drawer = createDrawerNavigator();

export default DrawerNavigation;
