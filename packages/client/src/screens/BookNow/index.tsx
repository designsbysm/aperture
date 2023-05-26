import React from 'react';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { Button, StyleSheet, Text, View } from 'react-native';

import { ScreenParams } from '../../navigation/Params';

const BookNowScreen: React.FC<Props> = ({ navigation }) => {
  const styles = createStyles();

  return (
    <View style={styles.container}></View>
  );
};

type Props = NativeStackScreenProps<ScreenParams, 'BookNow'>;

const createStyles = () => StyleSheet.create({
  container: {},
});

export default BookNowScreen;
