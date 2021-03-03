import { NgZone } from '@angular/core';

export interface Todo {
  id: string | number;
  title: string;
  description?: string;
  priority_level?: number;
  completed: boolean;
  create_at?: string;
  update_at?: string;
}
