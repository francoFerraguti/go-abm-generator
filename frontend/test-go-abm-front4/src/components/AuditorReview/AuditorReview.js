import React from "react";
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import { ValidatorForm } from 'react-form-validator-core';
import { TextValidator} from 'react-material-ui-form-validator';
import AddIcon from 'material-ui/svg-icons/content/add';
import NotificationSystem from 'react-notification-system';

import Endpoint from './Endpoints.js';
import AuditorReviewTable from "./AuditorReviewTable";

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
    AuditorReview: [],
    openDialog: false,
    form: INITIAL_FORM,
    fetchingData: false,
    errorMessage: false,
};

class AuditorReview extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleChange = this._handleChange.bind(this);
        this._handleSubmit = this._handleSubmit.bind(this);
        this._openDialog = this._openDialog.bind(this);
    }

    getAuditorReview(){
        this.setState({ fetchingData:true });
        Endpoint.getAllAuditorReview()
            .then(response => {
                let AuditorReview = response.data;
                this.setState({ AuditorReview, fetchingData: false });
            })
            .catch(error => {
                 this.setState({ fetchingData: false, errorMessage: true });
            })
    }

    componentDidMount(){
        this.getAuditorReview();
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
        let AuditorReview = {
            name: this.state.form.name,
        };

        Endpoint.createAuditorReview(AuditorReview)
            .then(response => {
                this.updateInfo();
                this._handleClose();
                this._addNotification('AuditorReview', 'The new AuditorReview has been created successfully', 'success');
            })
            .catch(error => {
                this._addNotification('AuditorReview', 'An error happened', 'error');
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
        this.getAuditorReview();
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
                        title="New AuditorReview"
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
                 {this.state.AuditorReview.length > 0
                    ? <AuditorReviewTable
                    AuditorReview = {this.state.AuditorReview}
                    updateInfo = {this.updateInfo}
                    fetchingData = {this.state.fetchingData}
                    _addNotification = {this._addNotification}
                />
                    : <div></div>}
            </div>
        );
    }
}

export default AuditorReview;