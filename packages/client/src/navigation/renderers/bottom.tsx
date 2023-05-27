import React from 'react';
import { createMaterialBottomTabNavigator } from '@react-navigation/material-bottom-tabs';
import { faBug } from '@fortawesome/pro-solid-svg-icons';

import FontAwesomeIcon from '../../components/UI/FA-Icon';
import { NavigationScreen } from '../screens/types';

const renderScreen = (screen: NavigationScreen) => {
  return (
    <Tabs.Screen
      component={screen.component}
      key={screen.name}
      name={screen.name}
      options={{
        tabBarIcon: ({ color }) => (
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

const Tabs = createMaterialBottomTabNavigator();

export default renderScreen;
