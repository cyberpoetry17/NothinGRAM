import { app } from './base';
import {BrowserRouter, Link, Redirect, Route, Switch, useHistory} from 'react-router-dom'
import axios from 'axios';
import React,{useState,useEffect} from 'react'
import { Like } from './Like';
import {Dislike} from './Dislike';
import jwt_decode from 'jwt-decode';
import "../styles/post-style.css";

export default function Comment({comments,posteduser}){

    var [commentuser,setCommentUser] = React.useState("")

    useEffect(()=>GetUserByUserId(),[])
    
    const GetUserByUserId = () =>{
        axios.get('http://localhost:8004/username/'+posteduser).then((response)=>{
            const data = response.data;
            setCommentUser(data)
        })
        .catch(()=>{alert('didnt retrieve user')});
    }

    return(
        <div>
            <p style={{backgroundColor:"white"}}>
                    <span style={{fontWeight:"500",marginRight:"6px"}}>
                        {commentuser}
                    </span>
            {comments}
            </p>
        </div> 
    )
}