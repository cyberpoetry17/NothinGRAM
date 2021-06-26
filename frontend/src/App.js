// import React, { useEffect } from 'react'
import AddImg from './components/AddImg'



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
      <AddImg/>
    </>
  );
}

export default App;
