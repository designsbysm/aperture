import SignupScreen from '../../screens/Signup';
import AddressScreen from '../../screens/Signup/Address';
import ContactScreen from '../../screens/Signup/Contact';
import DemographicsScreen from '../../screens/Signup/Demographics';
import InsuranceScreen from '../../screens/Signup/Insurance';
import PrivacyPolicyScreen from '../../screens/Signup/PrivacyPolicy';
import TermsOfServiceScreen from '../../screens/Signup/TermsOfService';
import { NavigationScreen } from './types';

const screens: NavigationScreen[] = [
  {
    component: SignupScreen,
    name: 'Signup',
  },
  {
    component: AddressScreen,
    name: 'SignupAddress',
  },
  {
    component: ContactScreen,
    name: 'SignupContact',
  },
  {
    component: DemographicsScreen,
    name: 'SignupDemographics',
  },
  {
    component: InsuranceScreen,
    name: 'SignupInsurance',
  },
  {
    component: PrivacyPolicyScreen,
    name: 'SignupPrivacyPolicy',
  },
  {
    component: TermsOfServiceScreen,
    name: 'SignupTermsOfService',
  },
];

export default screens;
