import React from 'react';
import { FlatList, ListRenderItemInfo, Platform, StyleSheet } from 'react-native';

import Card from './Card';
import { Appointment } from '../../models';
import { EM } from '../UI/Layout';

const AppointmentListComponent: React.FC<{
  appointments: Appointment[];
}> = ({ appointments }) => {
  const styles = createStyles();
  const renderItem = (data: ListRenderItemInfo<Appointment>) => {
    const addLast = Platform.OS !== 'web' && (data.index + 1 === appointments.length);

    return (
      <Card
        {...data.item}
        style={[
          styles.card,
          addLast && styles.last,
        ]}
      />
    );
  };

  return (
    <FlatList
      data={appointments}
      renderItem={renderItem}
      style={styles.container}
    />
  );
};

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

export default AppointmentListComponent;
