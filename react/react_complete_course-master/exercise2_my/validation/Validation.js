import React from 'react';

const validation = (props) => {
    let validation = ''
    if (props.leng) {
        if (props.leng < 5) validation = 'Text too short'
        else validation = 'Text long enough'    
    }

    return (
        <div>
            <p>{validation}</p>
        </div>
    )
}

export default validation