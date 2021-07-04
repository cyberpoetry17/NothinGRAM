import React from 'react'
import {Form, Container,Col,Button} from 'react-bootstrap';
import DatePick from '../components/DateOfBirth.js'
import {serviceConfig} from '../applicationSettings.js'
import 'react-datepicker/dist/react-datepicker.css'
class RegisterUser extends React.Component{
    constructor(props){
        super(props);

        this.state = {
            _name:"",
            _surname:"",
            _username:"",
            _dateOfBirth: new Date(),
            _phone:"",
            _web:"",
            _email: "",
            _password: "",
            _repeatPassword: "",
            _bio:"",
            _verify:"",
            _role:"",
            _notify:"",
            _private:"",
            _taggable:"",
            _gender:"",
            message: ""
        }
        this.child = React.createRef();
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleGender = this.handleGender.bind(this);

    }

    handleChange(e){
        const { id, value } = e.target;
        this.setState({ [id]: value });
    }
    
 
    handleGender = (event) => {
        this.setState({_gender: event.target.value})
      }
    handleSubmit(e) {
        e.preventDefault();
        

        const { _password, _repeatPassword } = this.state;
        
        if(_password.trim() !== _repeatPassword.trim()){
            alert('Passwords do not match!');
            return;
        }
        console.log("working")
        this.register();
    }
    componentDidMount(){
       
    }
    
    register(){
        
        const {_name,
            _surname,
            _username,
            _dateOfBirth,
            _phone,
            _web,
            _email,
            _password,
            _bio,
            _verify,
            _role,
            _notify,
            _private,
            _taggable,
            _gender,} = this.state;

        const registerRequest = {
            email: _email,
            password: _password,
            name: _name,
            surname: _surname,            
            phoneNumber: _phone,
            username: _username,
            bio: _bio,
            verify: _verify,
            role: _role,
            notifications: _notify,
            private:_private,
            taggable: _taggable,
            gender: _gender,
            web: _web,
            date: _dateOfBirth 
        }

        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(registerRequest)
        };

        fetch(`${serviceConfig.baseURL}/register`, requestOptions)
        .then(response => {
            if (!response.ok) {
                return Promise.reject(response);
            }
            this.props.history.push('/');
        })
        .catch(response => {
            const promise = Promise.resolve(response.json);
            promise.then(data => {
                alert(data.message);
            })            
        })
    }
    
   
    render(){
        const {_email, _password,_username,_surname,_phone,_dateOfBirth,_repeatPassword,_name} = this.state;
        
        return(
            <Container style={{position: "relative"}}>
                    <div 
                    className="register-div"
                    style={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center"}}

                    >
                        <h2 style={{textAlign:"center"}}>NothinGRAM</h2> 
                        <Form onSubmit={this.handleSubmit}>
                        <Form.Row>
                        <Form.Group as={Col} md="6">
                              
                                <Form.Control
                                    required
                                    id="_email"
                                    value={_email}
                                    type="email"
                                    placeholder="Email"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            <Form.Group as={Col} md="6">
                              
                                <Form.Control
                                    required
                                    id="_username"
                                    value={_username}
                                    type="text"
                                    placeholder="Username"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                        <Form.Row>
                        <Form.Group as={Col} md="6">
                                
                                <Form.Control
                                    required
                                    id="_password"
                                    type="password"
                                    value={_password}
                                    placeholder="Password"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            <Form.Group as={Col} md="6">
                        
                                <Form.Control
                                    required
                                    id="_repeatPassword"
                                    value={_repeatPassword}
                                    type="password"
                                    placeholder="Password"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                        <Form.Row>
                        <Form.Group as={Col} md="4">
                               
                                <Form.Control
                                    required
                                    id="_name"
                                    value={_name}
                                    type="text"
                                    placeholder="Name"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            <Form.Group as={Col} md="4">
                             
                                <Form.Control
                                    required
                                    id="_surname"
                                    value={_surname}
                                    type="text"
                                    placeholder="Surname"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            <Form.Group as={Col} md="4">
                              
                                <Form.Control
                                    required
                                    id="_phone"
                                    value={_phone}
                                    type="text"
                                    placeholder="Phone"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                       <DatePick 
                       id="_dateOfBirth"
                       value={_dateOfBirth}
                       onChange={this.handleChange}
                       ></DatePick>
           
                       <div>
                       {/* <div>

                       <input type="radio" value="Male"  
                      checked={_gender === 'Male'}
                      onChange={this.handleGender}
                       
                       /> Male

                       <input type="radio" value="Female" name="gender" 
                        checked={_gender === 'Female'}
                        onChange={this.handleGender}/> Female

                       </div> */}
                       </div>

                        <div className="text-center">
                                <Button variant="primary" type="submit">
                                    Submit
                                </Button>
                        </div>
                    </Form>
                </div>
            </Container>
        );
    }
}
export default RegisterUser;