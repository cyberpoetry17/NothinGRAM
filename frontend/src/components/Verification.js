import React from 'react'
import { Form, Container, Row, Col, Button} from 'react-bootstrap';
import Select from 'react-select'
import { serviceConfig } from "../applicationSettings.js";
import { withRouter } from "react-router";

export class Verification extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            _name: "",
            _surname: "",
            _username: "",
            _category: 0,
        };

        //this.child = React.createRef();
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    handleSubmit(e) {
        e.preventDefault();
        console.log("---SUBMIT HANDLER---");
        this.createVerificationRequest();
    }

    handleChange(e) {
        const { id, value } = e.target;
        this.setState({ [id]: value });
    }

    componentDidMount() {}

    createVerificationRequest() {
        const {
            _name,
            _surname,
            _username,
            _category,
          } = this.state;

        const verificationRequest = {
            name: _name,
            surname: _surname,
            username: _username,
            category: _category,
        };

        const verificationRequestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(verificationRequest),
          };

        fetch(`${serviceConfig.baseURL}/verification`, verificationRequestOptions)
        .then((response) => {
          if (!response.ok) {
            return Promise.reject(response);
          }
          this.props.history.push("/verification");
        })
        .catch((response) => {
          if (response.status === 400) {
            alert("Verification request of this type already exists.");
          }
        });
    }

    render(){
        const options = [
            { value: 'chocolate', label: 'Influencer' },
            { value: 'sports', label: 'Sports' },
            { value: 'news/media', label: 'NewsOrMedia' },
            { value: 'business', label: 'Business' },
            { value: 'brand', label: 'Brand' },
            { value: 'organization', label: 'Organization' }
          ]
          const {
            _name,
            _surname,
            _username,
            _category,
          } = this.state;


        return(
            <div>
            <h1 className="d-flex justify-content-center">Verification form</h1>
            <Container  width="20">
                <Row>
                    <Col></Col>
                    <Col>
                        <Form onSubmit={this.handleSubmit} className="justify-content-md-center">
                        <Row>
                        <Col>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Enter your name </Form.Label>
                                <Form.Control
                                    required
                                    //id="name"
                                    value={_name}
                                    type="name"
                                    placeholder="Enter name"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            </Col>
                            <Col>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Enter your surname</Form.Label>
                                <Form.Control
                                    required
                                    id="surname"
                                    value={_surname}
                                    type="surname"
                                    placeholder="Enter surname"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            </Col>
                            <Col>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Enter your username</Form.Label>
                                <Form.Control
                                    required
                                    id="username"
                                    value={_username}
                                    type="username"
                                    placeholder="Enter username"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                            </Col>
                        </Row>
                            <Select
                                required
                                id="category"
                                type="category"
                                options={options}
                                value={_category}
                                onChange={this.handleChange}
                            />
                            <Button variant="primary" type="submit">Send verification request</Button>
                        </Form>
                    </Col>
                    <Col></Col>
                </Row>
                
            </Container>
        </div>
        )
    }
}
export default Verification