import { app } from '../base';
import {BrowserRouter, Link, Redirect, Route, Switch, useHistory} from 'react-router-dom'
import {Carousel} from 'react-bootstrap';
import axios from 'axios';
import React,{useState,useEffect} from 'react'
import jwt_decode from 'jwt-decode';
import "../../styles/post-style.css";
import Comment from '../Post/Comment';
import CommentInput from '../Post/CommentInput';
import { Ellipsis } from 'react-bootstrap/esm/PageItem';

export default function Post({userid,postid,picpath,privatepost,tokenInfo,description,location}){
    
    var [username,setUsername] = React.useState();
    var [likes,setLikes] = React.useState(0);
    var [dislikes,setDislikes] = React.useState(0);
    var [comments,setComments] = React.useState();
    var [locationdesc,setLocationdesc] = React.useState();
    var [tags,setTags] = React.useState();
    const [media, setMedia] = useState([])
    const [loaded, setLoaded]= useState(false)
    const [firstTime, setFirstTime] = useState(true)
    let history = useHistory();
    var tokenInfo;

    useEffect(()=>{
        GetLikesForPost()
        GetDislikesForPost()
        GetCommentsForPost()
        GetMediaForPost()
        GetLocationForPostByLocationId()
        GetTagsForPostByLocationId()
    },[])

    useEffect(()=>{
        if(firstTime){
            setFirstTime(false)
            return
        }
        setLoaded(true)
        console.log(description)
        console.log("media",media)
    },[media])

    React.useEffect(()=>GetUsernameByUserId(),[]);

    const GetMediaForPost = ()=>{
        var Url ='http://localhost:8080/api/post/GetMediaForPost?PostId='+postid
        axios({
            method : 'get',
            url :Url,
            data:JSON.stringify(postid)
        }).then(res=>{
            setMedia({...media, media : res.data})
        });
    }

    const GetLocationForPostByLocationId = () =>{
        axios.get('http://localhost:8080/api/post/locationforpost/'+location).then((response) =>{
            if(response.data.country != "dumb")
                setLocationdesc("@"+response.data.city+","+response.data.country);
        });
    }

    const GetTagsForPostByLocationId = () =>{
        axios.get('http://localhost:8080/api/post/tagsforpost/'+postid).then((response) =>{
            if(response.data != null)
                setTags(response.data.map((tag)=>("#"+tag)));
        });
    }

    const GetUsernameByUserId = () =>{
        return axios.get('http://localhost:8080/api/user/username/'+userid).then((response) =>{
            setUsername(response.data.substring(0,(response.data.length)));
        });
    }

    const GetCommentsForPost = () =>{
        axios.get('http://localhost:8080/api/post/getcommentsforpost/'+postid).then((response)=>{
            setComments(response.data)
        });
    }
    
    const GetLikesForPost = () =>{                                                         
        return axios.get('http://localhost:8080/api/post/alllikesforpost/'+postid).then((response)=>{
            setLikes(response.data)
        });
    }

    const GetDislikesForPost = () =>{                                                        
        return axios.get('http://localhost:8080/api/post/alldislikesforpost/'+postid).then((response)=>{
            setDislikes(response.data)
        });
    }

    const CheckIfUserLikedPost = () =>{
        axios({method:'post',url:'http://localhost:8080/api/post/checkiflikedbyuser',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then((response)=>{
            if(response.data == false){             //moralo je ovde umesto u likethispost()
                axios({method:'post',url:'http://localhost:8080/api/post/createlike',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then(()=>GetLikesForPost());
            }else if(response.data == true){
                alert("You already liked this post");
            }
        });
    
    }

    const CheckIfUserDislikedPost = () => {
        axios({method:'post',url:'http://localhost:8080/api/post/checkifdislikedbyuser',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then((response)=>{
            if(response.data == false){
                axios({method:'post',url:'http://localhost:8080/api/post/createdislike',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then(()=>GetDislikesForPost());
            }else if (response.data == true){
                alert("You already disliked this post");
            }
        });
    }

    const DoINeedToRemoveDislike = () => {
        axios({method:'post',url:'http://localhost:8080/api/post/checkifdislikedbyuser',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then((response)=>{
            if(response.data == true){
                axios({method:'post',url:'http://localhost:8080/api/post/deletedislike',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then(()=>GetDislikesForPost());
            }
        });
    }

    const DoINeedToRemoveLike = () => {
        axios({method:'post',url:'http://localhost:8080/api/post/checkiflikedbyuser',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then((response)=>{
            if(response.data == true){
                axios({method:'post',url:'http://localhost:8080/api/post/deletelike',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then(()=>GetLikesForPost());
            }
        });
    }

    const LikeThisPost = () =>{
        try{
            tokenInfo = jwt_decode(localStorage.getItem('token'));
        }catch(e){
            tokenInfo = "";
            console.error(e);
        }
        if(tokenInfo!=""){
            DoINeedToRemoveDislike();
            CheckIfUserLikedPost();
        }else{
            alert("You are not logged in.You will be redirected to login.");
            history.push('/login');
        }
    }

    const DislikeThisPost = () =>{
        try{
            tokenInfo = jwt_decode(localStorage.getItem('token'));
        }catch(e){
            tokenInfo = "";
            console.error(e);
        }
        if(tokenInfo!=""){
            DoINeedToRemoveLike();
            CheckIfUserDislikedPost();
        }else{
            alert("You are not logged in.You will be redirected to login.");
            history.push('/login');
        }
    }

    const CheckIfUserReportedPost = () =>{
        axios({method:'post',url:'http://localhost:8080/api/post/checkifreportedbyuser',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then((response)=>{
            if(response.data == false){
                axios({method:'post',url:'http://localhost:8080/api/post/reportpost',headers:{},data:JSON.stringify({userid:tokenInfo.UserID,postid})}).then(()=>{alert("Post has been reported.")});
            }else if(response.data == true){
                alert("You already reported this post");
            }
        });
    }

    const ReportPost = () => {
        try{
            tokenInfo = jwt_decode(localStorage.getItem('token'));
        }catch(e){
            tokenInfo = "";
            console.error(e);
        }
        if(tokenInfo!=""){
            CheckIfUserReportedPost();
        }else{
            alert("You are not logged in.You will be redirected to login.");
            history.push('/login');
        }
    }
    
    const isItVideo =  (element)=>{
        if(element.Path.includes("mp4")){
            return true
        }
        return false
    }

    const render = () =>{
        return(
            <>
            <div className="post__center">
            <div className="post">
                <div className="post__header">
                    <div className="post__headerLeft">
                        <Link to={"/profile/"+username}><h1>{username}</h1></Link>
                    </div>
                    <button className="report_but" onClick={ReportPost}>Report</button>
                </div>
            <div className="post__body">
                {loaded?
                    <Carousel>

                    {media.media.map(el=>(
                        el.Type== 0 ? 
                        <Carousel.Item margin="auto">
                                <img className="d-block w-100" height="400" width="300" src={el.Link} alt="my pic"/>
                        </Carousel.Item>
                        :
                        <Carousel.Item margin="auto">
                            <video  className="d-block w-100" height="400" width="300" controls>
                            <source lassName="d-block"  src={el.Link} type="video/mp4"/>
                            Your browser does not support the video tag.
                            </video>
                        </Carousel.Item>
                        
                    
                    ))}
                    </Carousel>
                :
                <p>No pic !!!</p>
                }
                
            </div>
            <div className="post__headerLeft">
                <h5>{username}</h5><h5 style={{marginLeft:"8px",fontWeight:'normal'}}>{description}</h5>
                <h5 style={{marginLeft:"8px"}}>{locationdesc}</h5><h5 style={{marginLeft:"8px"}}>{tags}</h5>
            </div>
            <div className="post__header">
                <button className="like_but" onClick={LikeThisPost}>Like</button><p>{likes}</p>
                <button className="dislike_but" onClick={DislikeThisPost}>Dislike</button><p>{dislikes}</p>
            </div>
            <div >
                <CommentInput postid={postid} getcoms={GetCommentsForPost}/>
                <p>Comments</p>
                        {comments ? (
                            comments?.map((comment,i)=>(
                                <div className="feed" key={i}>
                                    <Comment comments={comment.Comment} posteduser={comment.UserId}></Comment>
                                </div>
                            ))
                        ):<></>}
            </div>
        </div>
        </div>
        </>
        )
    }
        return(
            render()
    );

}