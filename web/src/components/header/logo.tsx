import React from 'react';
import "./hader.less"
import { HomeOutlined } from "@ant-design/icons"
import { history } from 'umi';
const Logo = () => {
  const toHome = () => {
    history.push('/index/home');
  }
  return (
    <div className="logo" onClick={toHome}>
      <HomeOutlined />
      <span>U世界的V梦想</span>
    </div>
  )
}
export default Logo
