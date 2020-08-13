import React from 'react';
import './index.less';
import LeftFooter from './left-footer';
import Avatar from './avatar';
import Navigation from './navigation';
import Logo from '../../../components/header/logo';
import OtherNavigation from "./other-navigation.tsx"
const LeftMenus = () => {
  return (
    <aside className="aside">
      <Logo />
      <Avatar />
      <Navigation />
      <OtherNavigation />
      <LeftFooter />
    </aside>
  );
};
export default LeftMenus;
