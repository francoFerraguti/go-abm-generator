import React from "react";
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import { ValidatorForm } from 'react-form-validator-core';
import { TextValidator} from 'react-material-ui-form-validator';
import AddIcon from 'material-ui/svg-icons/content/add';
import NotificationSystem from 'react-notification-system';

import Endpoint from './Endpoints.js';
import PhotoTable from "./PhotoTable";

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
    Photo: [],
    openDialog: false,
    form: INITIAL_FORM,
    fetchingData: false,
    errorMessage: false,
};

class Photo extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleChangeUrl = this._handleChangeUrl.bind(this);
this._handleChangeLat = this._handleChangeLat.bind(this);
this._handleChangeLon = this._handleChangeLon.bind(this);

        this._handleSubmit = this._handleSubmit.bind(this);
        this._openDialog = this._openDialog.bind(this);
    }

    getPhoto(){
        this.setState({ fetchingData:true });
        Endpoint.getAllPhoto()
            .then(response => {
                let Photo = response.data;
                this.setState({ Photo, fetchingData: false });
            })
            .catch(error => {
                 this.setState({ fetchingData: false, errorMessage: true });
            })
    }

    componentDidMount(){
        this.getPhoto();
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
        let Photo = {
            url: this.state.url,
lat: this.state.lat,
lon: this.state.lon,

        };

        Endpoint.createPhoto(Photo)
            .then(response => {
                this.updateInfo();
                this._handleClose();
                this._addNotification('Photo', 'The new Photo has been created successfully', 'success');
            })
            .catch(error => {
                this._addNotification('Photo', 'An error happened', 'error');
                this._handleClose();
            })
    }

    _handleChangeUrl = (event) => {
    const target = event.target
    const value = target.value
    const name = target.name

    this.setState({
        url: value
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
        this.getPhoto();
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
                        title="New Photo"
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
    floatingLabelText='Url'
    onChange={this._handleChangeUrl}
    name='url'
    value={this.state.form.url}
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
                 {this.state.Photo.length > 0
                    ? <PhotoTable
                    Photo = {this.state.Photo}
                    updateInfo = {this.updateInfo}
                    fetchingData = {this.state.fetchingData}
                    _addNotification = {this._addNotification}
                />
                    : <div></div>}
            </div>
        );
    }
}

export default Photo;