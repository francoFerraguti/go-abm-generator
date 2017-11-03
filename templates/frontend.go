package templates

func GitIgnore() string {
	return `# Numerous always-ignore extensions

*.diff

*.err

*.orig

*.log

*.rej

*.swo

*.swp

*.vi

*~

*.sass-cache


# OS or Editor folders

.DS_Store

.cache

.project

.settings

.tags

.tmproj

.idea

nbproject

Thumbs.db

tags



# NPM packages folder.

node_modules/



# Brunch folder for temporary files.

tmp/



# Brunch output folder.



# Bower stuff.

bower_components/



# API Docs

docs/api/



# DotEnv

.env
`
}

func PackageJson() string {
	return `{
  "name": "todo-change",
  "version": "0.1.0",
  "private": true,
  "dependencies": {
    "axios": "^0.16.2",
    "jquery": "^3.2.1",
    "material-ui": "^0.19.1",
    "materialui-pagination": "0.0.6",
    "react": "^15.6.1",
    "react-confirm-alert": "^1.0.7",
    "react-dom": "^15.6.1",
    "react-form-validator-core": "^0.1.0",
    "react-google-maps": "^9.2.0",
    "react-loading-spinner": "^1.0.12",
    "react-material-ui-form-validator": "^1.0.0",
    "react-notification-system": "^0.2.15",
    "react-router": "^4.2.0",
    "react-router-dom": "^4.2.2",
    "react-scripts": "1.0.13",
    "underscore": "^1.8.3"
  },
  "scripts": {
    "start": "react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test --env=jsdom",
    "eject": "react-scripts eject"
  },
  "devDependencies": {
    "eslint-config-google": "^0.9.1"
  },
  "main": ".eslintrc.js",
  "keywords": [],
  "author": "",
  "license": "ISC",
  "description": ""
}
`
}

func IndexHtml() string {
	return `<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="theme-color" content="#000000">
    <link rel="manifest" href="%PUBLIC_URL%/manifest.json">
    <link rel="shortcut icon" href="%PUBLIC_URL%/todo-change.ico">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Arvo" rel="stylesheet">
    <title>App Name</title>
</head>
<body>
<noscript>
    You need to enable JavaScript to run this app.
</noscript>
<div id="root"></div>
</body>
</html>
`
}

func ManifestJson() string {
	return `{
  "short_name": "React App",
  "name": "Create React App Sample",
  "icons": [
    {
      "src": "todo-change.ico",
      "sizes": "192x192",
      "type": "image/jpeg"
    }
  ],
  "start_url": "./index.html",
  "display": "standalone",
  "theme_color": "#000000",
  "background_color": "#ffffff"
}
`
}

func RegisterServiceWorkerJs() string {
	return `// In production, we register a service worker to serve assets from local cache.

// This lets the app load faster on subsequent visits in production, and gives
// it offline capabilities. However, it also means that developers (and users)
// will only see deployed updates on the "N+1" visit to a page, since previously
// cached resources are updated in the background.

// To learn more about the benefits of this model, read https://goo.gl/KwvDNy.
// This link also includes instructions on opting out of this behavior.

const isLocalhost = Boolean(
  window.location.hostname === 'localhost' ||
    // [::1] is the IPv6 localhost address.
    window.location.hostname === '[::1]' ||
    // 127.0.0.1/8 is considered localhost for IPv4.
    window.location.hostname.match(
      /^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/
    )
);

export default function register() {
  if (process.env.NODE_ENV === 'production' && 'serviceWorker' in navigator) {
    // The URL constructor is available in all browsers that support SW.
    const publicUrl = new URL(process.env.PUBLIC_URL, window.location);
    if (publicUrl.origin !== window.location.origin) {
      // Our service worker won't work if PUBLIC_URL is on a different origin
      // from what our page is served on. This might happen if a CDN is used to
      // serve assets; see https://github.com/facebookincubator/create-react-app/issues/2374
      return;
    }

    window.addEventListener('load', () => {
      const swUrl = '${process.env.PUBLIC_URL}/service-worker.js';

      if (!isLocalhost) {
        // Is not local host. Just register service worker
        registerValidSW(swUrl);
      } else {
        // This is running on localhost. Lets check if a service worker still exists or not.
        checkValidServiceWorker(swUrl);
      }
    });
  }
}

function registerValidSW(swUrl) {
  navigator.serviceWorker
    .register(swUrl)
    .then((registration) => {
      registration.onupdatefound = () => {
        const installingWorker = registration.installing;
        installingWorker.onstatechange = () => {
          if (installingWorker.state === 'installed') {
            if (navigator.serviceWorker.controller) {
              // At this point, the old content will have been purged and
              // the fresh content will have been added to the cache.
              // It's the perfect time to display a "New content is
              // available; please refresh." message in your web app.
              console.log('New content is available; please refresh.');
            } else {
              // At this point, everything has been precached.
              // It's the perfect time to display a
              // "Content is cached for offline use." message.
              console.log('Content is cached for offline use.');
            }
          }
        };
      };
    })
    .catch((error) => {
      console.error('Error during service worker registration:', error);
    });
}

function checkValidServiceWorker(swUrl) {
  // Check if the service worker can be found. If it can't reload the page.
  fetch(swUrl)
    .then((response) => {
      // Ensure service worker exists, and that we really are getting a JS file.
      if (
        response.status === 404 ||
        response.headers.get('content-type').indexOf('javascript') === -1
      ) {
        // No service worker found. Probably a different app. Reload the page.
        navigator.serviceWorker.ready.then((registration) => {
          registration.unregister().then(() => {
            window.location.reload();
          });
        });
      } else {
        // Service worker found. Proceed as normal.
        registerValidSW(swUrl);
      }
    })
    .catch(() => {
      console.log(
        'No internet connection found. App is running in offline mode.'
      );
    });
}

export function unregister() {
  if ('serviceWorker' in navigator) {
    navigator.serviceWorker.ready.then((registration) => {
      registration.unregister();
    });
  }
}
`
}

func LogoSvg() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 841.9 595.3">
    <g fill="#61DAFB">
        <path d="M666.3 296.5c0-32.5-40.7-63.3-103.1-82.4 14.4-63.6 8-114.2-20.2-130.4-6.5-3.8-14.1-5.6-22.4-5.6v22.3c4.6 0 8.3.9 11.4 2.6 13.6 7.8 19.5 37.5 14.9 75.7-1.1 9.4-2.9 19.3-5.1 29.4-19.6-4.8-41-8.5-63.5-10.9-13.5-18.5-27.5-35.3-41.6-50 32.6-30.3 63.2-46.9 84-46.9V78c-27.5 0-63.5 19.6-99.9 53.6-36.4-33.8-72.4-53.2-99.9-53.2v22.3c20.7 0 51.4 16.5 84 46.6-14 14.7-28 31.4-41.3 49.9-22.6 2.4-44 6.1-63.6 11-2.3-10-4-19.7-5.2-29-4.7-38.2 1.1-67.9 14.6-75.8 3-1.8 6.9-2.6 11.5-2.6V78.5c-8.4 0-16 1.8-22.6 5.6-28.1 16.2-34.4 66.7-19.9 130.1-62.2 19.2-102.7 49.9-102.7 82.3 0 32.5 40.7 63.3 103.1 82.4-14.4 63.6-8 114.2 20.2 130.4 6.5 3.8 14.1 5.6 22.5 5.6 27.5 0 63.5-19.6 99.9-53.6 36.4 33.8 72.4 53.2 99.9 53.2 8.4 0 16-1.8 22.6-5.6 28.1-16.2 34.4-66.7 19.9-130.1 62-19.1 102.5-49.9 102.5-82.3zm-130.2-66.7c-3.7 12.9-8.3 26.2-13.5 39.5-4.1-8-8.4-16-13.1-24-4.6-8-9.5-15.8-14.4-23.4 14.2 2.1 27.9 4.7 41 7.9zm-45.8 106.5c-7.8 13.5-15.8 26.3-24.1 38.2-14.9 1.3-30 2-45.2 2-15.1 0-30.2-.7-45-1.9-8.3-11.9-16.4-24.6-24.2-38-7.6-13.1-14.5-26.4-20.8-39.8 6.2-13.4 13.2-26.8 20.7-39.9 7.8-13.5 15.8-26.3 24.1-38.2 14.9-1.3 30-2 45.2-2 15.1 0 30.2.7 45 1.9 8.3 11.9 16.4 24.6 24.2 38 7.6 13.1 14.5 26.4 20.8 39.8-6.3 13.4-13.2 26.8-20.7 39.9zm32.3-13c5.4 13.4 10 26.8 13.8 39.8-13.1 3.2-26.9 5.9-41.2 8 4.9-7.7 9.8-15.6 14.4-23.7 4.6-8 8.9-16.1 13-24.1zM421.2 430c-9.3-9.6-18.6-20.3-27.8-32 9 .4 18.2.7 27.5.7 9.4 0 18.7-.2 27.8-.7-9 11.7-18.3 22.4-27.5 32zm-74.4-58.9c-14.2-2.1-27.9-4.7-41-7.9 3.7-12.9 8.3-26.2 13.5-39.5 4.1 8 8.4 16 13.1 24 4.7 8 9.5 15.8 14.4 23.4zM420.7 163c9.3 9.6 18.6 20.3 27.8 32-9-.4-18.2-.7-27.5-.7-9.4 0-18.7.2-27.8.7 9-11.7 18.3-22.4 27.5-32zm-74 58.9c-4.9 7.7-9.8 15.6-14.4 23.7-4.6 8-8.9 16-13 24-5.4-13.4-10-26.8-13.8-39.8 13.1-3.1 26.9-5.8 41.2-7.9zm-90.5 125.2c-35.4-15.1-58.3-34.9-58.3-50.6 0-15.7 22.9-35.6 58.3-50.6 8.6-3.7 18-7 27.7-10.1 5.7 19.6 13.2 40 22.5 60.9-9.2 20.8-16.6 41.1-22.2 60.6-9.9-3.1-19.3-6.5-28-10.2zM310 490c-13.6-7.8-19.5-37.5-14.9-75.7 1.1-9.4 2.9-19.3 5.1-29.4 19.6 4.8 41 8.5 63.5 10.9 13.5 18.5 27.5 35.3 41.6 50-32.6 30.3-63.2 46.9-84 46.9-4.5-.1-8.3-1-11.3-2.7zm237.2-76.2c4.7 38.2-1.1 67.9-14.6 75.8-3 1.8-6.9 2.6-11.5 2.6-20.7 0-51.4-16.5-84-46.6 14-14.7 28-31.4 41.3-49.9 22.6-2.4 44-6.1 63.6-11 2.3 10.1 4.1 19.8 5.2 29.1zm38.5-66.7c-8.6 3.7-18 7-27.7 10.1-5.7-19.6-13.2-40-22.5-60.9 9.2-20.8 16.6-41.1 22.2-60.6 9.9 3.1 19.3 6.5 28.1 10.2 35.4 15.1 58.3 34.9 58.3 50.6-.1 15.7-23 35.6-58.4 50.6zM320.8 78.4z"/>
        <circle cx="420.9" cy="296.5" r="45.7"/>
        <path d="M520.5 78.1z"/>
    </g>
</svg>
`
}

func IndexJs() string {
	return `import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './components/App/App';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App />, document.getElementById('root')); 
registerServiceWorker();
`
}

func IndexCss() string {
	return `body {
  margin: 0;
  padding: 0;
  font-family: sans-serif;
}

.pagination {
  display: inline-block;
  padding-left: 0;
  margin: 20px 0;
  border-radius: 4px;
}
.pagination > li {
  display: inline;
}
.pagination > li > a,
.pagination > li > span {
  position: relative;
  float: left;
  padding: 6px 12px;
  margin-left: -1px;
  line-height: 1.42857143;
  color: #337ab7;
  text-decoration: none;
  background-color: #fff;
  border: 1px solid #ddd;
}
.pagination > li:first-child > a,
.pagination > li:first-child > span {
  margin-left: 0;
  border-top-left-radius: 4px;
  border-bottom-left-radius: 4px;
}
.pagination > li:last-child > a,
.pagination > li:last-child > span {
  border-top-right-radius: 4px;
  border-bottom-right-radius: 4px;
}
.pagination > li > a:hover,
.pagination > li > span:hover,
.pagination > li > a:focus,
.pagination > li > span:focus {
  z-index: 3;
  color: #23527c;
  background-color: #eee;
  border-color: #ddd;
}
.pagination > .active > a,
.pagination > .active > span,
.pagination > .active > a:hover,
.pagination > .active > span:hover,
.pagination > .active > a:focus,
.pagination > .active > span:focus {
  z-index: 2;
  color: #fff;
  cursor: default;
  background-color: #337ab7;
  border-color: #337ab7;
}
.pagination > .disabled > span,
.pagination > .disabled > span:hover,
.pagination > .disabled > span:focus,
.pagination > .disabled > a,
.pagination > .disabled > a:hover,
.pagination > .disabled > a:focus {
  color: #777;
  cursor: not-allowed;
  background-color: #fff;
  border-color: #ddd;
}
.pagination-lg > li > a,
.pagination-lg > li > span {
  padding: 10px 16px;
  font-size: 18px;
  line-height: 1.3333333;
}
.pagination-lg > li:first-child > a,
.pagination-lg > li:first-child > span {
  border-top-left-radius: 6px;
  border-bottom-left-radius: 6px;
}
.pagination-lg > li:last-child > a,
.pagination-lg > li:last-child > span {
  border-top-right-radius: 6px;
  border-bottom-right-radius: 6px;
}
.pagination-sm > li > a,
.pagination-sm > li > span {
  padding: 5px 10px;
  font-size: 12px;
  line-height: 1.5;
}
.pagination-sm > li:first-child > a,
.pagination-sm > li:first-child > span {
  border-top-left-radius: 3px;
  border-bottom-left-radius: 3px;
}
.pagination-sm > li:last-child > a,
.pagination-sm > li:last-child > span {
  border-top-right-radius: 3px;
  border-bottom-right-radius: 3px;
}

.pagination>.active>a{
  background-color: #00bcd4 !important;
  border-color: #00bcd4 !important;
}
.pagination>.active>a:hover{
  background-color: rgba(0, 188, 212, 0.64) !important;
  border-color: rgba(0, 188, 212, 0.64) !important;
}

.pagination {
  cursor: pointer;
}
.pagination>li>a, .pagination>li>span {
  color: #00bcd4;
}

.react-confirm-alert-button-group > button {
  background: #00bcd4 !important;
}
`
}

func AppTestJs() string {
	return `import ReactDOM from 'react-dom';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<App />, div);
});
`
}

func RestrictedJs() string {
	return `import React from 'react';

const unauthorizedPath = '/';

export default class RestrictedComponent extends React.Component {
  constructor(props) {
    super(props);

    this.roles = JSON.parse(localStorage.getItem('permissions'));
  }

  componentWillMount() {
    if (!this.isAuthorized()) {
      this.props.navigate(unauthorizedPath);
    }
  }

  isAuthorized() {
    return RestrictedComponent.isAuthorized(this.roles, this.props.needed);
  }

  static isAuthorized(roles, needed) {
    let flag = false;
    roles.forEach((rol) => {
      if (rol === needed) {
        flag = true;
        return;
      }
    });
    if (needed.length === 0) flag = true;
    return flag;
  }

  render() {
    return (
      <div>
        {this.isAuthorized() && this.props.children}
      </div>
    );
  }
}
`
}

func AppJs() string {
	return `import React from 'react';
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
                        <Route path="/home" component={Layout}/>

                        <Route component={Login}/>

                    </Switch>
                </BrowserRouter>
            </MuiThemeProvider>
        );
    }
}

export default App;
`
}

func AppCss() string {
	return `.App {
  text-align: center;
}

.App-header {
  background-color: #222;
  height: 150px;
  padding: 20px;
  color: white;
}

.user-button-icon {
  background-color: transparent;
  border: none;
  cursor: pointer; }
`
}

func AvatarJs() string {
	return `import React from 'react';
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
`
}

func DrawerJs() string {
	return `import React from 'react';
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
            	&&ROUTES_LIST&&
            </List>
          </MUIDrawer>
      );      
    }
  }
`
}

func LayoutJs() string {
	return `import React from 'react';
import {Redirect, Route, Switch} from 'react-router-dom';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import NavigationMenu from 'material-ui/svg-icons/navigation/menu';
import Avatar from './Avatar';
import Drawer from './Drawer';
import Main from './Main';
&&IMPORT_LIST&&

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
            	&&ROUTE_LIST&&

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
`
}

func MainJs() string {
	return `import React from 'react';

export default class Main extends React.Component {
  render() {
    return (
      <div></div>
    );
  }
}`
}

func LoginJs() string {
	return `import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import Loading from 'react-loading-spinner';
import CircularProgress from 'material-ui/CircularProgress';
import {white} from 'material-ui/styles/colors';
import Endpoint from './Api.js';
import {ValidatorForm} from 'react-form-validator-core';
import {TextValidator} from 'react-material-ui-form-validator';
import Dialog from 'material-ui/Dialog';
import {Redirect} from 'react-router-dom';

const spinnerStyle = {
    top: '50%',
    left: '50%',
    position: 'fixed',
};

const containerStyles = {
    loginContainer: {
        position: 'relative',
        height: '100vh',
        backgroundColor: 'rgb(250, 180, 200)',
    },
    loginPage: {
        height: '80%',
        position: 'absolute',
        top: '0',
        left: '0',
        right: '0',
        bottom: '0',
        textAlign: 'center',
        margin: 'auto',
    },
    titleLogin: {
        fontSize: '60px',
        color: 'white',
        marginTop: '20px',
        marginBottom: '20px',
        fontFamily: 'Arvo, serif',
    },
    subtitleLogin: {
        fontSize: '30px',
        color: 'white',
        marginTop: '20px',
        marginBottom: '20px',
        fontFamily: 'Arvo, serif',
    },
    imageLogin: {
        marginTop: '0px',
        marginBottom: '10px',
    },

    loginForm: {
        width: '60%',
        margin: 'auto',
        maxWidth: '450px',
    },

    logo: {
        width: '200px',
    },

    loginButton: {
        marginTop: '20px',
        marginBottom: '48px'
    },

    noAccountPrompt: {
        fontSize: '22px',
        color: 'white',
        fontFamily: 'Arvo, serif',
    },

};

const INITIAL_FORM = {
    username: '',
    password: '',
};

const INITIAL_STATE = {
    errorMessage: false,
    errorText: '',
    loading: false,
    form: INITIAL_FORM,
    authenticated: false,
};


class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleSubmit = this._handleSubmit.bind(this);
        this._handleClose = this._handleClose.bind(this);

        document.body.style.backgroundColor = "rgb(250, 180, 200)";
    }

    _handleSubmit() {
        let User = {
            username: this.state.form.username,
            password: this.state.form.password,
        };
        
        Endpoint.loginAccess(User)
            .then((response) => {
                localStorage.setItem('token', response.data);
                localStorage.setItem('username', this.state.form.username);
                let authenticated = true;
                this.setState({ authenticated });
            })
            .catch((error) => {
                if (error.response.status === 500) {
                    let errorText = 'Wrong credentials, please try again';
                    this.setState({errorMessage: true, errorText});
                }
            });
    }

    _handleClose = () => {
        this.setState({errorMessage: false});
    };

    isAuthenticated() {
        const token = localStorage.getItem('token');
        return token && token.length > 10;
    }

    render() {
        const isAlreadyAuthenticated = this.isAuthenticated();
        
        return (
            <div>
                {isAlreadyAuthenticated ? <Redirect to={{pathname: '/home'}}/> : (
                    <Loading isLoading={this.state.loading}
                         spinner={() => <CircularProgress id="loginSpinner" scale={1.5} centered={true} style={spinnerStyle} />}>
                    <div style={containerStyles.loginContainer}>
                        <div style={containerStyles.loginPage}>
                            <div style={containerStyles.titleLogin}>Nombre de la app</div>

                            {this.props.message &&
                            <p>{this.props.message}</p>
                            }
                            <div style={containerStyles.loginForm}>
                                <ValidatorForm
                                    ref="form"
                                    onSubmit={this._handleSubmit}
                                    onError={(errors) => console.log(errors)}
                                >
                                    <TextValidator
                                        floatingLabelText="Username"
                                        onChange={(e, value) => this.setState((prevState) => { return {form: {...prevState.form, username: value}}; })}
                                        name="username"
                                        value= {this.state.form.username}
                                        validators={['required']}
                                        errorMessages={['This field is required']}
                                    />
                                   <TextValidator
                                        floatingLabelText="Password"
                                        onChange={(e, value) => this.setState((prevState) => { return {form: {...prevState.form, password: value}}; })}
                                        type="password"
                                        name="password"
                                        value= {this.state.form.password}
                                        validators={['required']}
                                        errorMessages={['This field is required']}
                                    />
                                <br/>
                                <div>
                                    <RaisedButton
                                        style={containerStyles.loginButton}
                                        backgroundColor={white}
                                        labelColor={'#D50032'}
                                        label="Login"
                                        type="submit"
                                     />
                                </div>
                                </ValidatorForm>
                            </div>

                            <Dialog
                                title="Login Error!"
                                titleStyle={{textAlign: 'center'}}
                                modal={false}
                                open={this.state.errorMessage}
                                onRequestClose={this._handleClose}
                            >
                              {this.state.errorText}
                              <div
                                    style={{
                                    display: 'flex',
                                    justifyContent: 'flex-end',
                                }}>
                              <RaisedButton
                                        label="Cerrar"
                                        onClick={this._handleClose}
                                    />
                            </div>
                            </Dialog>
                        </div>
                    </div>
                </Loading>
                )}
           </div>
        );
    }
}

export default Login;
`
}

func LoginApiJs() string {
	return `import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    loginAccess(user) {
        return axios.post(URL_BASE + '/login', user, {});
    },
};
`
}

func MuiThemeJs() string {
	return `import getMuiTheme from 'material-ui/styles/getMuiTheme';
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
`
}

func ComponentJs() string {
	return `import React from "react";
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import { ValidatorForm } from 'react-form-validator-core';
import { TextValidator} from 'react-material-ui-form-validator';
import AddIcon from 'material-ui/svg-icons/content/add';
import NotificationSystem from 'react-notification-system';

import Endpoint from './Endpoints.js';
import &&COMPONENT_NAME&&Table from "./&&COMPONENT_NAME&&Table";

const styles = {
    buttonContainer: {
        position: "center",
        marginTop:"10px",
        marginBottom:"5px",
        textAlign: "center"
    },
    labelButton:{
        color: 'white'
    }
};

const customContentStyle = {
    width: '60%',
    maxWidth: 'none',
    borderRadius: '15px', 
    overflow: 'auto',
};

const INITIAL_FORM = {
    name: ''
};

const INITIAL_STATE = {
    &&COMPONENT_NAME&&: [],
    openDialog: false,
    form: INITIAL_FORM,
    fetchingData: false,
    errorMessage: false,
};

class &&COMPONENT_NAME&& extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        &&HANDLE_CHANGE_LIST&&
        this._handleSubmit = this._handleSubmit.bind(this);
        this._openDialog = this._openDialog.bind(this);
    }

    get&&COMPONENT_NAME&&(){
        this.setState({ fetchingData:true });
        Endpoint.getAll&&COMPONENT_NAME&&()
            .then(response => {
                let &&COMPONENT_NAME&& = response.data;
                this.setState({ &&COMPONENT_NAME&&, fetchingData: false });
            })
            .catch(error => {
                 this.setState({ fetchingData: false, errorMessage: true });
            })
    }

    componentDidMount(){
        this.get&&COMPONENT_NAME&&();
        this._notificationSystem = this.refs.notificationSystem;
    }

    _addNotification = (title, msg, level) =>{
        this._notificationSystem.addNotification({
            title: title,
            message: msg,
            level: level,
            position: 'tr'
        });
    };
   
    _handleSubmit() {
        let &&COMPONENT_NAME&& = {
            &&LIST_FIELDS_ASSIGNMENTS&&
        };

        Endpoint.create&&COMPONENT_NAME&&(&&COMPONENT_NAME&&)
            .then(response => {
                this.updateInfo();
                this._handleClose();
                this._addNotification('&&COMPONENT_NAME&&', 'The new &&COMPONENT_NAME&& has been created successfully', 'success');
            })
            .catch(error => {
                this._addNotification('&&COMPONENT_NAME&&', 'An error happened', 'error');
                this._handleClose();
            })
    }

    &&FIELDS_HANDLE_CHANGE_FUNCTIONS&&

    _openDialog = () => {
        this.setState({openDialog: true});
    };

    _handleClose = () => {
        this.setState({
            openDialog: false,
            form: INITIAL_FORM,
        });
    };

    _handleCloseError = () => {
        this.setState({
            errorMessage: false, 
        });
    };

    updateInfo = () =>{
        this.get&&COMPONENT_NAME&&();
    };

    render() {
        return (
            <div>
                <NotificationSystem ref="notificationSystem" />
                <div style={styles.buttonContainer}>
                   <RaisedButton 
                        backgroundColor={"#00bcd4"}
                        label="Create" 
                        onClick = {this._openDialog} 
                        labelStyle={styles.labelButton}
                        labelPosition="before" 
                        icon={<AddIcon />} 
                    >
                    </RaisedButton>
                     <Dialog
                        contentStyle={customContentStyle}
                        title="New &&COMPONENT_NAME&&"
                        titleStyle={{textAlign: "center"}}
                        modal={false}
                        open={this.state.openDialog}
                        onRequestClose={this._handleClose}
                    >
                        <div>
                            <ValidatorForm
                                ref="form"
                                onSubmit={this._handleSubmit}
                                onError={errors => console.log(errors)}
                            >   
                                <div>
                                    &&FIELDS_TEXT_VALIDATORS&&
                                </div>  
                                <div
                                    style={{
                                    display: 'flex',
                                    justifyContent: 'flex-end'
                                }}
                                >
                                    <RaisedButton
                                        label="Cancel"
                                        onClick={this._handleClose}
                                    />
                                    <RaisedButton
                                        style={styles.textAddButton}
                                        label="Create"
                                        type="submit"
                                        >
                                    </RaisedButton>
                                </div>
                            </ValidatorForm>
                        </div>
                    </Dialog>
                </div>
                 {this.state.&&COMPONENT_NAME&&.length > 0
                    ? <&&COMPONENT_NAME&&Table
                    &&COMPONENT_NAME&& = {this.state.&&COMPONENT_NAME&&}
                    updateInfo = {this.updateInfo}
                    fetchingData = {this.state.fetchingData}
                    _addNotification = {this._addNotification}
                />
                    : <div></div>}
            </div>
        );
    }
}

export default &&COMPONENT_NAME&&;`
}

func ComponentTableJs() string {
	return `import React from "react";
import Paper from 'material-ui/Paper';
import IconButton from 'material-ui/IconButton';
import Dialog from 'material-ui/Dialog';
import EditIcon from 'material-ui/svg-icons/editor/mode-edit';
import DeleteIcon from 'material-ui/svg-icons/action/delete';
import RaisedButton from 'material-ui/RaisedButton';
import { ValidatorForm } from 'react-form-validator-core';
import { TextValidator} from 'react-material-ui-form-validator';
import Loading from 'react-loading-spinner';
import CircularProgress from 'material-ui/CircularProgress';
import { confirmAlert } from 'react-confirm-alert';
import 'react-confirm-alert/src/react-confirm-alert.css'

import Endpoint from './Endpoints.js';
import Pagination from './Pagination';

import {
    Table,
    TableBody,
    TableHeader,
    TableHeaderColumn,
    TableRow,
    TableRowColumn
} from 'material-ui/Table';

const styleTable = {
    paddingLeft: 0,
    paddingRight: 0,
    textAlign: "center",
};
const spinnerStyle = {
    top: "50%",
    left: "50%",
    position: "fixed"
};
const styleRow = {
    textAlign: "center",
    overflow: 'visible',
};

const stylePaper = {
    height: '100%',
    width: '90%',
    marginLeft: '5%',
    marginRight: '5%',
    marginTop: '3%',
    paddingLeft:'3%',
    paddingRight: '3%',
    textAlign: 'center',
    display: 'inline-block',
};

const customContentStyle = {
    width: '60%',
    maxWidth: 'none',
    borderRadius: '15px', 
    overflow: 'auto',
};

const INITIAL_STATE = {
    pageOfItems: [],
    openDialog: false,
    &&FIELDS_LIST_INIT_STATE&&
    currentPage: 1,
};

class &&COMPONENT_NAME&&Table extends React.Component {

    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        &&HANDLE_CHANGE_LIST&&
        this._handleSubmit = this._handleSubmit.bind(this);
    }

    _renderHeaderTable = () =>{
        return(
            <TableHeader displaySelectAll={false} adjustForCheckbox={false}>
                <TableRow>
                      &&FIELDS_LIST_TABLE_HEADER&&
                     <TableHeaderColumn style={styleTable}>Acciones</TableHeaderColumn>
                </TableRow>
            </TableHeader>
        )
    };

    &&FIELDS_HANDLE_CHANGE_FUNCTIONS&&

    _handleSubmit() {
        let &&COMPONENT_NAME&& = {
            &&LIST_FIELDS_ASSIGNMENTS&&
        };
        let id&&COMPONENT_NAME&& = this.state.id;
       
        Endpoint.edit&&COMPONENT_NAME&&(id&&COMPONENT_NAME&&, &&COMPONENT_NAME&&)
            .then(response => {
                this.props._addNotification('Clase', 'The &&COMPONENT_NAME&& has been updated', 'success');
                this.props.updateInfo();
                this._handleClose();
            })
            .catch(error => {
                 this._handleClose();
                 this.props._addNotification('&&COMPONENT_NAME&&', 'Error updating this &&COMPONENT_NAME&&', 'error');
        })
    }

    delete&&COMPONENT_NAME&& = (id&&COMPONENT_NAME&&) =>{
        Endpoint.delete&&COMPONENT_NAME&&(id&&COMPONENT_NAME&&)
            .then(response => {
                this.props._addNotification('&&COMPONENT_NAME&&', 'The &&COMPONENT_NAME&& has been deleted', 'success');
                this.props.updateInfo();
            })
            .catch(error => {
                this._handleClose();
                this.props._addNotification('&&COMPONENT_NAME&&', 'Error deleting this &&COMPONENT_NAME&&', 'error');
            })
    };

    submit = (&&FIELDS_PARAMETER_LIST&&) => {
        confirmAlert({
            title: 'Delete &&COMPONENT_NAME&&',
            message: 'Do you really wish to delete this &&COMPONENT_NAME&&?',
            confirmLabel: 'Confirm',
            cancelLabel: 'Cancel',
            onConfirm: () => this.delete&&COMPONENT_NAME&&(id&&COMPONENT_NAME&&)
        })
    };

    showInfo&&COMPONENT_NAME&& = (&&FIELDS_PARAMETER_LIST&&) => {
        this.setState({
            &&FIELDS_SET_STATE_LIST&&
            openDialog: true
        });
    };
     
    _renderData = () =>{
        return(
            this.state.pageOfItems.map((row, index) =>(
                <TableRow key={index}>
                    &&FIELDS_TABLE_ROWS&&
                    <TableRowColumn style={styleRow}>
                        <IconButton tooltip="Edit" tooltipPosition = "top-left" onClick={() =>this.showInfo&&COMPONENT_NAME&&(&&FIELDS_ROW_LIST&&)}>
                            <EditIcon/>
                        </IconButton>
                        <IconButton tooltip="Delete" tooltipPosition = "top-right" onClick={() =>this.submit(&&FIELDS_ROW_LIST&&)}>
                            <DeleteIcon/>
                        </IconButton>
                    </TableRowColumn>
                </TableRow>
            ))
        )
    };

    onChangePage = (pageOfItems, currentPage) =>{
        this.setState({
            pageOfItems: pageOfItems,
            currentPage
        });
    };
    
    _handleClose = () => {
        this.setState({
            openDialog: false,
        });
    };

    render() {
        return (
            <Paper style={stylePaper} zDepth={1} >
                <Loading isLoading={this.props.fetchingData}
                         spinner={() => <CircularProgress id="loginSpinner" scale={1.5}
                                                      style={spinnerStyle}/>}>
                    <div>
                        <Table wrapperStyle={{overflow: 'visible'}} bodyStyle={{overflow: 'visible'}}>
                            {this._renderHeaderTable()}
                            <TableBody displayRowCheckbox={false}>
                                {this._renderData()}
                            </TableBody>
                        </Table>
                        <Pagination
                            items={this.props.&&COMPONENT_NAME&&}
                            onChangePage={this.onChangePage}
                            initialPage = {this.state.currentPage}
                        />
                    </div>
                </Loading>
                 <Dialog
                        contentStyle={customContentStyle}
                        title="Edit &&COMPONENT_NAME&&"
                        titleStyle={{textAlign: "center"}}
                        modal={false}
                        open={this.state.openDialog}
                        onRequestClose={this._handleClose}
                    >
                        <div >
                            <ValidatorForm
                                ref="form"
                                onSubmit={this._handleSubmit}
                                onError={errors => console.log(errors)}
                            >   
                                <div>
                                    &&FIELDS_TEXT_VALIDATORS&&
                                </div>  
                                <div
                                    style={{
                                    display: 'flex',
                                    justifyContent: 'flex-end'
                                }}
                                >
                                    <RaisedButton
                                        label="Cancel"
                                        onClick={this._handleClose}
                                    />
                                    <RaisedButton
                                        label="Edit"
                                        type="submit" />
                                </div>
                            </ValidatorForm>
                        </div>
                    </Dialog>
            </Paper>
        );
    }
}

export default &&COMPONENT_NAME&&Table;
`
}

func ComponentPaginationJs() string {
	return `import React  from 'react';
import PropTypes from 'prop-types';
import _ from 'lodash'

class Pagination extends React.Component {

    constructor(props) {
        super(props);
        this.state = { pager: {} };
    }

    componentWillMount() {
        // set page if items array isn't empty
        if (this.props.items && this.props.items.length) {
            this.setPage(this.props.initialPage);
        }
    }

    componentDidUpdate(prevProps, prevState) {
        // reset page if items array has changed
        if (this.props.items !== prevProps.items) {
            this.setPage(this.props.initialPage);
        }
    }

    setPage(page) {
        var items = this.props.items;
        var pager = this.state.pager;

        if (page < 1 || page > pager.totalPages) {
            return;
        }

        // get new pager object for specified page
        pager = this.getPager(items.length, page);

        // get new page of items from items array
        var pageOfItems = items.slice(pager.startIndex, pager.endIndex + 1);

        // update state
        this.setState({ pager: pager });

        // call change page function in parent component
        this.props.onChangePage(pageOfItems,page);
    }

    getPager(totalItems, currentPage, pageSize) {
        // default to first page
        currentPage = currentPage || 1;

        // default page size is 10
        pageSize = pageSize || 10;

        // calculate total pages
        var totalPages = Math.ceil(totalItems / pageSize);

        var startPage, endPage;
        if (totalPages <= 10) {
            // less than 10 total pages so show all
            startPage = 1;
            endPage = totalPages;
        } else {
            // more than 10 total pages so calculate start and end pages
            if (currentPage <= 6) {
                startPage = 1;
                endPage = 10;
            } else if (currentPage + 4 >= totalPages) {
                startPage = totalPages - 9;
                endPage = totalPages;
            } else {
                startPage = currentPage - 5;
                endPage = currentPage + 4;
            }
        }

        // calculate start and end item indexes
        var startIndex = (currentPage - 1) * pageSize;
        var endIndex = Math.min(startIndex + pageSize - 1, totalItems - 1);

        // create an array of pages to ng-repeat in the pager control
        var pages = _.range(startPage, endPage + 1);

        // return object with all pager properties required by the view
        return {
            totalItems: totalItems,
            currentPage: currentPage,
            pageSize: pageSize,
            totalPages: totalPages,
            startPage: startPage,
            endPage: endPage,
            startIndex: startIndex,
            endIndex: endIndex,
            pages: pages
        };
    }

    render() {
        var pager = this.state.pager;

        if (!pager.pages || pager.pages.length <= 1) {
            // don't display pager if there is only 1 page
            return null;
        }

        return (
            <ul className="pagination">
                <li className={pager.currentPage === 1 ? 'disabled' : ''}>
                    <a onClick={() => this.setPage(1)}>Primero</a>
                </li>
                <li className={pager.currentPage === 1 ? 'disabled' : ''}>
                    <a onClick={() => this.setPage(pager.currentPage - 1)}>Ant.</a>
                </li>
                {pager.pages.map((page, index) =>
                    <li key={index} className={pager.currentPage === page ? 'active' : ''}>
                        <a onClick={() => this.setPage(page)}>{page}</a>
                    </li>
                )}
                <li className={pager.currentPage === pager.totalPages ? 'disabled' : ''}>
                    <a onClick={() => this.setPage(pager.currentPage + 1)}>Siq.</a>
                </li>
                <li className={pager.currentPage === pager.totalPages ? 'disabled' : ''}>
                    <a onClick={() => this.setPage(pager.totalPages)}>Último</a>
                </li>
            </ul>
        );
    }
}

Pagination.defaultProps = {
    initialPage: 1
};

const propTypes = {
    items: PropTypes.array.isRequired,
    onChangePage: PropTypes.func.isRequired,
    initialPage: PropTypes.number
};
Pagination.propTypes = propTypes;

export default Pagination;`

}

func ComponentEndpointsJs() string {
	return `import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    &&CREATE_FUNCTION&&
    &&EDIT_FUNCTION&&
    &&DELETE_FUNCTION&&
    &&GET_ALL_FUNCTION&&
}`
}
