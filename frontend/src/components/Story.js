import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';
import {Link} from 'react-router-dom'
import "../styles/story.css";

export default function Story({postId,type,IdStory,UserId,size}) {
    //retriew media and render it
    //add userID to story
    const [media, setMedia] = useState(null);
    const [isLoaded,setIsLoaded] = useState(false);
    const [firstTime, setFirstTime] = useState(true)
    const [UserName,setUserName] = useState(null)

    useEffect(()=>{
        axios({
            method : 'get',
            url :'http://localhost:8005/GetMediaForStory?StoryId='+IdStory,
        }).then(res =>{
            console.log(res.data," media for story")
            setMedia(res.data)
        });
        console.log(UserId , "user id")
        axios({
            method : 'get',
            url :'http://localhost:8005/getusernamebyid/'+UserId,
        }).then(res =>{
            setUserName(res.data.substring(1,(res.data.length)-2));
        });
        if(size==1){
            document.getElementById(IdStory+"-div").style.borderColor = "white";
        }

        //document.getElementById("pic").style.borderRadius = "25%";
    },[])

    useEffect(()=>{
        if(firstTime){
            setFirstTime(false)
            return;
        }
        console.log(media);
        setIsLoaded(true);
    },[media])

    useEffect(()=>{
        var el = document.getElementById(IdStory);
        if(el!=null){
            el.style.borderRadius = "25%";
        }

        console.log(isLoaded);
    },[isLoaded])

    const click = (event) => {
        console.log("ovo cudo radii");
    };
    const renderImg = () => {
        if(size==1){
            return <div id={IdStory+"-div"} > <img id={IdStory} onClick={click} className="d-block w-100 h-100" src={media.Link} alt="my pic"/></div>
        }
        if(size== 2){
            return <div id={IdStory+"-div"} ><img id={IdStory} onClick={click} className="d-block w-100" height="200"  src={media.Link} alt="my pic"/></div>
        }
        return <div id={IdStory+"-div"} className="asdfsdfasasdas"><img id={IdStory} onClick={click} className="d-block" width="150" height="100" src={media.Link} alt="my pic"/></div>
    }
    return (
        <div>
            <Link to={"/profile/"+UserName}>{UserName}</Link>
            {/* <div id={IdStory+"-div"} className="asdfsdfasasdas"> */}
            {isLoaded === true? 
            [
                // (size == 1?
                //     <img onClick={click} className="d-block w-100 h-100" width="150" height="100" src={media.Link} alt="my pic"/>:
                //     <img onClick={click} className="d-block" width="150" height="100" src={media.Link} alt="my pic"/>
                // )
                renderImg()
            ]
            :
            <p>no pic </p>
            }
            {/* </div> */}
        </div>
    )
}
