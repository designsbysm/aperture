import React from 'react';
import { EM } from '@smaperture/layout';
import { Image, StyleSheet, View } from 'react-native';

const LogoComponent: React.FC = () => {
  const styles = createStyles();

  return (
    <View style={styles.container}>
      <Image
        source={require('../../../assets/logo.png')}
        style={styles.image}
      />
    </View>
  );
};

const createStyles = () => StyleSheet.create({
  container: {
    alignItems: 'center',
    marginBottom: EM * 2,
    marginTop: EM * 4,
  },
  image: {
    height: 90,
    width: 182,
  },
});

export default LogoComponent;
