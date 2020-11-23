import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LeftFooterComponent } from './left-footer.component';

describe('LeftFooterComponent', () => {
  let component: LeftFooterComponent;
  let fixture: ComponentFixture<LeftFooterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LeftFooterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LeftFooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
