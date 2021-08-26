import axios from "axios";
import { useState } from "react";
import { useEffect } from "react";
import Post from '../Post/Post'

export default function ReportedFeed() {
    var [reports,setReports] = useState()

    useEffect(()=>{
        GetAllPosts()
    },[])

    const GetAllPosts = () => {
        axios.get('http://localhost:8082/getallreported').then((responsenew)=>{
        const data = responsenew.data;
        if(data != null)
            setReports(data);
            console.log(data)
        })
        .catch(()=>{alert('didnt retrieve reports')})
    }

    const DeletePost = (postid) =>{
        axios.post('http://localhost:8082/deletepost/'+postid).then(()=>window.location.reload());
    }

    const DeleteProfile = (userid) =>{
        axios.post('http://localhost:8081/deleteprofile/'+userid).then(()=>window.location.reload());
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