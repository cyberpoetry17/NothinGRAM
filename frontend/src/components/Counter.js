import React, { Component } from 'react'
// rce 
//rconst
export class Counter extends Component {
    
    constructor(props) {
        super(props)
    
        this.state = {
             count: 0
        }
    }
    increment(){
        // this.setState({
        //     count:this.state.count+1
        // },
        // ()=>{
        //     console.log('after setState ',this.state.count)
        // })
        //when updating state based on previous state pass prevState and modify based on prevState
        //first parameter is prevState and secound is props if needed
        this.setState((prevState,props)=>({
            count: prevState.count+1
        }))
    }
    render() {
        return (
            <div>
                {this.state.count}
                <button onClick={()=>this.increment()}>Increment </button>            
            </div>
        )
    }
}

export default Counter
