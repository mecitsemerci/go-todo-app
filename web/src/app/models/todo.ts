import { NgZone } from '@angular/core';

export interface Todo {
  id: string;
  title: string;
  description: string;
  priority_level: number;
  completed: boolean;
  create_at: string;
  update_at: string;
}

export interface CreateTodo {
  title: string;
  description?: string;
  priority_level?: number;
}

export interface CreateTodoResult {
  todo_id: string;
}

export interface UpdateTodo {
  title: string;
  description: string;
  priority_level: number;
  completed: boolean;
}


