import BottomNavigation from '../Bottom';
import DrawerNavigation from '../Drawer';
import ForgotPasswordScreen from '../../screens/Login/ForgotPassword';
import LoginScreen from '../../screens/Login';
import SignupNavigation from '../Signup';
import { NavigationScreen } from './types';
import { useMobileNav } from '../../components/UI/Layout';

const screens: NavigationScreen[] = [
  {
    component: LoginScreen,
    name: 'Login',
  },
  {
    component: ForgotPasswordScreen,
    name: 'ForgotPassword',
  },
  {
    component: SignupNavigation,
    name: 'SignupNavigation',
  },
  {
    component: useMobileNav ? BottomNavigation : DrawerNavigation,
    name: 'Authenticated',
  },
];

export default screens;
