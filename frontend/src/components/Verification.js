import axios from 'axios';
import React,{useState,useEffect} from 'react'
import { app } from './base';
import {Form,Container,Row,Col,Button} from 'react-bootstrap';
import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';
import PlacesAutocomplete, {
    geocodeByAddress,
    getLatLng
  } from "react-places-autocomplete";
import Select from 'react-select'
import { auto } from '@popperjs/core';
import { send } from 'q';

export class Verification extends React.Component{
    
    render(){
        const options = [
            { value: 'chocolate', label: 'Influencer' },
            { value: 'sports', label: 'Sports' },
            { value: 'news/media', label: 'News/Media' },
            { value: 'business', label: 'Business' },
            { value: 'brand', label: 'Brand' },
            { value: 'organization', label: 'Organization' }
          ]
        return(
            <div>
            <h1 className="d-flex justify-content-center">Verification form</h1>
            <Container  width="20">
                <Row>
                    <Col></Col>
                    <Col>
                        <Form className="justify-content-md-center">
                        <Row>
                        <Col>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Enter your name </Form.Label>
                                <Form.Control  placeholder="Enter name " />
                            </Form.Group>
                            </Col>
                            <Col>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Enter your surname</Form.Label>
                                <Form.Control  placeholder="Enter surname" />
                            </Form.Group>
                            </Col>
                            <Col>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Enter your username</Form.Label>
                                <Form.Control  placeholder="Enter username" />
                            </Form.Group>
                            </Col>
                        </Row>
                            <Form.Group>
                                <Form.File id="exampleFormControlFile1" label="Input your face picture" />
                            </Form.Group>
                            <Select options={options}/>
                            <img width="100" height="100" alt="my pic"/><br/>
                            <Button onClick={send}>Send verification request</Button>
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