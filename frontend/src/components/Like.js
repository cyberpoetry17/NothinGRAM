import React from 'react';
import { app } from './base';
import {Link, Route} from 'react-router-dom';
import axios from 'axios';

export class Like extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
          likes:0
        };
      }

    
    componentDidMount(){
        //this.GetAllPosts();
    }
    
    // GetLikesForPost(postid){                                                         //metoda se ponasa cudno,kad se pozove izvrsi se milion puta iz nekog razloga
    //     axios.get('http://localhost:8005/alllikesforpost/'+postid).then((response)=>{
    //         const data = response.data;
    //         this.setState({likes:data});
    //         console.log(this.state.likes)
    //     })
    //     .catch(()=>{alert('didnt retrieve likes')});
    // }

    // GetAllPosts(){
    //     axios.get('http://localhost:8005/').then((response)=>{
    //         const data = response.data;
    //         this.setState({posts:data});
    //         console.log(this.state.posts)
    //     })
    //     .catch(()=>{alert('didnt retrieve ')});
    // }

    render(){
        const like = this.state.likes;
        return(
            <p>broj lajkova</p>
        )
    }

}
export default Like