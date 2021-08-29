import React from 'react'
import { Form, Container, Row, Col, Button, FormControl} from 'react-bootstrap';
import { serviceConfig } from "../applicationSettings.js";
import { app } from './base';
import { withRouter } from "react-router";

export class Verification extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            _name: "",
            _surname: "",
            _username: "",
            _category: 0,
            _mediapath: "",
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

    fileChange = async (e)=>{
        const file = e.target.files[0]
        const storageRef = app.storage().ref()
        const fileRef = storageRef.child(file.name)
        await fileRef.put(file)
        var a = await fileRef.getDownloadURL()
        console.log("file paths:",a)          
        this.setState({ _mediapath: a})
        console.log(this.state._mediapath)
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

    componentDidMount() {
        this.loadUserData();
    }

    loadUserData() {
        var token = localStorage.getItem("token");
        var user = null;

        const requestOpt = {
            method: "GET",
            headers: {
              "Content-Type": "aplication/json",
              Authorization: `Bearer ${token}`,
            },
            credentials: "same-origin",
        };
        fetch(`${serviceConfig.baseURL}/user`, requestOpt)
            .then((response) => response.json())
            .then((responseJson) => {
                user = responseJson;
                this.setState({
                    _name: user.name,
                    _surname: user.surname,
                    _username: user.username,
            });
        })
        .catch((error) => {
            console.error(error);
        });
    }

    createVerificationRequest() {
        const {
            _name,
            _surname,
            _username,
            _mediapath,
            _category,
          } = this.state;

        const verificationRequest = {
            name: _name,
            surname: _surname,
            username: _username,
            mediapath: _mediapath,
            category: _category,
        };

        console.log(JSON.stringify(verificationRequest))

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
          alert("Verification request sent. Returning to homepage.");
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
                        <Form.Row>
                            <Form.Group>
                                <Form.File id="exampleFormControlFile1" label="Select picture" onChange={this.fileChange} accept=".mp4,.jpg,.jpeg,.png"/>
                            </Form.Group>
                            <Form.Group>
                                <img width="100" height="100" src={this.state._mediapath}/>
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