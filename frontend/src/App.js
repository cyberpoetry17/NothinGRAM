import React from 'react'
import './App.css'
import Login from './components/Login'
import AddImg from './components/AddImg'
import Home from './components/Home';
import Post from './components/Post'
import Like from './components/Like';
import PostFeed from './components/PostFeed';
import Test from "./components/Test"
import AddPost  from './components/AddPost';
import Dislike from './components/Dislike';
import RegisterUser from './components/Register'
import Profile from './components/Profile'
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import 'bootstrap/dist/css/bootstrap.min.css'
import {Nav} from 'react-bootstrap';


export default function App() {
  return (

    <>
      <BrowserRouter>
          <Nav activeKey="/" >
            <Nav.Item>
              <Nav.Link href="/">HOME</Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link href="/pic">PICTURE</Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link href="/login">SIGN IN</Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link href="/register">SIGN UP!</Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link href="/posts">POSTS FEED</Nav.Link>
            </Nav.Item>
            <Nav.Item>
              <Nav.Link href="/addPost">ADD POST</Nav.Link>
            </Nav.Item>
          </Nav>
          {/* <div>
            <nav>
              <ul>
                <li>
                  <Link to="/">Home</Link>
                </li>
                <li>
                  <Link to="/pic">Picture</Link>
                </li>
                <li>
                  <Link to="/login">Sign in</Link>
                </li>
                <li>
                  <Link to="/register">Sign up!</Link>
                </li>
                <li>
                  <Link to="/posts">Post feed</Link>
                </li>
                <li>
                <Link to="/addPost">Add post</Link>
                </li>
              </ul>
            </nav>
          </div> */}

        <Switch>
          <Route path="/profile/:username" component={Profile}/>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/dislike">
            <Dislike/>
          </Route>
          <Route path="/like">
            <Like/>
          </Route>
          <Route path="/pic">
            <AddImg/>
          </Route>
          <Route path="/posts">
            <PostFeed/>
          </Route>
          <Route path="/addPost">
            <AddPost/>
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




