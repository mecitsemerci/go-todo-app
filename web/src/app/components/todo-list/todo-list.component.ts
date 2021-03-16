import { Component, OnInit } from '@angular/core';
import { TodoService } from '../../services/todo-service/todo.service';
import {
  NbToastrService,
  NbComponentStatus,
  NbMenuService,
  NbWindowService,
} from '@nebular/theme';
import { CreateTodo, Todo, UpdateTodo } from '../../models/todo';
import { FormControl, FormGroup } from '@angular/forms';
import { TodoEditComponent } from '../todo-edit/todo-edit.component';
@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss'],
})
export class TodoListComponent implements OnInit {
  todos: Todo[] = [];
  hideCompleted: boolean;
  loading = false;
  contextItem: any;
  contextItems = [
    { title: 'Edit', id: 'edit' },
    { title: 'Delete', id: 'delete' },
  ];

  addNewTodoForm = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    priority_level: new FormControl(''),
  });

  checkBoxStatus = ['basic', 'primary', 'warning', 'danger'];

  constructor(
    private todoService: TodoService,
    private toastrService: NbToastrService,
    private nbMenuService: NbMenuService,
    private windowService: NbWindowService
  ) {
    this.hideCompleted = false;
  }

  ngOnInit(): void {
    this.fetchTodos();

    this.nbMenuService.onItemClick().subscribe((event: any) => {
      if (event.tag.startsWith('action-context-menu')) {
        const todoID = event.tag.split('#').pop();
        if (event.item.id === 'delete') {
          this.deleteTodo(todoID);
        } else if (event.item.id === 'edit') {
          this.todoService.getTodo(todoID).subscribe(
            (res) => {
              this.openTodoEditModal((res as Todo).id);
            },
            (err) => {
              this.showError(err.error.message);
            }
          );
        }
      }
    });
  }

  getTodos(): Todo[] {
    return this.hideCompleted
      ? this.todos.filter((t) => !t.completed)
      : this.todos;
  }

  fetchTodos(): void {
    this.loading = true;
    this.todoService.getTodos().subscribe({
      next: (todos: any) => {
        this.todos = todos as Todo[];
      },
      error: (err) => {
        this.showError(err.error.message);
      },
      complete: () => {},
    });
    setTimeout(() => (this.loading = false), 100);
  }

  completedTodoPercentage(): number {
    if (this.todos === null || this.todos.length === 0) {
      return 0;
    }

    const percentage =
      (this.todos.filter((todo) => todo.completed).length / this.todos.length) *
      100;
    return Math.min(Math.max(parseFloat(percentage.toFixed()), 0), 100);
  }

  completeTask(event: any, todo: Todo): void {
    todo.completed = event.target.checked;

    const updateTodo: UpdateTodo = {
      title: todo.title,
      description: todo.description,
      completed: event.target.checked,
      priority_level: todo.priority_level,
    };

    this.todoService.updateTodo(todo.id, updateTodo).subscribe({
      next: (d) => {
        this.showInfo('The item has been updated');
      },
      error: (err) => {
        this.showError(err.error.message);
        console.error(err);
      },
      complete: () => this.fetchTodos(),
    });
  }

  enterSubmit(event: any): void {
    if (event.keyCode === 13) {
      this.addNewTodoSubmit(event);
    }
  }

  addNewTodoSubmit(event: any): void {
    event.preventDefault();
    this.addNewTodo();
  }

  addNewTodo(): void {
    const createTodo: CreateTodo = {
      title: this.addNewTodoForm.value.title,
      description: '',
      priority_level: 0,
    };

    this.todoService.addTodo(createTodo).subscribe({
      next: (d) => {
        if (d.todo_id) {
          this.showSuccess('The item has been added');
          this.clearForm();
          this.fetchTodos();
        }
      },
      error: (err) => {
        this.showError(err.error.message);
        console.error(err);
      },
      complete: () => {},
    });
  }

  clearForm(): void {
    this.addNewTodoForm.reset();
  }

  showToast(message: string, title: string, nbStatus: NbComponentStatus): void {
    setTimeout(
      () => this.toastrService.show(message, title, { status: nbStatus }),
      200
    );
  }

  showSuccess(message: string): void {
    this.showToast(message, 'Success', 'success');
  }
  showError(message: string): void {
    this.showToast(message, 'Error', 'danger');
  }
  showInfo(message: string): void {
    this.showToast(message, 'Info', 'info');
  }

  deleteTodo(id: string): void {
    this.todoService.deleteTodo(id).subscribe({
      next: (d) => {},
      error: (err) => {
        this.showError(err.error.message);
        console.error(err);
      },
      complete: () => this.fetchTodos(),
    });
  }

  openTodoEditModal(id: string): void {
    this.windowService
      .open(TodoEditComponent, {
        title: `Edit`,
        context: { todoID: id },
      })
      .onClose.subscribe(() => this.fetchTodos());
  }

  hideCompletedItems(): void {
    this.getTodos();
  }
}
