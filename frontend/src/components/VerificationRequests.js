import React from 'react'
import "../styles/post-style.css";
import axios from 'axios';
import { serviceConfig } from '../applicationSettings';

export default function VerificationRequests({verificationRequest}){

    const AcceptVerificationRequest = () => {
        const request = {name:verificationRequest.name,
                        surname:verificationRequest.surname,
                        username:verificationRequest.username,
                        category:verificationRequest.category,
                        status:verificationRequest.status,
                        mediapath:verificationRequest.mediapath
        }
        axios.post(`${serviceConfig.userURL}/acceptverification`, request).then(()=>window.location.reload())
    }

    const DeclineVerificationRequest = () => {
        const request = {name:verificationRequest.name,
                        surname:verificationRequest.surname,
                        username:verificationRequest.username,
                        category:verificationRequest.category,
                        status:verificationRequest.status,
                        mediapath:verificationRequest.mediapath
        }
        axios.post(`${serviceConfig.postURL}/declineverification/`, request).then(()=>window.location.reload())
    }

    return(
        <>
        <div className="post__center">
            <div className="post">
                <div className="post__header">
                    <div className="post__headerLeft">
                        <h3>{verificationRequest.username}</h3>
                        <h2 style={{marginLeft:"8px"}}>{verificationRequest.name}</h2>
                        <h2 style={{marginLeft:"8px"}}>{verificationRequest.surname}</h2>
                        <button style={{marginLeft:"8px"}} onClick={AcceptVerificationRequest}>Accept</button>
                        <button style={{marginLeft:"8px"}} onClick={DeclineVerificationRequest}>Decline</button>
                    </div>
                </div>
            </div>
        </div>
        </>
    )
}