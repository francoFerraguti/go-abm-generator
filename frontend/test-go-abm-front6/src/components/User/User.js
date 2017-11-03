import React from "react";
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import { ValidatorForm } from 'react-form-validator-core';
import { TextValidator} from 'react-material-ui-form-validator';
import AddIcon from 'material-ui/svg-icons/content/add';
import NotificationSystem from 'react-notification-system';

import Endpoint from './Endpoints.js';
import UserTable from "./UserTable";

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
    User: [],
    openDialog: false,
    form: INITIAL_FORM,
    fetchingData: false,
    errorMessage: false,
};

class User extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleChangeUsername = this._handleChangeUsername.bind(this);
        this._handleChangePassword = this._handleChangePassword.bind(this);
        this._handleChangeRefugeName = this._handleChangeRefugeName.bind(this);
        this._handleChangeRoleID = this._handleChangeRoleID.bind(this);
        this._handleChangeLat = this._handleChangeLat.bind(this);
        this._handleChangeLon = this._handleChangeLon.bind(this);

        this._handleSubmit = this._handleSubmit.bind(this);
        this._openDialog = this._openDialog.bind(this);
    }

    getUser(){
        this.setState({ fetchingData:true });
        Endpoint.getAllUser()
            .then(response => {
                let User = response.data;
                this.setState({ User, fetchingData: false });
            })
            .catch(error => {
                 this.setState({ fetchingData: false, errorMessage: true });
            })
    }

    componentDidMount(){
        this.getUser();
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
        let User = {
            username: this.state.username,
password: this.state.password,
refugeName: this.state.refugeName,
email: this.state.email,
roleID: this.state.roleID,
lat: this.state.lat,
lon: this.state.lon,

        };

        Endpoint.createUser(User)
            .then(response => {
                this.updateInfo();
                this._handleClose();
                this._addNotification('User', 'The new User has been created successfully', 'success');
            })
            .catch(error => {
                this._addNotification('User', 'An error happened', 'error');
                this._handleClose();
            })
    }

    _handleChangeUsername = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        username: value
    });
};
_handleChangePassword = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        password: value
    });
};
_handleChangeRefugeName = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        refugeName: value
    });
};
_handleChangeEmail = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        email: value
    });
};
_handleChangeRoleID = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        roleID: parseInt(value)
    });
};
_handleChangeLat = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        lat: parseFloat(value)
    });
};
_handleChangeLon = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        lon: parseFloat(value)
    });
};


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
        this.getUser();
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
                        title="New User"
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
                                    <TextValidator
    floatingLabelText='Username'
    onChange={this._handleChangeUsername}
    name='username'
    value={this.state.form.username}
/>
<TextValidator
    floatingLabelText='Password'
    onChange={this._handleChangePassword}
    name='password'
    value={this.state.form.password}
/>
<TextValidator
    floatingLabelText='RefugeName'
    onChange={this._handleChangeRefugeName}
    name='refugeName'
    value={this.state.form.refugeName}
/>
<TextValidator
    floatingLabelText='Email'
    onChange={this._handleChangeEmail}
    name='email'
    value={this.state.form.email}
/>
<TextValidator
    floatingLabelText='RoleID'
    onChange={this._handleChangeRoleID}
    name='roleID'
    value={this.state.form.roleID}
/>
<TextValidator
    floatingLabelText='Lat'
    onChange={this._handleChangeLat}
    name='lat'
    value={this.state.form.lat}
/>
<TextValidator
    floatingLabelText='Lon'
    onChange={this._handleChangeLon}
    name='lon'
    value={this.state.form.lon}
/>

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
                 {this.state.User.length > 0
                    ? <UserTable
                    User = {this.state.User}
                    updateInfo = {this.updateInfo}
                    fetchingData = {this.state.fetchingData}
                    _addNotification = {this._addNotification}
                />
                    : <div></div>}
            </div>
        );
    }
}

export default User;