import React, {useEffect, useState } from 'react'
import axios from "axios"

import './App.css'

function App() {
  const [tweets, setTweets] = useState([]);
  const [content, setContent] = useState("");
  useEffect(()=>{
    axios.get("http://localhost:8080/tweets").then((response) =>
    {
      setTweets(response.data);
    });
  },[])

  const postTweet = () => {
    axios.post("http://localhost:8080/tweet", {content})
    .then((response)=> {
      setTweets([...tweets, response.data]);
      setContent("")
    })
  }

  return (
    <>
      <div className='App'>
        <h1>Supa Tweet</h1>
        <div>
          <input type="text" value={content} onChange={(e) => setContent(e.target.value)}/>
          <button onClick={postTweet}>Post Tweet</button>
        </div>
        <ul>
          {tweets.map((tweet, index)=>{
            <li key={index}>{tweet.content}</li>
          })}
        </ul>

        
      </div>
      
     
      
    </>
  )
}

export default App
