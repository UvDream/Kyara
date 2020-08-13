import React from 'react';
import "./index.less"
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const User = () => {
    return (
        <div className="user">
            <IconFont type="uv-reorder-four-outline" style={{ fontSize: "18px" }} />
            <IconFont type="uv-chatbox-ellipses-outline" style={{ fontSize: "18px" }} />
            <IconFont type="uv-person-outline" style={{ fontSize: "18px" }} />
        </div>
    )
}
export default User;