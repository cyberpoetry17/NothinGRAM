import React from 'react'
import './App.css'
import Login from './components/user/Login';
import AddImg from './components/AddImg'
import Home from './components/Home';
// import UserHomepage from './components/UserHomepage';
// import Post from './components/Post'
import Like from './components/Post/Like';
import PostFeed from './components/feed/PostFeed';
//import Test from "./components/Test"
import AddPost from './components/Post/AddPost';
import Dislike from './components/Post/Dislike';
import RegisterUser from './components/user/Register'
import Update from './components/user/Update';
import Profile from './components/user/Profile';
import FollowerFeed from './components/feed/FollowerFeed'
import Verification from './components/user/Verification'
// import { version } from 'react-dom';
// import jwt_decode from 'jwt-decode';
import ProfileRequests from './components/user/ProfileRequests'
import {BrowserRouter, Route, Switch} from 'react-router-dom'
import 'bootstrap/dist/css/bootstrap.min.css'
import {Nav} from 'react-bootstrap';
import { version } from 'react-dom';
import jwt_decode from 'jwt-decode';
import UserInteractedContent from './components/user/UserInteractedContent';
import AddStory from './components/story/AddStory';
import AddCloseFollower from './components/user/AddCloseFollower';
import StoryForUser from './components/story/StoryForUser';
import SearchBar from './components/SearchBar';
import LocationSearch from './components/LocationSearch';
import TagSearch from './components/TagSearch';

export default function App() {

  return (

    <>

      <BrowserRouter>
          <Nav className="navbar" activeKey="/" >
              {window.localStorage.getItem('token') ? 
              <>
              <Nav.Item>
                <Nav.Link href={"/profile/"+jwt_decode(localStorage.getItem('token')).Username}>My profile</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/userfeed">MY FEED</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/posts">POSTS FEED</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <SearchBar/>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/" onClick={()=>{window.localStorage.removeItem('token');this.props.history.push('/');}}>SIGN OUT</Nav.Link>
              </Nav.Item>
              </>
              :
              <>
              <Nav.Item>
                <Nav.Link href="/">HOME</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="/posts">POSTS FEED</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <SearchBar/>
              </Nav.Item>
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
          <Route path="/search/locations/:location" render={(props) => (<LocationSearch location={props.match.params.location}/>)}/>
          <Route path="/search/tags/:tag" render={(props) => (<TagSearch location={props.match.params.tag}/>)}/>
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




