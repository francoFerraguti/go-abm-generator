import React from 'react';
import {Redirect, Route, Switch} from 'react-router-dom';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import NavigationMenu from 'material-ui/svg-icons/navigation/menu';
import Avatar from './Avatar';
import Drawer from './Drawer';
import Main from './Main';
import User from '../User/User.js'
import Role from '../Role/Role.js'
import Photo from '../Photo/Photo.js'
import RefugeAbility from '../RefugeAbility/RefugeAbility.js'
import AuditorReview from '../AuditorReview/AuditorReview.js'
import Pet from '../Pet/Pet.js'
import PetPhoto from '../PetPhoto/PetPhoto.js'


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
            	<Route path='/home/User' render={() => <User/>}/>
<Route path='/home/Role' render={() => <Role/>}/>
<Route path='/home/Photo' render={() => <Photo/>}/>
<Route path='/home/RefugeAbility' render={() => <RefugeAbility/>}/>
<Route path='/home/AuditorReview' render={() => <AuditorReview/>}/>
<Route path='/home/Pet' render={() => <Pet/>}/>
<Route path='/home/PetPhoto' render={() => <PetPhoto/>}/>


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
