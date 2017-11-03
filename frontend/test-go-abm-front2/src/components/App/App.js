import React from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import {MuiThemeProvider} from 'material-ui';
import Login from '../Login/Login';
import Layout from '../Layout/Layout';
import {muiTheme} from '../MuiTheme/MuiTheme';

class App extends React.Component {
    render() {
        return (
            <MuiThemeProvider muiTheme={muiTheme}>
                <BrowserRouter basename="/">
                    <Switch>

                        <Route path="/login" component={Login}/>
                        <Route path="/signup" component={Signup}/>
                        <Route path="/home" component={Layout}/>

                        <Route component={Login}/>

                    </Switch>
                </BrowserRouter>
            </MuiThemeProvider>
        );
    }
}

export default App;
