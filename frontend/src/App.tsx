import React, { useEffect, useState } from "react";

import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

import Todoitem from "./components/todoitem";

function App() {
    const [currentDate, setCurrentDate] = useState<string>("");

    useEffect(() => {
        const date = new Date();
        setCurrentDate(date.toLocaleDateString());
    }, []);

    return (
        <>
            <Navbar expand="lg" className="bg-body-tertiary navbar-background">
                <Container className="font-color">
                    <Navbar.Brand href="#home">Simple Todo</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link href="#home">ホーム</Nav.Link>
                            <Nav.Link href="#link">今日やることを決める</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>

            <Container className="text-center mt-5">
                <Row>
                    <Col>
                        <p>今日すること</p>
                        <Todoitem currentDate={currentDate} todayflag={true} />
                    </Col>
                    <Col>
                        <p>それ以外にすること</p>
                        <Todoitem currentDate={currentDate} todayflag={false} />
                    </Col>
                </Row>
            </Container>
        </>
    );
}

export default App;
