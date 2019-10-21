import React, { Component } from 'react';
import './App.css';
import Validation from './validation/Validation';
import Char from './char/Char';


class App extends Component {
  state = {
    input: '',
    inputLength: 0
  }

  inputChangeHandler = (event) => {
    console.log(event.target.value)
    console.log(event.target.value.length)
    this.setState({ input: event.target.value, inputLength: event.target.value ? event.target.value.length : 0 })
  }

  validationClickHandler = (event, idx) => {
    const input = this.state.input.slice(0, idx)  + this.state.input.slice(idx + 1)
    this.setState({ input, inputLength: input.length })
  }
  render() {
    let chars = null;
    if (this.state.inputLength) {
      chars = (
        <div>
          {this.state.input.split('').map((char, idx) => {
            return (<Char char={char} click={(event) => { this.validationClickHandler(event, idx)}}></Char>)
          })}
        </div>
      )
    }
    return (
      <div className="App">
        <input type="text" onChange={this.inputChangeHandler} value={this.state.input}/>
        <p>{this.inputLength}</p>
        <Validation leng={this.state.inputLength}></Validation>
        {chars}
        <ol>
          <li>Create an input field (in App component) with a change listener which outputs the length of the entered text below it (e.g. in a paragraph).</li>
          <li>Create a new component (=> ValidationComponent) which receives the text length as a prop</li>
          <li>Inside the ValidationComponent, either output "Text too short" or "Text long enough" depending on the text length (e.g. take 5 as a minimum length)</li>
          <li>Create another component (=> CharComponent) and style it as an inline box (=> display: inline-block, padding: 16px, text-align: center, margin: 16px, border: 1px solid black).</li>
          <li>Render a list of CharComponents where each CharComponent receives a different letter of the entered text (in the initial input field) as a prop.</li>
          <li>When you click a CharComponent, it should be removed from the entered text.</li>
        </ol>
        <p>Hint: Keep in mind that JavaScript strings are basically arrays!</p>
      </div>
    );
  }
}

export default App;
