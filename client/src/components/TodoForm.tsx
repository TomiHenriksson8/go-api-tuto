import { useState } from "react";
import { Link } from "react-router-dom";
import { logout } from "../utils/utils";

interface TodoFormProps {
  onCreate: (todoBody: string) => Promise<void>; // The function to create a new todo
  loading: boolean;
  error: string | null;
}

const TodoForm: React.FC<TodoFormProps> = ({ onCreate, loading, error }) => {
  const [todoBody, setTodoBody] = useState<string>("");

  const token = localStorage.getItem('token')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onCreate(todoBody);
    setTodoBody("");
  };

  return (
    <div className="py-3 w-full flex justify-center">
      <form onSubmit={handleSubmit} className="flex gap-1">
        <input
          type="text"
          value={todoBody}
          className="bg-gray-300 text-black h-12 w-[500px] rounded"
          placeholder="  add todo..."
          onChange={(e) => setTodoBody(e.target.value)}
        />

        <button className="px-6 py-3 bg-green-500 rounded" disabled={loading} type="submit">
          {loading ? "Adding..." : "Add"}
        </button>
      </form>
      {token && (
        <button className="bg-slate-400 border border-bg-slate-600 px-3 py-2 rounded-md hover:border-slate-400 text-[16px] cursor-pointer" onClick={logout}>Logout</button>
      )}

      {error && <div className="text-red-500 ">Error adding todo: {error}</div>}
    </div>
  );
};

export default TodoForm;
