import React from 'react';
import { Tabs } from 'antd';
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const Article = (
  <IconFont
    type="uv-thumbs-up-outline"
    style={{ fontSize: '18px', marginLeft: '12px' }}
  />
);
const Comment = (
  <IconFont
    type="uv-chatbubble-ellipses-outline"
    style={{ fontSize: '18px', marginLeft: '12px' }}
  />
);
const Random = (
  <IconFont
    type="uv-star-outline"
    style={{ fontSize: '18px', marginLeft: '12px' }}
  />
);
const { TabPane } = Tabs;
const TabList = () => {
  return (
    <Tabs defaultActiveKey="1" centered>
      <TabPane tab={Article} key="1">
        Content of Tab Pane 1
      </TabPane>
      <TabPane tab={Comment} key="2">
        Content of Tab Pane 2
      </TabPane>
      <TabPane tab={Random} key="3">
        Content of Tab Pane 3
      </TabPane>
    </Tabs>
  );
};
export default TabList;
