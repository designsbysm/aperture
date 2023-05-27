import React from 'react';
import { Button, Card, Colors, Divider, Paragraph } from 'react-native-paper';
import { StyleProp, StyleSheet, ViewStyle } from 'react-native';
import { faFlask, faHashtag, faLocationDot } from '@fortawesome/pro-solid-svg-icons';
import { format } from 'date-fns';
import { useNavigation } from '@react-navigation/native';

import CardItem from './CardItem';
import FontAwesomeIcon from '../UI/FA-Icon';
import { Appointment } from '../../models';
import { CardStatus, cardStatus } from '../../utils/appointments';
import { EM } from '../UI/Layout';

const AppointmentListCardComponent: React.FC<Props> = ({ id, date, patients, site, style }) => {
  const navigation = useNavigation();

  const tests = patients.map(patient => patient.tests.map(test => test.name))
    .flat()
    .sort()
    .reduce((prev, curr) => {
      if (!prev[curr]) {
        prev[curr] = {
          name: curr,
          total: 1,
        };
      } else {
        prev[curr].total += 1;
      }

      return prev;
    }, {});

  const results = patients.map(patient => patient.result)
    .flat()
    .sort();

  const status = cardStatus(date, results);
  const styles = createStyles(status);

  return (
    <Card
      elevation={4}
      style={[
        styles.container,
        styles.outline,
        style,
      ]}>
      <Card.Title
        left={() => <FontAwesomeIcon
          color={status.color}
          icon={status.icon}
          size={42}
        />}
        subtitle={format(new Date(date), 'MM/dd/yyyy hh:mm bb')}
        title={status.title}
        titleStyle={styles.title}
      />
      <Divider />
      <Card.Content style={styles.content}>
        <CardItem
          icon={faHashtag}
          iconColor={Colors.grey600}
          iconSize={24}>
          <Paragraph>{id}</Paragraph>
        </CardItem>
        <CardItem
          icon={faLocationDot}
          iconColor={Colors.grey600}
          iconSize={24}>
          <Paragraph>{site.name}</Paragraph>
        </CardItem>
        <CardItem
          icon={faFlask}
          iconColor={Colors.grey600}
          iconSize={24}>
          {Object.values(tests)
            .map(test => {
              return (
                <Paragraph key={test.name}>{`(${test.total}) ${test.name}`}</Paragraph>
              );
            })

          }
        </CardItem>
        <Card.Actions style={styles.buttons}>
          {status.id === 'upcoming' && <Button onPress={() => navigation.navigate('AppointmentPassport', {
            id,
          })}>Passport</Button>}
          {status.id !== 'cancelled' && status.id !== 'pending' && status.id !== 'upcoming' && <Button onPress={() => null}>Results</Button>}
          <Button onPress={() => navigation.navigate('AppointmentDetails', {
            id,
          })}>More Info</Button>
        </Card.Actions>
      </Card.Content>
    </Card>
  );
};

const createStyles = (status: CardStatus) => StyleSheet.create({
  buttons: {
    justifyContent: 'flex-end',
  },
  container: {
    borderRadius: EM / 2,
    borderWidth: 3,
  },
  content: {
    paddingBottom: 0,
    paddingHorizontal: 0,
  },
  outline: {
    borderColor: status.color,
  },
  title: {
    textTransform: 'capitalize',
  },
});

type Props = Appointment & {
  style?: StyleProp<ViewStyle>;
}

export default AppointmentListCardComponent;
