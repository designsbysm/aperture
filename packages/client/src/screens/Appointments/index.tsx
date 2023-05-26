import Error from '@smaperture/error';
import Loading from '@smaperture/loading';
import React, { useMemo } from 'react';
import { StyleSheet, View } from 'react-native';

import List from '../../components/Appointments/List';
import { useGetAppointmentsQuery } from '../../store/api';

const AppointmentListScreen = () => {
  const {
    data = [],
    isLoading,
    isError,
    error,
  } = useGetAppointmentsQuery(null);
  const styles = createStyles();

  const appointments = useMemo(() => data.slice()
    .filter(appointment => appointment)
    .sort((a, b) => a.date.localeCompare(b.date))
    .reverse(), [data]);

  if (isLoading) {
    return (
      <Loading />
    );
  } else if (isError) {
    return (
      <Error>{error.data?.toString() || error.error?.toString()}</Error>
    );
  }

  return (
    <View style={styles.container}>
      <List appointments={appointments} />
    </View>
  );
};

const createStyles = () => StyleSheet.create({
  container: {
    flex: 1,
  },
});

export default AppointmentListScreen;
