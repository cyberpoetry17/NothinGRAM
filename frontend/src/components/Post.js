import React from 'react';
import { app } from './base';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import { Like } from './Like';
import {Dislike} from './Dislike';
import "../styles/post-style.css";

export default function Post({userid,postid,picpath}){
    // componentDidMount(){
    //     this.GetAllPosts();
    // }
    
    const GetLikesForPost = (postid) =>{                                                         //jos nisam namestio skroz da
        axios.get('http://localhost:8005/alllikesforpost/'+postid).then((response)=>{
            const data = response.data;
            this.setState({likes:data});
            console.log(this.state.likes)
        })
        .catch(()=>{alert('didnt retrieve likes')});
    }

    // GetAllPosts(){
    //     axios.get('http://localhost:8005/').then((response)=>{
    //         const data = response.data;
    //         this.setState({posts:data});
    //         console.log(this.state.posts)
    //     })
    //     .catch(()=>{alert('didnt retrieve ')});
    // }

    const LikeThisPost = () =>{
        axios({method:'post',url:'http://localhost:8005/createlike',headers:{},data:JSON.stringify({userid:"00000000-0000-0000-0000-000000000030",postid})});
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
                    <button className="like_but">Dislike</button><p>17(isto metodom)</p>
                </div>
                <div className="post__body">
                    <img width="300" height="300" src={picpath} alt="my pic"/><br/>
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