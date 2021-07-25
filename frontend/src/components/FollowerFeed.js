import Post from "./Post";
import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import Stories from "./Stories";

export class UserInteractedContent extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          followerposts: [],
          followers:[]
        };
      }

    async componentDidMount(){
        await this.GetAllFollowerIds();
    }

    GetAllFollowerIds(){
        axios.get('http://localhost:8004/getallfollowedforloggeduser/'+jwt_decode(localStorage.getItem('token')).UserID).then((response)=>{
            
            response.data?.map((follow) =>(      //kao da nece da settuje state od dobavljenih podataka pa moram jednu u drugoj da pozivam
                axios.get('http://localhost:8005/allpostsbyuserid/'+follow).then((responsenew)=>{
                const data = responsenew.data;
                if(data != null)
                this.setState({followerposts:this.state.followerposts.concat(data)});
            })
            .catch(()=>{alert('didnt retrieve ')})
            ))
        }).catch(()=>{alert('You have not followed any other users.')})
    }

      render(){
        return(
            <>
        <Stories/>
        {this.state.followerposts?.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private} description={post.description} location = {post.LocationID}/>
        </div>
        ))}
        </>
        )
    }
}
export default UserInteractedContent