import Post from "../Post/Post";
import React, { useState } from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import { useEffect } from "react";
import Requests from "./Requests";
import {serviceConfig} from '../../applicationSettings'
import AgentRequests from "./AgentRequests";

export default function AgentConfirmationPage() {
    
    var [agents,setAgents] = useState()

    useEffect(()=>{
        GetAgentRequests()
    },[])

    const GetAgentRequests = () => {
        axios.get(`${serviceConfig.postURL}/getallagents`).then((responsenew)=>{
            const data = responsenew.data;
            if(data != null)
                setAgents(data);
                console.log(data)
            })
            .catch(()=>{alert('didnt retrieve agent requests')})
    }

    return(
        <>
            {agents?.map((agent,i) => (
            <div className="feed" key={i}>
                <AgentRequests agent={agent}/>
            </div>
            ))}
        </>
        );
}