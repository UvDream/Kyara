import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { HotArticleComponent } from './hot-article.component';

describe('HotArticleComponent', () => {
  let component: HotArticleComponent;
  let fixture: ComponentFixture<HotArticleComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ HotArticleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HotArticleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
