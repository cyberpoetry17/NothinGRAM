import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import {Form,Container,Row,Col,Button,Carousel} from 'react-bootstrap';
import Story from './Story';
import BootstrapSwitchButton from 'bootstrap-switch-button-react'
import {serviceConfig} from '../../applicationSettings.js'

export default function StoryForUser() {
    const [stories, setStories] = useState(null);

    useEffect(()=>{
        var userId = jwt_decode(localStorage.getItem('token')).UserID;
        axios({
            method : 'get',
            url :`${serviceConfig.postURL}/getUserStories/`+userId,
        }).then(res =>{
            console.log(res.data," media for story");
            setStories(res.data);
        });
    },[])
    useEffect(()=>{
        console.log("asd",stories);
    },[stories])


    const closeChanged = (id,e)=>{
        if(e == true){
            axios({
                method : 'post',
                url :`${serviceConfig.postURL}/AddToStoryHighlights/`+id
            });
        }else{

        }
        
    }

    return (
        <div>
            <Container>
            <div className="containerForStory">

            
            {stories?.map((s)=>(
              <div className="boxStory">
                <Story UserId={s.UserId} IdStory={s.IdStory} postId={s.PostID} type={s.Type} size="2" ShowOnStoryHighlights={s.ShowOnStoryHighlights}/> 
              </div>  
              
            ))}
            </div>
            </Container>
        </div>
    )
}
