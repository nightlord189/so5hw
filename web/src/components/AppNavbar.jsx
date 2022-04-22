import {Container, Navbar, Button} from "react-bootstrap";
import {authTokenAtom, usernameAtom, userTypeAtom} from "../store/store";
import {useAtom} from "jotai";

const AppNavbar = () => {
    const [, setAuthToken] = useAtom(authTokenAtom);
    const [, setUsername] = useAtom(usernameAtom);
    const [, setUserType] = useAtom(userTypeAtom);

    const onLogout = () => {
        setAuthToken('');
        setUsername('');
        setUserType('');
    }

    return (
        <Navbar bg="dark" expand="lg" variant="dark">
            <Container>
                <Navbar.Brand href="#home">SO5HW</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Button variant="primary" onClick={onLogout}>Logout</Button>
            </Container>
        </Navbar>
    )
}

export default AppNavbar;