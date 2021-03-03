import { Component, OnInit } from '@angular/core';
import { TodoService } from '../../services/todo-service/todo.service';
import { Todo } from '../../models/todo';
@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss'],
})
export class TodoListComponent implements OnInit {
  todos: Todo[] = [];
  loading = false;
  selectedItem = '0';
  constructor(private todoService: TodoService) {}

  ngOnInit(): void {
    this.fetchTodos();
  }

  fetchTodos(): void {
    this.loading = true;
    this.todoService
      .getTodos()
      .subscribe((todos) => (this.todos = todos as Todo[]));
    setTimeout(() => (this.loading = false), 100);
  }

  priorityChange(event:any): void {
    console.log('Prior:', event.target.value);
  }

  completeTask(event: any): void {
    console.log('Checked:', event.target.checked);
  }
}
