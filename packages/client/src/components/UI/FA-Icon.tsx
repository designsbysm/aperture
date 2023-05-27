import React from 'react';
import { FontAwesomeIcon as FontAwesomeReact } from '@fortawesome/react-fontawesome';
import { FontAwesomeIcon as FontAwesomeNative } from '@fortawesome/react-native-fontawesome';
import { IconDefinition } from '@fortawesome/fontawesome-common-types';
import { Platform, StyleProp, StyleSheet, View, ViewStyle } from 'react-native';

const Native: React.FC<{
  size: number;
}> = ({ size, ...props }) => {
  return (
    // @ts-expect-error: size is passed as a param
    <FontAwesomeNative
      {...props}
      size={size}
    />);
};

const Web: React.FC<{
  size: number;
}> = ({ size, ...props }) => {
  const styles = StyleSheet.flatten([
    {
      height: size,
      width: size,
    },
  ]);

  return (
    // @ts-expect-error: size is passed as a style
    <FontAwesomeReact
      {...props}
      style={styles}
    />);
};

const Icon: React.FC<{
  color: string;
  icon: IconDefinition;
  size: number;
  style?: StyleProp<ViewStyle>;
}> = ({ size, style, ...props }) => {
  return (
    <View style={style}>
      {Platform.OS === 'web' && <Web
        {...props}
        size={size}
      />}
      {Platform.OS !== 'web' && <Native
        {...props}
        size={size}
      />}
    </View>
  );
};

export default Icon;
