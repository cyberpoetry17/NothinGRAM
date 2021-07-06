import React from 'react'
import './App.css'
import Login from './components/Login'
import AddImg from './components/AddImg'
import Home from './components/Home';
import UserHomepage from './components/UserHomepage';
import Post from './components/Post'
import Like from './components/Like';
import PostFeed from './components/PostFeed';
import Test from "./components/Test"
import AddPost  from './components/AddPost';
import Dislike from './components/Dislike';
import RegisterUser from './components/Register'
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'

export default function App() {
  return (

    <>
      <BrowserRouter>
          <div>
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
          </div>

        <Switch>
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
          <Route path="/userhomepage" component={UserHomepage} />
          <Route path = "/">
            <Home/>
          </Route>
        </Switch>
      </BrowserRouter>

    </>


  );
}




