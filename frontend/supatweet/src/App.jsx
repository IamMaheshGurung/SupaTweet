import { useState } from "react";
import React from React;


import axios from "axios"


function App() {
    const[tweets, setTweets] = useState([]);
    const[content, setContent] = useState("");

    useEffect(() => {
        axios.get("http://localhost:8080/tweets")
        .then((response)=> {
            setTweets(response.data)
        });

    },[])
    const postTWeet = () =>{
        axios
        .post("http://localhost:8080/tweet", {content})
        .then((response)=>{
            setTweets([...tweets, response.data]);
            setContent("");
        });
    };
    return (
        <div className="App">
            <h1>SupaTweet</h1>
            <div>
                <input type="text" value={content} onChange={(e)=> setContent(e.target.value)}/>
                <button onClick={postTweet}>Post Tweet</button>
            </div>
            <ul>
                {tweets.map((tweet, index)=>(<li key={index}>{tweet.content}</li>))}
            </ul>
        </div>
    )

}

export default App;