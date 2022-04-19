/* eslint-disable max-len */
/* eslint-disable no-mixed-operators */
/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable no-useless-concat */
/* eslint-disable react/no-unescaped-entities */
import React from 'react';
import { observer } from 'mobx-react';
import {
    Form, Button, Tooltip, Container, Row, Col,
} from 'react-bootstrap';
import MainStore from '../store/main.js';

const Auth = () => {
    const handleChangeRate = (e) => {
        if (!e.target.validity.badInput) {
            MainStore.calculator.rate = e.target.value;
        }
    };

    const handleChangeTerm = (e) => {
        if (!e.target.validity.badInput) {
            MainStore.calculator.term = e.target.value;
        }
    };

    const handleChangeCapitalization = (e) => {
        MainStore.calculator.capitalization = e.target.checked;
    };

    const handleChangeAmount = (e) => {
        if (!e.target.validity.badInput) {
            MainStore.calculator.amount = Number(e.target.value);
        }
    };

    return (
        <Container className="justify-content-md-center">
            <Row className="mt-4">
                <Col xs lg="4"/>
                <Col className="mx-auto">
                    <h3>Auth</h3>
                    <Form>
                        <Form.Group className="mb-3" controlId="formEmail">
                            <Form.Label>Login</Form.Label>
                            <Form.Control type="email" placeholder="Enter login" />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control type="password" placeholder="Password" />
                        </Form.Group>

                        <Form.Group className="mb-3" controlId="formLoginType">
                            <Form.Label>Login type</Form.Label>
                            <Form.Select aria-label="Login type">
                                <option value="customer">Customer</option>
                                <option value="merchandiser">Merchandiser</option>
                            </Form.Select>
                        </Form.Group>

                        <Button variant="primary" type="submit">
                            Submit
                        </Button>
                    </Form>
                </Col>
                <Col xs lg="4"/>
            </Row>
        </Container>
    );
};

export default observer(Auth);