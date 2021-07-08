import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import Post from './Post';

export class Profile extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          posts: [],
          user:[]
        };
      }
      
    componentDidMount(){
        this.GetAllPostsForUser();
        this.GetUserByUserId();
    }

    GetUserByUserId(){
        axios.get('http://localhost:8004/getuserbyusername/'+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({user:data});
        })
        .catch(()=>{alert('didnt retrieve user')});
    }

    GetAllPostsForUser(){
        axios.get('http://localhost:8005/allpostsbyuserid/'+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({posts:data});
        })
        .catch(()=>{alert('didnt retrieve ')});
    }

    render(){
        const data = this.state.posts;
        const user = this.state.user;
        return(
            <>
            <img/>Bice slika<h1>{user.name}</h1>
        {data.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath}/>
        </div>
        ))}
        </>
        )
    }
}
export default Profile