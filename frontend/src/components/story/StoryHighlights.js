import React from 'react'
import {useState,useEffect} from 'react'
import jwt_decode from 'jwt-decode';
import axios from 'axios';
import Story from './Story';
import StoryGroup from './StoryGroup';
import {serviceConfig} from '../../applicationSettings.js'

export default function StoryHighlights({userId}) {
    const [stories, setStories] = useState(null);
    useEffect(()=>{
        // var userId = jwt_decode(localStorage.getItem('token')).UserID;
        axios({
            method : 'get',
            url :`${serviceConfig.postURL}/GetAllStoryHighlights/`+userId,
        }).then(res =>{
            console.log(res.data," media for story");
            setStories(res.data);
        });
    },[])
    return (
        <div className="containerStory">
            {stories?.map((s)=>(
              <div>
                  <div className="box"> <StoryGroup storyList={[s]} ForCloseFriends={false}/></div>
              </div>  
            ))}
        </div>
    )
}
