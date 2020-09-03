import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BlogMessageComponent } from './blog-message.component';

describe('BlogMessageComponent', () => {
  let component: BlogMessageComponent;
  let fixture: ComponentFixture<BlogMessageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BlogMessageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BlogMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
