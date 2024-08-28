import { useState } from "react";
import { Todo } from "../types/Todo";

export const useUpdateTodos = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const updateTodos = async (todo: Todo) => {
    setLoading(true);
    try {
      const url = `http://localhost:3000/api/todos/${todo._id}`;

      const updatedCompletedStatus = !todo.completed

      const response = await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            completed: updatedCompletedStatus
        })
      });

      if (!response.ok) {
        throw new Error("failed to update todo");
      }
      const data = await response.json();
      console.log('Todo updated', data)

    } catch (error) {
        console.error("Error updating todo", error)
        setError("Failed to update todo")
    } finally {
        setLoading(false);
    }
  };


  return {loading, error, updateTodos}
};
