import React from "react";
import { serviceConfig } from "../applicationSettings";
import { Container, Button } from "react-bootstrap";
import { Form } from "react-bootstrap";
import "../styles/Login.css";
// import { withRouter } from 'react-router'

class Update extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      _token: localStorage.getItem("token"),
      _name: "",
      _surname: "",
      _username: "",
      _dateOfBirth: new Date(),
      _phone: "",
      _web: "",
      _email: "",
      _password: "",
      _repeatPassword: "",
      _bio: "",
      _role: 0,
      _notify: false,
      checked: false,
      _private: false,
      _taggable: true,
      _gender: 0,
    };
    this.child = React.createRef();
    this.handleChange = this.handleChange.bind(this);
    this.handleGenderFemale = this.handleGenderFemale.bind(this);
    this.handleGenderMale = this.handleGenderMale.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.onChange = this.onChange.bind(this);
  }

  handleChange(e) {
    const { id, value } = e.target;
    this.setState({ [id]: value });
  }

  handleGenderMale() {
    this.setState({ _gender: 0 });
  }
  handleGenderFemale() {
    this.setState({ _gender: this.state._gender });

  }
  onChange = () => {
    this.setState({
      _gender: this.state._gender
    });
    console.log("GENDEEER")
    console.log(this.state._gender)
  };
  
  handleSubmit(e){
    e.preventDefault();
    const { _password, _repeatPassword } = this.state;
        
    if(_password.trim() !== _repeatPassword.trim()){
        alert('Passwords do not match!');
        return;
    }
    this.updateUser();
  }

  // componentWillMount(){
  //   this.renderMyData();
  // }

  renderMyData() {
    var token = localStorage.getItem("token");
    var user2 = null;

    const requestOpt = {
      method: "GET",
      headers: {
        "Content-Type": "aplication/json",
        Authorization: `Bearer ${token}`,
      },

      credentials: "same-origin", //,'access-control-allow-origin' : '*'
    };
    fetch(`${serviceConfig.baseURL}/user`, requestOpt)
      .then((response) => response.json())
      .then((responseJson) => {
        // this.setState({ user : responseJson })
        user2 = responseJson;
        this.setState({
          _name: user2.name,
          _surname: user2.surname,
          _username: user2.username,
          _email: user2.email,
          _phone: user2.phone,
          _password: user2.password,
          _repeatPassword: user2.password,
          _web: user2.web,
          _bio: user2.bio,
          _gender: user2.gender
        });
        //console.log(this.user)
      })
      .catch((error) => {
        console.error(error);
      });
  }
  componentDidMount() {
    this.renderMyData();
  }

  updateUser() {
    var token = localStorage.getItem("token");
    const {
      _name,
      _surname,
      _username,
      _dateOfBirth,
      _phone,
      _web,
      _email,
      _password,
      _bio,
      _role,
      _notify,
      _private,
      _taggable,
      _gender,
    } = this.state;

    const updateUser = {
      email: _email,
      password: _password,
      name: _name,
      surname: _surname,
      phone: _phone,
      username: _username,
      bio: _bio,
      role: _role,
      notifications: _notify,
      private: _private,
      taggable: _taggable,
      gender: _gender,
      web: _web,
      date: _dateOfBirth,
    };
    const requestOptions = {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(updateUser),
    };

    fetch(`${serviceConfig.baseURL}/update`, requestOptions)
      .then((response) => {
        if (!response.ok) {
          return Promise.reject(response);
        }
        this.props.history.push("/update");
        alert("Succesfully updated.")
      })
      .catch((response) => {
        if (response.status === 400) {
          alert("error!");
        }
      });
  }

  render() {
    const { _name, _surname, _username, _email, _phone, _password,_web,_repeatPassword,_bio,_gender} =
      this.state;

    return (
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

      <Container
        style={{
          position: "relative",
          top: "50%",
          transform: "translateY(30%)", 
        }}
      >
        {/* <h2 style={{textAlign:"center"}}>NothinGRAM</h2>  */}
        <Form style={{ textAlign: "center" }} onSubmit={this.handleSubmit}>
          <Form.Group>
            <Form.Control
              required
              id="_name"
              value={_name}
              type="text"
              pattern="[A-Za-z]+"
              placeholder="Name"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Control
              required
              id="_surname"
              value={_surname}
              type="text"
              pattern="[A-Za-z]+"
              placeholder="Surname"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Control
              required
              id="_username"
              value={_username}
              pattern="^[A-Za-z0-9]{0,10}$"
              type="text"
              placeholder="Username"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Control
              required
              id="_password"
              type="password"
              value={_password}
             // placeholder="Password"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Control
              required
              id="_repeatPassword"
              type="password"
              value={_repeatPassword}
              //placeholder="Repeat password"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            {" "}
            <Form.Control
              required
              id="_email"
              type="email"
              value={_email}
              placeholder="Email"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Control
              required
              id="_phone"
              type="text"
              value={_phone}
              pattern="[0][0-9]+"
              maxLength = {12}
              minLength = {10}
              placeholder="Phone"
              onChange={this.handleChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Control
              id="_web"
              type="text"
              pattern="^[A-Za-z0-9]{1,20}[.]+[A-za-z]{2,3}$"
              value={_web}
              placeholder="Web address"
              onChange={this.handleChange}
            />
          </Form.Group>

          <Form.Group>
            <Form.Control
              id="_bio"
              type="text"
              value={_bio}
              placeholder="Insert your bio"
              onChange={this.handleChange}
            />
          </Form.Group>

          <div>
              <div>
                <input
                  type="radio"
                  value={_gender}
                  checked={this.state._gender === 1}
                  onChange={this.onChange}
                />{" "}
                Male
                <input
                  type="radio"
                  checked={_gender}
                  value={this.state._gender === 2}
                  onChange={this.onChange}
                />{" "}
                Female
              </div>
            </div>
          <div className="buttonLogin">
            <Button variant="primary" type="submit">
              Update account
            </Button>
          </div>
        </Form>
      </Container>
    );
  }
}

export default Update;
