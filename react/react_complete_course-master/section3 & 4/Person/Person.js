import React from 'react';

import classes from './Person.css';

const person = (props) =>{
    const rnd = Math.random();

    if (rnd > 0.7) {
        throw new Error('Something went wrong');
    }
    // When using class-based components, it's this.props instead of props param
    return (
        <div className={classes.Person}>
            <p onClick={props.click}>I'm a { props.name } and I am { props.age } years old!</p>
            <p>{ props.children }</p>
            <input type="text" onChange={props.changed} value={props.name} />
        </div>
    )
    // return <p>I'm a Person and I am { Math.floor(Math.random() * 30) } years old!</p>
}

export default person;