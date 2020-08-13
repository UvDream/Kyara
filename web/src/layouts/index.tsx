import React, { useEffect } from 'react';
import './index.less';
import { Row, Col, BackTop } from 'antd';
import Headers from '../components/header/index';
import LeftMenus from './components/left-menus/index';
import TabList from './components/right-content/index';
import BlogMessage from './components/right-content/blog-message';
import RightContent from './components/right-content/index';
import RightFooter from './components/right-content/right-foot';
import '../styles/bootstrap.less';
import { initTheme } from '@/utils/theme';
export default (props: any) => {
  useEffect(() => {
    initTheme(true)
  }, [])
  return (
    <Row className="body">
      <Col xxl={18} xl={20} lg={22} md={24} sm={24} xs={24} className="main">
        <div className="main-left">
          <LeftMenus />
        </div>
        <div className="main-center">
          <Headers />
          <div style={{ display: 'flex' }}>
            {props.children}
            <RightContent />
          </div>
          <RightFooter />
          <BackTop />
        </div>
      </Col>
    </Row>
  );
};
