
import TodoForm from "./TodoForm";
import TodoList from "./TodoList";
import { useGetTodos } from "../hooks/useGetTodos";
import { useCreateTodo } from "../hooks/useCreateTodo";


const TodoPage = () => {
  const { todos, fetchTodos, loading: todosLoading, error: todosError } = useGetTodos();
  const { postTodo, loading: createLoading, error: createError } = useCreateTodo();

  const handleCreateTodo = async (todoBody: string) => {
    await postTodo(todoBody);
    fetchTodos()

  };

  return (
    <div>
      <TodoForm onCreate={handleCreateTodo} loading={createLoading} error={createError} />
      <TodoList todos={todos} loading={todosLoading} error={todosError} />
    </div>
  );
};

export default TodoPage;
