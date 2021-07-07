import React from 'react';
// import {Form, Button, Container,Image} from 'react-bootstrap';
// import {Link} from 'react-router-dom';
// import {serviceConfig} from '../applicationSettings.js'
// import logo from "../resources/nothingramBeli.png";
import { serviceConfig } from '../applicationSettings';
import {Container,Button} from 'react-bootstrap';
import {Form} from 'react-bootstrap';
import '../styles/Login.css';
// import { withRouter } from 'react-router'

class Update extends React.Component{
  constructor(props){
      super(props);

      this.state = {
        
        _token: localStorage.getItem('token'),
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
       // _verify:false,
        _role:1,
        _notify:false,
        checked: false,
        _private:false,
        _taggable:true,
        _gender: 0,
    


         
      }
      this.child = React.createRef();
      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);

  }

handleChange(e){
  const { id, value } = e.target;
        this.setState({ [id]: value });
}

handleSubmit(e){
  e.preventDefault();
  this.updateUser();
}

  // componentWillMount(){
  //   this.renderMyData();
  // }

  renderMyData(){
    var _token = localStorage.getItem("token")
    var user2 = null

    const RequestDTO = { token: _token}
        
        // const userLogged = { email: null,username: null,token: null}

        const requestOpt = {
            method: 'POST',
            headers:{'Content-Type': 'aplication/json'},
            body: JSON.stringify(RequestDTO),
            credentials: 'same-origin'//,'access-control-allow-origin' : '*'
        }
    fetch(`${serviceConfig.baseURL}/user`,requestOpt)
        .then((response) => response.json())
        .then((responseJson) => {
          // this.setState({ user : responseJson })
          user2 = responseJson
          this.setState({ _name : user2.name ,_surname: user2.surname,_username: user2.username,_email: user2.email,_phone: user2.phone,_password: user2.password})
          //console.log(this.user)
        })
        .catch((error) => {
          console.error(error);
        });
}
  componentDidMount(){
    this.renderMyData();
  }

  updateUser(){
    const {
      _token,
      _name,
      _surname,
      _username,
      _dateOfBirth,
      _phone,
      _web,
      _email,
      _password,
      _bio,
     // _verify,
      _role,
      _notify,
      _private,
      _taggable,
      _gender} = this.state;

  const updateUser = {
      token: _token,
      email: _email,
      password: _password,
      name: _name,
      surname: _surname,            
      phone: _phone,
      username: _username,
      bio: _bio,
     // verify: _verify,
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
    body: JSON.stringify(updateUser)};


   fetch(`${serviceConfig.baseURL}/update`, requestOptions)
        .then(response => {
            if (!response.ok) {
                return Promise.reject(response);
            }
            this.props.history.push('/update');
          
        })
        .catch(response => {
            if(response.status === 400){
              
            alert("error!");}
        })
    

  }

  render(){
      const {_name,_surname,_username,_email,_phone,_password} = this.state;

      return(
    //   <Container>
    //     <div>
    //         <h1>EDIT PROFILE</h1>
    //     </div>
    //     <form onSubmit={this.handleSubmit}>
    //         <label>Name:</label>
    //         <input type="text" value={_name} id="_name" onChange={this.handleChange} />
    //         {/* <textField></textField> */}
    //         <input type="submit" value="Submit" />
    //     </form>
    // </Container>

      <Container style={{position: "relative", top: "50%", transform: "translateY(30%)"}}>

                  {/* <h2 style={{textAlign:"center"}}>NothinGRAM</h2>  */}
                  <Form  style={{textAlign:"center"}} onSubmit={this.handleSubmit}>
                      <Form.Group >
                          <Form.Control
                              
                              id="_name"
                              value={_name}
                              type="text"
                              pattern="[A-Za-z]+"
                              onChange = {this.handleChange}
                          />
                      </Form.Group>
                      <Form.Group>
                          <Form.Control
                              
                              id="_surname"
                              value={_surname}
                              type="text"
                              pattern="[A-Za-z]+"
                              onChange = {this.handleChange}
                          />
                      </Form.Group>
                      <Form.Group>
                          <Form.Control
                              
                              id="_username"
                              value={_username}
                              pattern="^[A-Za-z0-9]{0,10}$"
                              type="text"
                              onChange = {this.handleChange}
                          />
                      </Form.Group>
                      <Form.Control
                                    
                                    id="_password"
                                    type="password"
                                    value={_password}
                                  
                                    onChange={this.handleChange}
                                />
                                  <Form.Control
                                    
                                    id="_email"
                                    type="_email"
                                    value={_email}
                                  
                                    onChange={this.handleChange}
                                />
                                 <Form.Control
                                    
                                    id="_phone"
                                    type="_phone"
                                    value={_phone}
                                  
                                    onChange={this.handleChange}
                                />
                      <div className="buttonLogin">
                                <Button variant="primary" type="submit">
                                    Update account
                                </Button>
                        </div>  
                
                  </Form>
                  

            
      </Container>);
    
  }
}

export default Update;