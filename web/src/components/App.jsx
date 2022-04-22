import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect,
} from 'react-router-dom';
import Auth from './Auth.jsx';
import Customer from './Customer.jsx';
import Merchandiser from './Merchandiser.jsx';
import {authTokenAtom, userTypeAtom} from "../store/store";
import { useAtom } from 'jotai'

const App = () => {
    const [authToken, ] = useAtom(authTokenAtom);
    const [userType, ] = useAtom(userTypeAtom);

    const isAuthorized = authToken !== '';

    let redirectUserPath = '/customer';
    if (userType === 'merchandiser') {
        redirectUserPath = '/merchandiser';
    }

    return (
        <Router>
            <Switch>
                <Route path="/auth">
                    {isAuthorized ? <Redirect to={redirectUserPath} /> : <Auth/>}
                </Route>
                <Route path="/customer">
                    {authToken === '' ? <Redirect to="/auth" /> : userType === 'customer' ? <Customer /> : <Redirect to={redirectUserPath} />}
                </Route>
                <Route path="/merchandiser">
                    {authToken === '' ? <Redirect to="/auth" /> : userType === 'merchandiser' ? <Merchandiser /> : <Redirect to={redirectUserPath} /> }
                </Route>
            </Switch>
        </Router>
    );
};

export default App;