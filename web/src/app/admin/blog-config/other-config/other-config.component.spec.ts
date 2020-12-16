import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OtherConfigComponent } from './other-config.component';

describe('OtherConfigComponent', () => {
  let component: OtherConfigComponent;
  let fixture: ComponentFixture<OtherConfigComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OtherConfigComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OtherConfigComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
