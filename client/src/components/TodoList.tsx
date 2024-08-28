import React from "react";
import { TodoListProps } from "../types/Todo";




const TodoList: React.FC<TodoListProps> = ({ todos, loading, error }) => {
  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="pt-4">
      <div className="border-b border-1 mb-4"></div>
      {todos && todos.length > 0 ? (
        <div className="text-center">
          {todos.map((todo) => (
            <div
              key={todo._id}
              className="flex justify-center mb-4"
            >
              <div className="flex items-center justify-between w-[600px] bg-gray-100 p-4 rounded-lg shadow">
                <div className="flex-1 text-left">
                  <p className="font-semibold">{todo.body}</p>
                </div>
                <div>
                  {todo.completed ? (
                    <button className="bg-green-500 text-white p-2 rounded w-[100px]">
                      Completed
                    </button>
                  ) : (
                    <button className="bg-red-500 text-white p-2 rounded w-[100px]">
                      Not Done
                    </button>
                  )}
                </div>
              </div>
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
