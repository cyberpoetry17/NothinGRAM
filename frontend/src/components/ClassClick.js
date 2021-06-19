import React, { Component } from 'react'

class ClassClick extends Component {
    clickHandler(){
        console.log('Clicked to the button')
    }
    render() {
        return (
            <div>
                <button onClick={this.clickHandler}>CLICK ME</button>
            </div>
        )
    }
}

export default ClassClick
