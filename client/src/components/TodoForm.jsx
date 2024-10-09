import axios from "axios";
import React, { useState } from "react";
import toast from "react-hot-toast";
import { BASE_URL } from "../TodoApp";

const TodoForm = ({getTodo}) => {
  const [body, setBody] = useState("");

  const createTodo = async () => {
    try {
      const result = await axios.post(BASE_URL + "/todo", {body:body});
      console.log("the result is: ", result);
      toast.success("Todo created successfully")
       setBody(''); // Clear input field
       getTodo()
    } catch (error) {
      console.log("Error create todo from frontend: ", error);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (body.trim()) {
      createTodo();
      
     
    }
  };
  return (
    <form onSubmit={handleSubmit} className="mb-4">
      <div className="input-group">
        <input
          type="text"
          className="form-control"
          value={body}
          onChange={(e) => setBody(e.target.value)}
          placeholder="Enter a new task"
        />
        <button className="btn btn-primary " type="submit">
          <i className="fa fa-plus"></i>
        </button>
      </div>
    </form>
  );
};

export default TodoForm;
