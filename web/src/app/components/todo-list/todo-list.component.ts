import { Component, Input, OnInit } from '@angular/core';
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
  selectedPriorityLevel;
  todos: Todo[] = [];
  loading = false;
  priorityLevels = [
    { value: 0, name: 'None' },
    { value: 1, name: 'Normal' },
    { value: 2, name: 'High' },
    { value: 3, name: 'Critical' },
  ];
  contextItem: any;
  contextItems = [
    { title: 'Details', id: 'details' },
    { title: 'Delete', id: 'delete' },
  ];

  addNewTodoForm = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    priority_level: new FormControl(''),
  });

  constructor(
    private todoService: TodoService,
    private toastrService: NbToastrService,
    private nbMenuService: NbMenuService,
    private windowService: NbWindowService
  ) {
    this.selectedPriorityLevel = this.priorityLevels[0];
  }

  ngOnInit(): void {
    this.fetchTodos();

    this.nbMenuService.onItemClick().subscribe((event: any) => {
      if (event.tag.startsWith('action-context-menu')) {
        let todoID = event.tag.split('-').pop();
        if (event.item.id === 'delete') {
          this.deleteTodo(todoID);
        } else if (event.item.id === 'details') {
          let result = this.todoService.getTodo(todoID).subscribe(
            (res) => {
              this.openTodoEditModal((res as Todo).id);
            },
            (err) => {
              this.showToast(err.error.message, 'danger');
            }
          );
        }
      }
    });
  }

  fetchTodos(): void {
    this.loading = true;
    this.todoService
      .getTodos()
      .subscribe((todos) => (this.todos = todos as Todo[]));
    setTimeout(() => (this.loading = false), 100);
  }

  priorityChange(event: any): void {
    console.log('Prior:', event.target.value);
  }

  completeTask(event: any, todo: Todo): void {
    todo.completed = event.target.checked;

    let updateTodo: UpdateTodo = {
      title: todo.title,
      description: todo.description,
      completed: event.target.checked,
      priority_level: todo.priority_level,
    };

    this.todoService.updateTodo(todo.id, updateTodo).subscribe(
      (res) => {},
      (err) => {
        this.showToast(err.error.message, 'danger');
      },
      () => this.fetchTodos()
    );
  }

  addNewTodoSubmit(): void {
    let createTodo: CreateTodo = {
      title: this.addNewTodoForm.value.title,
      description: this.addNewTodoForm.value.description,
      priority_level: this.addNewTodoForm.value.priority_level.value,
    };

    this.todoService.addTodo(createTodo).subscribe(
      (res) => {
        if (res.todo_id) {
          this.showToast('Success', 'success');
          this.fetchTodos();
        }
      },
      (err) => {
        this.showToast(err.error.message, 'danger');
      }
    );
  }

  showToast(message: string, status: NbComponentStatus) {
    this.toastrService.show(message, 'Message', { status: status });
  }

  deleteTodo(id: string): void {
    this.todoService.deleteTodo(id).subscribe(
      (res) => {},
      (err) => {
        this.showToast(err.error.message, 'danger');
      },
      () => this.fetchTodos()
    );
  }

  openTodoEditModal(id: string) {
    this.windowService.open(TodoEditComponent, { title: `Window ${id}` });
  }
}
