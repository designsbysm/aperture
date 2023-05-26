import DetailsScreen from '../../screens/Appointments/Details';
import ListScreen from '../../screens/Appointments';
import PassportScreen from '../../screens/Appointments/Passport';
import SearchScreen from '../../screens/Appointments/Search';
import { NavigationScreen } from './types';

const screens: NavigationScreen[] = [
  {
    component: ListScreen,
    name: 'Appointments',
  },
  {
    component: DetailsScreen,
    name: 'AppointmentDetails',
    title: 'Details',
  },
  {
    component: PassportScreen,
    name: 'AppointmentPassport',
    title: 'Passport',
  },
  {
    component: SearchScreen,
    name: 'AppointmentSearch',
    title: 'Search',
  },
];

export default screens;
