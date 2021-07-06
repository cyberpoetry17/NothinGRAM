import axios from 'axios';
import React,{useState,useEffect} from 'react'
import { app } from './base';
import {Form,Container,Row,Col,Button} from 'react-bootstrap';

function AddPost() {
    const [post, setPost] = useState({description:'', picpath:'', private:true, UserID:"0f608e5a-1e79-4dd3-ba7e-fe99c81e6fe2"})
    const ispisi = ()=>{
        console.log(JSON.stringify(post))
    }

    const setPrivate= ()=>{
        setPost({...post,private:true})
    }

    const setPublic= ()=>{
        setPost({...post,private:false})
    }

    const fileChange = async (e)=>{
        const file = e.target.files[0]
        //console.log("ispis",app.storage().ref())
        const storageRef = app.storage().ref()
        const fileRef = storageRef.child(file.name)
        await fileRef.put(file)
        var a = await fileRef.getDownloadURL()
        console.log(a)
        setPost({...post,picpath:a})
    }

    const add = ()=>{
        var a = JSON.stringify(post)
        axios({
            method : 'post',
            url :'http://localhost:8005/createpost',
            data:a,
        });
    }

    return (
        <div>
            <label className="d-flex justify-content-center">ADD POST</label>
            <Container  width="20">
                <Row>
                    <Col></Col>
                    <Col>
                        <Form className="justify-content-md-center">
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Description</Form.Label>
                                <Form.Control  placeholder="Enter description" onChange={e => setPost({...post, description:e.target.value})}/>
                            </Form.Group>
                            <Form.Group>
                                <Form.File id="exampleFormControlFile1" label="Example file input" onChange={fileChange}/>
                            </Form.Group>
                            <Form.Group>
                                {['radio'].map((type) => (
                                    <div key={`inline-${type}`} className="mb-3">
                                    <Form.Check inline label="PRIVATE" name="group1" type={type} id={`inline-${type}-1`} onChange={setPrivate}/>
                                    <Form.Check inline label="NOT PRIVATE" name="group1" type={type} id={`inline-${type}-2`} onChange= {setPublic}/>
                                    </div>
                                ))}
                            </Form.Group>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Location</Form.Label>
                                <Form.Control  placeholder="Enter location" />
                            </Form.Group>
                            <Form.Group controlId="exampleForm.ControlInput1">
                                <Form.Label>Tag</Form.Label>
                                <Form.Control  placeholder="Enter tag" />
                            </Form.Group>
                            
                            <img width="100" height="100" src={post.picpath} alt="my pic"/><br/>
                            <Button onClick={add}>ADD POST</Button>
                        </Form>
                    </Col>
                    <Col></Col>
                </Row>
                
            </Container>
            
            {/* <input 
            type= 'text' 
            value= {post.description}
            onChange={e => setPost({...post, description:e.target.value})}
            />
            <div>
                <label><input type="radio" name="private"  value='private' onChange={setPrivate}/> Private</label>
                <label><input type="radio" name="private"  value='not private' onChange={setPublic}/>Not private</label>
            </div>
            <input type="file" onChange={fileChange} /><br/>
            <img width="100" height="100" src={post.picpath} alt="my pic"/><br/>
            <button onClick={ispisi}>JSON</button>
            <button onClick={add}>ADD POST</button> */}
        </div>
    )
}

export default AddPost
