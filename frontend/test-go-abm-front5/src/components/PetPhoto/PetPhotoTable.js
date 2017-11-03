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
    idPet: 0,
    idPhoto: 0,

    currentPage: 1,
};

class PetPhotoTable extends React.Component {

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
    <TableHeaderColumn style={styleTable}>IdPet</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>IdPhoto</TableHeaderColumn>

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
        let PetPhoto = {
            name: this.state.name,
        };
        let idPetPhoto = this.state.id;
       
        Endpoint.editPetPhoto(idPetPhoto, PetPhoto)
            .then(response => {
                this.props._addNotification('Clase', 'The PetPhoto has been updated', 'success');
                this.props.updateInfo();
                this._handleClose();
            })
            .catch(error => {
                 this._handleClose();
                 this.props._addNotification('PetPhoto', 'Error updating this PetPhoto', 'error');
        })
    }

    deletePetPhoto = (idPetPhoto) =>{
        Endpoint.deletePetPhoto(idPetPhoto)
            .then(response => {
                this.props._addNotification('PetPhoto', 'The PetPhoto has been deleted', 'success');
                this.props.updateInfo();
            })
            .catch(error => {
                this._handleClose();
                this.props._addNotification('PetPhoto', 'Error deleting this PetPhoto', 'error');
            })
    };

    submit = (idPetPhoto, idPetPetPhoto, idPhotoPetPhoto) => {
        confirmAlert({
            title: 'Delete PetPhoto',
            message: 'Do you really wish to delete this PetPhoto?',
            confirmLabel: 'Confirm',
            cancelLabel: 'Cancel',
            onConfirm: () => this.deletePetPhoto(idPetPhoto)
        })
    };

    showInfoPetPhoto = (idPetPhoto, idPetPetPhoto, idPhotoPetPhoto) => {
        this.setState({
            id: idPetPhoto,
idPet: idPetPetPhoto,
idPhoto: idPhotoPetPhoto,

            openDialog: true
        });
    };
     
    _renderData = () =>{
        return(
            this.state.pageOfItems.map((row, index) =>(
                <TableRow key={index}>
                    <TableRowColumn style={styleRow}>{row.Id}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.IdPet}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.IdPhoto}</TableRowColumn>

                    <TableRowColumn style={styleRow}>
                        <IconButton tooltip="Edit" tooltipPosition = "top-left" onClick={() =>this.showInfoPetPhoto(row.Id, row.IdPet, row.IdPhoto)}>
                            <EditIcon/>
                        </IconButton>
                        <IconButton tooltip="Delete" tooltipPosition = "top-right" onClick={() =>this.submit(row.Id, row.IdPet, row.IdPhoto)}>
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
                            items={this.props.PetPhoto}
                            onChangePage={this.onChangePage}
                            initialPage = {this.state.currentPage}
                        />
                    </div>
                </Loading>
                 <Dialog
                        contentStyle={customContentStyle}
                        title="Edit PetPhoto"
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
    floatingLabelText='IdPet'
    onChange={this._handleChange}
    name='idPet'
    value={this.state.form.idPet}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='IdPhoto'
    onChange={this._handleChange}
    name='idPhoto'
    value={this.state.form.idPhoto}
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

export default PetPhotoTable;
