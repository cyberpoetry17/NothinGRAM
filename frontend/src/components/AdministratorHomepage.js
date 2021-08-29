import React from 'react';
import {Form, Button, Container,Image} from 'react-bootstrap';
import {Link} from 'react-router-dom';
import {withRouter} from 'react-router';
import {serviceConfig} from '../applicationSettings.js'
import logo from "../resources/nothingramBeli.png";
import '../styles/Login.css';

export class AdministratorHomepage extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            _username: "",
        };
    }

    render() {
        return(
            <h1>admin page</h1>
        )
    }
}

export default withRouter(AdministratorHomepage);