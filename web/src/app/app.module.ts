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
  NbAlertModule,
  NbToastrService,
  NbToastrModule,
  NbContextMenuModule,
  NbMenuService,
  NbMenuModule,
  NbWindowModule,
  NbWindowService,
  NbProgressBarModule,
  NbRadioModule,
} from '@nebular/theme';
import { NbEvaIconsModule } from '@nebular/eva-icons';
import { TodoListComponent } from './components/todo-list/todo-list.component';
import { TodoService } from './services/todo-service/todo.service';
import { ReactiveFormsModule } from '@angular/forms';
import { TodoEditComponent } from './components/todo-edit/todo-edit.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
@NgModule({
  declarations: [AppComponent, TodoListComponent, TodoEditComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    NbThemeModule.forRoot({ name: 'default' }),
    NbMenuModule.forRoot(),
    NbWindowModule.forRoot({}),
    NbToastrModule.forRoot({}),
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
    NbAlertModule,
    ReactiveFormsModule,
    NbContextMenuModule,
    NbProgressBarModule,
    NgbModule,
    NbRadioModule,
  ],
  providers: [
    TodoService,
    NbSidebarService,
    NbToastrService,
    NbMenuService,
    NbWindowService,
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
