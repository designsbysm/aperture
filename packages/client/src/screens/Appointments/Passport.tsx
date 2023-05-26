import React from 'react';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StyleSheet, Text, View } from 'react-native';

import { ScreenParams } from '../../navigation/Params';

const AppointmentPassportScreen: React.FC<Props> = ({ navigation }) => {
  const styles = createStyles();

  return (
    <View style={styles.container}></View>
  );
};

type Props = NativeStackScreenProps<ScreenParams, 'AppointmentPassport'>;

const createStyles = () => StyleSheet.create({
  container: {},
});

export default AppointmentPassportScreen;
