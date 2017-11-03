import getMuiTheme from 'material-ui/styles/getMuiTheme';
import {black, orange800} from 'material-ui/styles/colors';

const colors = {
    appBarColor: 'rgb(50,50,50)',
    textBarColor: black,
    primaryButtonColor: orange800,
    secondaryButtonColor: '#d33520',
    secondaryTextButtonColor: black,
    labelCheckBoxColor: orange800,
    textFieldColor: '#fffdea',
    drawerColor: 'rgb(50,50,50)',
};

export const muiTheme = getMuiTheme({
    palette: {
        accent1Color: colors.primaryButtonColor,
    },
    raisedButton: {
        primaryButtonColor: orange800,
        secondaryButtonColor: orange800,
        textColor: orange800,
    },
    textField: {
        focusColor: colors.textFieldColor,
    },
});
