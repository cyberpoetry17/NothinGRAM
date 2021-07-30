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
        axios.get('http://localhost:8004/getpublicuserids').then((response)=>{
            response.data?.map((userid) =>(
                axios.get('http://localhost:8005/allpostsbyuserid/'+userid).then((responsenew)=>{
                const data = responsenew.data;
                if(data != null)
                    this.setState({posts:this.state.posts.concat(data)});
                })
                .catch(()=>{alert('didnt retrieve ')})
            ))
        }).catch(()=>{alert('You have not followed any other users.')})
    }
    
    render(){
        const data = this.state.posts;
        return(
            <>
                {data.map((post,i) => (
                <div className="feed" key={i}>
                    <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private} description={post.description} location = {post.LocationID} />
                </div>
                ))}
            </>
        )
    }
}
export default PostFeed