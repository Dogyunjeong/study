import React from 'react';

import classes from './Cockpit.css';

const cockpit = (props) => {
        // let classes = ['red', 'bold'].join(' '); // "red bold"
    const assignedClasses = [];
    let btnClass = '';

    if (props.showPersons) {
        btnClass = classes.Red;
    }
    if ( props.persons.length <= 2 ) {
        assignedClasses.push( classes.red );
    }
    if ( props.persons.length <= 1 ) {
        assignedClasses.push( classes.bold );
    }
    return (
        <div className={classes.Cockpit}>
            <h1>{props.appTitle}</h1>
            <p className={assignedClasses.join(' ')}>This is really Working</p>
            {/* this is inefficient way */}
            <button
            className={btnClass}
            // style={style}
            onClick={props.clicked}>Toggle Persons</button>
        </div>
    )
};

export default cockpit;
