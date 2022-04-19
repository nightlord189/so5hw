import React from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
} from 'react-router-dom';
import Auth from './Auth.jsx';

const App = () => (
    <Router>
        <Switch>
            <Route path="/">
                <Auth />
            </Route>
        </Switch>
    </Router>
);

export default App;