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
    selector: 'app-landing',
    templateUrl: './landing.component.html',
    styleUrls: ['./landing.component.scss']
})
export class LandingComponent implements OnInit {
    show = true;
    matched_tutors!: TutorType[];

    constructor(private httpClient: HttpClient) {
        this.httpClient = httpClient;
      }

  ngOnInit(): void {
    this.httpClient.post<TutorMatchedRecords>('http://localhost:8090/api/tutors', {
      primary_subject: '', secondary_subject: 'Civil'
    }).subscribe(data => { 
        this.matched_tutors = data.matched_tutors;
     })
  }

}
