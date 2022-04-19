import React from 'react';
import ReactDOM from 'react-dom';
// eslint-disable-next-line import/extensions
import { configure } from 'mobx';
import App from './components/App.jsx';

const init = async () => {
    configure({
        enforceActions: 'never',
    });

    ReactDOM.render(
        <App />,
        document.getElementById('root'),
    );
};

export default init;