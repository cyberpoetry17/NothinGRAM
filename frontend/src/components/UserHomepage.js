import React, { Component } from 'react';
import {withRouter} from 'react-router';
import jwt_decode from 'jwt-decode';
import validateJWS from './JWTValidation';
// import jwt from 'json-web-token';

class UserHomepage extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
            tokenInfo: jwt_decode(localStorage.getItem('token'))
        }
    }

    render() {
        if (validateJWS(this.state.tokenInfo)) {
            alert('Please sign in to continue');
            this.props.history.push('/login');
        }

        return (
            <div>
                <h1>USER HOMEPAGE</h1>
                <h3>USER LOGGED: OK</h3>
                <h3>USERID: {this.state.tokenInfo.UserID}</h3>
                <h3>USERNAME: {this.state.tokenInfo.Username}</h3>
                <h3>EMAIL: {this.state.tokenInfo.Email}</h3>
                <h3>TOKEN EXPIRES: {this.state.tokenInfo.exp}</h3>
            </div>
        )
    }
}

export default withRouter(UserHomepage)
