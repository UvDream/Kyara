import React, { useState } from 'react';
import {
  VerticalRightOutlined,
  PlayCircleOutlined,
  PauseCircleOutlined,
  VerticalLeftOutlined,
} from '@ant-design/icons';
import './index.less';
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const AudioPlay = () => {
  const [status, setStatus] = useState(true)
  const [volume, setVolume] = useState(true)

  const beginFunc = (data: boolean) => {
    setStatus(data)
  }
  const volFunc = (data: boolean) => {
    setVolume(data)
  }
  return (
    <div className="audio">
      <div className="audio-top">
        <div className="audio-left imgRotate">
          <img
            src="https://dss2.bdstatic.com/6Ot1bjeh1BF3odCf/it/u=3604607153,123557897&fm=74&app=80&f=JPEG&size=f121,121?sec=1880279984&t=adfb89c635b6e6a43c06633a6299412f"
            alt=""
          />
        </div>
        <div className="audio-right">
          <marquee className="audio-right-title">这是音乐标题</marquee>
          <div className="audio-right-tool">
            <VerticalRightOutlined />
            {
              status ? <PauseCircleOutlined onClick={() => { beginFunc(false) }} /> : < PlayCircleOutlined onClick={() => { beginFunc(true) }} />
            }
            <VerticalLeftOutlined />
            {
              volume ? <IconFont type="uv-volume-high-outline" onClick={() => { volFunc(false) }} /> : <IconFont type="uv-volume-mute-outline" onClick={() => { volFunc(true) }} />
            }
          </div>
        </div>
      </div>
      <div className="audio-schedule"></div>
    </div>
  );

};
export default AudioPlay;
