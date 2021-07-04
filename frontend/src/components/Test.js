import React,{useState} from 'react'

export default function Test() {
    const [count, setCount] = useState(0) 
    const incrementFive = () => {
        for (let i=0 ;i<5 ; i++){
            setCount(count+1)
        }
    }
    return (
        <div>
            {/* not safe way  */}
            <button onClick={incrementFive}>Click {count}</button> 
            {/* safe way */}
            <button onClick={()=>setCount(count => count + 1)}>Click {count}</button>
        </div>
    )
}

// rfce