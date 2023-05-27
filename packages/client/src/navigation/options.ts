import { StyleSheet } from 'react-native';
import { Theme } from 'react-native-paper/lib/typescript/types';
import { useTheme } from 'react-native-paper';

import { useMobileNav } from '../components/UI/Layout';

export const useHeaderOptions = () => {
  const theme = useTheme();
  const styles = createStyles(theme);

  return {
    headerShown: useMobileNav,
    headerStyle: styles.header,
    headerTintColor: 'white',
  };
};

const createStyles = (theme: Theme) => StyleSheet.create({
  header: {
    backgroundColor: theme.colors.primary,
  },
});
