import React from 'react';
import { createDrawerNavigator } from '@react-navigation/drawer';
import { faBug } from '@fortawesome/pro-solid-svg-icons';

import FontAwesomeIcon from '../../components/UI/FA-Icon';
import { NavigationScreen } from '../screens/types';

const renderScreen = (screen: NavigationScreen) => {
  return (
    <Drawer.Screen
      component={screen.component}
      key={screen.name}
      name={screen.name}
      options={{
        drawerIcon: ({ color }) => (
          <FontAwesomeIcon
            color={color}
            icon={screen.icon || faBug}
            size={24}
          />
        ),
        title: screen.title,
      }}
    />
  );
};

const Drawer = createDrawerNavigator();

export default renderScreen;
