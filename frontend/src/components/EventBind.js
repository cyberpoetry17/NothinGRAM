import React, { Component } from 'react'

class EventBind extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
             message: 'Hello',
             stefan: '',
             dusan: ''
        }

        this.clickHandler = this.clickHandler.bind(this)
    }
    clickHandler(){
        this.setState({
            ...this.state,
            message: 'Goodbye!'
        })
    }
    // class property approach (<button onClick={this.clickHandler}>Click</button>) same as constructor bind approach
    // clickHandler = ()=> {
    //     this.setState({
    //         message: 'Goodbye!'
    //     })
    // }
    render() {
        return (
            <div>
                <div>{this.state.message}</div>
                {/* binding */}
                {/* <button onClick={this.clickHandler.bind(this)}>Click</button> */}
                {/* arrow function*/}
                {/* <button onClick={()=>this.clickHandler()}>Click</button> */}
                {/* bind inside consturctor */}
                <button onClick={this.clickHandler}>Click</button>
            </div>
        )
    }
}

export default EventBind
