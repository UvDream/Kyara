import React from 'react';
import './index.less';
const Avatar = () => {
  return (
    <div className="avatar">
      <div className="avatar-img">
        <img
          src="https://gitee.com/Uvdream/images/raw/master/images/20200724164740.png"
          alt=""
        />
      </div>
      <div className="avatar-person">
        <div className="avatar-person-name">wangzhongjie</div>
        <div className="avatar-person-introduction">让技术产生价值</div>
      </div>
    </div>
  );
};
export default Avatar;
