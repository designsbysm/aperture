import BookNowScreen from '../../screens/BookNow';
import { NavigationScreen } from './types';

const screens: NavigationScreen[] = [
  {
    component: BookNowScreen,
    name: 'BookNow',
    title: 'Book Now',
  },
];

export default screens;
