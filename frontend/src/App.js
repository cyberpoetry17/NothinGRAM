import React from 'react'
import './App.css'
import Login from './components/Login'
import AddImg from './components/AddImg'
import Home from './components/Home';
import RegisterUser from './components/Register'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";



export default function App() {
  return (
    <Router>
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
            
          </ul>
        </nav>

        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/pic">
            <AddImg/>
          </Route>
          <Route path="/login">
            <Login />
          </Route>
           <Route path="/register">
            <RegisterUser />
          </Route> 
          <Route path="/"> 
          <Home></Home></Route>
        </Switch>
      </div>
    </Router>

  );
}




