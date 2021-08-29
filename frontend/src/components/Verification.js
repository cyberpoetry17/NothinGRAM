import React from 'react'
import { Form, Container, Row, Col, Button, FormControl} from 'react-bootstrap';
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
        this.handleChangeInfluencer = this.handleChangeInfluencer.bind(this);
        this.handleChangeSports = this.handleChangeSports.bind(this);
        this.handleChangeNewsOrMedia = this.handleChangeNewsOrMedia.bind(this);
        this.handleChangeBusiness = this.handleChangeBusiness.bind(this);
        this.handleChangeBrand = this.handleChangeBrand.bind(this);
        this.handleChangeOrganization = this.handleChangeOrganization.bind(this);
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

    handleChangeInfluencer(e) {
        this.setState({ _category: 1})
    }

    handleChangeSports(e) {
        this.setState({ _category: 2})
    }

    handleChangeNewsOrMedia(e) {
        this.setState({ _category: 3})
    }

    handleChangeBusiness(e) {
        this.setState({ _category: 4})
    }

    handleChangeBrand(e) {
        this.setState({ _category: 5})
    }

    handleChangeOrganization(e) {
        this.setState({ _category: 6})
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
          this.props.history.push("/home");
        })
        .catch((response) => {
          if (response.status === 400) {
            alert("Verification request of this type already exists.");
          }
        });
    }

    render(){
        const {
            _name,
            _surname,
            _username,
            _category,
        } = this.state;

        return(
            <Container style={{position: "relative"}}>
                <h1 className="d-flex justify-content-center">Verification form</h1>
                <div
                    style={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center"
                    }}
                >
                    <Form onSubmit={this.handleSubmit}>
                        <Form.Row>
                            <Form.Group as={Col} md="6">
                                <FormControl
                                    required
                                    id="_name"
                                    value={_name}
                                    type="name"
                                    placeholder="Name"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                        <Form.Row>
                            <Form.Group as={Col} md="6">
                                <FormControl
                                    required
                                    id="_surname"
                                    value={_surname}
                                    type="surname"
                                    placeholder="Surname"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                        <Form.Row>
                            <Form.Group as={Col} md="6">
                                <FormControl
                                    required
                                    id="_username"
                                    value={_username}
                                    type="username"
                                    placeholder="Username"
                                    onChange={this.handleChange}
                                />
                            </Form.Group>
                        </Form.Row>
                        <div>
                            <div>
                                <input
                                    type="radio"
                                    value={1}
                                    checked={_category === 1}
                                    onChange={this.handleChangeInfluencer}
                                />{" "}
                                Influencer
                                <input
                                    type="radio"
                                    value={2}
                                    checked={_category === 2}
                                    onChange={this.handleChangeSports}
                                />{" "}
                                Sports
                                <input
                                    type="radio"
                                    value={3}
                                    checked={_category === 3}
                                    onChange={this.handleChangeNewsOrMedia}
                                />{" "}
                                News/Media
                                <input
                                    type="radio"
                                    value={4}
                                    checked={_category === 4}
                                    onChange={this.handleChangeBusiness}
                                />{" "}
                                Business
                                <input
                                    type="radio"
                                    value={5}
                                    checked={_category === 5}
                                    onChange={this.handleChangeBrand}
                                />{" "}
                                Brand
                                <input
                                    type="radio"
                                    value={6}
                                    checked={_category === 6}
                                    onChange={this.handleChangeOrganization}
                                />{" "}
                                Organization
                            </div>
                        </div>
                        <div>
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
export default withRouter(Verification)