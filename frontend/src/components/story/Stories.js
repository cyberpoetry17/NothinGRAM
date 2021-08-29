import React from 'react'
import {useState,useEffect} from 'react'
import axios from 'axios';
import Story from './Story';
import "../../styles/story.css";
import StoryGroup from './StoryGroup';
import jwt_decode from 'jwt-decode';

export default function Stories() {
    const [stories, setStories] = useState(null);
    const [closeStories, setCloseStories] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const [loaded2, setLoaded2] = useState(false);
    const [storiesMap, setStoriesMap] = useState({});
    const [firstTime, setFirstTime]= useState(true);
    useEffect(()=>{
        // loadStories();
        loadActiveStoriesFromFallowedUsers()
        loadCloseFriendsStory();
    },[])

    useEffect(()=>{
        console.log("stories:",stories);
    },[stories])

    const loadStories = () => {
        axios({
            method : 'get',
            url :'http://localhost:8080/getAllStories',
        }).then(res =>{
            console.log(res.data,"res.data");
            setStories(res.data);
        });
    }
    const loadCloseFriendsStory = () => {
        axios.get('http://localhost:8081/getAllCloseFollowersForUser/'+jwt_decode(localStorage.getItem('token')).UserID).then((response)=>{
            
            response.data?.map((follow) =>(
                axios.get('http://localhost:8080/GetCloseFrinedStoriesForUser/'+follow).then((responsenew)=>{
                const data = responsenew.data;
                if(data != null){
                    setCloseStories(data);
                }
            })
            .catch(()=>{alert('didnt retrieve ')})
            ))
        }).catch(()=>{alert('You have not followed any other users.')})
    }
    
    const loadActiveStoriesFromFallowedUsers=()=>{
        axios.get('http://localhost:8081/getallfollowedforloggeduser/'+jwt_decode(localStorage.getItem('token')).UserID).then((response)=>{
            
            response.data?.map((follow) =>(     
                axios.get('http://localhost:8080/GetActiveStoriesByUserId/'+follow).then((responsenew)=>{
                const data = responsenew.data;
                console.log(data)
                if(data != null){
                    if(stories ==null){
                        setStories(data)
                    }else{
                        var list = stories;
                        list.push(data);
                        setStories(list);
                    }
                }
            })
            .catch(()=>{alert('didnt retrieve ')})
            ))
        }).catch(()=>{alert('You have not followed any other users.')})
    }

    useEffect(()=>{
        if(firstTime){
            setFirstTime(false)
            return;
        }
        setLoaded(true);
    },[storiesMap])

    useEffect(async()=>{
        var map = storiesMap;
        if(stories != null){
            map=stories.reduce((acc, curr) => {
                const isArray = acc[curr.UserId];
                if (isArray) acc[curr.UserId].push(curr);
                else acc[curr.UserId] = [curr];
                return acc;
            }, map)
            console.log("mapa:",map)
            await setStoriesMap(map);
            await setLoaded(true);
        }
    },[stories])

    useEffect(async()=>{
        var map = storiesMap;
        if(closeStories != null){
            map=closeStories.reduce((acc, curr) => {
                const isArray = acc[curr.UserId+"close"];
                if (isArray) acc[curr.UserId+"close"].push(curr);
                else acc[curr.UserId+"close"] = [curr];
                return acc;
            },map)
            // console.log("mapa:",map)
            // var obj = Object.assign({},map,...storiesMap);
            // console.log(obj);   
            await setStoriesMap(map);
            await setLoaded2(true);
        }
    },[closeStories])

    const renderStoryGrop = (key) => {
        if(key.endsWith("close")){
            return <div className="box"> <StoryGroup storyList={storiesMap[key]} ForCloseFriends={true} /></div> 
        }
        return <div className="box"> <StoryGroup storyList={storiesMap[key]} ForCloseFriends={false}/></div>    
    }
    return (
        <div>
            <div className="containerStory">
                {loaded || loaded2?
                 Object?.keys(storiesMap).map(function (key) {
                // console.log('key: ', key);  // Returns key: 1 and key: 2
                return renderStoryGrop(key)
                    // <div className="box">    
                    //     <StoryGroup storyList={storiesMap[key]}/>
                    // </div>
                    
                }, this)
                :
                <p>loading..</p>
                }
            </div>
        </div>
    )
}
