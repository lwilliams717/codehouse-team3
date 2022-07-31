import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { HttpClient } from '@angular/common/http';

export interface User {
  first: string;
  last: string;
}

export interface CustomerType {
  text: string;
}


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {

  title = 'myapp';
  currentCustomer: CustomerType = {text: ''};
  displayedColumns: string[] = ['first', 'last'];
  users = new MatTableDataSource<User>();
  httpClient: HttpClient;

  constructor(httpClient: HttpClient) {
    this.httpClient = httpClient;
  }

  ngOnInit(): void {
    this.users.data = [
      { first: 'OneFirst', last: 'OneLast' },
      { first: 'TwoFirst', last: 'TwoLast' }
    ];
    
    this.httpClient.get('http://localhost:8090/').subscribe(resp => {
      console.log(resp);
    })
    
  }
}

