import logo from './logo.svg';
import './App.css';
import Greet from './components/Greet'
import Welcome from './components/Welcome'
import Hello from './components/Hello'
import React from 'react'
import Message from './components/Message'
import Counter from './components/Counter'
import FunctionClick from './components/FunctionClick'
import ClassClick from './components/ClassClick'
import EventBind from './components/EventBind'

function App() {
  return (
    <div className="App">
      {/* <Greet name='Mark' heroName='Batman'>
        <p>Children props</p>
      </Greet>
      <Greet name='Sara' heroName='Supermen'>
        <button>Action</button>
      </Greet>
      <Greet name='Lazar' heroName='Wonder Women'/>
      <Welcome name='Component name' heroName='Component heroName'/>
      <Hello/>
      <Message/>
      <Counter/>
      <FunctionClick/>
      <ClassClick/> */}
        {/*<EventBind/>*/}
      <Counter/>
    </div>
  );
}

export default App;
