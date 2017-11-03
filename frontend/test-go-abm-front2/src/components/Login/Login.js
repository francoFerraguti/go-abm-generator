import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import Loading from 'react-loading-spinner';
import CircularProgress from 'material-ui/CircularProgress';
import {white} from 'material-ui/styles/colors';
import Endpoint from './api.js';
import {ValidatorForm} from 'react-form-validator-core';
import {TextValidator} from 'react-material-ui-form-validator';
import Dialog from 'material-ui/Dialog';
import {Redirect} from 'react-router-dom';

const spinnerStyle = {
    top: '50%',
    left: '50%',
    position: 'fixed',
};

const containerStyles = {
    loginContainer: {
        position: 'relative',
        height: '100vh',
        backgroundColor: 'rgb(250, 180, 200)',
    },
    loginPage: {
        height: '80%',
        position: 'absolute',
        top: '0',
        left: '0',
        right: '0',
        bottom: '0',
        textAlign: 'center',
        margin: 'auto',
    },
    titleLogin: {
        fontSize: '60px',
        color: 'white',
        marginTop: '20px',
        marginBottom: '20px',
        fontFamily: 'Arvo, serif',
    },
    subtitleLogin: {
        fontSize: '30px',
        color: 'white',
        marginTop: '20px',
        marginBottom: '20px',
        fontFamily: 'Arvo, serif',
    },
    imageLogin: {
        marginTop: '0px',
        marginBottom: '10px',
    },

    loginForm: {
        width: '60%',
        margin: 'auto',
        maxWidth: '450px',
    },

    logo: {
        width: '200px',
    },

    loginButton: {
        marginTop: '20px',
        marginBottom: '48px'
    },

    noAccountPrompt: {
        fontSize: '22px',
        color: 'white',
        fontFamily: 'Arvo, serif',
    },

};

const INITIAL_FORM = {
    username: '',
    password: '',
};

const INITIAL_STATE = {
    errorMessage: false,
    errorText: '',
    loading: false,
    form: INITIAL_FORM,
    authenticated: false,
};


class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = INITIAL_STATE;
        this._handleSubmit = this._handleSubmit.bind(this);
        this._handleClose = this._handleClose.bind(this);

        document.body.style.backgroundColor = "rgb(250, 180, 200)";
    }

    _handleSubmit() {
        let User = {
            username: this.state.form.username,
            password: this.state.form.password,
        };
        
        Endpoint.loginAccess(User)
            .then((response) => {
                localStorage.setItem('token', response.data);
                localStorage.setItem('username', this.state.form.username);
                let authenticated = true;
                this.setState({ authenticated });
            })
            .catch((error) => {
                if (error.response.status === 500) {
                    let errorText = 'Wrong credentials, please try again';
                    this.setState({errorMessage: true, errorText});
                }
            });
    }

    _handleClose = () => {
        this.setState({errorMessage: false});
    };

    isAuthenticated() {
        const token = localStorage.getItem('token');
        return token && token.length > 10;
    }

    render() {
        const isAlreadyAuthenticated = this.isAuthenticated();
        
        return (
            <div>
                {isAlreadyAuthenticated ? <Redirect to={{pathname: '/home'}}/> : (
                    <Loading isLoading={this.state.loading}
                         spinner={() => <CircularProgress id="loginSpinner" scale={1.5} centered={true} style={spinnerStyle} />}>
                    <div style={containerStyles.loginContainer}>
                        <div style={containerStyles.loginPage}>
                            <div style={containerStyles.titleLogin}>Nombre de la app</div>

                            {this.props.message &&
                            <p>{this.props.message}</p>
                            }
                            <div style={containerStyles.loginForm}>
                                <ValidatorForm
                                    ref="form"
                                    onSubmit={this._handleSubmit}
                                    onError={(errors) => console.log(errors)}
                                >
                                    <TextValidator
                                        floatingLabelText="Username"
                                        onChange={(e, value) => this.setState((prevState) => { return {form: {...prevState.form, username: value}}; })}
                                        name="username"
                                        value= {this.state.form.username}
                                        validators={['required']}
                                        errorMessages={['This field is required']}
                                    />
                                   <TextValidator
                                        floatingLabelText="Password"
                                        onChange={(e, value) => this.setState((prevState) => { return {form: {...prevState.form, password: value}}; })}
                                        type="password"
                                        name="password"
                                        value= {this.state.form.password}
                                        validators={['required']}
                                        errorMessages={['This field is required']}
                                    />
                                <br/>
                                <div>
                                    <RaisedButton
                                        style={containerStyles.loginButton}
                                        backgroundColor={white}
                                        labelColor={'#D50032'}
                                        label="Login"
                                        type="submit"
                                     />
                                </div>
                                </ValidatorForm>
                            </div>

                            <Dialog
                                title="Login Error!"
                                titleStyle={{textAlign: 'center'}}
                                modal={false}
                                open={this.state.errorMessage}
                                onRequestClose={this._handleClose}
                            >
                              {this.state.errorText}
                              <div
                                    style={{
                                    display: 'flex',
                                    justifyContent: 'flex-end',
                                }}>
                              <RaisedButton
                                        label="Cerrar"
                                        onClick={this._handleClose}
                                    />
                            </div>
                            </Dialog>
                        </div>
                    </div>
                </Loading>
                )}
           </div>
        );
    }
}

export default Login;
