import FontAwesomeIcon from '@smaperture/fa-icon';
import React from 'react';
import { Button, Card, Colors, Divider, Paragraph, Subheading, Text, useTheme } from 'react-native-paper';
import { EM } from '@smaperture/layout';
import { StyleProp, StyleSheet, ViewStyle } from 'react-native';
import { faCalendarDay } from '@fortawesome/pro-regular-svg-icons';
import { faHashtag, faLocationDot, faNote } from '@fortawesome/pro-solid-svg-icons';
import { format } from 'date-fns';
import { useNavigation } from '@react-navigation/native';

import CardItem from '../../../components/Appointments/CardItem';
import { Appointment } from '../../../models/index';
import { cardStatus } from '../../../utils/appointments';

const AppointmentDetailEventComponent: React.FC<{
  appointment: Appointment;
  style?: StyleProp<ViewStyle>;
}> = ({ appointment, style }) => {
  const navigation = useNavigation();
  const { colors } = useTheme();
  const styles = createStyles();

  const results = appointment.patients.map(patient => patient.result)
    .flat()
    .sort();
  const appointmentStatus = cardStatus(appointment.date, results);
  if (appointmentStatus.title !== 'Cancelled' && appointmentStatus.title !== 'Pending' && appointmentStatus.title !== 'Up Coming') {
    appointmentStatus.title = 'Completed';
  }

  return (
    <Card
      elevation={4}
      style={style}>
      <Card.Title
        left={() => <FontAwesomeIcon
          color={Colors.grey600}
          icon={faCalendarDay}
          size={42}
        />}
        subtitle={appointmentStatus.title}
        title={format(new Date(appointment?.date), 'MM/dd/yyyy hh:mm bb')}
      />
      <Divider />
      <Card.Content style={styles.content}>
        <CardItem
          icon={faHashtag}
          iconColor={colors.accent}
          iconSize={24}>
          <Subheading>{appointment?.id}</Subheading>
        </CardItem>
        <CardItem
          icon={faLocationDot}
          iconColor={colors.accent}
          iconSize={24}>
          <Subheading>{appointment?.site.name}</Subheading>
          <Text style={styles.paragraph}>{appointment?.site.address}</Text>
          <Text style={styles.paragraph}>{`${appointment?.site.city}, ${appointment?.site.state} ${appointment?.site.zip}`}</Text>
        </CardItem>
        {appointment?.site.notes && <CardItem
          icon={faNote}
          iconColor={colors.accent}
          iconSize={24}>
          <Paragraph>{appointment?.site.notes}</Paragraph>
        </CardItem>}
        <Card.Actions style={styles.buttons}>
          <Button onPress={() => null}>Cancel</Button>
          <Button onPress={() => navigation.navigate('AppointmentPassport', {
            id: appointment.id,
          })}>Passport</Button>
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

export default AppointmentDetailEventComponent;
