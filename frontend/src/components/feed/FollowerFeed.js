import Post from "../Post/Post";
import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import Stories from "../story/Stories";

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
        axios.get('http://localhost:8080/api/user/getallfollowedforloggeduser/'+jwt_decode(localStorage.getItem('token')).UserID).then((response)=>{
            
            response.data?.map((follow) =>(     
                axios.get('http://localhost:8080/api/post/allpostsbyuserid/'+follow).then((responsenew)=>{
                const data = responsenew.data;
                console.log(data)
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