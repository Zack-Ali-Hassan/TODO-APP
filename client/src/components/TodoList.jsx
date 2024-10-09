import React from "react";
import TodoItem from "./TodoItem";
import axios from "axios";
import toast from "react-hot-toast";
import { BASE_URL } from "../TodoApp";

const TodoList = ({ todos, getTodo }) => {
  const deleteTodo = async (id) => {
    try {
      if (confirm(`Are you sure you want to delete ${id} ?`)) {
        const { data } = await axios.delete(
          BASE_URL + `/todo/${id}`
        );
        console.log("Result of deleting todo ", data);
        toast.success(data.msg);
        getTodo();
      }
    } catch (error) {
      console.log("Error deleting Todo from frontend: ", error.response.data);
      toast.error("Error....");
    }
  };
  const updateTodo = async (id) => {
    try {
      if (confirm(`Are you sure you want to update ${id} ?`)) {
        const { data } = await axios.patch(
          BASE_URL + `/todo/${id}`,{completed:true},
        );
        console.log("Result of updating todo ", data);
        toast.success(data.msg);
        getTodo();
      }
    } catch (error) {
      console.log("Error deleting Todo from frontend: ", error.response.data);
      toast.error("Error....");
    }
  };
  return (
    <ul className="list-group">
      {todos.map((todo, index) => (
        <TodoItem
          key={index}
          todo={todo}
          deleteTodo={() => deleteTodo(todo._id)}
          updateTodo={()=>updateTodo(todo._id)}
        />
      ))}
    </ul>
  );
};

export default TodoList;
