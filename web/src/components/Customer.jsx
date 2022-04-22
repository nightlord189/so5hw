import React, { useState, useEffect } from 'react';
import {Container, Row, Col} from "react-bootstrap";
import AppNavbar from './AppNavbar.jsx';

const Customer = () => {
    useEffect(() => {
        document.title = `Вы нажали 1 раз`;
    });

    return (
        <Container className="justify-content-md-center">
            <AppNavbar />
            <Row className="mt-4">
                <Col xs lg="4"/>
                <Col className="mx-auto">
                    <h3>Customer</h3>
                </Col>
            </Row>
        </Container>
    )
}

export default Customer;