// TODO:Add a result area to show the searching result 
import { useState } from 'react';
import './App.css';
import { SearchInput } from "../wailsjs/go/main/SearchEngine"

function App() {
    const [resultText, setResultText] = useState(["Please enter your search below 👇"]);
    const [text, setText] = useState('');
    const updateText = (e: any) => setText(e.target.value);
    const updateResultText = (result: Array<string>) => setResultText(result);

    function searchInput() {
        SearchInput(text).then(updateResultText)
    }

    return (
        <div id="App">
            <div className="result-area">
                <p id="result" className="result-label" style={{ whiteSpace: 'pre-wrap' }}>
                    {resultText.join("\n\n")}
                </p>
            </div>

            <div className='container-search'>
                <div id="input" className="input-box">
                    <input id="name" className="input" onChange={updateText} autoComplete="off" name="input" type="text" />
                    <button className="btn" onClick={searchInput}>Search</button>
                </div>
            </div>
        </div>
    )
}

export default App
