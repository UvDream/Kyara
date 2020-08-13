import React from 'react';
import TabList from './tab';
import BlogMessage from './blog-message';
import TagList from './tag-list';
import './index.less';
import '../../../styles/bootstrap.less';
const RightContent = () => {
  return (
    <div className="main-right">
      <TabList />
      <BlogMessage />
      <TagList />
    </div>
  );
};
export default RightContent;
