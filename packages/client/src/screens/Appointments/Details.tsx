import Error from '@smaperture/error';
import Loading from '@smaperture/loading';
import React from 'react';
import { EM } from '@smaperture/layout';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { Platform, ScrollView, StyleSheet } from 'react-native';

import CostCard from '../../components/Appointments/Details/CostCard';
import EventCard from '../../components/Appointments/Details/EventCard';
import ResultCard from '../../components/Appointments/Details/ResultCard';
import { ScreenParams } from '../../navigation/Params';
import { useGetAppointmentQuery } from '../../store/api';

const AppointmentDetailScreen: React.FC<Props> = ({ route }) => {
  const addLast = Platform.OS !== 'web';
  const id = route.params?.id;
  // const id = 21350;
  const styles = createStyles();

  const {
    data: appointment,
    isLoading,
    isError,
    error,
  } = useGetAppointmentQuery(id);

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
    <ScrollView style={styles.container}>
      <EventCard
        appointment={appointment}
        style={styles.card}
      />
      <ResultCard
        appointment={appointment}
        style={styles.card}
      />
      <CostCard
        appointment={appointment}
        style={[
          styles.card,
          addLast && styles.last,
        ]}
      />
    </ScrollView >
  );
};

type Props = NativeStackScreenProps<ScreenParams, 'AppointmentDetails'>;

const createStyles = () => StyleSheet.create({
  card: {
    marginTop: EM,
  },
  container: {
    marginHorizontal: EM,
    paddingBottom: EM,
  },
  last: {
    marginBottom: EM,
  },
});

export default AppointmentDetailScreen;
