import React from "react";
import { serviceConfig } from "../applicationSettings";
import { Container, Button } from "react-bootstrap";
import { Form } from "react-bootstrap";
import "../styles/Login.css";
// import { withRouter } from 'react-router'
class AdvancedUpdate extends React.Component {
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
      this.handleNotification = this.handleNotification.bind(this);
      this.handleTaggable = this.handleTaggable.bind(this);
      this.handlePrivate = this.handlePrivate.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
      
    }
  
    handleChange(e) {
      const { id, value } = e.target;
      this.setState({ [id]: value });
    }

    handleNotification(e) {
        this.setState({ _notify: !this.state._notify });
    }

    handleTaggable(e) {
        this.setState({ _taggable: !this.state._taggable });
    }

    handlePrivate(e) {
        this.setState({ _private: !this.state._private });
    }
  

    
    handleSubmit(e){
      e.preventDefault();
      this.updateUser();
    }
  
    componentWillMount(){
      this.renderMyData();
    }
  
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
            _gender: user2.gender,
            _role: user2.role,
            _notify: user2.notify,
            _private: user2.private,
            _taggable: user2.taggable

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
        credentials: 'same-origin',
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
      const { _private,_taggable,_notify} =
        this.state;
  
      return (
        <Container
          style={{
            position: "relative",
            top: "50%",
            transform: "translateY(30%)", 
          }}
        >
            <Form>
             <div>
                <input
                  type="checkbox"
                  id="_notify"
                  name="topping"
                  value="_notify"
                  checked={_notify}
                  onChange={this.handleNotification}
                />
                <label>Receive notifications</label>
              </div>
              </Form>
              <Form>
             <div>
                <input
                  type="checkbox"
                  id="_private"
                  name="topping"
                  value="_private"
                  checked={_private}
                  onChange={this.handlePrivate}
                />
                <label>Make profile private</label>
              </div>
              </Form>
              <Form>
             <div>
                <input
                  type="checkbox"
                  id="_taggable"
                  name="topping"
                  value="_taggable"
                  checked={_taggable}
                  onChange={this.handleTaggable}
                />
                <label>Make profile private</label>
              </div>
              </Form>



            <Form>
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
  
  export default AdvancedUpdate;
  