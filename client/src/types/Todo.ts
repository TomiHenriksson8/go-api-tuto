export interface Todo {
  _id: string;
  body: string;
  completed: boolean;
}

export interface TodoListProps {
  todos: Todo[];
  loading: boolean;
  error: string | null;
  onUpdateTodo: (updatedTodo: Todo) => void;
  onDeleteTodo: (todo: Todo) => void;
}
