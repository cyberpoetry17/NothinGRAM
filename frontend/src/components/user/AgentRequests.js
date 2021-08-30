import React,{useState,useEffect} from 'react'
import jwt_decode from 'jwt-decode';
import "../../styles/post-style.css";
import axios from 'axios';
import { serviceConfig } from '../../applicationSettings';

export default function AgentRequests({agent}){

    

    const AcceptRequest = () => {
        const agentData = {name:agent.name,surname:agent.surname,email:agent.email,username:agent.username,password:agent.password,date:agent.date,gender:agent.gender,
            phone:agent.phone,bio:agent.bio,web:agent.web,verify:agent.verify,private:agent.private,taggable:agent.taggable,notif:agent.notif,link:agent.link}
        axios.post(`${serviceConfig.userURL}/confirmagent`,agentData).then(()=>{
            DeleteRequest();
        }).then(()=>window.location.reload())
    }

    const DeleteRequest = () => {
        axios.post(`${serviceConfig.postURL}/removeAgent/`+agent.email).then(()=>window.location.reload())
    }

    return(
        <>
        <div className="post__center">
            <div className="post">
                <div className="post__header">
                    <div className="post__headerLeft">
                        <h3>{agent.name}</h3>
                        <h3 style={{marginLeft:"8px"}}>{agent.email}</h3>
                        <h3 style={{marginLeft:"8px"}}>{agent.link}</h3>
                        <button style={{marginLeft:"8px"}} onClick={AcceptRequest}>Accept</button>
                        <button style={{marginLeft:"8px"}} onClick={DeleteRequest}>Decline</button>
                    </div>
                </div>
            </div>
        </div>
        </>
    )
}