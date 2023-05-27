import React from 'react';
import { Image, StyleSheet, View } from 'react-native';

import { EM } from '../UI/Layout';

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
    height: 180,
    width: 180,
  },
});

export default LogoComponent;
