import Post from "./Post";
import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';

export class PostFeed extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          posts: [],
        };
      }

    componentDidMount(){
        this.GetAllPosts();
    }

    GetAllPosts(){
        axios.get('http://localhost:8005/').then((response)=>{
            const data = response.data;
            this.setState({posts:data});
            console.log(this.state.posts)
        })
        .catch(()=>{alert('didnt retrieve ')});
    }
    
    render(){
        const data = this.state.posts;
        const like = this.state.likes;
        console.log(data)
        return(
            <>
        {data.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID}/>
        </div>
        ))}
        </>
        )
    }
}
export default PostFeed