import { Injectable, ErrorHandler } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, map, retry } from 'rxjs/operators';
import {
  Todo,
  CreateTodo,
  CreateTodoResult,
  UpdateTodo,
} from '../../models/todo';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class TodoService {
  headers = { 'content-type': 'application/json' };

  constructor(private http: HttpClient) {}

  getTodos(): Observable<Todo[]> {
    return this.http.get<Todo[]>(`${environment.apiURL}/api/v1/todo`);
  }

  addTodo(todo: CreateTodo): Observable<CreateTodoResult> {
    return this.http.post<CreateTodoResult>(
      `${environment.apiURL}/api/v1/todo`,
      JSON.stringify(todo),
      {
        headers: this.headers,
      }
    );
  }

  updateTodo(todoID: string, todo: UpdateTodo): Observable<any> {
    return this.http.put<any>(
      `${environment.apiURL}/api/v1/todo/${todoID}`,
      JSON.stringify(todo),
      {
        headers: this.headers,
      }
    );
  }

  deleteTodo(todoID: string): Observable<any> {
    return this.http.delete<any>(
      `${environment.apiURL}/api/v1/todo/${todoID}`,
      {
        headers: this.headers,
      }
    );
  }

  getTodo(todoID: string): Observable<Todo> {
    return this.http.get<Todo>(`${environment.apiURL}/api/v1/todo/${todoID}`, {
      headers: this.headers,
    });
  }
}
