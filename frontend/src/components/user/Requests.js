import { app } from '../base';
import {BrowserRouter, Link, Redirect, Route, Switch, useHistory} from 'react-router-dom'
import {Carousel} from 'react-bootstrap';
import axios from 'axios';
import React,{useState,useEffect} from 'react'
import jwt_decode from 'jwt-decode';
import "../../styles/post-style.css";
import Comment from '../Post/Comment';
import CommentInput from '../Post/CommentInput';
import { Ellipsis } from 'react-bootstrap/esm/PageItem';
import {serviceConfig} from '../../applicationSettings'

export default function Requests({idfollower,requestid}){
    var [username,setUsername] = useState("")

    useEffect(()=>GetUserByUserId())

    const GetUserByUserId = () =>{
        axios.get(`${serviceConfig.userURL}/username/`+idfollower).then((response)=>{
            setUsername(response.data)
        })
        .catch(()=>{alert('didnt retrieve user')});
    }

    const AcceptRequest = () => {
        var userfollowing = jwt_decode(localStorage.getItem('token')).UserID;
        axios.get(`${serviceConfig.userURL}/getuseridandprivatebyusername/`+username).then((response)=>{
        const data = response.data
        console.log(data.Private)
        axios.post(`${serviceConfig.userURL}/getfollowstatus`,JSON.stringify({idfollower:userfollowing,iduser:idfollower})).then((responsestatus)=>{
        if(responsestatus.data == false){
            if(data.Private == false){
                axios.post(`${serviceConfig.userURL}/follow`,JSON.stringify({idfollower:idfollower,iduser:userfollowing}))     //zavrsi follow onoga ko je poslao i followuje nazad automatski
                axios.post(`${serviceConfig.userURL}/follow`,JSON.stringify({idfollower:userfollowing,iduser:idfollower}))
                axios.get(`${serviceConfig.userURL}/deleterequest/`+requestid)
                window.location.reload();
            }else{
                axios.post(`${serviceConfig.userURL}/createfollowrequest`,JSON.stringify({idfollower:userfollowing,idfollowed:idfollower}))
                axios.post(`${serviceConfig.userURL}/follow`,JSON.stringify({idfollower:idfollower,iduser:userfollowing}))
                axios.get(`${serviceConfig.userURL}/deleterequest/`+requestid)
                window.location.reload();
            }
        }else{
            axios.post(`${serviceConfig.userURL}/follow`,JSON.stringify({idfollower:idfollower,iduser:userfollowing}))
            axios.get(`${serviceConfig.userURL}/deleterequest/`+requestid)
            window.location.reload();
        }
        })
    })
    }

    const DeleteRequest = () =>{
        axios.get(`${serviceConfig.userURL}/deleterequest/`+requestid).then(()=>{window.location.reload();})
    }

    return(
        <>
        <div className="post__center">
            <div className="post">
                <div className="post__header">
                    <div className="post__headerLeft">
                        <h3>{username}</h3>
                        <button onClick={AcceptRequest}>Accept</button>
                        <button onClick={DeleteRequest}>Decline</button>
                    </div>
                </div>
            </div>
        </div>
        </>
    )

}