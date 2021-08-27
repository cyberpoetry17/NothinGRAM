import React from 'react';
// import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import Post from '../Post/Post';
import {BrowserRouter, Link, Redirect, Route, Switch, useHistory} from 'react-router-dom'
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import {useState,useEffect} from 'react'
import queryString from 'query-string';
import {Nav} from 'react-bootstrap';
import StoryHighlights from '../story/StoryHighlights';
import {serviceConfig} from '../../applicationSettings.js'

export class Profile extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          userid:"",
          posts: [],
          user:[],
          followed:false,
          isMyProfile:false
        };
      }
    async componentDidMount(){
        await this.IsProfileFollowedByLoggedUser();
        this.GetUserByUserId();
        this.GetAllPostsForUserDependingOnFollowage();
        
    }

    async IsProfileFollowedByLoggedUser(){
        if (window.localStorage.getItem('token') != null){
            var userfollowing = jwt_decode(localStorage.getItem('token')).UserID;
            await axios.get(`${serviceConfig.userURL}/getuseridandprivatebyusername/`+this.props.match.params.username).then((response)=>{
                const data = response.data
                axios.post(`${serviceConfig.userURL}/getfollowstatus`,JSON.stringify({idfollower:userfollowing,iduser:data.UserId})).then((response)=>{
                this.setState({followed:response.data})
                })
            })
        }  
    }

    GetUserByUserId(){
        axios.get(`${serviceConfig.userURL}/getuserbyusername/`+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({user:data});
        })
        .catch(()=>{alert('didnt retrieve user')});
    }

    GetAllPostsForUserDependingOnFollowage(){
        axios.get(`${serviceConfig.userURL}/getuseridandprivatebyusername/`+this.props.match.params.username).then((response)=>{
            this.setState({userid:response.data.UserId})
            if (window.localStorage.getItem('token') != null){
                if(response.data.UserId == jwt_decode(localStorage.getItem('token')).UserID){
                    this.setState({isMyProfile:true});
                }
                if((this.state.followed == true || this.props.match.params.username==jwt_decode(localStorage.getItem('token')).Username) ||
                (response.data.Private == false && this.state.followed==false && this.props.match.params.username!=jwt_decode(localStorage.getItem('token')).Username)){
                    axios.get(`${serviceConfig.postURL}/allpostsbyuserid/`+this.state.userid).then((response)=>{
                    const data = response.data;
                    this.setState({posts:data})
                    }).catch(()=>alert('didnt retrieve all posts for user'))
                }
            }
            else if (response.data.Private == false && this.state.followed==false){

                axios.get(`${serviceConfig.postURL}/allpostsbyuserid/`+this.state.userid).then((response)=>{
                    const data = response.data;
                    this.setState({posts:data})
                    }).catch(()=>alert('didnt retrieve all posts for user not logged'))
                }
        }).catch(()=>{alert('didnt retrieve user by username')});
    }

    FollowUser(event){
        if (window.localStorage.getItem('token') != null){
        var userfollowing = jwt_decode(localStorage.getItem('token')).UserID;
        axios.get(`${serviceConfig.userURL}/getuseridandprivatebyusername/`+this.props.match.params.username).then((response)=>{
            const data = response.data
            if(this.state.followed == false){
                if(data.Private == false){
                axios.post(`${serviceConfig.userURL}/follow`,JSON.stringify({idfollower:userfollowing,iduser:data.UserId})).then(
                ()=>{
                    alert('You have followed user with userid' + this.state.userid);
                    this.setState({followed:true});
                }).then(()=>this.GetAllPostsForUserDependingOnFollowage())
                }else{
                    axios.post(`${serviceConfig.userURL}/createfollowrequest`,JSON.stringify({idfollower:userfollowing,idfollowed:data.UserId})).then(
                ()=>{
                    alert('Follow request sent');
                })
                }
            }else if (this.state.followed == true){
                axios.post(`${serviceConfig.userURL}/unfollow`,JSON.stringify({idfollower:userfollowing,iduser:data.UserId})).then(
                ()=>{
                    alert('You have unfollowed user with userid' + this.state.userid);
                    this.setState({followed:false});
                }).then(()=>this.GetAllPostsForUserDependingOnFollowage())
            }
        }).catch(()=>{alert('didnt retrieve user by username')});
    }else{
        alert('You are not logged in.You will be redirected to the login page.');
        this.props.history.push('/login');
    }
    }

    render(){
        const data = this.state.posts;
        const user = this.state.user;
        const userid = this.state.userid;
        return(
            <>
            <div className="profile">
                    <div className="post__headerLeft">
                        <div className="profile__header">
                            <h1 style={{marginLeft:"8px"}}>{user.name}</h1>
                            {(window.localStorage.getItem('token') != null && this.props.match.params.username==jwt_decode(localStorage.getItem('token')).Username ) ?
                                <div style={{marginLeft:"8px",fontWeight:'normal'}}>
                                <BrowserRouter>
                                <Nav className="navbarprofile" activeKey="/" >
                                    <Nav.Item>
                                        <Nav.Link href="/addPost">Add post</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item>
                                        <Nav.Link href="/addStory">Add story</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item >
                                        <Nav.Link href={"/userinteracted/"+jwt_decode(localStorage.getItem('token')).Username}>Liked/disliked content</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item >
                                        <Nav.Link href="/verification">User Verification</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item >
                                        <Nav.Link href="/update">Update</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item >
                                        <Nav.Link href={"/requests/"+jwt_decode(localStorage.getItem('token')).Username}>Requests</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item >
                                        <Nav.Link href={"/closefollowerr"}>Close friends</Nav.Link>
                                    </Nav.Item>
                                    <Nav.Item >
                                        <Nav.Link href={"/stories"}>Stories</Nav.Link>
                                    </Nav.Item>
                                    {jwt_decode(localStorage.getItem('token')).Role === 1 ? 
                                     <Nav.Item >
                                     <Nav.Link href={"/agentrequests"}>Agent Requests</Nav.Link>
                                     </Nav.Item>
                                        :
                                     null 
                                    }
                                   
                                </Nav>
                                </BrowserRouter>
                                </div>
                            :
                                <button className="follow_but" onClick={this.FollowUser.bind(this)}>{this.state.followed ? "Unfollow" : "Follow"} </button>
                            }
                            
                        </div>
                    </div>
                </div>
            {(this.state.userid != "" && this.state.followed) || this.state.isMyProfile?
            <StoryHighlights userId={this.state.userid}/>:
            null
            }
        
        {data?.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private} description={post.description} location = {post.LocationID}/>
        </div>
        ))}
        </>
        )
    }
}
export default Profile