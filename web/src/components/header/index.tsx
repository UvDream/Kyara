import React from 'react';
import './hader.less';
import { Col, Row } from 'antd';
import Logo from './logo';
import SearchArea from './search';
import AudioPlay from './audio/index';
import User from './user/index';
const Header = () => {
  return (
    <div className="header">
      <Row style={{ height: '50px' }}>
        <Col xl={15} lg={15}
          sm={0} xs={0} className="header-search">
          <SearchArea />
        </Col>
        <Col xl={5} lg={5} sm={0} xs={0} className="header-video">
          <AudioPlay />
        </Col>

        <Col xl={4} sm={0} xs={0} className="header-user">
          <User />
        </Col>
      </Row>
    </div>
  );
};
export default Header;
