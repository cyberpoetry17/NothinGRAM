import React from 'react';
// import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import Post from './Post';
import {BrowserRouter, Link, Redirect, Route, Switch, useHistory} from 'react-router-dom'
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import {useState,useEffect} from 'react'
import queryString from 'query-string';

export class Profile extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          userid:"",
          posts: [],
          user:[],
          followed:false
        };
      }
    async componentDidMount(){
        await this.IsProfileFollowedByLoggedUser();
        this.GetUserByUserId();
        this.GetAllPostsForUserDependingOnFollowage();
        
    }

    async IsProfileFollowedByLoggedUser(){
        var userfollowing = jwt_decode(localStorage.getItem('token')).UserID;
        await axios.get('http://localhost:8004/getuseridbyusername/'+this.props.match.params.username).then((response)=>{
            const data = response.data
            axios.post('http://localhost:8004/getfollowstatus',JSON.stringify({idfollower:userfollowing,iduser:data})).then((response)=>{
            this.setState({followed:response.data})
            console.log(this.state.followed)
        })
    })
    }

    GetUserByUserId(){
        axios.get('http://localhost:8004/getuserbyusername/'+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({user:data});
        })
        .catch(()=>{alert('didnt retrieve user')});
    }

    GetAllPostsForUserDependingOnFollowage(){
        axios.get('http://localhost:8004/getuseridbyusername/'+this.props.match.params.username).then((response)=>{
            this.setState({userid:response.data})
            console.log(this.state.followed)
            if(this.state.followed == false){
                axios.get('http://localhost:8005/getnonprivateposts/'+response.data).then((response)=>{
                    const data = response.data;
                    this.setState({posts:data});
                }).catch(()=>{alert('didnt retrieve ')});
            }else if(this.state.followed == true){
                axios.get('http://localhost:8005/allpostsbyuserid/'+this.state.userid).then((response)=>{
                const data = response.data;
                this.setState({posts:data})
                })
            }
        }).catch(()=>{alert('didnt retrieve user by username')});
    }

    FollowUser(event){
        var userfollowing = jwt_decode(localStorage.getItem('token')).UserID;
        axios.get('http://localhost:8004/getuseridbyusername/'+this.props.match.params.username).then((response)=>{
            const data = response.data
            if(this.state.followed == false){
                axios.post('http://localhost:8004/follow',JSON.stringify({idfollower:userfollowing,iduser:data})).then(
                ()=>{
                    alert('You have followed user with userid' + this.state.userid);
                    this.setState({followed:true});
                }).then(()=>this.GetAllPostsForUserDependingOnFollowage())
            }else if (this.state.followed == true){
                axios.post('http://localhost:8004/unfollow',JSON.stringify({idfollower:userfollowing,iduser:data})).then(
                ()=>{
                    alert('You have unfollowed user with userid' + this.state.userid);
                    this.setState({followed:false});
                }).then(()=>this.GetAllPostsForUserDependingOnFollowage())
            }
        }).catch(()=>{alert('didnt retrieve user by username')});
    }

    render(){
        const data = this.state.posts;
        const user = this.state.user;
        const userid = this.state.userid;
        return(
            <>
            <div className="profile">
                    <div className="post__headerLeft">
                        <img src="" alt="" className="post__profilePic"/>slika
                        <div className="profile__header">
                            <h1 style={{marginLeft:"8px"}}>{user.name}</h1>
                            <button className="follow_but" onClick={this.FollowUser.bind(this)}>{this.state.followed ? "Unfollow" : "Follow"} </button>
                        </div>
                    </div>
                </div>
        {data.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private}/>
        </div>
        ))}
        </>
        )
    }
}
export default Profile