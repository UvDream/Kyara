import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LeftMenusComponent } from './left-menus.component';

describe('LeftMenusComponent', () => {
  let component: LeftMenusComponent;
  let fixture: ComponentFixture<LeftMenusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LeftMenusComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LeftMenusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
