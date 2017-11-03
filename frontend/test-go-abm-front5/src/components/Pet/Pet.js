import React from "react";
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import { ValidatorForm } from 'react-form-validator-core';
import { TextValidator} from 'react-material-ui-form-validator';
import AddIcon from 'material-ui/svg-icons/content/add';
import NotificationSystem from 'react-notification-system';

import Endpoint from './Endpoints.js';
import PetTable from "./PetTable";

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
    Pet: [],
    openDialog: false,
    form: INITIAL_FORM,
    fetchingData: false,
    errorMessage: false,
};

class Pet extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleChange = this._handleChange.bind(this);
        this._handleSubmit = this._handleSubmit.bind(this);
        this._openDialog = this._openDialog.bind(this);
    }

    getPet(){
        this.setState({ fetchingData:true });
        Endpoint.getAllPet()
            .then(response => {
                let Pet = response.data;
                this.setState({ Pet, fetchingData: false });
            })
            .catch(error => {
                 this.setState({ fetchingData: false, errorMessage: true });
            })
    }

    componentDidMount(){
        this.getPet();
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
        let Pet = {
            name: this.state.form.name,
        };

        Endpoint.createPet(Pet)
            .then(response => {
                this.updateInfo();
                this._handleClose();
                this._addNotification('Pet', 'The new Pet has been created successfully', 'success');
            })
            .catch(error => {
                this._addNotification('Pet', 'An error happened', 'error');
                this._handleClose();
            })
    }

    _handleChange = (event) => {
        const target = event.target;
        const value = target.value;
        const name = target.name;

        this.setState((prevState)=>{
            return {
                ...prevState,
                form: {
                    ...prevState.form,
                    [name]: value
                }
            }
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
        this.getPet();
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
                        title="New Pet"
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
                 {this.state.Pet.length > 0
                    ? <PetTable
                    Pet = {this.state.Pet}
                    updateInfo = {this.updateInfo}
                    fetchingData = {this.state.fetchingData}
                    _addNotification = {this._addNotification}
                />
                    : <div></div>}
            </div>
        );
    }
}

export default Pet;