import React from 'react';
import {Link} from 'react-router-dom';
import FlatButton from 'material-ui/FlatButton';
import Popover, {PopoverAnimationVertical} from 'material-ui/Popover';
import Menu from 'material-ui/Menu';
import MenuItem from 'material-ui/MenuItem';

const INITIAL_STATE = {
    open: false,
    username: 'User',
};

export default class Avatar extends React.Component {
  constructor(props) {
    super(props);
    this.state = INITIAL_STATE;
  }

  componentDidMount() {
      this.getAuthenticatedUsername();
  }

  handleTouchTap = (event) => {
    this.setState({
      open: true,
      anchorEl: event.currentTarget,
    });
  };

  handleRequestClose = () => {
    this.setState({
      open: false,
    });
  };

  getAuthenticatedUsername() {
    let username = localStorage.getItem('username');
    this.setState({username});
  }

  handleLogout() {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
  }

  render() {
    return (
      <div style={{marginTop: 5}}>
        <FlatButton
          style={{color: '#fff'}}
          onClick={this.handleTouchTap}
          label={this.state.username}/>
        <Popover
          open={this.state.open}
          anchorEl={this.state.anchorEl}
          anchorOrigin={{horizontal: 'left', vertical: 'bottom'}}
          targetOrigin={{horizontal: 'left', vertical: 'top'}}
          animation={PopoverAnimationVertical}
          onRequestClose={this.handleRequestClose}
        >
          <Menu>
            <Link to="/">
              <MenuItem primaryText="Log Out" onClick={() => { this.handleLogout(); this.handleRequestClose(); }}/>
            </Link>
          </Menu>
        </Popover>
      </div>
    );
  }
}
