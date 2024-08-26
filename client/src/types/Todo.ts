export interface Todo {
  _id: string;
  body: string;
  completed: boolean;
}

export interface TodoListProps {
    todos: Todo[];
    loading: boolean;
    error: string | null;
  }
