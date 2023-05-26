import React from 'react';
import { Button, StyleSheet, Text, View } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';

import { ScreenParams } from '../../navigation/Params';

const ForgotPasswordScreen: React.FC<Props> = ({ navigation }) => {
  const styles = createStyles();

  return (
    <View style={styles.container}></View>
  );
};

type Props = NativeStackScreenProps<ScreenParams, 'ForgotPassword'>;

const createStyles = () => StyleSheet.create({
  container: {},
});

export default ForgotPasswordScreen;
