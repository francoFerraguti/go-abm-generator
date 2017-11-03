import React from "react";
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
        id: 0,
    username: '',
    password: '',
    refugeName: '',
    email: '',
    roleID: 0,
    lat: '',
    lon: '',

    currentPage: 1,
};

class UserTable extends React.Component {

    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleChange = this._handleChange.bind(this);
        this._handleSubmit = this._handleSubmit.bind(this);
    }

    _renderHeaderTable = () =>{
        return(
            <TableHeader displaySelectAll={false} adjustForCheckbox={false}>
                <TableRow>
                          <TableHeaderColumn style={styleTable}>Id</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Username</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Password</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>RefugeName</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Email</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>RoleID</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Lat</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Lon</TableHeaderColumn>

                     <TableHeaderColumn style={styleTable}>Acciones</TableHeaderColumn>
                </TableRow>
            </TableHeader>
        )
    };

    _handleChange = (event) =>{
        const target = event.target;
        const value = target.value;

        this.setState({
            name: value,
        });
    };

    _handleSubmit() {
        let User = {
            name: this.state.name,
        };
        let idUser = this.state.id;
       
        Endpoint.editUser(idUser, User)
            .then(response => {
                this.props._addNotification('Clase', 'The User has been updated', 'success');
                this.props.updateInfo();
                this._handleClose();
            })
            .catch(error => {
                 this._handleClose();
                 this.props._addNotification('User', 'Error updating this User', 'error');
        })
    }

    deleteUser = (idUser) =>{
        Endpoint.deleteUser(idUser)
            .then(response => {
                this.props._addNotification('User', 'The User has been deleted', 'success');
                this.props.updateInfo();
            })
            .catch(error => {
                this._handleClose();
                this.props._addNotification('User', 'Error deleting this User', 'error');
            })
    };

    submit = (idUser, usernameUser, passwordUser, refugeNameUser, emailUser, roleIDUser, latUser, lonUser) => {
        confirmAlert({
            title: 'Delete User',
            message: 'Do you really wish to delete this User?',
            confirmLabel: 'Confirm',
            cancelLabel: 'Cancel',
            onConfirm: () => this.deleteUser(idUser)
        })
    };

    showInfoUser = (idUser, usernameUser, passwordUser, refugeNameUser, emailUser, roleIDUser, latUser, lonUser) => {
        this.setState({
            id: idUser,
username: usernameUser,
password: passwordUser,
refugeName: refugeNameUser,
email: emailUser,
roleID: roleIDUser,
lat: latUser,
lon: lonUser,

            openDialog: true
        });
    };
     
    _renderData = () =>{
        return(
            this.state.pageOfItems.map((row, index) =>(
                <TableRow key={index}>
                    <TableRowColumn style={styleRow}>{row.Id}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Username}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Password}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.RefugeName}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Email}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.RoleID}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Lat}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Lon}</TableRowColumn>

                    <TableRowColumn style={styleRow}>
                        <IconButton tooltip="Edit" tooltipPosition = "top-left" onClick={() =>this.showInfoUser(row.Id, row.Username, row.Password, row.RefugeName, row.Email, row.RoleID, row.Lat, row.Lon)}>
                            <EditIcon/>
                        </IconButton>
                        <IconButton tooltip="Delete" tooltipPosition = "top-right" onClick={() =>this.submit(row.Id, row.Username, row.Password, row.RefugeName, row.Email, row.RoleID, row.Lat, row.Lon)}>
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
                            items={this.props.User}
                            onChangePage={this.onChangePage}
                            initialPage = {this.state.currentPage}
                        />
                    </div>
                </Loading>
                 <Dialog
                        contentStyle={customContentStyle}
                        title="Edit User"
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
                                    <TextValidator
    floatingLabelText='Username'
    onChange={this._handleChange}
    name='username'
    value={this.state.username}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='Password'
    onChange={this._handleChange}
    name='password'
    value={this.state.password}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='RefugeName'
    onChange={this._handleChange}
    name='refugeName'
    value={this.state.refugeName}
/>
<TextValidator
    floatingLabelText='Email'
    onChange={this._handleChange}
    name='email'
    value={this.state.email}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='RoleID'
    onChange={this._handleChange}
    name='roleID'
    value={this.state.roleID}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='Lat'
    onChange={this._handleChange}
    name='lat'
    value={this.state.lat}
/>
<TextValidator
    floatingLabelText='Lon'
    onChange={this._handleChange}
    name='lon'
    value={this.state.lon}
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

export default UserTable;
