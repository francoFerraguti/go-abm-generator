import React from 'react';
import MUIDrawer from 'material-ui/Drawer';
import {List, ListItem} from 'material-ui/List';
import IconLinearScale from 'material-ui/svg-icons/editor/linear-scale';
import IconAssignmentTurnedIn from 'material-ui/svg-icons/action/assignment-turned-in';
import IconClass from 'material-ui/svg-icons/action/class';
import IconPlace from 'material-ui/svg-icons/maps/place';
import IconDeviceStorage from 'material-ui/svg-icons/device/storage';
import IconDashBoard from 'material-ui/svg-icons/action/dashboard';
import IconContactCalendar from 'material-ui/svg-icons/action/perm-contact-calendar';
import IconLayers from 'material-ui/svg-icons/maps/layers';
import IconBookMark from 'material-ui/svg-icons/action/bookmark';
import IconPages from 'material-ui/svg-icons/social/pages';
import IconGroup from 'material-ui/svg-icons/social/group';

export default class Drawer extends React.Component {
  constructor(props) {
    super(props);
  }

  createList = (list) => {
    let nestedItems = list.props.nestedItems;
    let showItem = false;

    nestedItems.forEach((item) => {
      if (item !== null) {
        showItem = true;
      }
    });

    if (showItem) return list;
    return null;
  };

  createItem = (id, title, subtitle, url, icon, permission) => {
    return (
        <ListItem
            key={id}
            primaryText={subtitle}
            leftIcon={icon}
            onClick={() => {
            this.setState({
              openAdminMenu: false,
              openFactMenu: false,
            });
            this.props.navigate(url);
            this.props.setDrawerOpen(false);
            this.props.setTitle(title);
          }}/>
    );
  };

  render() {
      return (
          <MUIDrawer
              open={this.props.open}
              docked={false}
              onRequestChange={(open) => this.props.setDrawerOpen(open)}
          >
            <List>
            	{this.createItem(0, 'User', 'User', '/home/User', <IconAssignmentTurnedIn/>, '')}
{this.createItem(1, 'Role', 'Role', '/home/Role', <IconAssignmentTurnedIn/>, '')}
{this.createItem(2, 'Photo', 'Photo', '/home/Photo', <IconAssignmentTurnedIn/>, '')}
{this.createItem(3, 'RefugeAbility', 'RefugeAbility', '/home/RefugeAbility', <IconAssignmentTurnedIn/>, '')}
{this.createItem(4, 'AuditorReview', 'AuditorReview', '/home/AuditorReview', <IconAssignmentTurnedIn/>, '')}
{this.createItem(5, 'Pet', 'Pet', '/home/Pet', <IconAssignmentTurnedIn/>, '')}
{this.createItem(6, 'PetPhoto', 'PetPhoto', '/home/PetPhoto', <IconAssignmentTurnedIn/>, '')}

            </List>
          </MUIDrawer>
      );      
    }
  }
}
