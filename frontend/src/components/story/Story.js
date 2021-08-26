import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';
import {Link} from 'react-router-dom'
import "../../styles/story.css";
import BootstrapSwitchButton from 'bootstrap-switch-button-react'
import {serviceConfig} from '../../applicationSettings.js'


export default function Story({postId,type,IdStory,UserId,size,ForCloseF,ShowOnStoryHighlights}) {
    //retriew media and render it
    //add userID to story
    const [media, setMedia] = useState(null);
    const [isLoaded,setIsLoaded] = useState(false);
    const [firstTime, setFirstTime] = useState(true)
    const [UserName,setUserName] = useState(null)

    useEffect(()=>{
        onLoad();
        if(size==1 && document.getElementById(IdStory+"-div") != null){
            document.getElementById(IdStory+"-div").style.borderColor = "white";
        }
        
    },[])
    
    const onLoad = () => {
        axios({
            method : 'get',
            url :`${serviceConfig.postURL}/GetMediaForStory?StoryId=`+IdStory,
        }).then(res =>{
            console.log(res.data," media for story")
            setMedia(res.data)
        });
        console.log(UserId , "user id")
        axios({
            method : 'get',
            url :`${serviceConfig.userURL}/username/`+UserId,
        }).then(res =>{
            setUserName(res.data.substring(1,(res.data.length)-2));
        });

    }
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
    const closeChanged = (e)=>{
        if(e == true){
            axios({
                method : 'post',
                url :`${serviceConfig.postURL}/AddToStoryHighlights/`+IdStory
            });
        }else{
            axios({
                method : 'post',
                url :`${serviceConfig.postURL}/RemoveFromStoryHighlights/`+IdStory
            });
        }
        
    }
    const renderImg = () => {
        if(ForCloseF == true){
            return <div id={IdStory+"-div"} className="greenBorder"><img id={IdStory} onClick={click} className="d-block" width="150" height="100" src={media.Link} alt="my pic"/></div>
        }
        if(size==1){
            return <div id={IdStory+"-div"} > <img id={IdStory} onClick={click} className="d-block w-100 h-100" src={media.Link} alt="my pic"/></div>
        }
        if(size== 2){
            return <div id={IdStory+"-div"} >
                    <img id={IdStory} onClick={click} className="d-block w-100" height="200"  src={media.Link} alt="my pic"/>
                    <div  className="outerDiv"> 
                    <BootstrapSwitchButton 
                        id={IdStory}
                        width={260}
                        checked={ShowOnStoryHighlights}
                        onlabel='show on story highlights'
                        onChange={closeChanged}
                        offlabel='dont show on story highlights'/>
                    </div>
                    </div>
        }
        return <div id={IdStory+"-div"} className="asdfsdfasasdas"><img id={IdStory} onClick={click} className="d-block" width="150" height="100" src={media.Link} alt="my pic"/></div>
    }
    return (
        <div>
            <Link to={"/profile/"+UserName}>{UserName}</Link>
            {isLoaded === true? 
                renderImg()
            :
            <p>no pic </p>
            }
        </div>
    )
}
