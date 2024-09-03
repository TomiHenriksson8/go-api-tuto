import { useState } from "react";
import { Todo } from "../types/Todo";

export const useCreateTodo = () => {
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [createTodo, setCreateTodo] = useState<Todo | null>(null);

  const postTodo = async (todoBody: string) => {
    const token = localStorage.getItem('token')
    setLoading(true);
    try {
      const response = await fetch("http://localhost:3000/api/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({ body: todoBody, completed: false }),
      });
      if (!response.ok) {
        throw new Error("Failed to create todo");
      }
      const data = await response.json();
      setCreateTodo(data);
    } catch (error) {
      setError((error as Error).message);
    } finally {
      setLoading(false);
    }
  };

  return { error, loading, postTodo, createTodo };
};
