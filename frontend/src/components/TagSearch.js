import axios from "axios";
import { useState } from "react";
import { useHistory, useParams } from "react-router-dom";
import Profile from "./user/Profile";
import Post from './Post/Post';
import { useEffect } from "react";
import {serviceConfig} from '../applicationSettings'

export default function TagSearch(){

    const {tag} = useParams()
    const [tagPosts,setTags] = useState()

    useEffect(() => {
        GetAllPosts();
    },[tag])

    const GetAllPosts = () =>{
        axios.get(`${serviceConfig.postURL}/postsbytags/`+tag).then((response) => {
            console.log(tag)
            setTags(response.data);
            console.log(response.data)
        });
    }

    return (
        <div>
        {tagPosts?.map((post,i) => (
            <div className="feed" key={i}>
                <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private}  description={post.description} location={post.LocationID}/>
            </div>
        ))}
        </div>
    )
}