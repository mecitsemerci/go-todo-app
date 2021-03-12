import { Component } from '@angular/core';
import { NbSidebarService } from '@nebular/theme';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'Todo App';

  constructor(private sidebarService: NbSidebarService){}

  toggle() {
    this.sidebarService.toggle(true);
    return false;
  }
}
