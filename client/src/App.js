import React, { useEffect,useState } from 'react'
import logo from './logo.svg'
import './App.css'

function App() {
  const [joke,setJoke] = useState('')
  useEffect (()=>{
    (async ()=>{
    const jokeTextResponse = await fetch('http://127.0.0.1:3000/joke')
    const jokeText = await jokeTextResponse.text();
    console.log(jokeText)
    setJoke(jokeText)
    })()
  },[])

  return (
    <div className="App" >
      {joke&&<p style={{"wordWrap": "break-word", width: "500px"}}>{joke}</p>}
    </div>
  )
}

export default App
