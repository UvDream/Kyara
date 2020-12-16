import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BlogAboutComponent } from './blog-about.component';

describe('BlogAboutComponent', () => {
  let component: BlogAboutComponent;
  let fixture: ComponentFixture<BlogAboutComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BlogAboutComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BlogAboutComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
