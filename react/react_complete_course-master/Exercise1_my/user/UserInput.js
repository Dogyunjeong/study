import React from 'react';

import './UserInput.css';
const userInput = (props) => {
    return (
        <div className="user-input">
            <input
                onChange={props.change}
                value={props.userName}
                type="text" />
        </div>
    )
};

export default userInput