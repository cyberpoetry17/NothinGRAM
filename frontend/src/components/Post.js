import React from 'react';
import { app } from './base';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';
import { Like } from './Like';
import Dislike from './Dislike';

export class Post extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
          posts: [],
        };
      }

    
    componentDidMount(){
        this.GetAllPosts();
    }
    
    // GetLikesForPost(postid){                                                         //metoda je ukleta,kad se pozove izvrsi se milion puta iz nekog razloga
    //     axios.get('http://localhost:8005/alllikesforpost/'+postid).then((response)=>{
    //         const data = response.data;
    //         this.setState({likes:data});
    //         console.log(this.state.likes)
    //     })
    //     .catch(()=>{alert('didnt retrieve likes')});
    // }

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
        return(
            <BrowserRouter>
            <>
            <ul>
                {data.map((post,i) => (
                <li key={i}>
                    <h3>Post id:{post.ID}</h3>,<h3>User id:{post.userid}</h3>,<Route path='/like' component={Like}/><Link to='/like'>Likes</Link>,<Route path='/dislike' component={Dislike}/><Link to='/dislike'>Dislikes</Link>
                </li>
                ))}
            </ul>
            </>
            </BrowserRouter>
            );
    }

}
export default Post