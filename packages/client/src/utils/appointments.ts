import { Colors } from 'react-native-paper';
import { IconDefinition } from '@fortawesome/fontawesome-common-types';
import { faCircleCheck, faCircleEllipsis, faCircleMinus, faCirclePlus, faCircleXmark, faClock } from '@fortawesome/pro-regular-svg-icons';

export const cardStatus = (date: string, results: string[]): CardStatus => {
  const unique = [...new Set(results)];

  let status = 'unknown';

  if (unique.length === 1) {
    status = unique[0];
  } else if (unique.includes('pending')) {
    status = 'pending';
  } else {
    status = 'completed';
  }

  if (status === 'pending') {
    const now = new Date()
      .toISOString();

    if (date > now) {
      status = 'upcoming';
    }
  }

  const title = statusToTitle(status);
  let icon = faClock;
  let color = Colors.grey600;

  switch (status) {
    case 'cancelled': {
      icon = faCircleXmark;
      color = Colors.grey900;
      break;
    }
    case 'completed': {
      icon = faCircleCheck;
      color = Colors.grey900;
      break;
    }
    case 'negative': {
      icon = faCircleMinus;
      color = Colors.green400;
      break;
    }
    case 'pending': {
      icon = faCircleEllipsis;
      break;
    }
    case 'positive': {
      icon = faCirclePlus;
      color = Colors.red400;
      break;
    }
  }

  return {
    color,
    icon,
    id: status,
    title,
  };
};

export const statusToTitle = (status: string) => {
  if (!status) {
    return 'Unkown';
  }

  let title = status.charAt(0)
    .toUpperCase() + status.slice(1);

  switch (status) {
    case 'upcoming': {
      title = 'Up Coming';
      break;
    }
  }

  return title;
};

export type CardStatus = {
  color: string;
  icon: IconDefinition;
  id: string;
  title: string;
}
