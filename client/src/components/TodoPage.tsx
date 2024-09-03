import TodoForm from "./TodoForm";
import TodoList from "./TodoList";
import { useGetTodos } from "../hooks/useGetTodos";
import { useCreateTodo } from "../hooks/useCreateTodo";
import { useEffect, useState } from "react";
import { Todo } from "../types/Todo";
import { useDeleteTodo } from "../hooks/useDeleteTodo";
import Welcome from "./Welcome";

const TodoPage = () => {
  const { todos, fetchTodos, loading: todosLoading, error: todosError } = useGetTodos();
  const { postTodo, loading: createLoading, error: createError } = useCreateTodo();
  const { deleteTodo, loading: deleteLoading, error: deleteError } = useDeleteTodo()

  const token = localStorage.getItem('token')

  const [localTodos, setLocalTodos] = useState<Todo[]>([]);

  useEffect(() => {
    console.log("Todos fetched from backend:", todos);
    setLocalTodos(todos);
  }, [todos]);

  const handleCreateTodo = async (todoBody: string) => {
    await postTodo(todoBody);
    fetchTodos();
  };

  const handleUpdateTodo = (updatedTodo: Todo) => {
    console.log("Updating todo in local state:", updatedTodo);
    const updatedTodos = localTodos.map((todo) =>
      todo._id === updatedTodo._id ? updatedTodo : todo
    );
    console.log("Updated todos array:", updatedTodos);
    setLocalTodos(updatedTodos);
  };

  const handleDeleteTodo = async (todo: Todo) => {
    await deleteTodo(todo)
    setLocalTodos((prevTodos => prevTodos.filter(t => t._id !== todo._id)))
  }

  return (

    <div>
      {token ? (
        <>
          <TodoForm onCreate={handleCreateTodo} loading={createLoading} error={createError} />
          <TodoList todos={localTodos} loading={todosLoading} error={todosError} onUpdateTodo={handleUpdateTodo} onDeleteTodo={handleDeleteTodo} />
        </>
      )  : (
        <>
          <Welcome />
        </>
      )}

    </div>
  );
};

export default TodoPage;
