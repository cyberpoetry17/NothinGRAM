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

function AddPost() {
    const [post, setPost] = useState({description:'', picpath:'', private:true, UserID:"0f608e5a-1e79-4dd3-ba7e-fe99c81e6fe2",city:"",country:"",LocationID:"",tags:[]})
    const [tagNames, setTagNames] = useState([])
    const [location,setLocation] = useState({city:"",address:"",country:""})
    // const [address,setAddress]= useState("")
    // const [coordinates, setCoordinates] = React.useState({
    //     lat: null,
    //     lng: null
    //   });

    const ispisi = ()=>{
        console.log(JSON.stringify(post))
    }

    const setPrivate= ()=>{
        setPost({...post,private:true})
    }

    const setPublic= ()=>{
        setPost({...post,private:false})
    }
    useEffect(()=>{
        // tags.forEach(element => {
        //     //setS
        // });
        console.log("kao radi")
        console.log(tagNames)
    },[tagNames])
    useEffect(()=>{
        // tags.forEach(element => {
        //     //setS
        // });
        console.log("lokacija")
        console.log(location)
    },[location])
    useEffect(()=>{
        // tags.forEach(element => {
        //     //setS
        // });
        console.log("post:")
        console.log(JSON.stringify(post))
        axios({
            method : 'post',
            url :'http://localhost:8005/createpost',
            data:JSON.stringify(post),
        });
    },[post.LocationID])
    useEffect(()=>{
        axios({
            method : 'get',
            url :'http://localhost:8005/getAllTags'
        }).then(res =>{
            setTagNames(res.data)
        });
    },[])
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
    const setTags = (e,newV)=>{
        setPost({...post,tags: newV})
    }
    const add = ()=>{
        axios({
            method : 'post',
            url :'http://localhost:8005/createlocation',
            data:JSON.stringify(location),
        }).then((res)=>{
            console.log(res.data)
            setPost({...post, LocationID:res.data})
        });
        
    }
    // const handleSelect = async value => {
    //     const results = await geocodeByAddress(value);
    //     const latLng = await getLatLng(results[0]);
    //     setAddress(value);
    //     setCoordinates(latLng);
    //   };

    return (
        <div>
            <h1 className="d-flex justify-content-center">ADD POST</h1>
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
                            {/* <PlacesAutocomplete
                                value={address}
                                onChange={setAddress}
                                onSelect={handleSelect}
                            >
                                {({ getInputProps, suggestions, getSuggestionItemProps, loading }) => (
                                <div>
                                    <p>Latitude: {coordinates.lat}</p>
                                    <p>Longitude: {coordinates.lng}</p>

                                    <input {...getInputProps({ placeholder: "Type address" })} />

                                    <div>
                                    {loading ? <div>...loading</div> : null}

                                    {suggestions.map(suggestion => {
                                        const style = {
                                        backgroundColor: suggestion.active ? "#41b6e6" : "#fff"
                                        };

                                        return (
                                        <div {...getSuggestionItemProps(suggestion, { style })}>
                                            {suggestion.description}
                                        </div>
                                        );
                                    })}
                                    </div>
                                </div>
                                )}
                            </PlacesAutocomplete> */}
                            <div>
                                <Row>
                                    <Col>
                                            <Form.Label>City</Form.Label>
                                            <Form.Control  placeholder="Enter description" onChange={e => setLocation({...location, city:e.target.value})}/>
                                    </Col>
                                    <Col>
                                        <Form.Group controlId="exampleForm.ControlInput1">
                                            <Form.Label>Country</Form.Label>
                                            <Form.Control  placeholder="Enter description" onChange={e => setLocation({...location, country:e.target.value})}/>
                                        </Form.Group>
                                    </Col>
                                    <Col>
                                        <Form.Group controlId="exampleForm.ControlInput1">
                                            <Form.Label>Address</Form.Label>
                                            <Form.Control  placeholder="Enter description" onChange={e => setLocation({...location, address:e.target.value})}/>
                                        </Form.Group>
                                    </Col>
                                </Row>
                            </div>
                            <Autocomplete
                            id="autoComplete1"
                            onChange={setTags}
                            multiple
                            options={tagNames}
                            getOptionLabel={(option) => option.TagName}
                            style={{ width: 300 }}
                            renderInput={(params) => <TextField {...params} label="Hashtag" variant="outlined" />}
                            />
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
