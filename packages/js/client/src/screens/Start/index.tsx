import axios from 'axios';
import React, { useEffect } from 'react';
import { StyleSheet, Text, View } from 'react-native';

import env from '../../environment';

const StartScreen = () => {
  const fetchHealth = async () => {
    const response = await axios.get(`${env.apiURL}/v1/healthcheck`);
    console.log('healthcheck', response.data);
  };

  useEffect(() => {
    fetchHealth();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Open up App.tsx to start working on your app!</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

export default StartScreen;
