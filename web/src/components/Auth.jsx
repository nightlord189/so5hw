/* eslint-disable max-len */
/* eslint-disable no-mixed-operators */
/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable no-useless-concat */
/* eslint-disable react/no-unescaped-entities */
import { useState } from 'react';
import { useFormik } from 'formik';
import { useAtom } from 'jotai'
import axios from 'axios';
import {
    Form, Button, Container, Row, Col, Alert
} from 'react-bootstrap';
import {authPath} from '../config';
import {authTokenAtom, usernameAtom, userTypeAtom} from "../store/store";

const Auth = () => {
    const [formError, setFormError] = useState('');
    const [, setAuthToken] = useAtom(authTokenAtom);
    const [, setUsername] = useAtom(usernameAtom);
    const [, setUserType] = useAtom(userTypeAtom);

    const formik = useFormik({
        initialValues: {
            username: '',
            password: '',
            loginType: 'customer',
        },
        onSubmit: async (values) => {
            if (values.username === '' || values.password === '') {
                setFormError('empty username or password')
                return
            }

            console.log(`login submit: ${JSON.stringify(values)}`);

            try {
                const response = await axios.post(authPath, {
                    username: values.username,
                    password: values.password,
                    type: values.loginType,
                });
                console.log(`login success: ${JSON.stringify(response.data)}`);
                setAuthToken(response.data);
                setUsername(values.username);
                setUserType(values.loginType);
            } catch (error) {
                console.log(`login failure: ${error}, status: ${error.response.status}`);
                if (error.response.status === 401) {
                    setFormError('bad credentials');
                } else {
                    setFormError(`error: ${error.response.data}`);
                }
            }
        },
    });

    return (
        <Container className="justify-content-md-center">
            <Row className="mt-4">
                <Col xs lg="4"/>
                <Col className="mx-auto">
                    <h3>Auth</h3>

                    <Form onSubmit={formik.handleSubmit}>
                        <Form.Group className="mb-3" controlId="formEmail">
                            <Form.Label>Login</Form.Label>
                            <Form.Control
                                value={formik.values.username}
                                onChange={formik.handleChange}
                                name = "username"
                                type="text" placeholder="Enter login" />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                                type="password"
                                name = "password"
                                value={formik.values.password}
                                onChange={formik.handleChange}
                                placeholder="Password" />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formLoginType">
                            <Form.Label>Login type</Form.Label>
                            <Form.Select
                                name = "loginType"
                                value={formik.values.loginType}
                                onChange={formik.handleChange}
                                aria-label="Login type">
                                <option value="customer">Customer</option>
                                <option value="merchandiser">Merchandiser</option>
                            </Form.Select>
                        </Form.Group>

                        {formError && (
                            <Alert variant="danger" dismissible onClose={() => setFormError('')}>
                                {formError}
                            </Alert>
                        )}
                        <Button variant="primary"
                                disabled={formik.isSubmitting}
                                type="submit">
                            Submit
                        </Button>
                    </Form>
                </Col>
                <Col xs lg="4"/>
            </Row>
        </Container>
    );
};

export default Auth;