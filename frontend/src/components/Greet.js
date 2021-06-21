import React from 'react'

// function Greet(){
//     return <h1>Greet hello !!!</h1>
// }

const Greet = (props) => {
    //way to destructur props and state
    const {name, heroName} = props
    
    console.log(props);
    return (
        <div>
            <h1>Hello {props.name}: {props.heroName} !</h1>
            {props.children}
        </div>
    )
}
export default Greet