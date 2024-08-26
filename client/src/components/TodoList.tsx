import React from "react";
import { TodoListProps } from "../types/Todo";




const TodoList: React.FC<TodoListProps> = ({ todos, loading, error }) => {
  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="pt-4">
      <div className="text-center text-[18px]">Your todos:</div>
      {todos && todos.length > 0 ? (
        <div className="text-center">
          {todos.map((todo) => (
            <div key={todo._id}>
              <p>Completed: {todo.completed ? "Yes" : "No"}</p>
              <p>{todo.body}</p>
              <p>ID: {todo._id}</p>
            </div>
          ))}
        </div>
      ) : (
        <div>No todos</div>
      )}
    </div>
  );
};

export default TodoList;
