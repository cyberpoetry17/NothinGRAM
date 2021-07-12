import axios from 'axios';
import React,{useState,useEffect} from 'react'
import { app } from './base';
import {Form,Container,Row,Col,Button} from 'react-bootstrap';
import TextField from '@material-ui/core/TextField';
import jwt_decode from 'jwt-decode';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogActions from '@material-ui/core/DialogActions';
import Autocomplete, { createFilterOptions } from '@material-ui/lab/Autocomplete';
import { useHistory } from "react-router-dom";
import PlacesAutocomplete, {
    geocodeByAddress,
    getLatLng
  } from "react-places-autocomplete";

const filter = createFilterOptions();

function AddPost() {
    const [post, setPost] = useState({description:'', picpath:'', private:true, UserID:"",city:"",country:"",LocationID:"",tags:[]})
    const [tagNames, setTagNames] = useState([])
    const [location,setLocation] = useState({city:"",address:"",country:""})
    const [open, toggleOpen] = useState(false);
    const [value, setValue] = useState(null);
    const [dialogValue, setDialogValue] = useState({ TagName: '',id :'',Posts:[]});
    const [firstTime, setFristTime]= useState(true)
    
    const handleClose = () => {
        setDialogValue({
            TagName: ''
        });
        toggleOpen(false);
        setValue([]);
      };
    
    const handleSubmit = (event) => {
        event.preventDefault();
        axios({
            method : 'post',
            url :'http://localhost:8005/addTag/',
            data:JSON.stringify({
                id:"00000000-0000-0000-0000-000000000000",
                TagName: dialogValue.TagName,
                Posts: []
            }),
        }).then(()=>{
            getAllTags()
        });
        handleClose();
    };
    var tokenInfo;
    let history = useHistory()

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
        console.log("tagovi izmenjeni")
        console.log(post.tags)
    },[post.tags])
    useEffect(()=>{
        console.log("kao radi")
        console.log(tagNames)
    },[tagNames])
    useEffect(()=>{
        console.log("lokacija")
        console.log(location)
    },[location])
    useEffect(()=>{
        if(firstTime){
            setFristTime(false)
            return
        }
        console.log("post:")
        console.log(JSON.stringify(post))
        axios({
            method : 'post',
            url :'http://localhost:8005/createpost',
            data:JSON.stringify(post),
        }).then(()=>{
            history.push("/posts")
        });
    },[post.LocationID])
    useEffect(()=>{
        if(localStorage.getItem('token')== null){
            alert("You need to log in to add post !!!")
            history.push("/login")
            return
        }
        tokenInfo = jwt_decode(localStorage.getItem('token'));
        if(tokenInfo==""){
            alert("You need to log in to add post !!!")
            history.push("/login")
            return
        }
        post.UserID = tokenInfo.UserID
        getAllTags()
    },[])
    const getAllTags = ()=>{
        axios({
            method : 'get',
            url :'http://localhost:8005/getAllTags'
        }).then(res =>{
            setTagNames(res.data)
        });
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
    const setTags = (e,newV)=>{
        setPost({...post,tags: newV})
    }
    const  add = ()=>{
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
                            onChange={(event, newValue) => {
                                console.log("novooo",newValue)
                                if(newValue.length == 0){
                                    return
                                }
                                if( newValue[newValue.length-1].id !== undefined ){
                                    setTags(event,newValue)
                                }
                                else if (typeof newValue === 'string') {
                                  // timeout to avoid instant validation of the dialog's form.
                                  setTimeout(() => {
                                    toggleOpen(true);
                                    setDialogValue({
                                      TagName: newValue,
                                      id:"00000000-0000-0000-0000-000000000000",
                                      post:[]
                                      
                                    });
                                  });
                                } else if (newValue && newValue[newValue.length-1].TagName) {
                                  toggleOpen(true);
                                  setDialogValue({
                                    TagName: newValue[newValue.length-1].TagName,
                                    id:"00000000-0000-0000-0000-000000000000",
                                    post:[]
                                  });
                                } else {
                                    setTags(newValue);
                                    console.log(newValue)
                                }
                                for( var i = 0; i<newValue.length; i++){
                                    if(newValue[i].id === undefined){
                                        newValue.splice(i, 1); 
                                    }
                                }
                            }}
                            filterOptions={(options, params) => {
                            const filtered = filter(options, params);
                    
                            if (params.inputValue !== '') {
                                filtered.push({
                                    TagName: params.inputValue
                                });
                            }
                    
                            return filtered;
                            }}
                            multiple
                            options={tagNames}
                            getOptionLabel={(option) => option.TagName}
                            style={{ width: 300 }}
                            renderInput={(params) => <TextField id="textField12" {...params} label="Hashtag" variant="outlined" />}
                            />
                            <img width="100" height="100" src={post.picpath} alt="my pic"/><br/>
                            <Button onClick={add}>ADD POST</Button>
                        </Form>
                    </Col>
                    <Col></Col>
                </Row>
                <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
                    <form onSubmit={handleSubmit}>
                    <DialogTitle id="form-dialog-title">Add a new hash tag</DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                        Did you miss any hash tag in our list? Please, add it!
                        </DialogContentText>
                        <TextField
                        autoFocus
                        margin="dense"
                        id="name"
                        value={dialogValue.TagName}
                        onChange={(event) => setDialogValue({ ...dialogValue, TagName: event.target.value })}//////
                        label="TagName"
                        type="text"
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose} color="primary">
                        Cancel
                        </Button>
                        <Button type="submit" color="primary">
                        Add
                        </Button>
                    </DialogActions>
                    </form>
                </Dialog>
            </Container>
        </div>
    )
}

export default AddPost
