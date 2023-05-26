import React, { useEffect, useState } from 'react';
import { Button, Text, TextInput, useTheme } from 'react-native-paper';
import { EM } from '@smaperture/layout';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StyleSheet, View } from 'react-native';
import { Theme } from 'react-native-paper/lib/typescript/types';

import env from '../../environment';
import Logo from '../../components/Login/Logo';
import { ScreenParams } from '../../navigation/Params';

const LoginScreen: React.FC<Props> = ({ navigation }) => {
  const [
    email,
    setEmail,
  ] = useState('');
  const [
    password,
    setPassword,
  ] = useState('');
  const theme = useTheme();
  const styles = createStyles(theme);

  const fetchHealth = async () => {
    const response = await fetch(`${env.apiURL}/v1/healthcheck`);
    const healthcheck = await response.text();

    console.log(healthcheck);
  };

  useEffect(() => {
    fetchHealth();
  }, []);

  return (
    <View style={styles.container}>
      <Logo />
      <TextInput
        label='Email'
        onChangeText={text => setEmail(text)}
        style={styles.input}
        // underlineColor='#000'
        value={email}
      />
      <TextInput
        label='Password'
        onChangeText={text => setPassword(text)}
        secureTextEntry={true}
        style={styles.input}
        value={password}
      />
      <View style={styles.buttons}>
        <Button
          contentStyle={styles.buttonContent}
          mode='contained'
          onPress={() => navigation.navigate('Authenticated')}
          style={styles.button}>
          Login
        </Button>
        <Button
          contentStyle={styles.buttonContent}
          onPress={() => navigation.navigate('ForgotPassword')}
          style={styles.button}>
          Lost Your Password?
        </Button>
      </View>
      <View style={styles.signupContainer}>
        <View style={styles.signup}>
          <Text>Don’t have an account?  </Text>
          <Button
            onPress={() => navigation.navigate('SignupNavigation')}>
            Signup?
          </Button>
        </View>
      </View>
    </View>
  );
};

type Props = NativeStackScreenProps<ScreenParams, 'Login'>;

const createStyles = (theme: Theme) => StyleSheet.create({
  button: {
    marginTop: EM,
  },
  buttonContent: {
    height: 64,
    justifyContent: 'center',
  },
  buttons: {
    marginTop: EM,
  },
  container: {
    flex: 1,
    marginBottom: EM,
    marginHorizontal: EM,
  },
  input: {
    backgroundColor: '#fff',
    marginTop: EM * 2,
  },
  signup: {
    alignItems: 'center',
    flexDirection: 'row',
  },
  signupContainer: {
    alignItems: 'center',
    flex: 1,
    justifyContent: 'flex-end',
  },
  // signupLink: {
  // color: theme.colors.primary,
  // fontWeight: '500',
  // textTransform: 'uppercase',
  // },
});

export default LoginScreen;
