import React from 'react';
// import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
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
            <div className="profile">
                    <div className="post__headerLeft">
                        <img src="" alt="" className="post__profilePic"/>slika
                        <div className="profile__header">
                            <h1 style={{marginLeft:"8px"}}>{user.name}</h1>
                            <button className="follow_but">Follow</button>
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