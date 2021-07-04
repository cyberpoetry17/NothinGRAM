// import React, { useEffect } from 'react'
import AddImg from './components/AddImg'
import {BrowserRouter, Link, Route, Switch} from 'react-router-dom'
import Home from './components/Home';
import Post from './components/Post'
import Like from './components/Like';
import Test from "./components/Test"
import AddPost  from './components/AddPost';


function App() {
  // const [fileUrl, setFileUrl] = React.useState(null)
  // const [users, setUsers] = React.useState([])
  // const onFileChange = async (e) =>{
  //     const file = e.target.files[0]
  //     console.log("ispis",app.storage().ref())
  //     const storageRef = app.storage().ref()
  //     const fileRef = storageRef.child(file.name)
  //     await fileRef.put(file)
  //     setFileUrl(await fileRef.getDownloadURL())
  // }

  // const onSubmit = (e) => {
  //   e.preventDefault()
  //   //const username = e.target.value;
  //   var username = document.getElementById("name").value;
  //   if(!username){
  //     return
  //   }
  //   db.collection("users").doc("UUvNYlsLmAaIuOKr10XH").set({
  //     name: username,
  //     avatar: fileUrl
  //   })
  // }
  // useEffect(() => {
  //   const fetchUsers = async () =>{
  //     const usersCollection = await db.collection('users').get()
  //     setUsers(usersCollection.docs.map(doc => {
  //       return doc.data()
  //     }))
  //   }
  //   fetchUsers();
  // }, [])
  return (
    <>
      <h1>This is out of router (static)</h1>
     

      <BrowserRouter>
        {/* this is menu bar  */}
        <div className="topnav">
          <Link to= "/" >HOME</Link><br/>
          <Link to="/pic">Add picture</Link><br />
          <Link to="/posts">Post feed</Link><br />
          <Link to="/addPost">Add post</Link><br />
        </div>
        <Switch>
          <Route path="/dislike">
            <Like/>
          </Route>
          <Route path="/like">
            <Like/>
          </Route>
          <Route path="/pic">
            <AddImg/>
          </Route>
          <Route path="/posts">
            <Post/>
          </Route>
          <Route path="/addPost">
            <AddPost/>
          </Route>
          <Route path = "/">
            <Home/>
          </Route>
        </Switch>
      </BrowserRouter>
      
    </>
  );
}

export default App;
