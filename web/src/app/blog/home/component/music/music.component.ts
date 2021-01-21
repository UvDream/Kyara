import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-music',
  templateUrl: './music.component.html',
  styleUrls: ['./music.component.less']
})
export class MusicComponent implements OnInit {
  public AudioStatus = false;
  public AudioVolume = 1;
  public Audio = new Audio();
  constructor() { }

  ngOnInit(): void {
    this.Audio.src = 'https://webfs.yun.kugou.com/202101211452/937bafd6c76b6aface90c6c16c47e5c5/G239/M00/0D/04/j4cBAF-87BmAQpz0ACkRT30fEHU644.mp3';
  }
  playFunc(item: boolean): void {
    this.AudioStatus = item;
    item ? this.Audio.play() : this.Audio.pause();
  }
  volumeFunc(item: boolean): void {
    console.log('静音');
    item ? this.Audio.volume = 0 : this.Audio.volume = 1;
    this.AudioVolume = this.Audio.volume;
  }

}
