import React from 'react';
import { Button, Card, Colors, Divider, Subheading, Text, useTheme } from 'react-native-paper';
import { StyleProp, StyleSheet, ViewStyle } from 'react-native';
import { faCircleDollar } from '@fortawesome/pro-regular-svg-icons';
import { faBuildings, faUser } from '@fortawesome/pro-solid-svg-icons';
import { useNavigation } from '@react-navigation/native';

import CardItem from '../../../components/Appointments/CardItem';
import FontAwesomeIcon from '../../UI/FA-Icon';
import { Appointment } from '../../../models/index';
import { EM } from '../../UI/Layout';

const AppointmentDetailcostComponent: React.FC<{
  appointment: Appointment;
  style?: StyleProp<ViewStyle>;
}> = ({ appointment, style }) => {
  const navigation = useNavigation();
  const { colors } = useTheme();
  const styles = createStyles();

  const tests = appointment.patients.map(patient => patient.tests)
    .flat();
  const total = tests.reduce((previous, current) => {
    return previous + current.cost;
  }, 0);

  const costs = tests.map(test => {
    return {
      ...test,
      key: `${test.paidBy}-${test.name}`,
    };
  })
    .sort((a, b) => a.key.localeCompare(b.key))
    .reduce((prev, curr) => {
      if (!prev[curr.key]) {
        prev[curr.key] = {
          cost: curr.cost,
          name: curr.name,
          paidBy: curr.paidBy.charAt(0)
            .toUpperCase() + curr.paidBy.slice(1),
          total: 1,
        };
      } else {
        prev[curr.key].total += 1;
        prev[curr.key].cost += curr.cost;
      }

      return prev;
    }, {});

  return (
    <Card
      elevation={4}
      style={style}>
      <Card.Title
        left={() => <FontAwesomeIcon
          color={Colors.grey600}
          icon={faCircleDollar}
          size={42}
        />}
        subtitle={`$${total.toFixed(2)}`}
        title='Cost'
      />
      <Divider />
      <Card.Content style={styles.content}>
        {Object.values(costs)
          .map(cost => {
            let icon = faBuildings;
            if (cost.paidBy === 'Patient') {
              icon = faUser;
            }

            return (
              <CardItem
                icon={icon}
                iconColor={colors.accent}
                iconSize={24}
                key={`${cost.paidBy}-${cost.name}`} >
                <Subheading>{cost.paidBy}</Subheading>
                <Text style={styles.paragraph}>{`(${cost.total}) ${cost.name} = $${cost.cost.toFixed(2)}`}</Text>
              </CardItem>

            );
          })}
        <Card.Actions style={styles.buttons}>
          <Button onPress={() => null}>Receipt</Button>
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

export default AppointmentDetailcostComponent;
