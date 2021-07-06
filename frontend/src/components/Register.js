import React from 'react'
import {Form, Container,Col,Button,Image} from 'react-bootstrap';
// import DatePick from '../components/DateOfBirth.js'
import {serviceConfig} from '../applicationSettings.js'
import 'react-datepicker/dist/react-datepicker.css'
import logo from "../resources/nothingramBeli.png";
// import {DatePickerComponent} from "@syncfusion/ej2-react-calendars";
import DatePick from '../components/DateOfBirth.js'
import { withRouter } from 'react-router'


// import required css from library
import "react-datepicker/dist/react-datepicker.css";


class RegisterUser extends React.Component{
    constructor(props){
        super(props);

        this.state = {
            _name:"",
            _surname:"",
            _username:"",
            _dateOfBirth: new Date(),
            _phone:"",
            _web:"web",
            _email: "",
            _password: "",
            _repeatPassword: "",
            _bio:"insert bio",
            _verify:false,
            _role:1,
            _notify:false,
            _private:false,
            _taggable:true,
            _gender: 0,
            message: ""
        }
        this.child = React.createRef();
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        // this.handleGender = this.handleGender.bind(this);
        this.handleGenderMale = this.handleGenderMale.bind(this);
        this.handleGenderFemale = this.handleGenderFemale.bind(this);
        this.handleNotification = this.handleNotification.bind(this);
       

    }

    handleChange(e){
        const { id, value } = e.target;
        this.setState({ [id]: value });
    }

    handleChangeDate(e){
        
        this.setState( {_dateOfBirth : e.target.Date});
    }
    
    
    handleGenderMale(){
        this.setState({ _gender: 0})
    }
    handleGenderFemale(){
        this.setState({ _gender: 1})
    }

    // handleGender(event){
    //     this.setState({ _gender: event.target.value})
    // }

    handleNotification (){
        this.setState({ _notify: true})
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
            _gender} = this.state;

        const registerRequest = {
            email: _email,
            password: _password,
            name: _name,
            surname: _surname,            
            phone: _phone,
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
            this.props.history.push('/login');
          
        })
        .catch(response => {
            if(response.status === 400){
            alert("User already exists!");}
        })
    }
    
   
    render(){
        const {_email, _password,_username,_surname,_phone, _repeatPassword,_name,_gender,_notify,_dateOfBirth} = this.state;
        //const dateValue =  new Date(new Date().getDay,new Date().getMonth,new Date().getMonth)
        
        return(
            <Container style={{position: "relative"}}>
                    <div 
                    className="register-div"
                    style={{
                       
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center"}}

                    ><Image 
                    className="photo"
                    src={logo}
                    alt="NOTHINGRAM"
                   
                   
                />
                       
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
                                    pattern="^[A-Za-z0-9]{0,10}$"
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
                                    placeholder="Repeat password"
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
                                    pattern="[A-Za-z]+"
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
                                    pattern="[A-Za-z]+"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            <Form.Group as={Col} md="5">      
                                <Form.Control
                                    required
                                    id="_phone"
                                    value={_phone}
                                    pattern="[0][0-9]+"
                                    maxLength = {12}
                                    minLength = {10}
                                    type="text"
                                    placeholder="Phone"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                        {/* <DatePickerComponent placeholder="Insert date" value={_dateOfBirth} onChange={this.handleChange}></DatePickerComponent> */}
                        <DatePick 
                       id="_dateOfBirth"
                       value={_dateOfBirth}
                       
                       placeholder="insert date"
                       ></DatePick>  
     
                        <div>
                        <Form.Group as={Col} controlId="formHorizontalCheck">
                        <Col sm={{ span: 10, offset: 2 }}>
                        <Form.Check label="Recieve notification" values={_notify}/>
                        </Col>
                        </Form.Group>

                        </div>
                       
                       <div>
                       <div>
                       
                       <input type="radio" value={0}
                      checked={_gender === 0}
                      onChange={this.handleGenderMale} 
                       /> Male
                       <input type="radio"   name="gender" value={1}
                        checked={_gender === 1}
                        onChange={this.handleGenderFemale}/> Female

                       </div>
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
export default withRouter(RegisterUser);
// export default RegisterUser;