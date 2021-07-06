import React, { Component } from 'react'
import jwt_decode from 'jwt-decode'

class UserHomepage extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
            token: localStorage.getItem('token')
        }
    }
    
    render() {
        const tokenInfo = jwt_decode(this.state.token);

        return (
            <div>
                <h1>USER HOMEPAGE</h1>
                <h3>USER LOGGED: OK</h3>
                <h3>USERID: {tokenInfo.UserID}</h3>
                <h3>USERNAME: {tokenInfo.Username}</h3>
                <h3>EMAIL: {tokenInfo.Email}</h3>
                <h3>TOKEN EXPIRES: {tokenInfo.exp}</h3>
            </div>
        )
    }
}

export default UserHomepage
