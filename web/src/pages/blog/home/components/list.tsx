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
        id: props.data.id,
      },
    });
  }
  return (
    <>
      {/*  置顶单图*/}
      {props.data.cover_type === '2' ? (
        <div
          className="list-img"
          style={{ backgroundImage: `url(${props.data.img_url})` }}
          onClick={toDetail}
        >
          {/*<img src={props.data.img_url} alt=""/>*/}
          <header className="list-header">
            <div className="list-title">
              <span>置顶</span>
              {props.data.title}
            </div>
            <div className="list-summary">{props.data.introduction}</div>
          </header>
        </div>
      ) : (
          ''
        )}

      {/*  Icon */}
      {props.data.cover_type === '1' ? (
        <div className="list-normal" onClick={toDetail}>
          <div className="list-normal-top">
            <div className="list-normal-top-icon">
              <IconFont type={props.data.icon} />
            </div>
            <h2 className="list-normal-top-title">{props.data.title}</h2>
            <p className="list-normal-top-summary text-ellipsis-two">
              {props.data.introduction}
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
      {props.data.cover_type === '3' ? (
        <div className="list-leftImg" onClick={toDetail}>
          {/* <img src={props.data.img_url} alt="" /> */}
          <div
            className="list-leftImg-img"
            style={{ backgroundImage: `url(${props.data.img_url})` }}
          ></div>
          <div className="list-leftImg-content">
            <h2>{props.data.title}</h2>
            <p className="list-leftImg-content-summary text-ellipsis-two">
              {props.data.introduction}
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
      {props.data.cover_type === '0' ? (
        <div className="list-text" onClick={toDetail}>
          <h2>{props.data.title}</h2>
          <p className="text-ellipsis-two list-text-summary">
            {props.data.introduction}
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
      {props.data.cover_type === '4' ? (
        <div className="list-topImg" onClick={toDetail}>
          <div
            className="list-topImg-img"
            style={{ backgroundImage: `url(${props.data.img_url})` }}
          ></div>
          <div className="list-topImg-content">
            <h2>{props.data.title}</h2>
            <p className="text-ellipsis-two">{props.data.introduction}</p>
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
