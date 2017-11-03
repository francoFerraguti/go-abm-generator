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
    url: '',
    lat: '',
    lon: '',

    currentPage: 1,
};

class PhotoTable extends React.Component {

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
    <TableHeaderColumn style={styleTable}>Url</TableHeaderColumn>
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
        let Photo = {
            name: this.state.name,
        };
        let idPhoto = this.state.id;
       
        Endpoint.editPhoto(idPhoto, Photo)
            .then(response => {
                this.props._addNotification('Clase', 'The Photo has been updated', 'success');
                this.props.updateInfo();
                this._handleClose();
            })
            .catch(error => {
                 this._handleClose();
                 this.props._addNotification('Photo', 'Error updating this Photo', 'error');
        })
    }

    deletePhoto = (idPhoto) =>{
        Endpoint.deletePhoto(idPhoto)
            .then(response => {
                this.props._addNotification('Photo', 'The Photo has been deleted', 'success');
                this.props.updateInfo();
            })
            .catch(error => {
                this._handleClose();
                this.props._addNotification('Photo', 'Error deleting this Photo', 'error');
            })
    };

    submit = (idPhoto, urlPhoto, latPhoto, lonPhoto) => {
        confirmAlert({
            title: 'Delete Photo',
            message: 'Do you really wish to delete this Photo?',
            confirmLabel: 'Confirm',
            cancelLabel: 'Cancel',
            onConfirm: () => this.deletePhoto(idPhoto)
        })
    };

    showInfoPhoto = (idPhoto, urlPhoto, latPhoto, lonPhoto) => {
        this.setState({
            id: idPhoto,
url: urlPhoto,
lat: latPhoto,
lon: lonPhoto,

            openDialog: true
        });
    };
     
    _renderData = () =>{
        return(
            this.state.pageOfItems.map((row, index) =>(
                <TableRow key={index}>
                    <TableRowColumn style={styleRow}>{row.Id}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Url}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Lat}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Lon}</TableRowColumn>

                    <TableRowColumn style={styleRow}>
                        <IconButton tooltip="Edit" tooltipPosition = "top-left" onClick={() =>this.showInfoPhoto(row.Id, row.Url, row.Lat, row.Lon)}>
                            <EditIcon/>
                        </IconButton>
                        <IconButton tooltip="Delete" tooltipPosition = "top-right" onClick={() =>this.submit(row.Id, row.Url, row.Lat, row.Lon)}>
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
                            items={this.props.Photo}
                            onChangePage={this.onChangePage}
                            initialPage = {this.state.currentPage}
                        />
                    </div>
                </Loading>
                 <Dialog
                        contentStyle={customContentStyle}
                        title="Edit Photo"
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
    floatingLabelText='Url'
    onChange={this._handleChange}
    name='url'
    value={this.state.form.url}
/>
<TextValidator
    floatingLabelText='Lat'
    onChange={this._handleChange}
    name='lat'
    value={this.state.form.lat}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='Lon'
    onChange={this._handleChange}
    name='lon'
    value={this.state.form.lon}
    validators={['required']}
    errorMessages={['This field is required']}
/>

                                    <TextValidator
                                        floatingLabelText="Nombre"
                                        onChange={this._handleChange}
                                        name="name"
                                        value = {this.state.name}
                                        validators={['required']}
                                        errorMessages={['Este campo es requerido']}
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

export default PhotoTable;
