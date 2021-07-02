import React from 'react'
import './App.css';
//import AddImg from './components/AddImg'




// function App() {
//   // const [fileUrl, setFileUrl] = React.useState(null)
//   // const [users, setUsers] = React.useState([])
//   // const onFileChange = async (e) =>{
//   //     const file = e.target.files[0]
//   //     console.log("ispis",app.storage().ref())
//   //     const storageRef = app.storage().ref()
//   //     const fileRef = storageRef.child(file.name)
//   //     await fileRef.put(file)
//   //     setFileUrl(await fileRef.getDownloadURL())
//   // }

//   // const onSubmit = (e) => {
//   //   e.preventDefault()
//   //   //const username = e.target.value;
//   //   var username = document.getElementById("name").value;
//   //   if(!username){
//   //     return
//   //   }
//   //   db.collection("users").doc("UUvNYlsLmAaIuOKr10XH").set({
//   //     name: username,
//   //     avatar: fileUrl
//   //   })
//   // }
//   // useEffect(() => {
//   //   const fetchUsers = async () =>{
//   //     const usersCollection = await db.collection('users').get()
//   //     setUsers(usersCollection.docs.map(doc => {
//   //       return doc.data()
//   //     }))
//   //   }
//   //   fetchUsers();
//   // }, [])
//   return (
//     <>
//        <Login/>
//     </>
//   );
// }

//import React from "react";
import Login from './components/Login'
import AddImg from './components/AddImg'
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
              <Link to="/image">addImage</Link>
            </li>
            <li>
              <Link to="/login">Login</Link>
            </li>
          </ul>
        </nav>

        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/image">
            <AddImg/>
          </Route>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/">
           
          </Route>
        </Switch>
      </div>
    </Router>
  );
}




