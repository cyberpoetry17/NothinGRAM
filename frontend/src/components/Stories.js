import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';
import Story from './Story';

export default function Stories() {
    const [stories, setStories] = useState(null);
    
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
        console.log("stories:",stories)
    },stories)

    const click = (event) => {
        console.log("stories:",stories)
    };
    
    return (
        <div>
            {stories?.map(s=>(
                <Story IdStory={s.IdStory} postId={s.postId} type={s.type} /> 
                
            ))}
            {/* <button onClick={click}>Click me</button> */}
        </div>
    )
}
