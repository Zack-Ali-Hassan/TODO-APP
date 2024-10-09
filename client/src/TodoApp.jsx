import React, { useEffect, useState } from "react";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";
import axios from "axios";
export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:5555/api" : "/api"
const TodoApp = () => {
  const [todos, setTodos] = useState([]);
  const getTodo = async () => {
    try {
      const { data } = await axios.get(BASE_URL + "/todos");
      setTodos(data);
      console.log("Get all Todos: ", data);
    } catch (error) {
      console.log("Error display from frontend: ", error);
    }
  };
  useEffect(() => {
    getTodo();
  }, []);


  return (
    <>
      <div className="container mt-4">
        <h1 className="text-center mb-4 fw-bold fs-2">
          Appkan waa mid sahlan oo loogu talagalay maareynta hawlaha to-do.
        </h1>
        <TodoForm getTodo={getTodo}/>
        <TodoList todos={todos} getTodo={getTodo}/>
      </div>
    </>
  );
};

export default TodoApp;
