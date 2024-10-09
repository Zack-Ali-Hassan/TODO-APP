import React from 'react';

const TodoItem = ({ todo, deleteTodo, updateTodo }) => {
  return (
    <li className={`list-group-item d-flex justify-content-between align-items-center ${todo.completed ? 'list-group-item-success' : ''}`}>
      <span style={{ textDecoration: todo.completed ? 'line-through' : 'none' }}>
        {todo.body}
      </span>
      <div>
        <button className="btn btn-success btn-sm me-2" onClick={updateTodo}>
          {todo.completed ? 'Undo' : 'Complete'}
        </button>
        <button className="btn btn-danger btn-sm" onClick={deleteTodo}>Delete</button>
      </div>
    </li>
  );
};

export default TodoItem;
