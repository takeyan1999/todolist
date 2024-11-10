import React, { useState, useEffect } from "react";
import InputGroup from "react-bootstrap/InputGroup";
import { Button } from "react-bootstrap";

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
            try {
                const response = await axios.get("http://localhost:8080/todos");
                // console.log(response.data); // 取得した全データをログに出力
                // console.log(response.data[0]); // 最初の項目をログに出力
                setTodoItems(response.data); // データをstateにセット
                // console.log(TodoItems); // stateにセットされた後のTodoItemsをログに出力
            } catch (error) {
                console.error("Error fetching todos:", error);
            }
        })();
    }, []);

    const creatTodo = () => {
        const todo = {
            task: "string",
            priority: 3,
            status: true,
            deadline: "2024-11-15",
        };

        const url = axios
            //todoのところにリストが入る。
            .post("http://localhost:8080/todos", todo)

            //then=成功した場合の処理
            .then(() => {
                console.log(url);
                console.log("good");
            })
            //catch=エラー時の処理
            .catch((err) => {
                console.log("err:", err);
            });
    };

    const deleteTodo = (deleteid: number) => {
        console.log("delete");
        console.log(deleteid);
        const url = axios
            .delete("http://localhost:8080/todos?id=" + deleteid)
            .then(() => {
                console.log("削除ID:", url);
            })
            .catch((err) => {
                console.log("err:", err);
            });
    };

    return (
        <>
            <Button onClick={creatTodo}>create todo</Button>
            {TodoItems.map((TodoItems) => (
                <React.Fragment key={TodoItems.id}>
                    {appProps.todayflag
                        ? new Date(TodoItems.deadline) <= new Date(appProps.currentDate) && (
                              <InputGroup className="mb-3">
                                  <InputGroup.Checkbox aria-label="Checkbox for following text input" />
                                  {new Date(TodoItems.deadline) < new Date(appProps.currentDate) ? (
                                      <span className="form-control text-danger">
                                          {TodoItems.task}
                                          {TodoItems.deadline}
                                      </span>
                                  ) : (
                                      <span className="form-control">{TodoItems.task}</span>
                                  )}
                                  <Button className="bg-white border" onClick={() => deleteTodo(TodoItems.id)}>
                                      <i className="bi bi-x-lg text-danger"></i>
                                  </Button>
                              </InputGroup>
                          )
                        : new Date(TodoItems.deadline) > new Date(appProps.currentDate) && (
                              <InputGroup className="mb-3">
                                  <InputGroup.Checkbox aria-label="Checkbox for following text input" />
                                  <span className="form-control">{TodoItems.task}</span>
                                  <Button className="bg-white border" onClick={() => deleteTodo(TodoItems.id)}>
                                      <i className="bi bi-x-lg text-danger"></i>
                                  </Button>
                              </InputGroup>
                          )}
                </React.Fragment>
            ))}
        </>
    );
};

export default Todoitem;
