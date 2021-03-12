import { Component, OnInit } from '@angular/core';
import {
  NbComponentStatus,
  NbToastrService,
  NbWindowRef,
} from '@nebular/theme';
import { TodoService } from 'src/app/services/todo-service/todo.service';
import { FormControl, FormGroup } from '@angular/forms';
import { UpdateTodo } from 'src/app/models/todo';

@Component({
  selector: 'app-todo-edit',
  templateUrl: './todo-edit.component.html',
  styleUrls: ['./todo-edit.component.scss'],
})
export class TodoEditComponent implements OnInit {
  todoID = '';

  options = [
    { value: '0', label: 'None' },
    { value: '1', label: 'Normal' },
    { value: '2', label: 'High' },
    { value: '3', label: 'Critical' },
  ];

  updateTodoForm = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    priority_level: new FormControl('0'),
    completed: new FormControl(''),
  });


  constructor(
    protected windowRef: NbWindowRef,
    private todoService: TodoService,
    private toastrService: NbToastrService
  ) {}

  ngOnInit(): void {
    if (this.todoID === '') {
      return;
    }

    this.todoService.getTodo(this.todoID).subscribe({
      next: (todo) => {
        this.updateTodoForm.setValue({
          title: todo.title,
          description: todo.description,
          priority_level: todo.priority_level.toString(),
          completed: todo.completed,
        });
      },
      error: (err) => {
        this.showError(err.error.message);
        console.log(err);
      },
      complete: () => {},
    });
  }

  updateTodoSubmit(): void {
    console.log('submitted', this.updateTodoForm.getRawValue());
    const updatedTodo: UpdateTodo = {
      title: this.updateTodoForm.value.title,
      description: this.updateTodoForm.value.description,
      priority_level: +this.updateTodoForm.value.priority_level,
      completed: this.updateTodoForm.value.completed,
    };

    this.todoService.updateTodo(this.todoID, updatedTodo).subscribe({
      next: (d) => {
        this.showSuccess('The item has been updated successfully.');
      },
      error: (err) => {
        this.showError(err.error.message);
      },
      complete: () => {},
    });
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
}
