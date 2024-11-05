import React, { useState, useEffect } from "react";
import InputGroup from "react-bootstrap/InputGroup";

import axios from "axios";

type TodoList = {
    id: number;
    task: string;
    priority: number;
    status: boolean;
    deadline: string;
};

type appProps = {
    currentDate: string;
    todayflag: boolean;
};

const Todoitem = (appProps: appProps) => {
    const [TodoItems, setTodoItems] = useState<TodoList[]>([
        { id: 0, task: "", priority: 0, status: false, deadline: "" },
    ]);

    useEffect(() => {
        (async () => {
            const data = await axios.get("http://localhost:8080");
            console.log(data.data);
            console.log(data.data[0]);
            setTodoItems(data.data);
            console.log(TodoItems);
        })();
    }, []);

    return (
        <>
            {TodoItems.map((TodoItems) => (
                <React.Fragment key={TodoItems.id}>
                    {appProps.todayflag
                        ? new Date(TodoItems.deadline).toLocaleDateString() === appProps.currentDate && (
                              <InputGroup className="mb-3">
                                  <InputGroup.Checkbox aria-label="Checkbox for following text input" />
                                  <span className="form-control">{TodoItems.task}</span>
                              </InputGroup>
                          )
                        : new Date(TodoItems.deadline).toLocaleDateString() !== appProps.currentDate && (
                              <InputGroup className="mb-3">
                                  <InputGroup.Checkbox aria-label="Checkbox for following text input" />
                                  <span className="form-control">{TodoItems.task}</span>
                              </InputGroup>
                          )}
                </React.Fragment>
            ))}
        </>
    );
};

export default Todoitem;
