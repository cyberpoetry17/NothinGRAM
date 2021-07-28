import React from 'react'
import './App.css'
import Login from './components/Login'
import AddImg from './components/AddImg'
import Home from './components/Home';
// import UserHomepage from './components/UserHomepage';
// import Post from './components/Post'
import Like from './components/Like';
import PostFeed from './components/PostFeed';
//import Test from "./components/Test"
import AddPost  from './components/AddPost';
import Dislike from './components/Dislike';
import RegisterUser from './components/Register'
import Update from './components/Update';
import Profile from './components/Profile';
import FollowerFeed from './components/FollowerFeed'
import Verification from './components/Verification'
// import { version } from 'react-dom';
// import jwt_decode from 'jwt-decode';
import ProfileRequests from './components/ProfileRequests'
import {BrowserRouter, Route, Switch} from 'react-router-dom'
import 'bootstrap/dist/css/bootstrap.min.css'
import {Nav} from 'react-bootstrap';
import { version } from 'react-dom';
import jwt_decode from 'jwt-decode';
import UserInteractedContent from './components/UserInteractedContent';
import AddStory from './components/AddStory';
import AddCloseFollower from './components/AddCloseFollower';
import StoryForUser from './components/StoryForUser';


export default function App() {

  return (

    <>

      <BrowserRouter>
          <Nav className="navbar" activeKey="/" >
              {window.localStorage.getItem('token') ? 
              <Nav.Item>
              <Nav.Link href={"/profile/"+jwt_decode(localStorage.getItem('token')).Username}>My profile</Nav.Link>
              </Nav.Item>
              :
              <Nav.Item>
              <Nav.Link href="/">HOME</Nav.Link>
              </Nav.Item>
              }
              {window.localStorage.getItem('token') ? 
              <>
              <Nav.Item>
              <Nav.Link href="/userfeed">MY FEED</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/posts">POSTS FEED</Nav.Link>
              </Nav.Item>
              </>
              :
              <Nav.Item>
                <Nav.Link href="/posts">POSTS FEED</Nav.Link>
              </Nav.Item>
              }
              {window.localStorage.getItem('token') ? 
              <Nav.Item>
              <Nav.Link href="/" onClick={()=>{window.localStorage.removeItem('token');this.props.history.push('/');}}>SIGN OUT</Nav.Link>
              </Nav.Item>
              :
              <>
              {/* <Nav.Item>
                <Nav.Link href="/update">EDIT USER</Nav.Link>
              </Nav.Item> */}
              <Nav.Item>

              <Nav.Link href="/login">SIGN IN</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/register">SIGN UP!</Nav.Link>
              </Nav.Item>
              </>
              }
          </Nav>

        <Switch >
          <Route path="/userinteracted/:username" component={UserInteractedContent}/>
          <Route path="/requests/:username" component={ProfileRequests}/>
          <Route path="/userfeed" component={FollowerFeed}/>
          <Route path="/closefollowerr" component={AddCloseFollower}/>
          <Route path="/verification/" component={Verification}/>
          <Route path="/profile/:username" component={Profile}/>
          <Route className="main" path="/posts">
            <PostFeed/>
          </Route>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/update">
            <Update />
          </Route>
          <Route path="/dislike">
            <Dislike/>
          </Route>
          <Route path="/like">
            <Like/>
          </Route>
          {/* <Route path="/pic">
            <AddImg/>
          </Route> */}
          <Route  path="/addPost">
            <AddPost/>
          </Route>
          <Route  path="/addStory">
            <AddStory/>
          </Route>
          <Route  path="/stories">
            <StoryForUser/>
          </Route>

          <Route path="/register">
            <RegisterUser />
          </Route> 
          <Route path = "/">
            <Home/>
          </Route>

        </Switch>

      </BrowserRouter>
    </>


  );
}




