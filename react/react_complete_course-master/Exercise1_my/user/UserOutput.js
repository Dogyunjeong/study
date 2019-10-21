import React from 'react';

const userOutput = (props) => {
    const userOutputStyle = {
        width: '60%',
        border: '1px solid #ccc',
    }
    return (
        <div
            className="user-output"
            sytles={userOutputStyle}>
            <p>User Name: {props.userName}</p>
            <p>s</p>
        </div>
    )
}

export default userOutput;