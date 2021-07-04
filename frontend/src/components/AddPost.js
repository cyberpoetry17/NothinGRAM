import axios from 'axios';
import React,{useState,useEffect} from 'react'
import { app } from './base';

function AddPost() {
    const [post, setPost] = useState({description:'', picpath:'', private:true, UserID:"0f608e5a-1e79-4dd3-ba7e-fe99c81e6fe2"})
    const ispisi = ()=>{
        console.log(JSON.stringify(post))
    }

    const setPrivate= ()=>{
        setPost({...post,private:true})
    }

    const setPublic= ()=>{
        setPost({...post,private:false})
    }

    const fileChange = async (e)=>{
        const file = e.target.files[0]
        //console.log("ispis",app.storage().ref())
        const storageRef = app.storage().ref()
        const fileRef = storageRef.child(file.name)
        await fileRef.put(file)
        var a = await fileRef.getDownloadURL()
        console.log(a)
        setPost({...post,picpath:a})
    }

    const add = ()=>{
        var a = JSON.stringify(post)
        axios({
            method : 'post',
            url :'http://localhost:8005/createpost',
            data:a,
        });
        
    }

    return (
        <div>
            <input 
            type= 'text' 
            value= {post.description}
            onChange={e => setPost({...post, description:e.target.value})}
            />
            <div>
                <label><input type="radio" name="private"  value='private' onChange={setPrivate}/> Private</label>
                <label><input type="radio" name="private"  value='not private' onChange={setPublic}/>Not private</label>
            </div>
            <input type="file" onChange={fileChange} /><br/>
            <img width="100" height="100" src={post.picpath} alt="my pic"/><br/>
            <button onClick={ispisi}>JSON</button>
            <button onClick={add}>ADD POST</button>
        </div>
    )
}

export default AddPost
