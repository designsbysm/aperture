import { faCalendarDay, faCommentsQuestion, faPlus, faUser } from '@fortawesome/pro-solid-svg-icons';

import AppointmentsNavigation from '../Appointments';
import BookNowNavigation from '../BookNow';
import ProfileNavigation from '../Profile';
import SupportNavigation from '../Support';
import { NavigationScreen } from './types';

const screens: NavigationScreen[] = [
  {
    component: AppointmentsNavigation,
    icon: faCalendarDay,
    name: 'AppointmentsNavigation',
    title: 'Appointments',
  },
  {
    component: BookNowNavigation,
    icon: faPlus,
    name: 'BookNowNavigation',
    title: 'Book Now',
  },
  {
    component: SupportNavigation,
    icon: faCommentsQuestion,
    name: 'SupportNavigation',
    title: 'Support',
  },
  {
    component: ProfileNavigation,
    icon: faUser,
    name: 'ProfileNavigation',
    title: 'Profile',
  },
];

export default screens;
