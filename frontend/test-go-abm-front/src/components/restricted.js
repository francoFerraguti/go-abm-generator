import React from 'react';

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
