import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import Post from './Post';

export class Profile extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          posts: [],
        };
      }

    componentDidMount(){
        this.GetAllPostsForUser();
    }

    GetAllPostsForUser(){
        axios.get('http://localhost:8005/allpostsbyuserid/'+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({posts:data});
            console.log(this.state.posts)
        })
        .catch(()=>{alert('didnt retrieve ')});
    }

    render(){
        const data = this.state.posts;
        return(
            <>
            <p>Ovde bi trebala komponenta koja je profil,slika,ime etc.</p>
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