import axios from "axios";
import { useState } from "react";
import { useHistory, useParams } from "react-router-dom";
import Profile from "./Profile";
import Post from "./Post";
import { useEffect } from "react";

export default function LocationSearch(){

    const {location} = useParams()
    const [locationPosts,setLocationPosts] = useState()

    useEffect(() => {
        GetAllPosts();
    },[location])

    const GetAllPosts = () =>{
        axios.get('http://localhost:8005/postsbylocation/'+location).then((response) => {
            setLocationPosts(response.data);
            console.log(response.data)
        });
    }

    return (
        <div>
        {locationPosts?.map((post,i) => (
            <div className="feed" key={i}>
                <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private}  description={post.description} location={post.LocationID}/>
            </div>
        ))}
        </div>
    )
}