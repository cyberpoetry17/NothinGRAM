import React from 'react'

const Hello = ()=>{
    // return (
    //     <div>
    //         <h1>hello men !!!</h1>
    //     </div>
    // )
     
    //JSX code
    return React.createElement('div',null,React.createElement(
        'h1',
        {id:'h1_Id', className: 'dummyCllass'},
        'Hello men !'
    ));
}

export default Hello