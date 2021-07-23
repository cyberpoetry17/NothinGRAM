import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';

export default function Story({postId,type,IdStory}) {
    //retriew media and render it
    //add userID to story
    const [media, setMedia] = useState(null);
    const [isLoaded,setIsLoaded] = useState(false);
    const [firstTime, setFirstTime] = useState(true)

    useEffect(()=>{
        axios({
            method : 'get',
            url :'http://localhost:8005/GetMediaForStory?StoryId='+IdStory,
        }).then(res =>{
            console.log(res.data," media for story")
            setMedia(res.data)
        });
    },[])

    useEffect(()=>{
        if(firstTime){
            setFirstTime(false)
            return;
        }
        console.log(media);
        setIsLoaded(true);
    },[media])

    return (
        <div>
            <p>Story</p>
            {isLoaded === true? 
            <img className="d-block w-100" width="100" height="100" src={media.Link} alt="my pic"/>
            :
            <p>no pic</p>
            }
        </div>
    )
}
