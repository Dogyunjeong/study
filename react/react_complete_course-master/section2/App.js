import React, { Component } from 'react';
import './App.css';
import Person from './Person/Person';

class App extends Component {
  // only working in class component
  // if state property is changed, it will trigger to update UI
  state = {
    persons: [
      { id: 'afafa1', name: 'Max', age: 28 },
      { id: 'afafa2', name: 'Manu', age: 29 },
      { id: 'afaf2a', name: 'Stephanie', age: 26 },
    ],
    otherState: 'some other value',
    showPersons: false,
  }

  // Handler will notice that it is event handler
  switchNameHandler = (newName) => {
    // console.log('Was clicked!');
    // Don't do this: this.state.persons[0].name = 'MaxiMilian';

    // this will change state and props which is passed in jsx
    this.setState({ persons: [
      { name: newName, age: 28 },
      { name: 'Manu', age: 29 },
      { name: 'Stephanie', age: 27 },
    ]})
  }

  nameChangedHandler = (event, id) => {
    const personIndex = this.state.persons.findIndex(p => {
      return p.id === id;
    })

    // const person = this.state.persons[personIndex];
    
    // immutable way
    const person = {
      ...this.state.persons[personIndex]
    };
    // const person = Object.assign({}, this.state.persons[personIndex])

    person.name = event.target.value;

    const persons = [...this.state.persons];
    persons[personIndex] = person;


    this.setState({persons: persons})
  }
  
  deletePersonHandler = (personIndex) => {
    // this is not good as unpredictable
    // const persons = this.state.persons;

    // Should update state with immutable pattern
    // this is better way when manipulating array
    // const persons = this.state.persons.slice()
    // alternative way
    const persons = [...this.state.persons];
    persons.splice(personIndex, 1);
    this.setState({persons: persons});
  }

  togglePersonsHandler = () => {
    const doesShow = this.state.showPersons;
    this.setState({
      showPersons: !doesShow,
    })
  }

  // importing React allow to use render()
  // Whenever re-render render funtion is working
  render() {
    const style = {
      backgroundColor: 'white',
      font: 'inherit',
      border: '1px solid blue',
      padding: '8px',
      cursor: 'pointer',
    }

    let persons = null;

    if (this.state.showPersons) {
      persons = (
        <div>
          {this.state.persons.map((person, index) => {
            return <Person
              click={this.deletePersonHandler.bind(this, index)}
              name={person.name}
              age={person.age}
              // without key, react will re render whole when array is changed,
              // with key only the element will change
              key={person.id}
              changed={(event) => this.nameChangedHandler(event, person.id)}/>
          })}
        </div>
      );
    }
    // return jsx it will work as below of return statement
    return (
      <div className="App">
        <h1>Hi, I'm a React App</h1>
        <p>This is really Working</p>
        {/* this is inefficient way */}
        <button
          style={style}
          onClick={this.togglePersonsHandler}>Toggle Persons</button>
        { persons }
      </div>
    );
    // return React.createElement('div', {className: 'App'}, React.createElement('h1', null, 'Hi, I\'m a ReactApp!!'));
  }
}

export default App;