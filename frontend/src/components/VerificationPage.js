import React, { useState } from 'react';
import axios from 'axios';
import { useEffect } from "react";
import {serviceConfig} from '../applicationSettings'
import VerificationRequests from './VerificationRequests';

export default function VerificationPage() {
    
    var [verificationRequests, setVerificationRequests] = useState()

    useEffect(()=>{
        GetVerificationRequests()
    },[])

    const GetVerificationRequests = () => {
        axios.get(`${serviceConfig.baseURL}/waitlistedverificationrequests`).then((responsenew)=>{
            const data = responsenew.data;
            if(data != null)
                setVerificationRequests(data);
                console.log(data)
            })
            .catch(()=>{alert('Could not retrieve waitlisted verification requests.')})
    }

    return(
        <>
            {verificationRequests?.map((request, i) => (
            <div className="verReq" key={i}>
                <VerificationRequests request={request}/>
            </div>
            ))}
        </>
        );
}