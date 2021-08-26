import Post from "../Post/Post";
import React from 'react';
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import axios from 'axios';

export class UserInteractedContent extends React.Component{

    constructor(props) {
        super(props);
        this.state = {
          likedposts: [],
          dislikedposts:[],
          userid:""
        };
      }

    async componentDidMount(){
        await this.GetUserIdByUsername()
        this.GetAllPosts();
    }

    async GetUserIdByUsername(){
        await axios.get('http://localhost:8081/getuseridandprivatebyusername/'+this.props.match.params.username).then((response)=>{
            const data = response.data;
            this.setState({userid:data.UserId});
        })
    }

    GetAllPosts(){
        axios.get('http://localhost:8082/getlikedbyuser/'+this.state.userid).then((response)=>{
            const data = response.data;
            this.setState({likedposts:data});
        })
        .catch(()=>{alert('didnt retrieve liked posts')});
        axios.get('http://localhost:8082/getdislikedbyuser/'+this.state.userid).then((response)=>{
            const data = response.data;
            this.setState({dislikedposts:data});
        })
        .catch(()=>{alert('didnt retrieve disliked posts')});
    }
    
    render(){
        const data = this.state.likedposts;
        const dataDisl = this.state.dislikedposts;
        return(
            <>
            <h1>Liked posts</h1>
        {data?.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private}  description={post.description} location={post.LocationID}/>
        </div>
        ))}
            <h1>Disliked posts</h1>
        {dataDisl?.map((post,i) => (
        <div className="feed" key={i}>
            <Post userid={post.userid} postid={post.ID} picpath={post.picpath} privatepost={post.private} description={post.description} location={post.LocationID}/>
        </div>
        ))}
        </>
        )
    }
}
export default UserInteractedContent