import React from 'react';
import "./hader.less"
import { Popover, Input } from "antd"
import { SlidersOutlined, SearchOutlined } from "@ant-design/icons"
import Trend from './trend/index'
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const { Search } = Input;
const SearchArea = () => {
  return (
    <div className="search">
      <Popover placement="bottomLeft" content={Trend} trigger="click">
        <IconFont type="uv-pulse-outline" className="search-icon" />
      </Popover>
      <div className="search-input">
        <input type="text" placeholder="请输入搜索内容" />
        <SearchOutlined className="search-input-icon" />
      </div>
    </div>
  )
}
export default SearchArea
