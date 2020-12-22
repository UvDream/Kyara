import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { BlogMessageComponent } from './blog-message.component';

describe('BlogMessageComponent', () => {
  let component: BlogMessageComponent;
  let fixture: ComponentFixture<BlogMessageComponent>;

  beforeEach(waitForAsync(() => {
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
