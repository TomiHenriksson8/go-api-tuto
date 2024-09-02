import React, { useState } from "react";
import { Todo, TodoListProps } from "../types/Todo";
import { useUpdateTodos } from "../hooks/useUpdateTodo";

const TodoList: React.FC<TodoListProps> = ({ todos, loading, error, onUpdateTodo, onDeleteTodo }) => {
  const { updateTodos } = useUpdateTodos();

  const [selectedTodo, setSelectedTodo] = useState<Todo | null>(null)

  const handleClick = async (todo: Todo) => {
    try {
      await updateTodos(todo);
      onUpdateTodo({ ...todo, completed: !todo.completed });
    } catch (error) {
      console.error("Error updating todo", error);
    }
  };

  const handleClickDelete = async (selectedTodo: Todo) => {
    try {
      await onDeleteTodo(selectedTodo)
      setSelectedTodo(null)
    } catch (error) {
      console.error("Error deleting todo ", error)
    }
  }

  const todoInfo = (todo: Todo) => {
    console.log('info', todo)
    setSelectedTodo(todo)
  }

  const closeModal = () => {
    setSelectedTodo(null)
  }

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="pt-4">
      <div className="border-b border-1 mb-4"></div>
      {todos && todos.length > 0 ? (
        <div className="text-center">
          {todos.map((todo) => (
            <div key={todo._id} className="flex justify-center mb-4">
              <div onClick={() => todoInfo(todo)} className="flex items-center justify-between w-[600px] bg-gray-100 p-4 rounded-lg shadow">
                <div className="flex-1 text-left">
                  <p className="font-semibold">{todo.body}</p>
                </div>
                <div>
                  <button
                    onClick={(e) => {
                      e.stopPropagation()
                      handleClick(todo)
                    }}
                    className={`text-white p-2 rounded w-[100px] ${
                      todo.completed ? "bg-green-500" : "bg-red-500"
                    }`}
                  >
                    {todo.completed ? "Completed" : "Not Done"}
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div>No todos</div>
      )}
      {selectedTodo && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-80 z-50">
          <div className="bg-white rounded-lg shadow-lg p-6 w-full max-w-md flex flex-col">
            <h2 className="text-xl font-semibold mb-4 text-center">{selectedTodo.body}</h2>
            <div className="flex flex-row justify-between pb-3">
              <p className="mb-4"><strong>Completed:</strong> {selectedTodo.completed ? "Yes" : "No"}</p>
              <button
                onClick={() => handleClickDelete(selectedTodo)}
                className="bg-red-500 border border-bg-slate-600 px-3 py-2 rounded-md hover:border-slate-400 text-[16px] cursor-pointer"
              >
                Delete
              </button>
            </div>
            <button
              onClick={closeModal}
              className="bg-slate-400 border border-bg-slate-600 px-3 py-2 rounded-md hover:border-slate-400 text-[16px] cursor-pointer"
            >
              Close
            </button>

          </div>
        </div>
      )}
    </div>
  );
};

export default TodoList;
