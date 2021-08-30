import React from 'react'
import "../styles/post-style.css";
import axios from 'axios';
import { serviceConfig } from '../applicationSettings';

export default function VerificationRequests({request}){

    const AcceptVerificationRequest = () => {
        const requestData = {name:request.name,
                        surname:request.surname,
                        username:request.username,
                        category:request.category,
                        status:request.status,
                        mediapath:request.mediapath
        }
        axios.post(`${serviceConfig.baseURL}/acceptverification`, requestData).then(()=>window.location.reload())
    }

    const DeclineVerificationRequest = () => {
        const requestData = {name:request.name,
                        surname:request.surname,
                        username:request.username,
                        category:request.category,
                        status:request.status,
                        mediapath:request.mediapath
        }
        axios.post(`${serviceConfig.baseURL}/declineverification`, requestData).then(()=>window.location.reload())
    }

    return(
        <>
        <div className="post__center">
            <div className="post">
                <div className="post__header">
                    <div className="post__headerLeft">
                        <img width="100" height="100" src={request.mediapath}/>
                        <h3>{request.username}</h3>
                        <h5 style={{marginLeft:"8px"}}>{request.name}</h5>
                        <h5 style={{marginLeft:"8px"}}>{request.surname}</h5>
                        <button style={{marginLeft:"8px"}} onClick={AcceptVerificationRequest}>Accept</button>
                        <button style={{marginLeft:"8px"}} onClick={DeclineVerificationRequest}>Decline</button>
                    </div>
                </div>
            </div>
        </div>
        </>
    )
}