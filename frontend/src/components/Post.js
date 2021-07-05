import React from 'react';
import { app } from './base';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import { Like } from './Like';
import {Dislike} from './Dislike';
import "../styles/post-style.css";

export default function Post({userid,postid}){

    var [liked,setState] = React.useState(false);

    const GetLikesForPost = (postid) =>{                                                         //jos nisam namestio skroz da
        axios.get('http://localhost:8005/alllikesforpost/'+postid).then((response)=>{
            const data = response.data;
            this.setState({likes:data});
        })
        .catch(()=>{alert('didnt retrieve likes')});
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

    const LikeThisPost = () =>{
        CheckIfUserLikedPost();
    }

    const DislikeThisPost = () =>{
        CheckIfUserDislikedPost();
    }

    return(
        <>
        <div className="post">
            <div className="post__header">
                <div className="post__headerLeft">
                    <h3>{userid}(bice slika i ime)</h3>
                    <h3 style={{marginLeft:"8px"}}>{postid}(vrv se sklanja)</h3>
                </div>
                <button className="like_but" onClick={LikeThisPost}>Like</button><p>17(ubaciti metodom)</p>
                <button className="dislike_but" onClick={DislikeThisPost}>Dislike</button><p>17(isto metodom)</p>
            </div>
            <div className="post__body">
                <img className="postImg"/>OVDE DODATI SLIKU
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