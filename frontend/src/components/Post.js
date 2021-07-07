import React from 'react';
import { app } from './base';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import { Like } from './Like';
import {Dislike} from './Dislike';
import "../styles/post-style.css";

export default function Post({userid,postid,picpath}){
    
    var [username,setUsername] = React.useState();
    var [likes,setLikes] = React.useState(0);
    var [dislikes,setDislikes] = React.useState(0);

    const GetUsernameByUserId = () =>{
        return axios.get('http://localhost:8005/getusernamebyid/'+userid).then((response) =>{
            setUsername(response.data.substring(1,(response.data.length)-2));
        });
    }
    
    const GetLikesForPost = () =>{                                                         //jos nisam namestio skroz da
        return axios.get('http://localhost:8005/alllikesforpost/'+postid).then((response)=>{
            setLikes(response.data)
        });
    }

    const GetDislikesForPost = () =>{                                                         //jos nisam namestio skroz da
        return axios.get('http://localhost:8005/alldislikesforpost/'+postid).then((response)=>{
            setDislikes(response.data)
        });
    }

    const CheckIfUserLikedPost = () =>{
        axios({method:'post',url:'http://localhost:8005/checkiflikedbyuser',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})}).then((response)=>{
            if(response.data == false){             //moralo je ovde umesto u likethispost()
                axios({method:'post',url:'http://localhost:8005/createlike',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})});
            }else if(response.data == true){
                alert("You already liked this post");
            }
        });

    }

    const CheckIfUserDislikedPost = () => {
        axios({method:'post',url:'http://localhost:8005/checkifdislikedbyuser',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})}).then((response)=>{
            if(response.data == false){
                axios({method:'post',url:'http://localhost:8005/createdislike',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})});
            }else if (response.data == true){
                alert("You already disliked this post");
            }
        });
    }

    const DoINeedToRemoveDislike = () => {
        axios({method:'post',url:'http://localhost:8005/checkifdislikedbyuser',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})}).then((response)=>{
            if(response.data == true){
                axios({method:'post',url:'http://localhost:8005/deletedislike',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})});
            }
        });
    }

    const DoINeedToRemoveLike = () => {
        axios({method:'post',url:'http://localhost:8005/checkiflikedbyuser',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})}).then((response)=>{
            if(response.data == true){
                axios({method:'post',url:'http://localhost:8005/deletelike',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})});
            }
        });
    }

    const LikeThisPost = () =>{
        DoINeedToRemoveDislike();
        CheckIfUserLikedPost();
    }

    const DislikeThisPost = () =>{
        DoINeedToRemoveLike();
        CheckIfUserDislikedPost();
    }

    React.useEffect(()=>GetUsernameByUserId(),[])
    React.useEffect(()=>GetLikesForPost(),[])
    React.useEffect(()=>GetDislikesForPost(),[])
        return(
            <>
            <div className="post">
                <div className="post__header">
                    <div className="post__headerLeft">
                        <Link to={"/profile/"+userid}>{username}</Link>
                        <h3 style={{marginLeft:"8px"}}>{postid}(vrv se sklanja)</h3>
                    </div>
                <button className="like_but" onClick={LikeThisPost}>Like</button><p>{likes}</p>
                <button className="dislike_but" onClick={DislikeThisPost}>Dislike</button><p>{dislikes}</p>
            </div>
            <div className="post__body">
                <img className="postImg" src={picpath} width="100" height="400"/>OVDE DODATI SLIKU
            </div>
            <div >
                <p>Comments</p>
                <p style={{backgroundColor:"white"}}>
                    <span style={{fontWeight:"500",marginRight:"6px"}}>
                        Ovde ubaciti ime metodom
                    </span>
                        Komentar isto dobaviti
                </p>
            </div>
        </div>
        </>
    );

}