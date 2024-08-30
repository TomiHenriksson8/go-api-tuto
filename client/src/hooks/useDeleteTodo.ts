import { useState } from "react";
import { Todo } from "../types/Todo";

export const useDeleteTodo = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const deleteTodo = async (todo: Todo) => {
    setLoading(true)
    setError(null)
    try {
      const url = `http://localhost:3000/api/todos/${todo._id}`
      const response = await fetch(url, {
        method: 'DELETE',
      })
      if (!response.ok) {
        throw new Error('Error deleting todo')
      }
    } catch (error) {
      setError((error as Error).message)
    } finally {
      setLoading(false)
    }
  };

  return {loading, error, deleteTodo}
};
