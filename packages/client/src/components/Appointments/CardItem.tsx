import React from 'react';
import { Divider } from 'react-native-paper';
import { IconDefinition } from '@fortawesome/fontawesome-common-types';
import { StyleSheet, View } from 'react-native';

import FontAwesomeIcon from '../UI/FA-Icon';
import { EM } from '../UI/Layout';

const ItemHeaderComponent: React.FC<{
  icon: IconDefinition;
  iconColor: string;
  iconSize: number;
}> = ({ children, icon, iconColor, iconSize }) => {
  const styles = createStyles();

  return (
    <View style={styles.container}>
      <FontAwesomeIcon
        color={iconColor}
        icon={icon}
        size={iconSize}
        style={styles.icon}
      />
      <View style={styles.items}>
        <View style={styles.text}>
          {children}
        </View>
        <Divider />
      </View>
    </View>
  );
};

const createStyles = () => StyleSheet.create({
  container: {
    flexDirection: 'row',
  },
  icon: {
    alignItems: 'center',
    marginRight: EM,
    paddingLeft: EM,
    paddingVertical: EM / 2,
    width: EM * 3.5,
  },
  items: {
    flex: 1,
  },
  text: {
    paddingRight: EM,
    paddingVertical: EM / 2,
  },
});

export default ItemHeaderComponent;
