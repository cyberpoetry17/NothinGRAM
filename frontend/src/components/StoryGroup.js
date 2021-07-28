import React from 'react'
import {useState,useEffect} from 'react'
import Story from './Story'
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogActions from '@material-ui/core/DialogActions';
import {Form,Container,Row,Col,Button,Carousel} from 'react-bootstrap';

export default function StoryGroup(storyList) {
    const [stories, setStories] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const [open, toggleOpen] = useState(false);
    useEffect(()=>{
        setStories(storyList.storyList)
        console.log("storyList:",storyList.storyList)
    },[])

    useEffect(()=>{
        if(stories== null){
            return;
        }
        setLoaded(true);
    },[stories])

    const openDialog = () => {
        toggleOpen(true);
    }

    const handleClose = () => {
        toggleOpen(false);
      };

    return (
        <div>
            {/* <Carousel>
            {loaded?
                stories.map(s=>(
                    
                        <Carousel.Item margin="auto">
                            <Story UserId={s.UserId} IdStory={s.IdStory} postId={s.PostID} type={s.Type}/> 
                        </Carousel.Item>
                    
                )):
                <p>loading..</p>
            }
            </Carousel> */}
            {/* <button onClick={openDialog}>Open </button> */}
            {loaded?
                <div onClick={openDialog}>
                    <Story onClick={openDialog} UserId={stories[0].UserId} IdStory={stories[0].IdStory} postId={stories[0].PostID} type={stories[0].Type}/> 
                </div>
                 :
                 <p>loading..</p>
            }
           
            {/* Dialog */}
            <Dialog open={open} aria-labelledby="form-dialog-title">
                <Carousel>
                    {loaded?
                        stories.map(s=>(
                            
                                <Carousel.Item margin="auto">
                                    <Story UserId={s.UserId} IdStory={s.IdStory} postId={s.PostID} type={s.Type} size="1"/> 
                                </Carousel.Item>
                            
                        )):
                        <p>loading..</p>
                    }
                </Carousel>
                <DialogActions>
                    <Button onClick={handleClose} color="primary">
                    Close
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    )
}
