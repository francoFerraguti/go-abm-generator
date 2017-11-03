import React from 'react';
import {Redirect, Route, Switch} from 'react-router-dom';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import NavigationMenu from 'material-ui/svg-icons/navigation/menu';
import Avatar from './Avatar';
import Drawer from './Drawer';
import Main from './Main';
%%IMPORT_LIST%%

const containerStyles = {
    layout: {
        backgroundColor: '#D50032',
    },
};

export default class Layout extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            responsiveMenu: false,
            loading: false,
            drawerOpen: false,
            title: 'Menu',
        };
    }

    isAuthenticated() {
          const token = localStorage.getItem('token');
          return token && token.length > 10;
    }

    navigate = (url) => {
        this.props.history.push(url);
    };

    setTitle = (title) => {
        this.setState({
          title: title,
        });
    };

  toggleDrawer = () => this.setState({drawerOpen: !this.state.drawerOpen});

  setDrawerOpen = (open) => this.setState({drawerOpen: open});

  render() {
    const isAlreadyAuthenticated = this.isAuthenticated();

    return (
      <div>
        {!isAlreadyAuthenticated ? <Redirect to={{pathname: '/'}}/> : (
        <div>
          <AppBar
            title={this.state.title}
            iconElementLeft={<IconButton onClick={this.toggleDrawer}><NavigationMenu/></IconButton>}
            iconElementRight={<Avatar username = {this.props.location.username}/>}
            style={containerStyles.layout}
          />
           <div className="main-content">
            <Switch>
            	%%ROUTE_LIST%%

                <Route path="/home" render={() => <Main/>}/>
            <Redirect to="/"/>
            </Switch>
          </div>
          <Drawer
            open={this.state.drawerOpen}
            toggle={this.toggleDrawer}
            setDrawerOpen={this.setDrawerOpen}
            setTitle={this.setTitle}
            navigate={this.navigate}
          />
        </div>
        )}
      </div>
    );
  }
}
