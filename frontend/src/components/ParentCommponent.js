import React, { Component } from 'react'

export class ParentCommponent extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             parentName: 'Parent'
        }
        this.greetParent = this.greetParent.bind(this)
    }
    
    greetParent(){
        alert('Hello ${this.state.parenName}')
        //alert('Hello'+this.state.parentName)
    }
    render() {
        return (
            <div>
                
            </div>
        )
    }
}

export default ParentCommponent
