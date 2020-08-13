import React from 'react';
import './index.less';
import { createFromIconfontCN } from '@ant-design/icons';
import { history } from 'umi';
const IconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});

/**
 * @param: type
 * 置顶单图 top
 * icon文字 icon
 * 正常文字 normal
 * 左图右文 leftImg
 * 上图下文 topImg
 */
const List = (props: any) => {
  const toDetail = () => {
    history.push({
      pathname: '/index/detail',
      query: {
        a: props.id,
      },
    });
  }
  return (
    <>
      {/*  置顶单图*/}
      {props.data.type === 'top' ? (
        <div
          className="list-img"
          style={{ backgroundImage: `url(${props.data.img})` }}
          onClick={toDetail}
        >
          {/*<img src={props.data.img} alt=""/>*/}
          <header className="list-header">
            <div className="list-title">
              <span>置顶</span>
              {props.data.title}
            </div>
            <div className="list-summary">{props.data.summary}</div>
          </header>
        </div>
      ) : (
          ''
        )}

      {/*  Icon */}
      {props.data.type === 'icon' ? (
        <div className="list-normal" onClick={toDetail}>
          <div className="list-normal-top">
            <div className="list-normal-top-icon">
              <IconFont type={props.data.icon} />
            </div>
            <h2 className="list-normal-top-title">{props.data.title}</h2>
            <p className="list-normal-top-summary text-ellipsis-two">
              {props.data.summary}
            </p>
          </div>
          <div className="list-normal-split"></div>
          <div className="list-normal-bottom">
            <IconFont
              type="uv-person-outline"
              className="list-normal-bottom-iconO"
            />
            <span>{props.data.auth}</span>
            <IconFont
              type="uv-time-outline"
              className="list-normal-bottom-icon"
            />
            <span>{props.data.time}</span>
            <IconFont
              type="uv-chatbubble-outline"
              className="list-normal-bottom-icon"
            />
            <span>{props.data.comment}条评论</span>
          </div>
        </div>
      ) : (
          ''
        )}
      {/* 左图右文 */}
      {props.data.type === 'leftImg' ? (
        <div className="list-leftImg" onClick={toDetail}>
          {/* <img src={props.data.img} alt="" /> */}
          <div
            className="list-leftImg-img"
            style={{ backgroundImage: `url(${props.data.img})` }}
          ></div>
          <div className="list-leftImg-content">
            <h2>{props.data.title}</h2>
            <p className="list-leftImg-content-summary text-ellipsis-two">
              {props.data.summary}
            </p>
            <div className="list-normal-split"></div>
            <div className="list-normal-bottom ">
              <IconFont
                type="uv-person-outline"
                className="list-normal-bottom-iconO"
              />
              <span>{props.data.auth}</span>
              <IconFont
                type="uv-time-outline"
                className="list-normal-bottom-icon"
              />
              <span>{props.data.time}</span>
              <IconFont
                type="uv-chatbubble-outline"
                className="list-normal-bottom-icon"
              />
              <span>{props.data.comment}条评论</span>
            </div>
          </div>
        </div>
      ) : (
          ''
        )}
      {/* 正常文字 */}
      {props.data.type === 'normal' ? (
        <div className="list-text" onClick={toDetail}>
          <h2>{props.data.title}</h2>
          <p className="text-ellipsis-two list-text-summary">
            {props.data.summary}
          </p>
          <div className="list-normal-split"></div>
          <div className="list-normal-bottom">
            <IconFont
              type="uv-person-outline"
              className="list-normal-bottom-iconO"
            />
            <span>{props.data.auth}</span>
            <IconFont
              type="uv-time-outline"
              className="list-normal-bottom-icon"
            />
            <span>{props.data.time}</span>
            <IconFont
              type="uv-chatbubble-outline"
              className="list-normal-bottom-icon"
            />
            <span>{props.data.comment}条评论</span>
          </div>
        </div>
      ) : (
          ''
        )}
      {/*  上图下文*/}
      {props.data.type === 'topImg' ? (
        <div className="list-topImg" onClick={toDetail}>
          <div
            className="list-topImg-img"
            style={{ backgroundImage: `url(${props.data.img})` }}
          ></div>
          <div className="list-topImg-content">
            <h2>{props.data.title}</h2>
            <p className="text-ellipsis-two">{props.data.summary}</p>
            <div className="list-normal-split"></div>
            <div className="list-normal-bottom">
              <IconFont
                type="uv-person-outline"
                className="list-normal-bottom-iconO"
              />
              <span>{props.data.auth}</span>
              <IconFont
                type="uv-time-outline"
                className="list-normal-bottom-icon"
              />
              <span>{props.data.time}</span>
              <IconFont
                type="uv-chatbubble-outline"
                className="list-normal-bottom-icon"
              />
              <span>{props.data.comment}条评论</span>
            </div>
          </div>
        </div>
      ) : (
          ''
        )}
    </>
  );
};
export default List;
