import React, { Component } from 'react'
import { app } from './base';

// const db = app.firestore()
export class AddImg extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
             fileUrl: "asd"
        }
    }

    onFileChange = async (e) =>{
        e.preventDefault()
        const file = e.target.files[0]
        //console.log("ispis",app.storage().ref())
        const storageRef = app.storage().ref()
        const fileRef = storageRef.child(file.name)
        await fileRef.put(file)
        var a = await fileRef.getDownloadURL()
        console.log(a)
        this.setState(()=>({
             fileUrl: a
        }))
    }

    onSubmit = (e) => {
        e.preventDefault()
        console.log(this.state.fileUrl)
        //const username = e.target.value;
        // var username = document.getElementById("name").value;
        // if(!username){
        // return
        // }
        // db.collection("users").doc("UUvNYlsLmAaIuOKr10XH").set({
        // name: username,
        // avatar: fileUrl
        // })
    }
    render() {
        return (
            <div>
                <form onSubmit={this.onSubmit.bind(this)}>
                <input type="file" onChange={this.onFileChange.bind(this)} />
                <button>submit</button>
                </form>
                <ul>
                {/* {users.map(user => {
                    return (<li><img width="100" height="100" src= {user.avatar} alt={user.name}/>
                        <p>{user.name}</p></li>
                    )
                })} */}
                <li><img width="100" height="100" src={this.state.fileUrl} alt="my pic"/></li>
                </ul>
            </div>
        )
    }
}

export default AddImg
