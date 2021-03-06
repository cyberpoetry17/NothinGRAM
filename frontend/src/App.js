import React from 'react'
import './App.css'
import Login from './components/user/Login';
import AddImg from './components/AddImg'
import Home from './components/Home';
import PostFeed from './components/feed/PostFeed';
import AddPost from './components/Post/AddPost';
import RegisterUser from './components/user/Register'
import Update from './components/user/Update';
import Profile from './components/user/Profile';
import FollowerFeed from './components/feed/FollowerFeed'
import Verification from './components/user/Verification'
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
import ReportedFeed from './components/feed/ReportedFeed';
import AddAgent from './components/AddAgent';
import AgentConfirmationPage from './components/user/AgentConfirmationPage';
import VerificationPage from './components/VerificationPage'


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
              {jwt_decode(localStorage.getItem('token')).Role === 1 ?
              <>
              <Nav.Item>
                <Nav.Link href="/reportfeed">Reported Content</Nav.Link>
              </Nav.Item>
              </> : null
              }
              <Nav.Item>
                <Nav.Link href="/posts" onClick={()=>{window.localStorage.removeItem('token')}}>SIGN OUT</Nav.Link>
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
              <Nav.Item>
                <Nav.Link href="/addAgent">ADD AGENT</Nav.Link>
              </Nav.Item>
              </>
              }
          </Nav>

        <Switch >
          <Route path="/agentrequests" component={AgentConfirmationPage}/>
          <Route path="/addAgent" component={AddAgent}/>
          <Route path="/reportfeed" component={ReportedFeed}/>
          <Route path="/userinteracted/:username" component={UserInteractedContent}/>
          <Route path="/requests/:username" component={ProfileRequests}/>
          <Route path="/userfeed" component={FollowerFeed}/>
          <Route path="/closefollowerr" component={AddCloseFollower}/>
          <Route path="/verification" component={Verification}/>
          <Route path="/verificationrequests" component={VerificationPage}/>
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




