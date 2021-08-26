import Post from "../Post/Post";
import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import { useEffect } from "react";
import Requests from "./Requests";

export class UserInteractedContent extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
          requests:[],
          userid:"",
          username:""
        };
      }

    async componentDidMount(){
        await this.GetUserIdByUsername()
        this.GetAllRequests();
    }

    async GetUserIdByUsername(){
        await axios.get('http://localhost:8080/api/user/getuseridandprivatebyusername/'+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({userid:data.UserId});
        })
    }

    GetAllRequests(){
        axios.get('http://localhost:8080/api/user/getallrequests/'+this.state.userid).then((response)=>{
            const data = response.data;
            this.setState({requests:data});
        })
        .catch(()=>{alert('didnt retrieve requests')});
    }

    render(){
        const data = this.state.requests;
        return(
            <>
            
            <h1>Follow requests</h1>
        {data?.map((post,i) => (
        <div className="feed" key={i}>
            <Requests idfollower = {post.idfollower} requestid = {post.id} />
        </div>
        ))}
        </>
        )
    }
}
export default UserInteractedContent