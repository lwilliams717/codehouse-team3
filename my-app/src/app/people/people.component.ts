import { Component, OnInit } from '@angular/core';


import { HttpClient } from '@angular/common/http';

export interface TutorType {
    id: string,
    user_type: string,
    name: string,
    primary_subject: string,
    secondary_subject: string,
    description: string
  }
  
  export interface TutorMatchedRecords {
    matched_tutors: TutorType[]
  }


@Component({
    selector: 'app-people',
    templateUrl: './people.component.html',
    styleUrls: ['./people.component.scss']
})
export class PeopleComponent implements OnInit {
    show = true;
    tutors!: TutorType[];

    constructor(private httpClient: HttpClient) {
        this.httpClient = httpClient;
      }

  ngOnInit(): void {
    this.httpClient.post<TutorMatchedRecords>('http://localhost:8090/api/tutors', {
      primary_subject: '', secondary_subject: 'Physics'
    }).subscribe(data => { 
        this.tutors = data.matched_tutors;
     })
  }

}
