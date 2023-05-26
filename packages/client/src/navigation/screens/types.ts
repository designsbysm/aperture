import { ReactNode } from 'react';
import { IconDefinition } from '@fortawesome/fontawesome-common-types';

export type NavigationScreen = {
  component: ReactNode,
  icon?: IconDefinition,
  name: string;
  title?: string;
}
