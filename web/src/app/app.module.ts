import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  NbThemeModule,
  NbLayoutModule,
  NbButtonModule,
  NbSidebarModule,
  NbCardModule,
  NbListModule,
  NbCheckboxModule,
  NbSpinnerModule,
  NbIconModule,
  NbBadgeModule,
  NbToggleModule,
  NbSidebarService,
  NbActionsModule,
  NbButtonGroupModule,
  NbInputModule,
  NbSelectModule,
  NbAlertModule
} from '@nebular/theme';
import { NbEvaIconsModule } from '@nebular/eva-icons';
import { TodoListComponent } from './components/todo-list/todo-list.component';
import { TodoService } from './services/todo-service/todo.service';

@NgModule({
  declarations: [AppComponent, TodoListComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    NbThemeModule.forRoot({ name: 'corporate' }),
    NbLayoutModule,
    NbEvaIconsModule,
    NbButtonModule,
    NbSidebarModule,
    NbCardModule,
    NbListModule,
    NbCheckboxModule,
    NbSpinnerModule,
    NbIconModule,
    NbBadgeModule,
    NbToggleModule,
    NbActionsModule,
    NbButtonGroupModule,
    NbInputModule,
    NbSelectModule,
    NbAlertModule
  ],
  providers: [TodoService, NbSidebarService],
  bootstrap: [AppComponent],
})
export class AppModule {}
