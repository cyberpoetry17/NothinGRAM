import React from 'react';
import {Form, Button, Container,Image} from 'react-bootstrap';
import {Link} from 'react-router-dom';
import {serviceConfig} from '../applicationSettings.js'
import logo from "../resources/nothingramBeli.png";
import '../styles/Login.css';
import { withRouter } from 'react-router'


class Login extends React.Component{
    constructor(props){
        super(props);

        this.state = {
            _email: "",
            _password: "",
            message: ""
        }
        this.child = React.createRef();
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);

    }
 
  

    handleChange(e){
        const { id, value } = e.target;
        this.setState({ [id]: value });
    }

    handleSubmit(e) {
        e.preventDefault();
        this.login();
    }

    componentDidMount(){
      
    }

    login(){
        const {_email,_password} = this.state

        const LoginDTO = { email: _email, password: _password}
        // const userLogged = { email: null,username: null,token: null}

        const requestOpt = {
            method: 'POST',
            headers:{'Content-Type': 'aplication/json'},
            body: JSON.stringify(LoginDTO),
            //credentials: 'same-origin' ,'access-control-allow-origin' : '*'
        }

        fetch(`${serviceConfig.baseURL}/login`,requestOpt)
            .then(response => {
                if(!response.ok){
                    console.log("neuspelo");
                    return Promise.reject(response);   
                }
                console.log("USPELO");
                this.props.history.push('/home');
                return response.json();
            })
            // .then(response => {
            //     console.log(response.headers.get('set-cookie')); // undefined
            //     console.log(document.cookie); // nope
            //     return response.json();
            //   })
            // .then(response => {
            //     if (!response.ok) {
            //         return Promise.reject(response);
            //     }
               
              
            // })
            .then((data) => {
                console.log(data.token);
                if(data != null){
                    if(data.token != null){
                        localStorage.setItem('token', data.token)
                    }}
               
            })
            .catch(response => {
               if(response.status === 400)
               {alert("Please insert valid credentials")}
            })
            
            
            
           
    }
    
    render(){
        const {_email, _password} = this.state;

        return(
            <Container style={{position: "relative", top: "50%", transform: "translateY(41%)"}}>
                    <div className='login-div'>
                        <Image 
                            style={{marginLeft:"20%"}}
                            src={logo}
                            alt="NOTHINGRAM"
                            className="logo"
                           
                        />
                        {/* <h2 style={{textAlign:"center"}}>NothinGRAM</h2>  */}
                        <Form  style={{textAlign:"center"}} onSubmit={this.handleSubmit}>
                            <Form.Group>
                                <Form.Control
                                    required
                                    id="_email"
                                    value={_email}
                                    type="text"
                                    placeholder="Email"
                                    onChange = {this.handleChange}
                                />
                            </Form.Group>
                            <Form.Group>
                                {/* <Form.Label>
                                    Password
                                </Form.Label> */}
                                <Form.Control
                                    required
                                    id="_password"
                                    value={_password}
                                    type="password"
                                    placeholder="Password"
                                    onChange = {this.handleChange}
                                />
                            </Form.Group>
                        <div className="buttonLogin">
                                <Button variant="primary" type="submit">
                                    Log in
                                </Button>
                        </div>
                        </Form>
                        <div style={{textAlign:"center"}} className="text-center">
                            Don't have an account?
                            <br/>   
                            <Link to="/register">
                                Sign up
                            </Link>
                    </div>

                    </div>
            </Container>
        );
    }
}

export default withRouter(Login);