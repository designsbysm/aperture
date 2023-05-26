import { StyleSheet } from 'react-native';
import { Theme } from 'react-native-paper/lib/typescript/types';
import { useMobileNav } from '@smaperture/layout';
import { useTheme } from 'react-native-paper';

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
