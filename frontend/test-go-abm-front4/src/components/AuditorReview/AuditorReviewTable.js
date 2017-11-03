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
    idUser: 0,
    idRefuge: 0,
    idRefugeAbility: 0,
    score: 0,

    currentPage: 1,
};

class AuditorReviewTable extends React.Component {

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
    <TableHeaderColumn style={styleTable}>IdUser</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>IdRefuge</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>IdRefugeAbility</TableHeaderColumn>
    <TableHeaderColumn style={styleTable}>Score</TableHeaderColumn>

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
        let AuditorReview = {
            name: this.state.name,
        };
        let idAuditorReview = this.state.id;
       
        Endpoint.editAuditorReview(idAuditorReview, AuditorReview)
            .then(response => {
                this.props._addNotification('Clase', 'The AuditorReview has been updated', 'success');
                this.props.updateInfo();
                this._handleClose();
            })
            .catch(error => {
                 this._handleClose();
                 this.props._addNotification('AuditorReview', 'Error updating this AuditorReview', 'error');
        })
    }

    deleteAuditorReview = (idAuditorReview) =>{
        Endpoint.deleteAuditorReview(idAuditorReview)
            .then(response => {
                this.props._addNotification('AuditorReview', 'The AuditorReview has been deleted', 'success');
                this.props.updateInfo();
            })
            .catch(error => {
                this._handleClose();
                this.props._addNotification('AuditorReview', 'Error deleting this AuditorReview', 'error');
            })
    };

    submit = (idAuditorReview, idUserAuditorReview, idRefugeAuditorReview, idRefugeAbilityAuditorReview, scoreAuditorReview) => {
        confirmAlert({
            title: 'Delete AuditorReview',
            message: 'Do you really wish to delete this AuditorReview?',
            confirmLabel: 'Confirm',
            cancelLabel: 'Cancel',
            onConfirm: () => this.deleteAuditorReview(idAuditorReview)
        })
    };

    showInfoClass = (idAuditorReview, idUserAuditorReview, idRefugeAuditorReview, idRefugeAbilityAuditorReview, scoreAuditorReview) => {
        this.setState({
            id: idAuditorReview,
idUser: idUserAuditorReview,
idRefuge: idRefugeAuditorReview,
idRefugeAbility: idRefugeAbilityAuditorReview,
score: scoreAuditorReview,

            openDialog: true
        });
    };
     
    _renderData = () =>{
        return(
            this.state.pageOfItems.map((row, index) =>(
                <TableRow key={index}>
                    <TableRowColumn style={styleRow}>{row.Id}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.IdUser}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.IdRefuge}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.IdRefugeAbility}</TableRowColumn>
<TableRowColumn style={styleRow}>{row.Score}</TableRowColumn>

                    <TableRowColumn style={styleRow}>
                        <IconButton tooltip="Edit" tooltipPosition = "top-left" onClick={() =>this.showInfoAuditorReview(row.Id, row.IdUser, row.IdRefuge, row.IdRefugeAbility, row.Score)}>
                            <EditIcon/>
                        </IconButton>
                        <IconButton tooltip="Delete" tooltipPosition = "top-right" onClick={() =>this.submit(row.Id, row.IdUser, row.IdRefuge, row.IdRefugeAbility, row.Score)}>
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
                            items={this.props.AuditorReview}
                            onChangePage={this.onChangePage}
                            initialPage = {this.state.currentPage}
                        />
                    </div>
                </Loading>
                 <Dialog
                        contentStyle={customContentStyle}
                        title="Edit AuditorReview"
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
    floatingLabelText='Id'
    onChange={this._handleChange}
    name='id'
    value={this.state.form.id}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='IdUser'
    onChange={this._handleChange}
    name='idUser'
    value={this.state.form.idUser}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='IdRefuge'
    onChange={this._handleChange}
    name='idRefuge'
    value={this.state.form.idRefuge}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='IdRefugeAbility'
    onChange={this._handleChange}
    name='idRefugeAbility'
    value={this.state.form.idRefugeAbility}
    validators={['required']}
    errorMessages={['This field is required']}
/>
<TextValidator
    floatingLabelText='Score'
    onChange={this._handleChange}
    name='score'
    value={this.state.form.score}
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

export default AuditorReviewTable;
