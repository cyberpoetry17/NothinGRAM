import { app } from '../base';
import {BrowserRouter, Link, Redirect, Route, Switch, useHistory} from 'react-router-dom'
import axios from 'axios';
import React,{useState,useEffect} from 'react'
import jwt_decode from 'jwt-decode';
import "../../styles/post-style.css";
import { getAlgorithms } from 'json-web-token';
import {serviceConfig} from '../../applicationSettings'

export default function CommentInput({postid,getcoms}){
    var [comment,setComment] = React.useState("");
    let history = useHistory();

    const PostComment = () => {
        if (window.localStorage.getItem('token') != null){
            axios({method:'post',url:`${serviceConfig.postURL}/addComment/`,data:JSON.stringify({Comment:comment,UserId:jwt_decode(localStorage.getItem('token')).UserID,PostId:postid})}).then(()=>{
                setComment('');
            }).then(()=>{
                getcoms();
            });
        }else{
            alert("You are not logged in.You will be redirected to login.");
            history.push('/login');
        }
    }

    return <div className="comment_input">
        <textarea className="comment_textarea" name="" id="" cols="30" rows="1" placeholder="Write a comment.." value={comment} onChange={(e)=>setComment(e.target.value)}>

        </textarea>
        <button className="comment_but" onClick={PostComment}>Post</button>
    </div>
}