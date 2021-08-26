import axios from "axios";
import { useState } from "react";
import { useEffect } from "react";
import Post from '../Post/Post'
import {serviceConfig} from '../../applicationSettings.js'

export default function ReportedFeed() {
    var [reports,setReports] = useState()

    useEffect(()=>{
        GetAllPosts()
    },[])

    const GetAllPosts = () => {
        axios.get(`${serviceConfig.postURL}/getallreported`).then((responsenew)=>{
        const data = responsenew.data;
        if(data != null)
            setReports(data);
            console.log(data)
        })
        .catch(()=>{alert('didnt retrieve reports')})
    }

    const DeletePost = (postid) =>{
        axios.post(`${serviceConfig.postURL}/deletepost/`+postid).then(()=>window.location.reload());
    }

    const DeleteProfile = (userid) =>{
        axios.post(`${serviceConfig.userURL}/deleteprofile/`+userid).then(()=>window.location.reload());
    }

    return(
    <>
        {reports?.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private} description={post.description} location = {post.LocationID} />
            <div className="post__header">
            <button className="like_but" onClick={() => DeletePost(post.ID)}>Delet</button>
            <button className="like_but" onClick={() => DeleteProfile(post.userid)}>Delet profile</button>
            </div>
        </div>
        ))}
    </>
    );

}