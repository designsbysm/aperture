import FontAwesomeIcon from '@smaperture/fa-icon';
import React from 'react';
import { Button, Card, Colors, Divider, Subheading, Text } from 'react-native-paper';
import { EM } from '@smaperture/layout';
import { StyleProp, StyleSheet, ViewStyle } from 'react-native';
import { faFlask } from '@fortawesome/pro-regular-svg-icons';
import { useNavigation } from '@react-navigation/native';

import CardItem from '../../../components/Appointments/CardItem';
import { Appointment } from '../../../models/index';
import { cardStatus } from '../../../utils/appointments';

const AppointmentDetailResultComponent: React.FC<{
  appointment: Appointment;
  style?: StyleProp<ViewStyle>;
}> = ({ appointment, style }) => {
  const navigation = useNavigation();
  const styles = createStyles();

  const patients = appointment.patients.map(patient => ({
    ...patient,
    tests: patient.tests.map(test => test.name)
      .sort(),
  }))
    .sort((a, b) => a.name.localeCompare(b.name));

  return (
    <Card
      elevation={4}
      style={style}>
      <Card.Title
        left={() => <FontAwesomeIcon
          color={Colors.grey600}
          icon={faFlask}
          size={42}
        />}
        subtitle={appointment?.lab.name}
        title='Results'
      />
      <Divider />
      <Card.Content style={styles.content}>
        {patients.map(patient => {
          const status = cardStatus(appointment.date, [patient.result]);
          return (
            <CardItem
              icon={status.icon}
              iconColor={status.color}
              iconSize={24}
              key={patient.name}>
              <Subheading>{`${patient.name} –  ${status.title}`}</Subheading>
              {patient.tests.map(test => {
                return (
                  <Text
                    key={test}
                    style={styles.paragraph}>
                    {test}
                  </Text>
                );
              })}
            </CardItem>
          );
        })}
        <Card.Actions style={styles.buttons}>
          <Button onPress={() => null}>Results</Button>
        </Card.Actions>
      </Card.Content>
    </Card>
  );
};

const createStyles = () => StyleSheet.create({
  buttons: {
    justifyContent: 'flex-end',
    paddingTop: EM,
  },
  content: {
    paddingBottom: 0,
    paddingHorizontal: 0,
  },
  paragraph: {
    color: Colors.grey600,
  },
});

export default AppointmentDetailResultComponent;
