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
    refugeId: 0,
    familyId: 0,
    breed: '',
    color: '',
    age: 0,

    currentPage: 1,
};

class PetTable extends React.Component {

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
    <TableHeaderColumn style={styleTable}>RefugeId</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>FamilyId</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Breed</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Color</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Age</TableHeaderColumn>

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
        let Pet = {
            name: this.state.name,
        };
        let idPet = this.state.id;
       
        Endpoint.editPet(idPet, Pet)
            .then(response => {
                this.props._addNotification('Clase', 'The Pet has been updated', 'success');
                this.props.updateInfo();
                this._handleClose();
            })
            .catch(error => {
                 this._handleClose();
                 this.props._addNotification('Pet', 'Error updating this Pet', 'error');
        })
    }

    deletePet = (idPet) =>{
        Endpoint.deletePet(idPet)
            .then(response => {
                this.props._addNotification('Pet', 'The Pet has been deleted', 'success');
                this.props.updateInfo();
            })
            .catch(error => {
                this._handleClose();
                this.props._addNotification('Pet', 'Error deleting this Pet', 'error');
            })
    };

    submit = (idPet, refugeIdPet, familyIdPet, breedPet, colorPet, agePet) => {
        confirmAlert({
            title: 'Delete Pet',
            message: 'Do you really wish to delete this Pet?',
            confirmLabel: 'Confirm',
            cancelLabel: 'Cancel',
            onConfirm: () => this.deletePet(idPet)
        })
    };

    showInfoPet = (idPet, refugeIdPet, familyIdPet, breedPet, colorPet, agePet) => {
        this.setState({
            id: idPet,
refugeId: refugeIdPet,
familyId: familyIdPet,
breed: breedPet,
color: colorPet,
age: agePet,

            openDialog: true
        });
    };
     
    _renderData = () =>{
        return(
            this.state.pageOfItems.map((row, index) =>(
                <TableRow key={index}>
                    <TableRowColumn style={styleRow}>{row.Id}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.RefugeId}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.FamilyId}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Breed}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Color}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Age}</TableRowColumn>

                    <TableRowColumn style={styleRow}>
                        <IconButton tooltip="Edit" tooltipPosition = "top-left" onClick={() =>this.showInfoPet(row.Id, row.RefugeId, row.FamilyId, row.Breed, row.Color, row.Age)}>
                            <EditIcon/>
                        </IconButton>
                        <IconButton tooltip="Delete" tooltipPosition = "top-right" onClick={() =>this.submit(row.Id, row.RefugeId, row.FamilyId, row.Breed, row.Color, row.Age)}>
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
                            items={this.props.Pet}
                            onChangePage={this.onChangePage}
                            initialPage = {this.state.currentPage}
                        />
                    </div>
                </Loading>
                 <Dialog
                        contentStyle={customContentStyle}
                        title="Edit Pet"
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
    floatingLabelText='RefugeId'
    onChange={this._handleChange}
    name='refugeId'
    value={this.state.form.refugeId}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='FamilyId'
    onChange={this._handleChange}
    name='familyId'
    value={this.state.form.familyId}
/>
<TextValidator
    floatingLabelText='Breed'
    onChange={this._handleChange}
    name='breed'
    value={this.state.form.breed}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='Color'
    onChange={this._handleChange}
    name='color'
    value={this.state.form.color}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='Age'
    onChange={this._handleChange}
    name='age'
    value={this.state.form.age}
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

export default PetTable;
