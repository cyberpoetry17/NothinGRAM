import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';
import Story from './Story';
import "../styles/story.css";

export default function Stories() {
    const [stories, setStories] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const [storiesMap, setStoriesMap] = useState(null);
    const [firstTime, setFirstTime]= useState(true);
    useEffect(()=>{
        axios({
            method : 'get',
            url :'http://localhost:8005/getAllStories',
        }).then(res =>{
            console.log(res.data,"res.data")
            setStories(res.data)
        });
    },[])

    useEffect(()=>{
        if(firstTime){
            setFirstTime(false)
            return;
        }
        setLoaded(true);
    },[storiesMap])

    useEffect(()=>{
        console.log("stories:",stories)
        if(stories != null){
            var map=stories.reduce((acc, curr) => {
                const isArray = acc[curr.UserId];
                if (isArray) acc[curr.UserId].push(curr);
                else acc[curr.UserId] = [curr];
                return acc;
            }, {})
            setStoriesMap(map)
            console.log("mapa:",map)
        }
    },stories)
    
    return (
        <div>
            <div className="container">
                { loaded ?
                stories.map(s=>(
                    <div className="box">
                        <Story UserId={s.UserId} IdStory={s.IdStory} postId={s.PostID} type={s.Type} /> 
                    </div>
                )):
                <p>loading..</p>
                }
                {/* <button onClick={click}>Click me</button> */}
            </div>
        </div>
    )
}
