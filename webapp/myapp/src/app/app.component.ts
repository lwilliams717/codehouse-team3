import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';

export interface User {
  first: string;
  last: string;
}


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {

  title = 'myapp';
  currentCustomer = 'Test';
  displayedColumns: string[] = ['first', 'last'];
  users = new MatTableDataSource<User>();

  ngOnInit(): void {
    this.users.data = [
      { first: 'OneFirst', last: 'OneLast' },
      { first: 'TwoFirst', last: 'TwoLast' }
    ];
  }
}

