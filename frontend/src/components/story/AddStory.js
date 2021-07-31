import React from 'react'
import {Form,Container,Row,Col,Button,Carousel} from 'react-bootstrap';
import {useState,useEffect} from 'react'
import { app } from '../base';
import { useHistory } from "react-router-dom";
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import BootstrapSwitchButton from 'bootstrap-switch-button-react'

export default function AddStory() {

    const [MediaPath, setMediaPath] = useState("");
    const [IsItForCloseFriends, setForCloseFriends] = useState(false);
    let history = useHistory()

    const fileChange = async (e)=>{
        const file = e.target.files[0]
        const storageRef = app.storage().ref()
        const fileRef = storageRef.child(file.name)
        await fileRef.put(file)
        var a = await fileRef.getDownloadURL()
        console.log("file paths:",a)          
        setMediaPath(a)
    }
    const closeChanged = (e)=>{
        console.log(e);
        setForCloseFriends(e);
    }
    const addStory = (e)=>{
        var token = jwt_decode(localStorage.getItem('token'));
        const body = {
            MediaPath:MediaPath,
            UserId:token.UserID,
            IsOnlyForCloseFriends:IsItForCloseFriends
            // ovde dodaj IsItForCloseFriends i na beku isto primi u dto
        };
        axios({
            method : 'post',
            url :'http://localhost:8005/addStory',
            data:JSON.stringify(body)
        })
        history.push("/")
    }

    return (
        <div>
             <h1 className="d-flex justify-content-center">ADDING STORY</h1>
            <Container >
                <Row>
                    <Col></Col>
                    <Col>
                    <Form className="justify-content-md-center">
                        <Form.Group>
                            <Form.File  id="exampleFormControlFile1" label="Select picture or video" onChange={fileChange} accept=".mp4,.jpg,.jpeg,.png"/>
                        </Form.Group>
                        <Form.Group>
                            <img className="d-block w-100" height="380"  src={MediaPath} alt="MEDIA"/>
                        </Form.Group>
                    </Form>
                    <label>Close friends:  </label>
                    <br/>
                    <BootstrapSwitchButton
                        width={200}
                        checked={false}
                        onlabel='For close friends'
                        onChange={closeChanged}
                        offlabel='For all'/>
                        <br/>
                    <Button onClick={addStory}>ADD STORY</Button>
                    </Col>
                    <Col></Col>
                </Row>
            </Container>
            
        </div>
    )
}
